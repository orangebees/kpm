package main

import (
	"github.com/orangebees/kpm/pkg/kpm"
	"os"
)

func main() {
	err := kpm.CLI(os.Args...)
	if err != nil {
		println(err.Error())
		return
	}
}
