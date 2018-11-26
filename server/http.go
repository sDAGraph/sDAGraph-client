package server

import (
	//"lurcury/core/block"
	//"lurcury/core"
	"encoding/json"  
	"fmt"
	"io/ioutil"
	"net/http"
	"sDAGraph-client/server/route"
	//"lurcury/types"
	"time"
	"path/filepath"

)

func readFile(filename string) (map[string]string, error) {
    bytes, err := ioutil.ReadFile(filename)
    if err != nil {
        fmt.Println("ReadFile: ", err.Error())
        return nil, err
    }
    var j = map[string]string{}
    if err := json.Unmarshal(bytes, &j); err != nil {
        fmt.Println("Unmarshal: ", err.Error())
        return nil, err
    }
    return j, nil
}

func httpSet (){
	absPath, _ := filepath.Abs("../server/config.json")
	//fmt.Println("absPath: ",absPath)
	config,err :=readFile(absPath)//"config.json")
	if(err!=nil){
		fmt.Println("error:", err)
	}
	route.Router()
	fmt.Println("connect port"+config["port"])
	err2:= http.ListenAndServe(config["port"], nil)
	if err2 != nil {
		fmt.Println("error:", err2)
	}

}


func Server() { 
	httpSet()
	//httpSet2(coreStruct)
	time.Sleep(100 * time.Second)
}
