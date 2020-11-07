package main
import "fmt"

func main(){
 var a interface{}
 i:=2
 s:="Hello World"
 type person struct{
   name string
   id int64
 }

 p:=person{name:"Bapan Biswas",id:12}
 a=i
 fmt.Println("The value of a is",a)
 a=s
 fmt.Println("The value of a is",a)
 a=p
 fmt.Println("The value of a is",a)
 k:=a.(person)
 fmt.Println(k)
}