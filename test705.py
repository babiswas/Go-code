package main
import "net/http"
import "encoding/json"

func getuser(w http.ResponseWriter,r *http.Request){
  w.Write([]bytes("Hello World"))
}

func main(){
  http.HandleFunc("/user",getuser)
  http.ListenAndServe(":3001",nil)
}