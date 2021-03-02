# First is the build stage, to build the binary
FROM golang:1.16.0-alpine3.13 AS build

WORKDIR /app

COPY . .

# Compiling the code
RUN go build -o /bin/capstone-server .

# Second and final stage, to copy across only the binary from the build stage
FROM golang:1.16.0-alpine3.13

# Keep the working directory as /app even in the production stage
WORKDIR /app

# Copy the binary from the build stage
COPY --from=build /bin/capstone-server .

# Expose port 8080 for the web server to run
EXPOSE 8080

# Run the web server binary
CMD ["./capstone-server"]