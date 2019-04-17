FROM golang:1.12.1-alpine3.9 as builder
WORKDIR /go/src/github.com/vishen/always-break
COPY main.go main.go
RUN CGO_ENABLED=0 go build 

FROM scratch
WORKDIR /app
COPY --from=builder /go/src/github.com/vishen/always-break/always-break always-break
ENTRYPOINT ["/app/always-break"]
