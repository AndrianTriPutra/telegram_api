FROM golang:1.19 AS builder
WORKDIR /app
COPY ./../go.mod .
COPY ./../go.sum .
RUN go mod download
COPY ./../ .
RUN GOOS="linux" GOARCH=amd64 CGO_ENABLED=0 go build -o telegram_api *.go

FROM alpine:latest
WORKDIR /app
ARG ENVCONSUL_VERSION=0.6.2
RUN apk --no-cache add curl ca-certificates \
 && curl https://releases.hashicorp.com/envconsul/${ENVCONSUL_VERSION}/envconsul_${ENVCONSUL_VERSION}_linux_amd64.tgz | tar -xz
RUN apk add --no-cache tzdata

COPY --from=builder /app/telegram_api .
COPY --from=builder /app/config/.*.yaml /app/config/

CMD ["/app/telegram_api"]