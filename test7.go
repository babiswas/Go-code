package main
import "fmt"

type teacher struct{
   firstname string
   lastname string
   subject [] string
}

type student struct{
  firstname string
  lastname string
  subject [] string
}

type person interface{
   display()
}

func (t teacher)display(){
   str1:=" "
   for i:=0;i<len(t.subject);i++{
       str1+=t.subject[i]
       str1+=" "
   }
   fmt.Println(t.firstname+" "+t.lastname+" "+str1)
}


func (s student)display(){
   str1:=" "
   for i:=0;i<len(s.subject);i++{
       str1+=s.subject[i]
       str1+=" "
   }
   fmt.Println(s.firstname+" "+s.lastname+" "+str1)
}


func test(p person){
   switch p.(type){
     case teacher:
         fmt.Println(p.(teacher).firstname)
         fmt.Println(p.(teacher).lastname)
         p.display()
     case student:
         fmt.Println(p.(student).firstname)
         fmt.Println(p.(student).lastname)
         p.display()
     default:
         fmt.Println("Good Bye")
   }
}

func main(){
   t:=teacher{firstname:"Teacher",lastname:"Classteacher",subject:[]string{"Science","Commerce"}}
   s:=student{firstname:"Student",lastname:"ClassStudent",subject:[]string{"Physics","Math"}}
   test(t)
   test(s)
   var k person
   k=t
   fmt.Println("Displaying power of interfaces")
   fmt.Println(k)
   k=s
   fmt.Println(k)
}