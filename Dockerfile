FROM golang:1.21-alpine3.17 as builder
# Define The current working directory
WORKDIR /app
RUN apk update; \
    apk add --no-cache \
    git \
    curl \
    vim \
    make 
COPY .env ./
COPY . ./

# Download Go modules
COPY go.mod ./
COPY go.sum ./
RUN go mod download
# use The arm64 as the CPU architecture
ENV CGO_ENABLED=0
ENV GOOS=linux
# Build App
RUN go build -a -v main.go

# Define App Image
FROM golang:1.21-alpine3.17 as release
# Copy App binary to Image
COPY --from=builder /app/main /app/
COPY --from=builder /app/.env /app/
RUN chmod +x /app/main
WORKDIR /app
RUN mkdir uploads
CMD ["./main", "server"]