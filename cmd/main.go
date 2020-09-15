package main

import (
	"github.com/manjoc/dingtalk/cmd/dingtalk"
	"log"
)

func main() {
	log.SetFlags(0)
	dingtalk.Execute()
}
