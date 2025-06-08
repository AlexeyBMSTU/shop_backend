FROM ubuntu:20.04

LABEL authors="a.naryzhnyy"

RUN apt-get update && apt-get install -y \
    curl \
    git \
    build-essential \
    ca-certificates && \
    rm -rf /var/lib/apt/lists/*

ENV GO_VERSION=1.24.4
RUN curl -fsSL https://go.dev/dl/go${GO_VERSION}.linux-amd64.tar.gz -o go.tar.gz \
    && tar -C /usr/local -xzf go.tar.gz \
    && rm go.tar.gz

ENV GOPATH=/go
ENV PATH="/usr/local/go/bin:${GOPATH}/bin:${PATH}"

WORKDIR /app

# Copy go.mod and go.sum to /app
COPY go.mod ./
COPY go.sum ./

# Download dependencies
RUN go mod download

# Copy the contents of the src directory to /app
COPY src/ ./

# Check the contents of the /app directory
RUN ls -la /app

# Compile the application
RUN go build -o app .

EXPOSE 10000

ENTRYPOINT ["/app/app"]