package main
import "net/http"
import "encoding/json"
import "fmt"
import  "github.com/jinzhu/gorm"
import _ "github.com/jinzhu/gorm/dialects/postgres"
import "log"


type User struct{
   Name string 
   Email string 
}

func getuser(w http.ResponseWriter,r *http.Request){
   db,err:=dbinit()
   if err!=nil{
      log.Println("Error Occured")
   }
   user:=[]User{}
   db.Find(&user)
   data,err:=json.Marshal(user)
   w.Write([]byte(data))
}

func  dbinit()(db *gorm.DB,err error){
   db,err=gorm.Open("postgres","user=test password=Learner#12 dbname=productuser sslmode=disable")
	if err!=nil{
		fmt.Println("Error occured")
	}
   return db,err
}

func main(){
   http.HandleFunc("/user",getuser)
   http.ListenAndServe(":5001",nil)
}