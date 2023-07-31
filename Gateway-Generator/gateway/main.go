// Code generated by hertz generator.

		package main
		
		import (
			"os"
		
			"github.com/cloudwego/hertz/pkg/app/server"
		)
		
		func main() {
			url := "0.0.0.0:" + os.Args[1]
			h := server.Default(server.WithHostPorts(url))
		
			register(h)
			h.Spin()
		}
		