FROM golang:1.14.4-alpine AS go_builder
ADD . /project
WORKDIR /project
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -ldflags "-s -w" -o ./server ./cmd/main.go

EXPOSE 3000
CMD ["./server"]
