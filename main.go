package main

import "main/pkg"

func main() {
	service := &pkg.Service{
		Name: "Test",
	}

	service.Hello()
}
