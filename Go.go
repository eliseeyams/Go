package myStrings

//All Go programs start with a package declaration.
//The main package is a special package that tells the Go
//compiler that this is an executable program, rather than
//a library. The main function is the entry point of the
//program, where the execution starts.

// fmt means format. It is a package that implements
// formatted I/O with functions analogous
// to C's printf and scanf. The functions in
// this package take a format string and a
// variable number of arguments, and return the
// formatted string or write it to an output stream.

/*var tasksItems = []string{"Go to the gym", "Buy groceries", "Read a book"}

func main() {

	//fmt.Println("####### Welcome to our Todolist App! ######")
	http.HandleFunc("/hello-go", helloUser)
	http.HandleFunc("/show-tasks", showTasks)

	http.ListenAndServe(":8085", nil)

}

func showTasks(writer http.ResponseWriter, request *http.Request) {
	for index, task := range tasksItems {
		fmt.Fprintf(writer, "%d. %s\n", index+1, task)
	}
	//fmt.Fprintln(writer, tasksItems)
}

func helloUser(writer http.ResponseWriter, request *http.Request) {
	var greeting = "Hello User, Welcome to our Todolist App!"
	fmt.Fprintln(writer, greeting)
}

func printTasks(taskItems []string) {
	fmt.Println("List of my Todos")
	for index, task := range taskItems {
		//fmt.Println("TasksI:", index, task)
		fmt.Printf("%d. %s\n", index+1, task)
	}
}

func addTask(taskIetms []string, newTask string) []string {
	var updateTaskItems = append(taskIetms, newTask)
	return updateTaskItems
}*/

/*
import (

	"errors"
	"net/http"

	//	"github.com/gin-gonic"
	"github.com/gin-gonic/gin"

	_ "net/http"

)

	type todo struct {
		ID        string `json:"id"`
		Item      string `json:"title"`
		Completed bool   `json:"completed"`
	}

	var todos = []todo{
		{ID: "1", Item: "Go to the gym", Completed: false},
		{ID: "2", Item: "Buy groceries", Completed: false},
		{ID: "3", Item: "Read a book", Completed: false},
	}

	func getTodos(context *gin.Context) {
		context.IndentedJSON(http.StatusOK, todos)
	}

	func addTdodo(context *gin.Context) {
		var newTodo todo

		if err := context.BindJSON(&newTodo); err != nil {
			return
		}
		todos = append(todos, newTodo)
		context.IndentedJSON(http.StatusCreated, newTodo)
	}

	func getTodo(context *gin.Context) {
		id := context.Param("id")
		todo, err := getTodoById(id)

		if err != nil {
			context.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
			return
		}
		context.IndentedJSON(http.StatusOK, todo)
	}

	func getTodoById(id string) (*todo, error) {
		for i, t := range todos {
			if t.ID == id {
				return &todos[i], nil
			}
		}
		return nil, errors.New("todo not found")
	}

func toggleTodoStatus(context *gin.Context) {

		id := context.Param("id")
		todo, err := getTodoById(id)

		if err != nil {
			context.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
			return
		}
		todo.Completed = !todo.Completed
		context.IndentedJSON(http.StatusOK, todo)
	}

	func main() {
		router := gin.Default()
		router.GET("/todos", getTodos)
		router.GET("/todos/:id", getTodo)
		router.PATCH("/todos/:id", toggleTodoStatus)
		router.POST("/todos", addTdodo)
		router.Run("localhost:9090")
	}
*/

func Reverse(s string) string {
	result := ""
	for _, v := range s {
		result = string(v) + result
	}
	return result
}
