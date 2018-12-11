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
	fmt.Println("GetAllNews")
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

	fmt.Println("GetNews")
	db, session := sDAGraph_mongo.GetDB(mongoIp,mongoName)
	if (req.Method == "GET") {
	    param := req.FormValue("param")
	    value := req.FormValue("value")
	    fmt.Println("param:",param)
	    fmt.Println("val2:",value)
	   
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
                    respondWithError(res, http.StatusBadRequest, "Invalid data")
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
	} else{
            defer req.Body.Close()
	    var newsdata params.NewsData

            if err := json.NewDecoder(req.Body).Decode(&newsdata); err != nil {
                respondWithError(res, http.StatusBadRequest, "Invalid request payload")
                return
            }

            if (req.Method == "PUT"){
                err := sDAGraph_mongo.UpdatebyID(db, mongoCollection, newsdata)
                if err  != nil {
                    respondWithError(res, http.StatusBadRequest, "Invalid data")
                    return
                }

	    }else if (req.Method == "DELETE"){
                err := sDAGraph_mongo.Delete(db, mongoCollection, newsdata)
                if err  != nil {
                    respondWithError(res, http.StatusBadRequest, "Invalid data")
                    return
                }
            }

	    respondWithJson(res, http.StatusOK, map[string]string{"result": "success"})
	}
	session.Close()

    }

    InsertNews := func (res http.ResponseWriter, req *http.Request){
        res.Header().Add("Access-Control-Allow-Origin","*")
	fmt.Println("InsertNews")

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

    InsertNewsFile := func (res http.ResponseWriter, req *http.Request){
        res.Header().Add("Access-Control-Allow-Origin","*")
	fmt.Println("InsertNewsFile")

        defer req.Body.Close()
        var newsdata params.NewsFile

        if err := json.NewDecoder(req.Body).Decode(&newsdata); err != nil {
            respondWithError(res, http.StatusBadRequest, "Invalid request payload")
            return
        }
        fmt.Println("name:",newsdata.Abspath)
	fmt.Println("abspath",newsdata.Abspath + newsdata.Name)

        //newsdata.ID = bson.NewObjectId()
        db, session := sDAGraph_mongo.GetDB(mongoIp,mongoName)
        in := newsdata
        //插入
        if err := sDAGraph_mongo.InsertFile(db,mongoCollection,in); err != nil {
            respondWithError(res, http.StatusInternalServerError, err.Error())
            return
        }

        session.Close()
        respondWithJson(res, http.StatusCreated, newsdata)
    }

    DownloadNewsFile := func (res http.ResponseWriter, req *http.Request){
        res.Header().Add("Access-Control-Allow-Origin","*")
	fmt.Println("DownloadNewsFile")
        defer req.Body.Close()
        var newsdata params.NewsFile

        if err := json.NewDecoder(req.Body).Decode(&newsdata); err != nil {
            respondWithError(res, http.StatusBadRequest, "Invalid request payload")
            return
        }
        fmt.Println("name:",newsdata.Abspath)

        //newsdata.ID = bson.NewObjectId()
        db, session := sDAGraph_mongo.GetDB(mongoIp,mongoName)
        in := newsdata
        //插入
        if err := sDAGraph_mongo.DloadFile(db,mongoCollection,in); err != nil {
            respondWithError(res, http.StatusInternalServerError, err.Error())
            return
        }

        session.Close()
        respondWithJson(res, http.StatusCreated, newsdata)
    }

    ReadNewsFile := func (res http.ResponseWriter, req *http.Request){
        res.Header().Add("Access-Control-Allow-Origin","*")
	fmt.Println("ReadNewsFile")
        //newsdata.ID = bson.NewObjectId()
        db, session := sDAGraph_mongo.GetDB(mongoIp,mongoName)

        data := sDAGraph_mongo.ReadAllFile(db,mongoCollection)
        
	session.Close()
        respondWithJson(res, http.StatusOK,data)
    }

    DeleteNewsFile := func (res http.ResponseWriter, req *http.Request){
        res.Header().Add("Access-Control-Allow-Origin","*")
	
	fmt.Println("DeleteNewsFile")
        defer req.Body.Close()
        var newsdata params.NewsFile

        if err := json.NewDecoder(req.Body).Decode(&newsdata); err != nil {
            respondWithError(res, http.StatusBadRequest, "Invalid request payload")
            return
        }
        fmt.Println("name:",newsdata.Abspath)

        //newsdata.ID = bson.NewObjectId()
        db, session := sDAGraph_mongo.GetDB(mongoIp,mongoName)
        in := newsdata
        //插入
        if err := sDAGraph_mongo.DeleteFile(db,mongoCollection,in); err != nil {
            respondWithError(res, http.StatusInternalServerError, err.Error())
            return
        }

        session.Close()
        respondWithJson(res, http.StatusCreated, newsdata)
    }

    http.HandleFunc("/getAllNews",GetAllNews)
    http.HandleFunc("/getNewsold", GetNewsold)
    http.HandleFunc("/getNews", GetNews)

    http.HandleFunc("/insertNews", InsertNews)

    http.HandleFunc("/insertNewsFile", InsertNewsFile)
    http.HandleFunc("/downloadNewsFile", DownloadNewsFile)
    http.HandleFunc("/readNewsFile", ReadNewsFile)
    http.HandleFunc("/deleteNewsFile", DeleteNewsFile)
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
