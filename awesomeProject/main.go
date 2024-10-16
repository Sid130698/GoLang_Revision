package main

import "fmt"

func main() {
	var reviseGo = "Revise Go quickly"
	var learnByCreating = "Learn by creating a new application"
	var reviseJava = "Prepare for Java Interview"
	var taskItems = []string{reviseGo, learnByCreating , reviseJava}
	fmt.Println("Hello to our TodoList Application! ")
	// fmt.Println("Tasks:" ,taskItems)
	taskItems= addList(taskItems, "Clear the interview")
	printList(taskItems)
}

func printList(myList []string){
	for index, task := range myList{
        fmt.Printf("%d -> %s\n",index, task);
    }
}

func addList(myList []string, newItem string) []string{
	updatedList := append(myList, newItem);
	return updatedList
}