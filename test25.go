package main
import "fmt"

func first(){
  fmt.Println("first function")
}

func Second(){
  fmt.Println("Second Function")
}

func Third(){
  fmt.Println("Third Function")
}

func main(){
  defer Second()
  first()
  for i:=0;i<3;i++{
     Third()
  }
}