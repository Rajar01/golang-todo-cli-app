package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

type Todo struct {
	title  string
	isDone bool
}

var todos = []Todo{
	{"Todo 1", false},
	{"Todo 2", false},
	{"Todo 3", false},
}

func clearTerminal() {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	default:
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		return
	}
}

func addTodo() {
	clearTerminal()

	println("Input todo data")
	print("Todo title : ")

	reader := bufio.NewReader(os.Stdin)
	newTodoTitle, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	newTodoTitle = strings.TrimSpace(newTodoTitle)

	newTodo := Todo{title: newTodoTitle, isDone: false}
	todos = append(todos, newTodo)
	println("New todo added")

	time.Sleep(3 * time.Second)
}

func listAllTodo() {
	clearTerminal()

	for k, v := range todos {
		fmt.Printf("%d. Title : %s, isDone : %t\n", k+1, v.title, v.isDone)
	}

	time.Sleep(5 * time.Second)
}

func removeTodo() {
	clearTerminal()

	for k, v := range todos {
		fmt.Printf("%d. Title : %s, isDone : %t\n", k+1, v.title, v.isDone)
	}

	var todoIndex int
	print("Enter todo number to be remove : ")
	if _, err := fmt.Scanln(&todoIndex); err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	todos = append(todos[:todoIndex-1], todos[todoIndex:]...)
	println("Success deleted todo")

	time.Sleep(3 * time.Second)
}

func updateTodo() {
	reader := bufio.NewReader(os.Stdin)

	clearTerminal()

	for k, v := range todos {
		fmt.Printf("%d. Title : %s, isDone : %t\n", k+1, v.title, v.isDone)
	}

	var todoIndex int
	print("Enter todo number to be update : ")
	if _, err := fmt.Scanln(&todoIndex); err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	clearTerminal()

	print("Todo title : ")

	updatedTodoTitle, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	updatedTodoTitle = strings.TrimSpace(updatedTodoTitle)

	print("Todo status (true/false): ")
	var updatedTodoStatus bool
	isDone, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	isDone = strings.TrimSpace(isDone)

	if isDone == "true" {
		updatedTodoStatus = true
	}

	if isDone == "false" {
		updatedTodoStatus = false
	}

	updatedTodo := Todo{title: updatedTodoTitle, isDone: updatedTodoStatus}
	todos[todoIndex-1] = updatedTodo

	println("Success updated todo")

	time.Sleep(3 * time.Second)
}

func main() {
	for {
		clearTerminal()

		println("Welcome to Todo CLI Application")
		println("1. Add todo")
		println("2. List all todo")
		println("3. Remove todo from list")
		println("4. Update todo")
		println("5. Exit")

		var userChoice string
		print("Enter menu number : ")
		if _, err := fmt.Scanln(&userChoice); err != nil {
			return
		}

		switch userChoice {
		case "1":
			addTodo()
		case "2":
			listAllTodo()
		case "3":
			removeTodo()
		case "4":
			updateTodo()
		case "5":
			println("Thanks you for using Todo CLI Application")
			return
		default:
			println("Don't have such option, please try again")
		}
	}
}
