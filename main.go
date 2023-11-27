package main

import (
	"github.com/fish895623/bilf/route"
)

// NOTE Query about gorm https://gorm.io/docs/query.html
func main() {
	route.Setup().Run(":8081")
}
