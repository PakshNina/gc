FROM golang:latest as builder

WORKDIR /src
COPY . .

RUN GOOS=linux go build -o app ./cmd/4_gogc_1000/
FROM golang:latest
WORKDIR /root/
COPY --from=builder /src/app .
EXPOSE 8080
CMD ["./app"]