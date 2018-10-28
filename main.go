package main

import "fmt"
import "net/http"

func main(){
  server := http.Server { Addr: ":8080" }

  http.HandleFunc("/hello", helloAPI)
  
  fmt.Println("Starting the server in 8080 port" )

  if err:= server.ListenAndServe() ; err!= nil {
    fmt.Printf("Hey we got an error while serving requests : %s \n", err)
  }
}


func helloAPI(res http.ResponseWriter, req *http.Request) {
  res.WriteHeader(http.StatusOK)
  fmt.Fprintln(res, "Everything is working fine and you got a 200 OK")
}
