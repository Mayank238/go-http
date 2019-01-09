package main

import (
    "fmt"
    "net/http"
    "strings"
    "log"
)


func main() {
    http.HandleFunc("/", sayhelloName) // set router
   
    http.HandleFunc("/a", sayHii)
        err := http.ListenAndServe(":9090", nil) // set listen port
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}

func sayHii(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, `<h1> Aniket </h1>`)
}


func sayhelloName(w http.ResponseWriter, r *http.Request) {
    // r.ParseForm()  // parse arguments, you have to call this by yourself
    // fmt.Println(r.Form)  // print form information in server side
    // fmt.Println("path", r.URL.Path)
    // fmt.Println("scheme", r.URL.Scheme)
    // fmt.Println(r.Form["url_long"])
    // for k, v := range r.Form {
    //     fmt.Println("key:", k)
    //     fmt.Println("val:", strings.Join(v, ""))
    // }
    // w.Write([]byte("aaaaa"))
    fmt.Fprintf(w, "hello" +   
          strings.Trim(r.URL.Path, "/") +
        "third line") // send data to client side
}