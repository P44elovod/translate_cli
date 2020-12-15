package main

import (
	"translate_cli/internal/app"
	"translate_cli/util"
)

func main() {

	err := app.Start()

	util.FailOnError(err, "Failed on API starting")

}
