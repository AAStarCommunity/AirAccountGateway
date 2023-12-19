package main

import (
	"AirAccountGateway/conf"
	"AirAccountGateway/internal/controllers"
	"fmt"
)

func main() {
	// server模式运行
	c := conf.Get()
	fmt.Println("node:" + c.Node.Host)
	fmt.Println(fmt.Sprintf("port: %d", c.Service.Port))

	port := c.Service.Port
	_ = controllers.SetRouters().Run(fmt.Sprintf(":%d", port))
}
