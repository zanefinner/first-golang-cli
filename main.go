package main

import (
	"fmt"

	"github.com/zanefinner/first-golang-cli/greetings"
)

func main() {
	fmt.Println("Test")
	greetings.SayHello()
	greetings.SayHey()
}
