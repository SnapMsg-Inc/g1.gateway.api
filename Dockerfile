FROM golang:latest

WORKDIR /usr/snapmsg-gateway

COPY go.mod go.mod 

# installing gin and swagger
RUN go get -u github.com/gin-gonic/gin
RUN go get -u firebase.google.com/go/v4@latest
RUN go get -u google.golang.org/api/option
RUN go get -u github.com/swaggo/gin-swagger
RUN go get -u github.com/swaggo/files
RUN go install github.com/swaggo/swag/cmd/swag@latest

COPY . .

RUN swag init

ENV USERS_URL=https://users-ms-marioax.cloud.okteto.net

EXPOSE 3000

CMD ["go", "run", "main.go"] 
