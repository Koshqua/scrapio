FROM golang:latest
WORKDIR /go/src/scrapi
COPY . . 
RUN go mod download
RUN go install
CMD ["scrapio"]
