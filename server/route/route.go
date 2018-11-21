package route

import (
    //"fmt"
    //"encoding/json"
    //"io/ioutil"
    "net/http"
    "sDAGraph-client/db"
)

func Router(){
    GetNews := func (res http.ResponseWriter, req *http.Request){
        res.Header().Add("Access-Control-Allow-Origin","*")
        res.Write([]byte("123"))
    }

    GetIndex := func (res http.ResponseWriter, req *http.Request){
        res.Header().Add("Access-Control-Allow-Origin","*")
        res.Write([]byte("456"))
    }

    InsertIndexdata := func (res http.ResponseWriter, req *http.Request){
        res.Header().Add("Access-Control-Allow-Origin","*")
        res.Write([]byte("456"))
    }

    http.HandleFunc("/getNews", GetNews)
    http.HandleFunc("/getIndex", GetIndex)

    http.HandleFunc("/insertIndexdata", InsertIndexdata)
}
