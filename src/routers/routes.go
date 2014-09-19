// routes
package routers

import (
	//	"fmt"
	"controllers"
	"net/http"
)

/*
func main() {
	fmt.Println("Hello World!")
}
*/

func Init() {
	http.HandleFunc("/", controllers.IndexController)
	http.HandleFunc("/new", controllers.CreateController)
	http.HandleFunc("/paste/", controllers.ShowController)
}
