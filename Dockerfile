FROM golang:1.17.2

RUN apt -y update && apt -y install build-essential
RUN apt -y install nano

COPY . /
WORKDIR /

RUN go mod download
RUN go build -o more-tech-back cmd/main.go

CMD ["./more-tech-back"]
EXPOSE 8080