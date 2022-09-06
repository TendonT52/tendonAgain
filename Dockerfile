FROM golang:latest

WORKDIR /app

# Download necessary Go modules
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
RUN go build
EXPOSE 8080
CMD ["./tendonAPI"]