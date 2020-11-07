package main
import "fmt"

func main(){
  var employee=map[string]string{}
  fmt.Println(employee)
  fmt.Printf("%T\n",employee)
  employee["firstname"]="Bapan"
  employee["lastname"]="Biswas"
  employee["roll"]="XXXyy"
  fmt.Println(employee)
}