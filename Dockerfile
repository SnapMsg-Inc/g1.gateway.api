FROM golang:latest

WORKDIR /usr/src

COPY go.mod go.mod 

# installing gin and swagger
RUN go get -u github.com/gin-gonic/gin
RUN go get -u github.com/swaggo/gin-swagger
RUN go get -u github.com/swaggo/files
RUN go install github.com/swaggo/swag/cmd/swag@latest
#ENV PATH "$PATH:$(go env GOPATH)/bin"

COPY . .

RUN swag init

EXPOSE 3000

CMD ["go", "run", "main.go"] 
