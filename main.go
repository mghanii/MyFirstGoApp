
package main
import (
"log"
"net/http"
"github.com/mohamedabdelghani/myfirstgoapp/routes"
"github.com/gorilla/mux"
)

func main() {
  r := mux.NewRouter()
	routes.RegisterPersonRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}


