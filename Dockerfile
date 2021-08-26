FROM golang:1.6.0 as builder

WORKDIR /go/src/github.com/oyamoh-brian/tv-service-videos

COPY . .

RUN go get -u github.com/golang/dep/cmd/dep
RUN dep init && dep ensure
RUN CGO_ENABLED=0 GOOS=linux go build  -o tv-service-videos -a -installsuffix cgo main.go dummy-videos.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from=builder /go/src/github.com/oyamoh-brian/tv-service-videos/tv-service-videos .

CMD ["./consignment-service"]