package main
import "net/http"

type User struct{
   Name string
   Email string
}

func (u *User)ServeHTTP(w http.ResponseWriter,r *http.Request){
   w.Write([]byte(u.Name+u.Email))
}

func main(){
   http.Handle("/user",&User{Name:"Bapan",Email:"bapan@gmail.com"})
   http.ListenAndServe(":9000",nil)
}