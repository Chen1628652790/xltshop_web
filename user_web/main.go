package main

import (
	"github.com/xlt/shop_web/user_web/initialize"
	"log"
)

func main() {
	engine := initialize.InitRouter()
	if err := engine.Run(":8021"); err != nil {
		log.Fatal("engine.Run failed, err:", err.Error())
	}
}
