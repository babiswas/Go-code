package main
import "fmt"

func main(){
  var employee=make(map[string]int)
  employee["mark"]=12
  employee["star"]=13
  employee["cat"]=14
  fmt.Println(employee)
  fmt.Println(len(employee))
  var employee1=make(map[string]int)
  fmt.Println(len(employee1))
}