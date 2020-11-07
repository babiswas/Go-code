package main
import "fmt"

func main(){
  var employee=map[string]int{"test1":1,"test2":2,"test3":3}
  fmt.Println(employee)
  delete(employee,"test1")
  fmt.Println(employee)  
}