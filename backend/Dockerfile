FROM golang:1.11.5 as builder

WORKDIR /myapp
COPY . .
RUN go build -o doraku-server


FROM ubuntu:18.04

WORKDIR /myapp
COPY --from=builder /myapp/doraku-server /myapp

COPY database /myapp/database

CMD ["./doraku-server"]