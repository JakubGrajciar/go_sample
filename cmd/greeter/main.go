package main

import (
	"fmt"
	"greeter/pkg/greeter"
)

func main() {
	// Create greeter
	g, err := greeter.NewGreeter("Hi")
	if err != nil {
		panic(err)
	}

	// Greet Me, expect 'Hi Me'
	ret, err := g.Greet("Me")
	if err != nil {
		panic(err)
	}
	fmt.Println(ret)

	// Greet "", expect error
	ret, err = g.Greet("")
	if err != nil {
		fmt.Println("ERROR:", err.Error())
	}
	fmt.Println(ret)

	// Print counters
	cnt := g.GetCount()
	fmt.Println("cnt:", cnt)
	errCnt := g.GetErrCount()
	fmt.Println("err:", errCnt)

	// Create new greater without factory
	gg := greeter.Greeter{
		Greeting: "I am",
	}

	// Greet Me, expect "I am Me"
	ret, err = gg.Greet("Me")
	if err != nil {
		panic(err)
	}
	fmt.Println(ret)

	// Print counters
	cnt = gg.GetCount()
	fmt.Println("cnt:", cnt)
	errCnt = gg.GetErrCount()
	fmt.Println("err:", errCnt)
}
