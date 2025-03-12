FROM golang:1.24

WORKDIR /todo

COPY . ./
RUN go mod download

RUN make build 

EXPOSE 80

CMD [ "./bin/main" ]