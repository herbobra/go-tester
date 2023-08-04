/*
package main

import "fmt"

// import "fmt"

	func main() {
		var a = "initial"
	    fmt.Println(a)

	    var b, c int = 1, 2
	    fmt.Println(b, c)

	    var d = true
	    fmt.Println(d)

	    var e int
	    fmt.Println(e)

	    f := "apple"
	    fmt.Println(f)

}
*/
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)


func IndexHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params){
	w.WriteHeader(200)
	w.Write([]byte("Работает"))
}


func main() {
	fmt.Println("hi!")
	router:=httprouter.New()
	router.GET("/", IndexHandler)
	log.Fatal(http.ListenAndServe(":8080", router))
    
}