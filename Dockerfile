# Use an official Golang runtime as a parent image
FROM golang:latest

# Set the working directory to /app
WORKDIR .
# Copy the current directory contents into the container at /app
COPY . .

# Download and install any required dependencies
RUN go mod download

# Build the Go app
RUN go build .

# Expose port 8080 for incoming traffic
EXPOSE 4002

# Define the command to run the app when the container starts
CMD ["/portfolio-backend"]