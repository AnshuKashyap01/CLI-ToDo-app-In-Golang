package main

import "fmt"

func main() {

	todos := Todos{}
	todos.add("Buy Milk")
	todos.add("Buy bread")
	todos.add("Buy toast")
	fmt.Printf("%v\n", todos)
	todos.delete(0)
	fmt.Printf("%+v\n",todos) //+ will also add title,completed etc. in output
	
}