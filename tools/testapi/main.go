package main

import (
	"github.com/sh-miyoshi/doraku/testapi/cmd"
)

func main() {
	cmd.InitConfig()
	cmd.Execute()
}
