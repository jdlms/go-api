FROM golang

# Install wget and Dockerize
RUN apt-get update && apt-get install -y wget \
    && wget https://github.com/jwilder/dockerize/releases/download/v0.6.1/dockerize-linux-amd64-v0.6.1.tar.gz \
    && tar -C /usr/local/bin -xzvf dockerize-linux-amd64-v0.6.1.tar.gz \
    && rm dockerize-linux-amd64-v0.6.1.tar.gz

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN go build -o main ./cmd/server/main.go

EXPOSE 8080

# Use Dockerize to wait for the DB
CMD ["dockerize", "-wait", "tcp://db:5432", "-timeout", "60s", "/app/main"]
