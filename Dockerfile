# build stage
FROM golang:1.17-alpine AS builder
WORKDIR /app

COPY . .

RUN export GOPROXY=https://goproxy.cn && go mod tidy
RUN go build -o main main.go

# run stage
FROM alpine
WORKDIR /app
COPY --from=builder /app/main .
COPY app.env .

EXPOSE 8180
CMD [ "/app/main" ]