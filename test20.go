package main
import "fmt"

func main(){
   m:=make(map[string]string)
   fmt.Println(m)
   m=map[string]string{"Bapan":"Biswas","vikash":"Dubey","test1":"test2"}
   fmt.Println(m)
   a:=make([] string,0,10)
   for key,_:=range m{
      fmt.Println(key)
      a=append(a,key)
   }
   fmt.Println(a)
   fmt.Println(len(a))
}