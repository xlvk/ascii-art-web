FROM golang:1.19

LABEL Authors = "Fatima Abbas, Ebrahim"
LABEL Name = "ascii-art-web-dockerize"

COPY . .
WORKDIR /ascii-art-web/projects/httpserver
RUN go build -o /go-ascii-art-web
EXPOSE 2003
CMD ["/go-ascii-art-web"]
