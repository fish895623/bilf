package main

import (
	"github.com/fish895623/bilf/route"
)

func main() {
	route.SetupRouter().Run(":8081")
}
