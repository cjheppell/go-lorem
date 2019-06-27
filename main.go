package main

import (
	"github.com/cjheppell/go-lorem/cmd"
)

func main() {
	goLoremCmd := cmd.NewGoLoremCommand()
	goLoremCmd.Execute()
}
