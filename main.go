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
	fmt.Println("-for1 start------------- ")
	for index, number := range numbers{
	fmt.Println(index, number)
	}
	fmt.Println("-for2 start------------- ")
	for i:=0;i<len(numbers);i++{
	fmt.Println(i)
	}
	total:=0
	fmt.Println("-for3 start------------- ")
	for _,number := range numbers{
	total += number
	}
	return total
}

func canIDrink(age int)bool{
	return age >= 18
}


func canIDrinkInKorea(age int)bool{
	if koreanAge := age+2;koreanAge<20{
		return false
	}
	return true
}

func canIDrinkinswitch(age int)bool{
	switch {
	case age < 10:
		return false
	case age == 18:
		return true
	case age > 50:
		return false
	}
	return false
}

func canIDrinkinswitchInKorea(age int)bool{
	switch koreanAge := age+2; koreanAge{
	case 10:
		return false
	case 18:
		return true
	}
	return false
}

func main() {
	fmt.Println(multiply(2,2))
	totalLenght_1_3, _ :=lenAndUpper_1_3("nico")
	fmt.Println(totalLenght_1_3)
	repeatMe("jinsuh","nico","Lynn","test")
	fmt.Println("-------------- ")
	fmt.Println("1-4")
	totalLenght_1_4, _ :=lenAndUpper_1_4("nico")
	fmt.Println(totalLenght_1_4)
	fmt.Println("-------------- ")
	fmt.Println("1-5")
	result:=superAdd(7,8,9,10,11,12)
	fmt.Println(result)
	fmt.Println("---------------")
	fmt.Println("1-6")
	fmt.Println(canIDrink(16))
	fmt.Println(canIDrinkInKorea(16))
	fmt.Println("---------------")
	fmt.Println("1-7")
	fmt.Print(canIDrinkinswitch(18))
	fmt.Print(canIDrinkinswitchInKorea(18))
	fmt.Println("---------------")
	fmt.Println("1-8")
	a:=2
	b:=&a
	fmt.Println(&a,b)
	fmt.Println(&a,&b)
	fmt.Println(&a,b)
	fmt.Println(a,b)
	fmt.Println(*b)
	*b=20
	fmt.Println(a)
	fmt.Println("---------------")
	fmt.Println("1-9")
	names:= []string{"nico","lynn","dal"}
	names = append(names,"flynn")
	fmt.Println(names)

}
