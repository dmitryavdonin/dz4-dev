kubectl create -f 01-postgres-secret.yaml
kubectl apply -f 02-postgres-db-pv.yaml
kubectl apply -f 03-postgres-db-pvc.yaml
kubectl apply -f 04-postgres-db-deployment.yaml
kubectl apply -f 05-postgres-db-service.yaml
