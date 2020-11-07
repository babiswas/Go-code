package main

import "net/http"
import "os"
import "log"
import "io"
import "strings"

func main(){
  http.HandleFunc("/",func(w http.ResponseWriter,r *http.Request){
    f,err:=os.Open("./"+r.URL.Path)
    if err!=nil{
       w.WriteHeader(http.StatusInternalServerError)
       log.Println(err)
    }
    defer f.Close()
    var contentType string
    switch{
      case strings.HasSuffix(r.URL.Path,"go"):
        contentType="text/go"
      case strings.HasSuffix(r.URL.Path,"py"):
        contentType="text/py"
    }
    w.Header().Add("Content-Type",contentType)
    io.Copy(w,f)
  })
  http.ListenAndServe(":8000",nil)
}