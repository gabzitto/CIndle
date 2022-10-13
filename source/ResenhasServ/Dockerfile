FROM golang:1.19-alpine

WORKDIR /app

COPY . ./

RUN go mod tidy
RUN go mod download

RUN go build -o res-serv

CMD [ "./res-serv" ]