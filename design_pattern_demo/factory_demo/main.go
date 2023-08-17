package main

import (
	"factory_demo/simple_factory"
	"factory_demo/single_factory"
	"fmt"
)

func main() {
	ak47, _ := simple_factory.GetGun("ak47")
	printDetails(ak47)
	// ak := new(simple_factory.AK47)

	for i := 0; i < 30; i++ {
		go single_factory.GetInstance()
	}
	fmt.Scanln()
}
func printDetails(g simple_factory.IGun) {
	fmt.Println(g.GetName())
	fmt.Println(g.GetPower())
}
