
FROM golang:1.23-alpine3.20 AS builder

WORKDIR /app

RUN apk add --no-cache protobuf protobuf-dev git make

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN make generate-proto
RUN make build

EXPOSE 50051

CMD [ "./main" ]
