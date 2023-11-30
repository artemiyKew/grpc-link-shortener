package main

import (
	"github.com/artemiyKew/grpc-link-shortener/internal/app"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Fatal(app.Run())
}
