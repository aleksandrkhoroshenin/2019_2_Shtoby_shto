FROM alpine AS base
RUN apk add --no-cache curl wget

FROM golang:1.12 AS go-builder
WORKDIR /go/app
COPY . /go/app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/app/trello-api /go/app/src/trello-api.go

FROM base
COPY --from=go-builder /go/app/trello-api /trello-api
CMD ["/trello-api"]