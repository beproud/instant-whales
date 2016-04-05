build::
		go get github.com/gin-gonic/gin
		go get github.com/fsouza/go-dockerclient
		go get github.com/Sirupsen/logrus
		GOOS=linux GOARCH=amd64 go build
		docker-compose build
run::
		./docker-compose up
