package main

import (
	"github.com/EricOgie/egobank/konstants"
	myapp "github.com/EricOgie/egobank/myApp"
	"github.com/EricOgie/egobank/mylogger"
)

func main() {

	mylogger.Info(konstants.AppStart)
	myapp.StartApp()
}
