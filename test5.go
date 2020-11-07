package main
import (
   "fmt"
   "strconv"
)

func main(){
   var a int
   a=12
   str1:="Ramu has "+strconv.Itoa(a)+" "+"Goats"
   fmt.Println(str1)
   var b int
   b=6
   str1="Ramu has "+strconv.Itoa(a-b)+" "+"Goats"
   fmt.Println(str1)
   str1="Ramu has "+string(6)+" "+"Cows"
   fmt.Println(str1)
}
	