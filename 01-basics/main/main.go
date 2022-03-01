package main

import (
	"comcast-india-go/01-basics/numbers"
	"fmt"
)

func main(){
	num := 121347
	result := numbers.IsPrime(num)
	if result{
		fmt.Printf("%d Number is prime ",num)
	}else{
		fmt.Printf("%d Number is nor prime ",num)
	}
}