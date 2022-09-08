# Specify base image for Go API
FROM golang:1.18

# Specify that we need to execute any commands in directory
WORKDIR /src

# Copy everything from this project into the filesystem of the container.
COPY . .

# Obtain package needed to run redis commands.
RUN go get github.com/go-redis/redis

# Compile the binary EXE for our app.
RUN go build -o main ./cmd/server.go

EXPOSE 8000

# Start it
CMD [ "./cmd/server" ]
