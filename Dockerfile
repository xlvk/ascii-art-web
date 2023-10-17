FROM golang:1.19

LABEL Authors = "Fatima Abbas, Ebrahim"
LABEL Name = "ascii-art-web-dockerize"

COPY . /ascii-art-web
WORKDIR /ascii-art-web
RUN go build -o server.go
EXPOSE 2003
CMD ./server.go
