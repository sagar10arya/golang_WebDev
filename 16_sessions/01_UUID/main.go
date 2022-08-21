package main
import (
	"fmt"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

// "	github.com/satori/go.uuid"
func main(){
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("session-1")
	if err != nil {
		id := uuid.NewV4()
		cookie = &http.Cookie{
			Name:  "session-1",
			Value: id.String(),
			//Secure:true,
			Path: "/",
		}
		http.SetCookie(w, cookie)

	}
	fmt.Println(cookie)
}
