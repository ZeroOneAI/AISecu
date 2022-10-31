IMAGE_NAME=zerooneai/projectmanager

IMG=${IMAGE_NAME}:${VERSION}

update:
	make build push
	kubectl apply -f deploy/kubernetes/

build:
	docker buildx build --platform linux/amd64 . -t ${IMG}

push:
	docker push ${IMG}
