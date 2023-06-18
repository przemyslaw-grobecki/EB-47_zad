# Dockerfile for Server Application
FROM golang:1.20

# Set the working directory
WORKDIR /app
# Copy server application code
COPY 4_zad/src .
# Install dependencies

RUN go get


# Set environment variables, if any
ENV PORT=1323

# Expose the server port
EXPOSE $PORT
# Specify the command to run the server application
CMD ["go", "run", "server.go"]