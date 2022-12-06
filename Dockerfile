FROM golang:1.19
WORKDIR /usr/src/app
COPY ./ ./
RUN go mod download && go mod verify
COPY . .
RUN go build -v -o ./op/
EXPOSE 8080
CMD ["/usr/src/app/op/CarParking"]