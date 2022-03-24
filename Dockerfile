ARG GOLANG_DOCKER_TAG=1.17.8-alpine3.15
ARG ALPINE_DOCKER_TAG=3.15

FROM golang:$GOLANG_DOCKER_TAG as builder

RUN apk update && apk upgrade && apk add --no-cache make

WORKDIR /build
COPY . .

RUN make build OUT_PATH=/build/bin/gherkingen

FROM alpine:$ALPINE_DOCKER_TAG

WORKDIR /app
COPY --from=builder /build/bin/gherkingen /app/gherkingen

ENTRYPOINT [ "/app/gherkingen" ]
