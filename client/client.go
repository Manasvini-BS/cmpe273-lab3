
/*@Author :Manasvini Banavara Suryanarayana
*@SJSU ID : 010102040
*CMPE 273 Lab #3
*/
package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "encoding/json"
    "strconv"
   
)


type Resp struct {
   Key int `json:"key"`
   Value string `json:"value"`
}

var KeyValueMap = map[int]string{}
var ServerSlotMap = map[string]int{}

func hashing(a int)int {
    return (a%3)*10
}
func getVal(a int){

    hashval := hashing(9)
    var port2 string
    for i3, k3 := range ServerSlotMap{
        if hashval == k3 {
                port2 = i3
                
                break

        }
    }
    str := "http://localhost:"+port2+"/keys/"+strconv.Itoa(9)
        
          client1 := &http.Client{}
          req, err111 := http.NewRequest("GET",str, nil)
          if err111 != nil {
            fmt.Println(err111)
          }
          
          
          resp, err2 := client1.Do(req)
          
          
        if err2 != nil {
            fmt.Println(err2)
        }
        defer resp.Body.Close()
       
        
        body1, err3 := ioutil.ReadAll(resp.Body)
        if err3 != nil {
            fmt.Println(err3)
        }
      
        
        var getvalue interface{}
        err4 := json.Unmarshal(body1, &getvalue)
        if err4 != nil {
            fmt.Println(err4)
        }
       

        res1 := getvalue.(map[string]interface{})
        
        t := Resp{
            Key: int(res1["key"].(float64)),
            Value: res1["value"].(string),
        }
        respjson, err5 := json.Marshal(t)
           if err5 != nil {
                fmt.Println(err5)
            }
            fmt.Println("\n---------------------------------------------")
            fmt.Println("Response from server for key given :")
        fmt.Println(string(respjson))
            fmt.Println("---------------------------------------------")

}

func main() {
    KeyValueMap[1]="a"
    KeyValueMap[2]="b"
    KeyValueMap[3]="c"
    KeyValueMap[4]="d"
    KeyValueMap[5]="e"
    KeyValueMap[6]="f"
    KeyValueMap[7]="g"
    KeyValueMap[8]="h"
    KeyValueMap[9]="i"
    KeyValueMap[10]="j"
             
    servers := []string{"3000","3001","3002"}
    

    
    //creates server- slot mapping
    for i,k := range servers{
        if i==i {}
        x,err := strconv.Atoi(k)
        if err != nil {
            fmt.Println(err)
          }
        
        ServerSlotMap[k] = hashing(x)
        
    }
    // redirecting insert in to the respective server slot
    for i1 ,k1 := range KeyValueMap{
       
       var port string
       
        result := hashing(i1)
        for i3, k3 := range ServerSlotMap{
            if result == k3 {
                port = i3
               
                break

            }
        }
       
       

        str := "http://localhost:"+port+"/keys/"+strconv.Itoa(i1)+"/"+k1
        fmt.Println("Request sent to the server: ",str)
       
          client1 := &http.Client{}
          req, err111 := http.NewRequest("PUT",str, nil)
          if err111 != nil {
            fmt.Println(err111)
          }
          
          //req.Header.Set("Content-Type", "application/json")
          resp, err2 := client1.Do(req)
          
        
        if err2 != nil {
            fmt.Println(err2)
        }
        defer resp.Body.Close()
        fmt.Println("Response Code : ",resp.Status)
        
       



    }
        
    //calling get method for retirveing specific key : 9

    getVal(9)

    
}