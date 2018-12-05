package route

import (
    "fmt"
    //"strconv"
    "encoding/json"
    "io/ioutil"
    "net/http"
    "sDAGraph-client/db"
    "sDAGraph-client/params"
    "gopkg.in/mgo.v2/bson"
)

type User struct {
    Id string
}

func Router(selversion string){
    fmt.Println("selversion",selversion)
    c := params.Chain()
    mongoIp := c.Version.Sue[selversion].MongoIp
    mongoName := c.Version.Sue[selversion].MongoName
    mongoCollection := c.Version.Sue[selversion].MongoCollection

    GetAllNews := func (res http.ResponseWriter, req *http.Request){
	res.Header().Add("Access-Control-Allow-Origin","*")
	db, session := sDAGraph_mongo.GetDB(mongoIp,mongoName)
	users, err := sDAGraph_mongo.FindAll(db,mongoCollection)
	session.Close()
	if err != nil {
		respondWithError(res, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(res, http.StatusOK, users)
    }

    GetNewsold := func (res http.ResponseWriter, req *http.Request){
        res.Header().Add("Access-Control-Allow-Origin","*")
        res.Header().Add("Content-Type", "application/json; charset=utf-8")
        
        b, _ := ioutil.ReadAll(req.Body)
        defer req.Body.Close()
	fmt.Println("b:",string(b))
	var newsdata User
        json.Unmarshal(b, &newsdata)

	//db, session := sDAGraph_mongo.GetDB(mongoIp,mongoName)
	in2 := bson.M{"id": "0xaaazz833"}
	fmt.Println("in2:",in2)
	/*esult2 := sDAGraph_mongo.FindOne(db,mongoCollection,in2)
        fmt.Println("final:",result2)
	respBody, _ := json.Marshal(result2)
	session.Close()
	res.Write([]byte(respBody))*/
    }

    GetNews := func (res http.ResponseWriter, req *http.Request){
        res.Header().Add("Access-Control-Allow-Origin","*")
        
	if (req.Method == "GET") {
	    param := req.FormValue("param")
	    value := req.FormValue("value")
	    fmt.Println("param:",param)
	    fmt.Println("val2:",value)
	    db, session := sDAGraph_mongo.GetDB(mongoIp,mongoName)
	    if (param == "id") {
	    	user, err := sDAGraph_mongo.FindbyID(db, mongoCollection, value)
                if err != nil {
                    respondWithError(res, http.StatusBadRequest, "Invalid ID")
                    return
                }
		respondWithJson(res, http.StatusOK, user)
	    }else{
        	in2 := bson.M{param : value}
        	fmt.Println("in2:",in2)
        	user, err := sDAGraph_mongo.FindOne(db,mongoCollection,in2)
                if err != nil {
                    respondWithError(res, http.StatusBadRequest, "Invalid ID")
                    return
                }
		respondWithJson(res, http.StatusOK, user)
	    }
	    /*if err != nil {
                respondWithError(res, http.StatusBadRequest, "Invalid ID")
                return
	    }
	    respondWithJson(res, http.StatusOK, user)
	    */
	    session.Close()
	}

    }

    CreateNews := func (res http.ResponseWriter, req *http.Request){
        res.Header().Add("Access-Control-Allow-Origin","*")

        defer req.Body.Close()
        var newsdata params.NewsData

	if err := json.NewDecoder(req.Body).Decode(&newsdata); err != nil {
		respondWithError(res, http.StatusBadRequest, "Invalid request payload")
		return
	}
	fmt.Println("name:",newsdata.Name)

	newsdata.ID = bson.NewObjectId()
	db, session := sDAGraph_mongo.GetDB(mongoIp,mongoName)
	in := newsdata
	//插入
        if err := sDAGraph_mongo.Insert(db,mongoCollection,in); err != nil {
		respondWithError(res, http.StatusInternalServerError, err.Error())
		return
	}

	session.Close()
	respondWithJson(res, http.StatusCreated, newsdata)
    }

    http.HandleFunc("/getAllNews",GetAllNews)
    http.HandleFunc("/getNewsold", GetNewsold)
    http.HandleFunc("/getNews", GetNews)

    http.HandleFunc("/createNews", CreateNews)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
