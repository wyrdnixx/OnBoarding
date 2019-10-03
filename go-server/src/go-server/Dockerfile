FROM golang:1.10

# workdir inside the container
WORKDIR /go/src/github.com/wyrdnixx/go-server/

COPY . .

RUN ["go", "get", "github.com/githubnemo/CompileDaemon"]
RUN ["go", "get", "github.com/tkanos/gonfig"]
RUN ["go", "get", "github.com/go-sql-driver/mysql"]

#RUN ["go", "get", "github.com/labstack/echo"]
#RUN ["go", "get", "github.com/labstack/echo/middleware"]
###


ENTRYPOINT CompileDaemon -log-prefix=false -build="go build ./src/go-server/" -command="./go-server"
