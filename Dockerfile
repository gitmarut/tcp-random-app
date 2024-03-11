# base image
FROM golang:1.20.3-alpine3.17 as base
WORKDIR /builder

ENV GO111MODULE=on CGO_ENABLED=0

COPY go.mod /builder/
RUN go mod download

COPY . .
RUN go build -o /builder/main /builder/main.go

# runner image
FROM gcr.io/distroless/static:latest
WORKDIR /app
COPY --from=base /builder/main main

EXPOSE 17777
CMD ["/app/main"]
