package controllers

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"user-service/models"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"

	"gorm.io/gorm"
)

type UserController struct {
	DB *gorm.DB
}

var GetUsersLatency = prometheus.NewHistogramVec(
	prometheus.HistogramOpts{
		Name:    "http_request_get_users_duration_seconds",
		Help:    "Latency of get_users request in second.",
		Buckets: prometheus.LinearBuckets(0.01, 0.05, 10),
	},
	[]string{"status"},
)

func NewUserController(DB *gorm.DB) UserController {
	return UserController{DB}
}

func init() {

	prometheus.MustRegister(GetUsersLatency)
}

// Create User
func (uc *UserController) CreateUser(ctx *gin.Context) {
	var payload *models.CreateUserRequest

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	now := time.Now()
	newUser := models.User{
		Username:  payload.Username,
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Phone:     payload.Phone,
		CreatedAt: now,
		UpdatedAt: now,
	}

	result := uc.DB.Create(&newUser)
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicate key") {
			ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "User with that Username already exists"})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": newUser})
}

// Update User
func (uc *UserController) UpdateUser(ctx *gin.Context) {
	userId := ctx.Param("userId")

	var payload *models.UpdateUser
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	var updatedUser models.User
	result := uc.DB.First(&updatedUser, "id = ?", userId)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No user with that userId exists"})
		return
	}
	now := time.Now()
	userToUpdate := models.User{
		Username:  payload.Username,
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Phone:     payload.Phone,
		CreatedAt: updatedUser.CreatedAt,
		UpdatedAt: now,
	}

	uc.DB.Model(&updatedUser).Updates(userToUpdate)

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": updatedUser})
}

// Get User by ID
func (uc *UserController) FindUserById(ctx *gin.Context) {
	userId := ctx.Param("userId")

	var user models.User
	result := uc.DB.First(&user, "id = ?", userId)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No user with that ID exists"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": user})
}

// Get all users
func (uc *UserController) FindUsers(ctx *gin.Context) {

	// prepare metrics
	var status string
	timer := prometheus.NewTimer(prometheus.ObserverFunc(func(v float64) {
		GetUsersLatency.WithLabelValues(status).Observe(v)
	}))
	defer func() {
		timer.ObserveDuration()
	}()

	var page = ctx.DefaultQuery("page", "1")
	var limit = ctx.DefaultQuery("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var users []models.User
	results := uc.DB.Limit(intLimit).Offset(offset).Find(&users)
	if results.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "results": len(users), "data": users})
}

// Delete User by ID
func (uc *UserController) DeleteUser(ctx *gin.Context) {
	userId := ctx.Param("userId")

	result := uc.DB.Delete(&models.User{}, "id = ?", userId)

	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No user with that ID exists"})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
