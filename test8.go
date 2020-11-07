package main
import "fmt"

func main(){
  var str1 []string=[]string{"hello1","hello2","hello3"}
  str1=append(str1,"hello4","hello5","hello6")
  fmt.Println(str1)
}