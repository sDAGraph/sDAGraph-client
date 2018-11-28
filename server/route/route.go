package route

import (
    "fmt"
    "encoding/json"
    "io/ioutil"
    "net/http"
    "sDAGraph-client/db"
    "gopkg.in/mgo.v2/bson"
    "path/filepath"
    "sDAGraph-client/params"
)

type User struct {
    Name string
    Id string
    Number uint
}

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

func Router(selversion string){
    fmt.Println("selversion",selversion)
    GetNews := func (res http.ResponseWriter, req *http.Request){
        res.Header().Add("Access-Control-Allow-Origin","*")
        res.Header().Add("Content-Type", "application/json; charset=utf-8")

	absPath, _ := filepath.Abs("../server/route/dbconfig.json")
        fmt.Println("absPath: ",absPath)
        dbconfig,err :=readFile(absPath)
        if(err!=nil){
            fmt.Println("error:", err)
        }
	c := params.Chain()
	fmt.Println(c)
        fmt.Println("version",c.Version.Sue[selversion].MongoIp)

        db, session := sDAGraph_mongo.GetDB(dbconfig["ip"], dbconfig["dbname"])
	//db, session := sDAGraph_mongo.GetDB("mongodb://192.168.51.202:27017", "sDAG")
	//in := bson.M{"Name":"test1","Id":"0xasdf2341","Number":13}
	in := bson.M{"id": "0xasdf2341"}
	fmt.Println("in:",in)
	result2 := sDAGraph_mongo.FindOne(db,dbconfig["sessionname"],in)
        fmt.Println("final:",result2)
	respBody, _ := json.Marshal(result2)
	session.Close()
	res.Write([]byte(respBody))
    }

    GetIndex := func (res http.ResponseWriter, req *http.Request){
        res.Header().Add("Access-Control-Allow-Origin","*")
        absPath, _ := filepath.Abs("../server/route/dbconfig.json")
        fmt.Println("absPath: ",absPath)
        config,err :=readFile(absPath)
	if(err!=nil){
            fmt.Println("error:", err)
        }
	fmt.Println("ip:",config["ip"])
	res.Write([]byte("456"))
    }

    InsertIndexdata := func (res http.ResponseWriter, req *http.Request){
        res.Header().Add("Access-Control-Allow-Origin","*")

        b, _ := ioutil.ReadAll(req.Body)
        defer req.Body.Close()
        var newsdata User
        json.Unmarshal(b, &newsdata)
	fmt.Println(newsdata.Name)

        absPath, _ := filepath.Abs("../server/route/dbconfig.json")
        fmt.Println("absPath: ",absPath)
        dbconfig,err :=readFile(absPath)
        if(err!=nil){
            fmt.Println("error:", err)
        }
	db, session := sDAGraph_mongo.GetDB(dbconfig["ip"], dbconfig["dbname"])

	in := User{Name:"test1",Id:"0xasdf2341",Number:13}
	//插入
        result := sDAGraph_mongo.Insert(db,dbconfig["sessionname"],in)
        fmt.Println(result)
	session.Close()
	res.Write([]byte("result"))
    }

    http.HandleFunc("/getNews", GetNews)
    http.HandleFunc("/getIndex", GetIndex)

    http.HandleFunc("/insertIndexdata", InsertIndexdata)
}
