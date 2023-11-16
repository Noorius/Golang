#FROM golang:1.20
#
#RUN apt-get update && apt-get install -y
#
#WORKDIR /HW2
#
## pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
#COPY go.mod go.sum ./
#RUN go mod download && go mod verify
#
## Copy the entire project
#COPY . .
#
## Build the main package
#RUN go build -o main ./cmd/api/
#
## Set the default command to run the application
#CMD ["./cmd/api"]
#------------------------------
#FROM golang:1.20
#
#RUN apt-get update && apt-get install -y
#
#WORKDIR /app
#COPY ./go.mod ./go.sum ./
#RUN go mod download && go mod verify
#RUN go get -v -u github.com/golang-migrate/migrate/v4
#
#COPY . .
#
#RUN go build -o main ./cmd/api/
#
#EXPOSE 8080
#
#CMD ["migrate -path ./migrations -database postgres://numismatics:numismatics@db/numismatics?sslmode=disable up"]
#CMD ["./main"]


# Use the official Go image as a base image
FROM golang:1.20 as build-stage

# Set the working directory inside the container
WORKDIR /HW2

# Copy go.mod and go.sum files to the working directory
COPY go.mod go.sum ./
COPY . .

RUN apt-get update
RUN apt-get -y install postgresql-client
RUN chmod +x wait-for-postgres.sh

# Download and install Go module dependencies
RUN go mod download

# Build the Go application
RUN go build -o main ./cmd/api/

EXPOSE 8080
EXPOSE 5436
EXPOSE 5432

# Command to run the executable
CMD ["./main"]
