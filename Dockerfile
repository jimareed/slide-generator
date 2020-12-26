FROM golang:1.13-alpine AS builder
RUN apk --update add git

WORKDIR /go/src/github.com/jimareed/app/
COPY . /go/src/github.com/jimareed/app/
RUN rm -f goapp
RUN go get github.com/gorilla/mux
RUN go get github.com/jimareed/slides
RUN go get github.com/jimareed/drawing
RUN go get github.com/rs/cors
RUN go build -o ./goapp

FROM alpine:3.7

EXPOSE 8080
COPY --from=builder /go/src/github.com/jimareed/app /usr/local/bin/
RUN chown -R nobody:nogroup /usr/local/bin/goapp && chmod +x /usr/local/bin/goapp
USER nobody
CMD goapp -input /usr/local/bin/slides