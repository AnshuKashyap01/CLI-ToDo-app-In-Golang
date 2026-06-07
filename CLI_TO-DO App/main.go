package main



// git status
// git add .
// git commit -m "Implement CLI todo app"
// git push origin main

func main() {

	todos := Todos{}
	storage := NewStorage[Todos]("todos.json")
	storage.Load(&todos)
	cmdFlags:=NewCmdFlags()
	cmdFlags.Execute(&todos)
	//todos.add("Buy toast")
	todos.print()
	storage.save(todos)
	
	//fmt.Printf("%+v\n",todos) //+ will also add title,completed etc. in output

}