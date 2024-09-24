package main

import (
	"github.com/sirupsen/logrus"
	"petProject/internal/app"
)

const configDirectoryPath = "config"

func main() {
	logrus.Info("starting http server")
	app.Run(configDirectoryPath)

}
