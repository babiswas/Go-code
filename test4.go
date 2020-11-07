package main
import (
    "fmt"
    "bufio"
    "os"
    "strconv"
)

func main(){
   var a int64
   var b int64
   scanner:=bufio.NewScanner(os.Stdin)
   fmt.Println("Enter the value of a")
   scanner.Scan()
   a,_=strconv.ParseInt(scanner.Text(),10,64)
   fmt.Println("Enter the value of b")
   scanner.Scan()
   b,_=strconv.ParseInt(scanner.Text(),10,64)
   fmt.Println("the sum of a and b is:",a+b)  
}