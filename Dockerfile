FROM golang:rc-buster

WORKDIR /go/src/app

COPY go.mod ./
COPY go.sum ./

RUN apt update && apt install ffmpeg -y
RUN go mod download

COPY ./ ./

RUN go build . && rm -rf go.* api lib models pkg && go clean

EXPOSE 8080
CMD ["./qualifighting.backend.de"]
