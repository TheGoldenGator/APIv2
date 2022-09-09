# Specify base image for Go API
FROM golang:1.18

# Specify that we need to execute any commands in directory
WORKDIR /src

# Copy everything from this project into the filesystem of the container.
COPY . .

# Compile the binary EXE for our app.
RUN go build -o ./cmd/main ./cmd/server.go

EXPOSE 8000

# Start it
CMD [ "./cmd/main" ]
