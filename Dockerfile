ARG GOLANG_DOCKER_TAG=1.18.3-alpine3.15
ARG ALPINE_DOCKER_TAG=3.15

FROM golang:$GOLANG_DOCKER_TAG as builder

RUN apk update && apk upgrade && apk add --no-cache make

WORKDIR /build
COPY . .

ARG APPLICATION_VERSION=unknown
RUN make build OUT_PATH=/build/bin/gherkingen VERSION=$APPLICATION_VERSION

FROM alpine:$ALPINE_DOCKER_TAG

WORKDIR /app
COPY --from=builder /build/bin/gherkingen /app/gherkingen

ENTRYPOINT [ "/app/gherkingen" ]
