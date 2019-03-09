ARG GOLANG_TARGET=golang:1.12
ARG TARGET=multiarch/alpine:armhf-v3.9
FROM ${GOLANG_TARGET} as builder

ENV APP_GOPATH $GOPATH/src/github.com/pirate/app/

WORKDIR $APP_GOPATH

RUN go get github.com/gorilla/mux
#compile app
COPY *.go $APP_GOPATH
RUN CGO_ENABLED=0 GOARM=7 GOOS=linux go build -a --installsuffix cgo --ldflags="-s" -o controller

#resulting app
FROM ${TARGET} as final
COPY --from=builder /go/src/github.com/pirate/app/controller /app/controller 
WORKDIR /app
ENTRYPOINT ["/app/controller"]
