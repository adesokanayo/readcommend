FROM golang:1.15.3-alpine AS build_base

ENV CGO_ENABLED=1
ENV GO111MODULE=on
RUN apk add --no-cache git  git gcc g++

WORKDIR /src
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# Build the Go app
RUN go build -o ./out/app ./cmd/api/main.go

# Start fresh from a smaller image
FROM alpine:3.12
RUN apk add ca-certificates

WORKDIR /app

COPY --from=build_base /src/out/app /app/readcommendapi
COPY --from=build_base /src/data /app/data

RUN chmod +x readcommendapi

# This container exposes port 5000 
EXPOSE 5000

# Run the binary program produced by `go install`
ENTRYPOINT ./readcommendapi
