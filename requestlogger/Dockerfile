FROM golang:alpine as builder
WORKDIR /src
COPY . .
RUN go mod init github.com/rconway/goutils/requestlogger
RUN go build -o requestlogger

FROM alpine
WORKDIR /app
COPY --from=builder /src/requestlogger /app/requestlogger
EXPOSE 80/tcp
ENTRYPOINT [ "/app/requestlogger" ]
