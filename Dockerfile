FROM golang:1.24

WORKDIR /todo

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN make build 

EXPOSE 80

CMD [ "./bin/main" ]