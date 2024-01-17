FROM golang:latest

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# working directory
WORKDIR /app

# Copy and download depedencies using go mod
COPY    go.mod .
COPY    go.sum .

# copy folder into container
COPY . .

# Build the application
RUN go build -o main .

# command to running executable file
CMD ["./main"]