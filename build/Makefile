update:
	make build push deploy

deploy:
	@ EXECUTE_FILE=deploy.sh "$(dir ${MAKEFILE_LIST})executor.sh"

build:
	@ EXECUTE_FILE=build.sh "$(dir ${MAKEFILE_LIST})executor.sh"

push:
	@ EXECUTE_FILE=push.sh "$(dir ${MAKEFILE_LIST})executor.sh"

.PHONY: deploy build push
