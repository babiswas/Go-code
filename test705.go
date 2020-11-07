package main
import "net/http"

func getuser(w http.ResponseWriter,r *http.Request){
  w.Write([]byte("Hello World"))
}

func main(){
  http.HandleFunc("/user",getuser)
  http.ListenAndServe(":3001",nil)
}