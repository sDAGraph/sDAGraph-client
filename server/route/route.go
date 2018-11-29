package route

import (
    "fmt"
    "encoding/json"
    "io/ioutil"
    "net/http"
    "sDAGraph-client/db"
    "gopkg.in/mgo.v2/bson"
    "sDAGraph-client/params"
)

type User struct {
    Id string
}

func Router(selversion string){
    fmt.Println("selversion",selversion)
    c := params.Chain()
    mongoIp := c.Version.Sue[selversion].MongoIp
    mongoName := c.Version.Sue[selversion].MongoName
    mongoSession := c.Version.Sue[selversion].MongoSession

    GetNews := func (res http.ResponseWriter, req *http.Request){
        res.Header().Add("Access-Control-Allow-Origin","*")
        res.Header().Add("Content-Type", "application/json; charset=utf-8")

        b, _ := ioutil.ReadAll(req.Body)
        defer req.Body.Close()
	fmt.Println("b:",string(b))
	var newsdata User
        json.Unmarshal(b, &newsdata)

    
	db, session := sDAGraph_mongo.GetDB(mongoIp,mongoName)
	in2 := bson.M{"id": "0xaaazz833"}
	fmt.Println("in2:",in2)
	in:= newsdata
	fmt.Println("in:",in)
	result2 := sDAGraph_mongo.FindOne(db,mongoSession,in)
        fmt.Println("final:",result2)
	respBody, _ := json.Marshal(result2)
	session.Close()
	res.Write([]byte(respBody))
    }

    GetIndex := func (res http.ResponseWriter, req *http.Request){
        res.Header().Add("Access-Control-Allow-Origin","*")
        
	fmt.Println("ip:",mongoIp)
	res.Write([]byte("456"))
    }

    InsertData := func (res http.ResponseWriter, req *http.Request){
        res.Header().Add("Access-Control-Allow-Origin","*")

        b, _ := ioutil.ReadAll(req.Body)
        defer req.Body.Close()
        var newsdata params.NewsData
        json.Unmarshal(b, &newsdata)
	fmt.Println(newsdata.Name)
 
	db, session := sDAGraph_mongo.GetDB(mongoIp,mongoName)
	//in := User{Name:"test2",Id:"0xasdf2342",Number:17}
	in := newsdata
	//插入
        result := sDAGraph_mongo.Insert(db,mongoSession,in)
        fmt.Println(result)
	session.Close()
	res.Write([]byte("result"))
    }

    http.HandleFunc("/getNews", GetNews)
    http.HandleFunc("/getIndex", GetIndex)

    http.HandleFunc("/insertData", InsertData)
}
