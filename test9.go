package main
import "fmt"

func main(){
  var a[] string=[]string{"hello1","hello2"}
  var b[] string=[]string{"hello3","hello4"}
  c:=append(a,b...)
  fmt.Println(c)
}