build::
		go get github.com/gin-gonic/gin
		go get github.com/docker/engine-api/
		go get github.com/docker/go-connections/
		go get github.com/Sirupsen/logrus
		GOOS=linux GOARCH=amd64 go build
		docker-compose build
run::
		docker-compose up
