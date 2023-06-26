FROM reg01.chocofood.kz/docker-proxy/library/golang:1.20-alpine as builder
WORKDIR /src

COPY vrpSimple/ .

RUN GOOS=linux GOARCH=amd64 GO111MODULE=on go build -v -o ./build/run main.go

FROM reg01.chocofood.kz/docker-proxy/library/alpine
WORKDIR /app
RUN echo "Asia/Almaty" > /etc/timezone

COPY --from=builder /src/build/run .
ENTRYPOINT ["/app/run"]