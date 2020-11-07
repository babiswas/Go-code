package main
import "fmt"

func main(){
  var employee=make(map[string]int)
  employee["mark"]=10
  employee["sandy"]=30
  fmt.Println(employee)
  employeelist:=make(map[string]int)
  employeelist["mark"]=10
  employeelist["sandy"]=20
  fmt.Println(employeelist)
}