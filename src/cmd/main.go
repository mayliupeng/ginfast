package main

import (
	"ginfast/src/classes"
	"github.com/gin-gonic/gin"
)

func main() {
	r:=gin.New()
	classes.NewIndexClass(r).Build()
	r.Run(":8080")
}