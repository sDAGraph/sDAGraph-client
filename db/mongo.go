package sDAGraph_mongo

import(
	"fmt"
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

/*
func FindAll(db *mgo.Database, collection string, content interface{}) ([]bson.M){//interface{}){
        c := db.C(collection)
        var users []bson.M
        c.Find(content).All(&users)
        return users
}
*/

