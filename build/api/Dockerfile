ARG IMAGE_NAME
ARG PROD_IMAGE

#################################################
# Stage base for all stage
#################################################
FROM ${IMAGE_NAME} as base

ARG ENVIRONMENT
ARG PORT

ARG APP_NAME
ARG SERVICE

WORKDIR /${APP_NAME}
COPY . ./

EXPOSE ${PORT}

#################################################
# Stage for developer
#################################################
FROM base as local

ARG APP_NAME
ARG SERVICE

RUN apk --update add --no-cache bash gcc musl-dev git mercurial openssh curl && \
   go get -u -v \
   github.com/pilu/fresh

WORKDIR /${APP_NAME}/cmd/${SERVICE}

CMD fresh

#################################################
# Stage for builder
#################################################
FROM base as builder

ARG APP_NAME
ARG SERVICE

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /${APP_NAME}/cmd/${SERVICE}
RUN go build -o go-api

#################################################
# Stage for Test
#################################################
FROM builder as test

ARG APP_NAME
ARG SERVICE

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /${APP_NAME}
RUN go get -u golang.org/x/lint/golint && go vet ./... && golint ./... \
   && go test -timeout 120s ./... -coverprofile /go-code-cover \
   -coverpkg=go-api/internal/...,go-api/pkg/...

# Coverage Stage
FROM scratch AS cover
COPY --from=test /go-code-cover /go-code-cover

# ################################################
# Stage for PROD
# ################################################
FROM ${PROD_IMAGE} as prod

ARG APP_NAME
ARG SERVICE

ARG ENVIRONMENT=prod
ARG PORT
ARG WORKDIR

COPY --from=builder /${APP_NAME}/configs ./configs
COPY --from=builder /${APP_NAME}/internal/app/db/migrations ./migrations
COPY --from=builder /${APP_NAME}/cmd/${SERVICE}/adminapi ./

EXPOSE ${PORT}

RUN apk --update add tzdata \
   && cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime \
   && apk del tzdata \
   && rm -rf /var/cache/apk/*

CMD ./go-api -dirConfig ./configs -logLevel ${LOGGER_LEVEL} -logOutput ${LOGGER_OUTPUT} \
   -logFormat ${LOGGER_FORMAT} -logLocaltion ${LOGGER_LOCATION} \
   -env ${ENVIRONMENT} -port ${PORT} -workdir ${WORKDIR} \
   -dbHost ${MYSQL_HOST} -dbName ${MYSQL_DATABASE}  -dbPort ${MYSQL_PORT} \
   -dbUser ${MYSQL_USER} -dbPass ${MYSQL_PASSWORD} -dirMigration ./migrations