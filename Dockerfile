FROM golang:1.10

WORKDIR /go/src/github.com/emman27/profilepics

RUN go get github.com/golang/dep/cmd/dep

COPY Gopkg.lock Gopkg.toml ./
RUN dep ensure -v -vendor-only

COPY imaging imaging
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go install -ldflags="-w -s" -v github.com/emman27/profilepics

FROM alpine:latest
WORKDIR /app
RUN apk --no-cache add ca-certificates
COPY --from=0 /go/bin/profilepics ./profilepics
EXPOSE 8080
CMD ["./profilepics"]
