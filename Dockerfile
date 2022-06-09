FROM golang:1.18.3-alpine3.16 AS build
# Support CGO and SSL
RUN apk --no-cache add gcc g++ make
RUN apk add git
RUN apk add npm
WORKDIR /go/src/app
COPY . .
RUN go get github.com/mattn/go-sqlite3
RUN GOOS=linux go build -ldflags="-s -w" -o ./bin/server ./main.go
RUN cd react && npm install && npm run build




FROM alpine:3.10
RUN apk --no-cache add ca-certificates
WORKDIR /usr/bin
COPY --from=build /go/src/app/bin /go/bin
COPY --from=build /go/src/app/react/build /go-eub-yt
EXPOSE 8080
ENTRYPOINT /go/bin/server --port 8080