clean:
	cd apps/photo-storage && npm run clean;
	cd apps/photo-filter && ([ -e photo-filter ] && rm photo-filter);
	cd apps/web-client && npm run clean;

install:
	cd apps/photo-storage && npm install -q;
	cd apps/web-client && npm install -q;

lint:
	cd apps/photo-storage && npm run lint .;
	cd apps/web-client && npm run lint .;

test:
	cd apps/photo-storage && npm test;
	cd apps/web-client && npm test;

build:
	docker-compose build

tag:
	docker tag photo-filter:latest 635034346210.dkr.ecr.us-east-1.amazonaws.com/photo-filter:latest
	docker tag photo-storage:latest 635034346210.dkr.ecr.us-east-1.amazonaws.com/photo-storage:latest
	docker tag web-client:latest 635034346210.dkr.ecr.us-east-1.amazonaws.com/web-client:latest

push:
	docker push 635034346210.dkr.ecr.us-east-1.amazonaws.com/photo-filter:latest
	docker push 635034346210.dkr.ecr.us-east-1.amazonaws.com/photo-storage:latest
	docker push 635034346210.dkr.ecr.us-east-1.amazonaws.com/web-client:latest

clean-images:
	docker images -q -f "dangling=true" | xargs -I {} docker rmi {}

deploy:
	kubectl apply -f ./apps/photo-filter/k8s/deployment.yaml
	kubectl apply -f ./apps/photo-filter/k8s/service.yaml
	kubectl apply -f ./apps/photo-storage/k8s/deployment.yaml
	kubectl apply -f ./apps/photo-storage/k8s/service.yaml
	kubectl apply -f ./apps/web-client/k8s/deployment.yaml
	kubectl apply -f ./apps/web-client/k8s/service.yaml

undeploy:
	kubectl delete -f ./apps/web-client/k8s/service.yaml
	kubectl delete -f ./apps/web-client/k8s/deployment.yaml
	kubectl delete -f ./apps/photo-filter/k8s/service.yaml
	kubectl delete -f ./apps/photo-filter/k8s/deployment.yaml
	kubectl delete -f ./apps/photo-storage/k8s/service.yaml
	kubectl delete -f ./apps/photo-storage/k8s/deployment.yaml

.PHONY: clean install lint test build tag push clean deploy