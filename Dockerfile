FROM golang:1.22-alpine


RUN touch /tmp/runner-build-errors.log

# Install air for code live reloading
RUN go install github.com/cosmtrek/air@v1.49.0

RUN mkdir /app

WORKDIR /app


COPY . .

# Build the binary with the neccessary environment variables
RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /houdini ./cmd

# Set the binary as executable
RUN chmod +x /houdini


# Expose port 8086
EXPOSE 8086

# Run the entrypoint script
CMD [ "air", "-c" ,".air.toml"]