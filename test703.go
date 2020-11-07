package main
import "fmt"
import "strconv"
import "encoding/json"

type User struct{
  Name string
  Rollnumber int
  Email string
}


type Product struct{
  Name string `json:"name"`
}

products:=make([]Product,4,6)


func (u* User)userinit(name,email string,roll int)*User{
   u.Name=name
   u.Email=email
   u.Rollnumber=roll
   return u
}



func (u *User)display()*User{
   fmt.Println(u.Name+" "+u.Email+" "+strconv.Itoa(u.Rollnumber))
   return u
}


func NewProduct(w http.ResponseWriter,r *http.Request){
  switch r.Method{
  case http.MethodPost:
      var newproduct  Product
      body,err:=ioutil.ReadAll(r.Body)
      if err!=nil{
        w.WriteHeader(http.StatusBadRequest)
      }
      body,err:=json.Unmarshal(body,&newproduct)
      if err!=nil{
        w.WriteHeader(http.StatusBadRequest)
      }
      products=append(products,body)
      fmt.Println(products)
  }
}


func main(){
  u:=User{"Bapan",12,"bapan@gmail.com"}
  u.display().userinit("Tapan","tapan@gmail.com",13).display().userinit("Tapajit","tapajit@gmail.com",14).display()
  body,err:=json.Marshal(u)
  if err!=nil{
     fmt.Println("Error occured")
  }
  fmt.Println(string(body))
}