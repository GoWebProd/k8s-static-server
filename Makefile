IMAGE = "ghcr.io/GoWebProd/k8s-static-server"
TAG = "v0.0.1"

image:
	docker build --push -t $(IMAGE):$(TAG) .