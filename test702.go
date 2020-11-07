package main
import "fmt"
import "io"
import "strings"
import "log"



func main(){
  r:=strings.NewReader("abcdefg")
  buf:=make([]byte,4)
  if _,err:=io.ReadFull(r,buf);err!=nil{
     log.Fatal(err)
  }
  fmt.Println(buf)
  if _,err:=io.ReadFull(r,buf);err!=nil{
     log.Fatal(err)
  }
  mbuf:=make([]byte,4)
  if _,err:=io.ReadFull(r,mbuf);err!=nil{
     log.Fatal(err)
  }
  fmt.Println(mbuf)
}