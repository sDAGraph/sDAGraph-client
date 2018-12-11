package sDAGraph_mongo

import(
	"fmt"
	"os"
    	"io"
	"path/filepath"
        
	"gopkg.in/mgo.v2"
        "gopkg.in/mgo.v2/bson"
	"sDAGraph-client/params"
)
/*
type User struct {
    Name string
    Id string
    Number uint
}*/

func GetDB(ip string, dbName string) (*mgo.Database,*mgo.Session) {
	session, err := mgo.Dial(ip)
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
	db := session.DB(dbName)
	return db, session
}

func DefaultGetDB() *mgo.Database {
	session, err := mgo.Dial("mongodb://192.168.51.212:27017")
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
	db := session.DB("testdatabase")
	return db
}

func Insert(db *mgo.Database, collection string, content interface{}) (error){
	c := db.C(collection)
	err := c.Insert(content)
	return err
}

func Update(db *mgo.Database, collection string, target interface{}, content interface{}) {
        c := db.C(collection)
	err := c.Update(&target, &content)
        if err != nil {
                fmt.Println(err)
        }
}

func FindOne(db *mgo.Database, collection string, content interface{}) (interface{}, error){
        c := db.C(collection)
	var users params.NewsData//bson.M
	err := c.Find(content).One(&users)
	return users, err
}

func UpdatebyID(db *mgo.Database, collection string, content params.NewsData) (error){
	c := db.C(collection)
	fmt.Println("id:",content.ID)
	err := c.UpdateId(content.ID, &content)
	return err
}

func FindbyID(db *mgo.Database, collection string, id string) (params.NewsData, error){
        c := db.C(collection)
        var users params.NewsData//bson.M
        //bsonid := bson.ObjectIdHex(id)
        //fmt.Println("bsonId1:",bsonid)
	//err := c.FindId(bsonid).One(&users)
	err := c.FindId(bson.ObjectIdHex(id)).One(&users)
	return users, err
}

func FindAll(db *mgo.Database, collection string) ([]params.NewsData, error){//interface{}){
        c := db.C(collection)
        var users []params.NewsData
	err := c.Find(bson.M{}).All(&users)
        return users, err
}

func Delete(db *mgo.Database, collection string, content params.NewsData) (error){
	c := db.C(collection)
	err := c.Remove(&content)
	return err
}

//create file
func InsertFile(db *mgo.Database, collection string, content string) (error){
	_, fileName := filepath.Split(content)
	
	c := db.GridFS(collection)

	user, err := c.Create(fileName)
	fmt.Println("create err:",err)

	out, err2 := os.OpenFile(content, os.O_RDWR, 0666)
	fmt.Println("open err:",err2)
	
	_,err = io.Copy(user, out)
	err = user.Close()
	err = out.Close()

	return err
}


func DloadFile(db *mgo.Database, collection string, content string) (error){
        _, fileName := filepath.Split(content)

	c := db.GridFS(collection)
        user, _ := c.Open(fileName)
	//fmt.Println("user:",user)
    	out, _ := os.OpenFile(content, os.O_CREATE| os.O_RDWR, 0666)
	_,err := io.Copy(out, user)
	//fmt.Println("Openf err:",err)
    	err = user.Close()
	err = out.Close()

    	return err
}

type fileinfo struct {
    //文件大小
    Length int32
    //md5
    Md5 string
    //文件名
    Filename string
}

/*
func FindAll(db *mgo.Database, collection string, content interface{}) ([]bson.M){//interface{}){
        c := db.C(collection)
        var users []bson.M
        c.Find(content).All(&users)
        return users
}
*/

