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

# install datadog statsd
RUN go get -u github.com/DataDog/datadog-go/v5/statsd

RUN go mod tidy

COPY . .

RUN swag init

ENV USERS_URL=users-api:3001
ENV POSTS_URL=posts-api:3001
ENV MESSAGES_URL=messages-api:3001

ENV STATSD_HOST=datadog-agent
ENV STATSD_PORT=8125

ENV SRV_ADDR=localhost:3001
EXPOSE 3001

CMD ["go", "run", "main.go"] 
