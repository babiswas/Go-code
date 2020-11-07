package main
import (
    "fmt"
    "sort"
)

func main(){
   m:=map[string]int{"India":20,"China":30,"USA":40,"France":60}
   keys:=make([]string,0,len(m))
   for k:=range m{
      keys=append(keys,k)
   }
   sort.Strings(keys)
   for _,k:=range keys{
     fmt.Println(k,m[k])
   }
}