package main

import "fmt"

func main() {
	
	var row = 5
	for i :=0 ;i <row;i++{
		for j:=0;j<=i;j++{
			fmt.Println("*")
		}
		fmt.Println("")
	}
}