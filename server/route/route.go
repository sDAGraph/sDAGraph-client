package route


//const _24K = (1 << 10) * 24// 24 MB

import (
    "fmt"
    //"strconv"
    "os"
    "io"
    "encoding/json"
    "io/ioutil"
    "net/http"
	//"path/filepath"
    "path"
	"net/url" 
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
	fmt.Println("mongoIP:",mongoIp)

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

    TestFile := func (res http.ResponseWriter, req *http.Request){
        res.Header().Add("Access-Control-Allow-Origin","*")
	    var in params.Img

        /*err := req.ParseMultipartForm(2 << 10)
		if err != nil {
			respondWithError(res, http.StatusInternalServerError, "FILE_TOO_BIG_2MB")
			return
		}*/

    	f, h, err := req.FormFile("file")
    	if err != nil {
            fmt.Println("error1: ", err)
			return
    	}
	
    	defer f.Close()
		/*
		fileBytes, err := ioutil.ReadAll(f)
		if err != nil {
			respondWithError(res, http.StatusInternalServerError, "INVALID_FILE")
			return
		}

		filetype := http.DetectContentType(fileBytes)
		if filetype != "image/jpeg" && filetype != "image/jpg" &&
			filetype != "image/gif" && filetype != "image/png" &&
			filetype != "application/pdf" {
			respondWithError(res, http.StatusInternalServerError, "INVALID_FILE_TYPE")
			return
		}
*/
        in.ImgName = h.Filename
        in.ImgUrl = params.UPLOAD_PATH + h.Filename

        fmt.Println("imgFormat:",in.ImgName)
        fmt.Println("img.ImgUrl:",in.ImgUrl)

    	h2, err:= os.Create(in.ImgUrl)
    	if err != nil {
        	fmt.Println(err)
        	return
    	}
    	defer h2.Close()

    	if _, err := io.Copy(h2, f); err != nil {
			fmt.Println("error2: ", err)
			return
    	}

        respondWithJson(res, http.StatusCreated, in)
    }

	DLTestFile := func (res http.ResponseWriter, req *http.Request){
		res.Header().Add("Access-Control-Allow-Origin","*")

		var in params.Img
		Filename := req.FormValue("filename")
        in.ImgName = Filename
        in.ImgUrl = params.UPLOAD_PATH + Filename

		file, err := os.OpenFile(in.ImgUrl, os.O_RDWR, 0666)
		fileName := path.Base(in.ImgUrl)
		fileName = url.QueryEscape(fileName) // 防止中文乱码
		res.Header().Add("content-disposition", "attachment; filename=\""+fileName+"\"")
		//res.Header().Add("Content-Type", "application/octet-stream")

		fmt.Println("img.ImgUrl:",in.ImgUrl)	
		//file, err := os.OpenFile(in.ImgUrl, os.O_RDWR, 0666)
		if err != nil{
            respondWithError(res,http.StatusInternalServerError, err.Error())
            return
		}
		defer file.Close()
		_, copyerr := io.Copy(res, file)
		if copyerr != nil {
			respondWithError(res,http.StatusInternalServerError, copyerr.Error())
			return
		}

        respondWithJson(res, http.StatusCreated, "OK")
	}

    http.HandleFunc("/getAllNews",GetAllNews)
    http.HandleFunc("/getNewsold", GetNewsold)
    http.HandleFunc("/getNews", GetNews)

    http.HandleFunc("/insertNews", InsertNews)

    http.HandleFunc("/insertNewsFile", InsertNewsFile)
    http.HandleFunc("/downloadNewsFile", DownloadNewsFile)
    http.HandleFunc("/readNewsFile", ReadNewsFile)
    http.HandleFunc("/deleteNewsFile", DeleteNewsFile)

    http.HandleFunc("/testFile", TestFile)
	http.HandleFunc("/dltestFile",DLTestFile)
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
