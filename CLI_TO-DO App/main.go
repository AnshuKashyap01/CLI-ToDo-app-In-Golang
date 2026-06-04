package main



// git status
// git add .
// git commit -m "Implement CLI todo app"
// git push origin main

func main() {

	todos := Todos{}
	todos.add("Buy Milk")
	todos.add("Buy bread")
	//todos.add("Buy toast")
	todos.print()
	
	//fmt.Printf("%+v\n",todos) //+ will also add title,completed etc. in output

}