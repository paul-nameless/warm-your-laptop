########################################
FROM golang:1.14.7-alpine3.12 AS builder

WORKDIR /go/src/home/
COPY . .

RUN go build -ldflags="-s -w" -o /go/bin/main

############
FROM scratch
COPY --from=builder /go/bin/main /main
CMD ["/main"]