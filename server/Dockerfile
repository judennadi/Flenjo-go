FROM golang:1.19-alpine
WORKDIR /go/src/github.com/judennadi/flenjo-go/server
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
EXPOSE 5000
RUN go build -o server .
CMD [ "./server" ]