ARG GOLANG_TARGET=golang:1.12-alpine3.9
ARG TARGET=scratch
FROM ${GOLANG_TARGET} as builder

ENV APP_GOPATH $GOPATH/src/github.com/pirate/app/
ENV GO111MODULE=on

WORKDIR $APP_GOPATH

RUN apk add git
COPY go.* $APP_GOPATH
RUN go mod download
#compile app
COPY *.go $APP_GOPATH
RUN CGO_ENABLED=0 GOARM=7 GOOS=linux go build -a --installsuffix cgo --ldflags="-s" -o controller

#resulting app
FROM ${TARGET} as final
COPY --from=builder /go/src/github.com/pirate/app/controller /app/controller 
WORKDIR /app
ENTRYPOINT ["/app/controller"]
