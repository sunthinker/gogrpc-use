package main

import "grpc-demo/sc"

func main() {
	ser := &sc.ServiceEntity{}
	cli := &sc.ClientEntity{}

	go ser.InitService(":8080")

	cli.InitClient(":8080")
	cli.ClientTest("123456789")
	cli.ClientTest("987654321")
	cli.ClientTest("")
}
