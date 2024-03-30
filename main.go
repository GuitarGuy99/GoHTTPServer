package main

import( 
"fmt"
"os"
"log"
"net/http"
"github.com/joho/godotenv"
"github.com/go-chi/chi"
)

func main(){
 fmt.Println("Hello World")
 godotenv.Load()
 PortString := os.Getenv("PORT")


 if PortString == ""{
    log.Fatal("PORT is not set in the .env file")
 }
 router := chi.NewRouter()
 srv := &http.Server{
     Handler: router,
     Addr:  ":" + PortString,
     }

 log.Printf("Server staring on port %v",PortString)
 err := srv.ListenAndServe()
 if err != nil {
   log.Fatal(err)
 }
 fmt.Print("Port:",PortString)

}
