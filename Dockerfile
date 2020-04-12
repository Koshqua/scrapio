FROM golang:latest
WORKDIR /go/src/scrapio
COPY . . 
RUN go mod download
RUN go install
CMD ["scrapio"]
