package main

import (
	"fmt"
	"strings"
)

func multiply(a,b int) int {
	return a*b
}

func lenAndUpper_1_3(name string)(int, string){
	return len(name),strings.ToUpper(name)
}


func lenAndUpper_1_4(name string)(length int,uppercase string){
	// defer는 function이 값을 return하고 나면 실행됨
	defer fmt.Println("I'm done")
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func repeatMe(words ...string){
	fmt.Println(words)
}

func superAdd(numbers ...int)int{
for index, number := range numbers{
	fmt.Println(index, number)
}
for i:=0;i<len(numbers);i++{
	fmt.Println(i)
}
total:=0
for _,number := range numbers{
	total += number
}
return total
}

func main() {
	fmt.Println(multiply(2,2))
	totalLenght_1_3, _ :=lenAndUpper_1_3("nico")
	fmt.Println(totalLenght_1_3)
	repeatMe("jinsuh","nico","Lynn","test")
	fmt.Println("1-4")
	totalLenght_1_4, _ :=lenAndUpper_1_4("nico")
	fmt.Println(totalLenght_1_4)
	fmt.Println("1-5")
	result:=superAdd(7,8,9,10,11,12)
	fmt.Println(result)
}
