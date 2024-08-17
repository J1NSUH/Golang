package main

import (
	"Golang/something" // something 패키지를 가져옵니다
	"fmt"
)

func main() {
	fmt.Println("Hello world")
	something.SayHello() // something 패키지의 함수를 호출합니다
}
