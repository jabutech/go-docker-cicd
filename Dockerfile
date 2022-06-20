# Use image golang
FROM golang:1.17-alpine

# Work directory app
WORKDIR /app

# Copy project into workdir /app
COPY . .

# Build go app
RUN go build -o hello-world

#  Run binary file
CMD ./hello-world


