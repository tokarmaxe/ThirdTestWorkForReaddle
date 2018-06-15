package main

import (
	"fmt"
)

func main(){
	var errors int = 0
	var inputValue int = 0;
	var i int = 0;
	for errors<3{
		fmt.Println("Write, please, Fibonacci number ", i)
		fmt.Scanf("%d",&inputValue)
		if(inputValue == fib(i)){
		}else{
			fmt.Println("Correct answer: ", fib(i))
			errors++
		}
		i++
	}
}
func fib(n int) (int){
	if n==0{
		return 0
	}else if n==1{
		return 1
	}else if n==2{
		return 1
	}
	return fib(n-2)+fib(n-1)
}