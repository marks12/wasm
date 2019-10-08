package main

import (
	"fmt"
)

func main() {

	type Elements struct {
		Some string
		Some2 int
	}

	type Elements2 struct {
		Some string
		Some2 int
	}

	type Elements3 struct {
		Some string
		Some2 int
	}

	n1 := Elements{Some: "Some to string1", Some2: 100500}
	n2 := Elements2{Some: "Some to string2", Some2: 100500}
	n3 := Elements3{Some: "Some to string3", Some2: 100500}

	println("Hello World!!!", n1.Some)

	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Println(n3)


	//keepDoingSomething()
}