#FROM ubuntu:latest
#LABEL authors="Мария"
#
#ENTRYPOINT ["top", "-b"]

FROM golang:1.21

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -o house ./main.go

CMD ["./house"]