FROM golang:1.22

RUN apt-get update && apt-get install -y make

RUN apt-get update && apt install -y protobuf-compiler
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

RUN mkdir order_processing_service
WORKDIR /order_processing_service

COPY . .
RUN go mod download

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN make proto

EXPOSE 50000