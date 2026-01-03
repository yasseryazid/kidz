FROM golang:1.25.5-alpine

WORKDIR /app

# Install git (dibutuhkan go get) & air (hot reload)
RUN apk add --no-cache git \
    && go install github.com/air-verse/air@latest

# Copy dependency files first (cache-friendly)
COPY go.mod go.sum ./
RUN go mod download

# Copy all source code
COPY . .

EXPOSE 8080

CMD ["air"]
