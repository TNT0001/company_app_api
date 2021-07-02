#Image
PROXY_IMAGE=nginx:1.17.4-alpine
GO_IMAGE=golang:1.13.5-alpine3.11
DB_IMAGE=mysql:5.7
REDIS_IMAGE=redis:5.0.6-alpine
PROD_IMAGE=alpine:3.11.2

#default value
APPLICATION=go-api
SERVICE=api
APP_PORT=5000
DOMAIN_UPSTREAM=app
NAMESERVERS=127.0.0.11
BUILD_TARGET=local
DB_PORT=3306
NGINX_PORT=81
MYSQL_ROOT_PASSWORD=password
MYSQL_DATABASE=go_test
MYSQL_USER=go_test
MYSQL_PASSWORD=go_test
WORKDIR=/${APPLICATION}

CUR_DIR=$(shell pwd)
VOLUME_APP=${CUR_DIR}:/${APPLICATION}

#name image and container
DB_CONT_NAME=${APPLICATION}-db-cont
APP_IMAGE_NAME=${APPLICATION}/go-${SERVICE}
APP_CONT_NAME=${APPLICATION}-go-${SERVICE}-cont
NGINX_IMAGE_NAME=${APPLICATION}/nginx:latest
NGINX_CONT_NAME=${APPLICATION}-nginx-cont

#args for make
ifdef env
ENVIRONMENT=${env}
else
ENVIRONMENT=local
endif

ifdef target
BUILD_TARGET=${target}
else
BUILD_TARGET=local
endif

ifdef db_open_port
DB_PORT_OPTION=-p ${db_open_port}:${DB_PORT}
else
DB_PORT_OPTION=-p ${DB_PORT}:${DB_PORT}
endif

ifdef app_open_port
APP_PORT_OPTION=-p ${app_open_port}:${APP_PORT}
else
APP_PORT_OPTION=-p ${APP_PORT}:${APP_PORT}
endif

ifdef nginx_open_port
NGINX_PORT_OPTION=-p ${nginx_open_port}:${NGINX_PORT}
else
NGINX_PORT_OPTION=-p ${NGINX_PORT}:${NGINX_PORT}
endif

ifdef version_image
VERSION=${version_image}
else
VERSION=0.0.1
endif

# Build args for Dockerfiles
BUILD_APP_ARGS=--build-arg IMAGE_NAME=${GO_IMAGE} --build-arg ENVIRONMENT=${ENVIRONMENT} --build-arg PORT=${APP_PORT} --build-arg SERVICE=${SERVICE} --build-arg APP_NAME=${APPLICATION} --build-arg PROD_IMAGE=${PROD_IMAGE} --build-arg WORKDIR=${WORKDIR}
BUILD_PROXY_ARGS=--build-arg IMAGE_NAME=${PROXY_IMAGE} --build-arg DOMAIN_UPSTREAM=${DOMAIN_UPSTREAM} --build-arg NAMESERVERS=${NAMESERVERS} --build-arg PORT=${APP_PORT}

#ENV for Database
DB_ENVS=-e MYSQL_ROOT_PASSWORD=${MYSQL_ROOT_PASSWORD} -e MYSQL_DATABASE=${MYSQL_DATABASE} -e MYSQL_USER=${MYSQL_USER} -e MYSQL_PASSWORD=${MYSQL_PASSWORD}
#database
db-cont:
	@echo ":::Run database container:::"
	docker run -d --env-file=.env ${DB_PORT_OPTION} --name ${DB_CONT_NAME} ${DB_ENVS} -d ${DB_IMAGE}

remove-db-cont:
	@echo ":::Remove database container"
	-docker stop ${DB_CONT_NAME}
	-docker rm ${DB_CONT_NAME}

#app
app-image:
	@echo ":::Building go app image:::"
	docker build --target ${BUILD_TARGET} --rm -f build/api/Dockerfile ${BUILD_APP_ARGS} -t ${APP_IMAGE_NAME}:latest .

app-cont:
	@echo ":::Run go app container:::"
	docker run -v ${VOLUME_APP} -d --env-file=.env ${APP_PORT_OPTION} --link ${DB_CONT_NAME}:db --name ${APP_CONT_NAME} ${APP_IMAGE_NAME}:latest
	# docker run -v ${VOLUME_APP} -d --env-file=.env ${APP_PORT_OPTION} --name ${APP_CONT_NAME} ${APP_IMAGE_NAME}:latest

remove-app-cont:
	@echo ":::Remove app container"
	-docker stop ${APP_CONT_NAME}
	-docker rm ${APP_CONT_NAME}

#proxy
proxy-image:
	@echo ":::Build proxy image:::"
	docker build --rm -f build/nginx/Dockerfile ${BUILD_PROXY_ARGS} -t ${NGINX_IMAGE_NAME} .

proxy-cont:
	@echo ":::Run proxy cont:::"
	docker run -d --env-file=.env ${NGINX_PORT_OPTION} --link ${APP_CONT_NAME}:app --name ${NGINX_CONT_NAME} ${NGINX_IMAGE_NAME}

remove-proxy-cont:
	@echo ":::Remove proxy container"
	-docker stop ${NGINX_CONT_NAME}
	-docker rm ${NGINX_CONT_NAME}

#make
build-db: remove-db-cont db-cont
build-app: remove-app-cont app-image app-cont
build-proxy: remove-proxy-cont proxy-image proxy-cont
first-launch: build-db build-app build-proxy

start:
	@echo ":::Start container:::"
	docker start ${APP_CONT_NAME} ${NGINX_CONT_NAME}
stop:
	@echo ":::Stop container:::"
	docker stop ${APP_CONT_NAME} ${NGINX_CONT_NAME}

# make build-prod env=prod version_image=x.x.x
build-prod: 
	@echo ":::Building go app image for PROD:::"
	docker build --target prod --rm -f build/api/Dockerfile ${BUILD_APP_ARGS} -t nri_admin_api:${VERSION} .
	docker save go-api:${VERSION}>nri_admin_api_${VERSION}.tar

test-prod: 
	@echo ":::Building go app image for TEST:::"
	docker build --target test --rm -f build/api/Dockerfile ${BUILD_APP_ARGS} -t nri_admin_api:Test .
	docker rmi -f nri_admin_api:Test