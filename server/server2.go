/*@Author : Manasvini Banavara Suryanarayana
*SJSU ID : 010102040
*CMPE 273 Lab#3
*/
package main

import (
    "fmt"
    "./httprouter"
    "net/http"
    "strconv"
    "encoding/json"
)

type Response1 struct {
   Key int `json:"key"`
   Value string `json:"value"`
}
type Response2 struct {
   Arr []Response1 `json:"Response"`
}

var KeyValueMap = map[int]string{}

func putmethod(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
    
    id := p.ByName("key_id")
    val := p.ByName("value")
    idval,err := strconv.Atoi(id)
    if err!=nil{
      fmt.Println("error occured in conversion  ")
    }
    KeyValueMap[idval]=val

    rw.WriteHeader(200)



}
func getAllvalue(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
  Responseret := make([]Response1, len(KeyValueMap))
    i := 0
    for key,value := range KeyValueMap {
      var temp Response1 
      temp.Key = key
      temp.Value = value
      Responseret[i]=temp
      i++
    }

    //constructing struct for sending back response body
    test := Response2{
        Arr: Responseret,
    }

    //converting response body struct to json format
   respjson, err2 := json.Marshal(test)
   if err2 != nil {
        fmt.Println("error occured 2")
    }
     
    rw.Header().Set("Content-Type","application/json")
    rw.WriteHeader(200)
    //sending back response
    fmt.Fprintf(rw, "%s", respjson)
}
func getValue(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
   

     id :=  p.ByName("key_id")
     
     idval,err := strconv.Atoi(id)
    if err!=nil{
      fmt.Println("error occured in conversion  ")
    }

    value := KeyValueMap[idval]

    //constructing struct for sending back response body
    test := Response1{
        Key: idval,
        Value: value,
    }

    //converting response body struct to json format
   respjson, err2 := json.Marshal(test)
   if err2 != nil {
        fmt.Println("error occured 2")
    }
     
    rw.Header().Set("Content-Type","application/json")
    rw.WriteHeader(200)
    //sending back response
    fmt.Fprintf(rw, "%s", respjson)
     
}
  
func main() {
    mux := httprouter.New()
    mux.GET("/keys/:key_id", getValue)
    mux.PUT("/keys/:key_id/:value", putmethod)
    mux.GET("/keys", getAllvalue)
    server := http.Server{
            Addr:        "0.0.0.0:3000",
            Handler: mux,
    }

    server.ListenAndServe()
}