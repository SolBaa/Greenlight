
FROM golang:1.16.5-alpine3.13 as build


# General setup
ENV SERVICE_NAME=greenlight
ENV APP_DIR=/go/src/github.com/solbaa/${SERVICE_NAME}
RUN apk add make git build-base
RUN go get github.com/silenceper/gowatch
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
WORKDIR ${APP_DIR}
COPY . .
RUN go mod tidy
RUN go build

# Deployable Image
FROM alpine:3.12

# Add binary to bin directory
ENV SERVICE_NAME=greenlight
ENV BUILDER_APP_DIR=/go/src/github.com/solbaa/${SERVICE_NAME}
WORKDIR /app
COPY --from=build ${BUILDER_APP_DIR}/${SERVICE_NAME} .
COPY --from=build ${BUILDER_APP_DIR} ./migrations/

ENTRYPOINT ["./marvik"]
