up: postgres && prometheus

prometheus:
	@helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
	@helm repo add stable https://charts.helm.sh/stable
	@helm repo update
	@helm install prom prometheus-community/kube-prometheus-stack -f prometheus.yaml --atomic

old-install-nginx-ingress:
	@helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx && helm repo update && helm install ingress-nginx ingress-nginx/ingress-nginx  --set port=8080

install-ingress:
	@helm repo add nginx-stable https://helm.nginx.com/stable && helm repo update && helm install ingress-nginx ingress-nginx/ingress-nginx --set port=8080

dependency-build:
	@cd app-chart && helm dependency build
app:
	@cd app-chart && helm install app . -f values.yaml


remove:
	@helm uninstall app

grafana-forward:
	@kubectl port-forward service/app-grafana 9001:80

prometheus-forward:
	@kubectl port-forward service/prometheus-operated 9091:9090	

loadtest-run:
	@k6 run --duration 9000s --vus 30 k6-script.js