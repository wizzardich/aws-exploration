FROM golang:1.17 as builder

WORKDIR /app

COPY . ./

RUN go build

FROM debian:stable-slim

COPY --from=builder /app/aws-sample /aws-sample

RUN chmod +x aws-sample

EXPOSE 8080

ENTRYPOINT /aws-sample