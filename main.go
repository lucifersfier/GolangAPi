package main

import (
	"github.com/gin-gonic/gin"
	"errors"
	"net/http"
)

type todo struct{
	//how data structure will look using struct
	//define properties and IDs for unique identification each todo item
	ID 			string `JSON:"id"`
	Item 		string `JSON:"item"`
	Completed 	bool `JSON:"completed"`


}

var todos = []todo{
	{ID: "1", Item :"Clean room",Completed: true},
	{ID: "2", Item :"Dancing",Completed: true},
	{ID: "3", Item :"Reading",Completed: true},
	{ID: "4", Item :"learning",Completed: true},

}

func getTodos(context *gin.Context){
	context.IndentedJSON(http.StatusOK, todos)
}

func addTodo(context *gin.Context){
	var newTodo todo

	if err := context.BindJSON(&newTodo); err != nil {
		return 
	}
	todos = append(todos,newTodo)
	context.IndentedJSON(http.StatusCreated, newTodo)
}


 

 func getTodo(context *gin.Context){
	id := context.Param("id")
	todo, err := getTodoById(id)

	if err != nil{
		context.IndentedJSON(http.StatusNotFound, gin.H{"message":"todo not found"})
		return 
	}
	context.IndentedJSON(http.StatusOK, todo)
 }

 func getTodoById(id string)(*todo,error){
	for i, t := range todos{
		if t.ID == id{
			return &todos[i], nil
		}
	}

   return nil, errors.New("todo not foound")

}

func toggletodoStatus(context *gin.Context){
	id := context.Param("id")
	todo, err := getTodoById(id)

	if err != nil{
		context.IndentedJSON(http.StatusNotFound, gin.H{"message":"todo not found"})
		return 
	}

	todo.Completed = !todo.Completed

	context.IndentedJSON(http.StatusOK, todo)
}


func main(){
	router := gin.Default() //roouter is server
	router.GET("/todos",getTodos)
	router.GET("/todos/:id",getTodo)
	router.PATCH("/todos/:id",toggletodoStatus)
	router.POST("/todos",addTodo)
	router.Run("localhost:9090")
}