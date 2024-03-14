up:
	docker-compose build
	docker-compose up

down:
	docker-compose down

apply:
	
	kubectl apply -f myapp-deployment.yaml
	kubectl apply -f db-deployment.yaml
	kubectl apply -f db-configmap.yaml


forward:
	kubectl port-forward service/myapp-service 8080:8080

remove-containers:
	kubectl delete deployment myapp-deployment
	kubectl delete deployment db-deployment
	kubectl delete service myapp-service
	kubectl delete service db-service
get:
	kubectl get pods
	kubectl get deployments
	kubectl get services
	kubectl get secret db-secret

build-image:
	docker build -t vdesh/app:latest .

push-image:
	docker push vdesh/app:latest