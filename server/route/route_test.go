package route

import "testing"
import "fmt"
import "sDAGraph-client/db"
import "gopkg.in/mgo.v2/bson"
import "sDAGraph-client/params"

func TestSignbatch(t *testing.T) {
        // write some test
	d, _ := sDAGraph_mongo.GetDB("mongodb://192.168.51.202:27017","sDAG")
	dd, _ := sDAGraph_mongo.FindbyID(d, "test", "5c06324c1df425d54e55eb71")
	fmt.Println("dd:",dd)
        
	var user params.NewsData
        user.ID = bson.ObjectIdHex("5c06324c1df425d54e55eb71")
        user.Name = "tet5"
        user.Number = 99

        sDAGraph_mongo.Delete(d, "test", user)
        dd3,_ := sDAGraph_mongo.FindbyID(d, "test", "5c06324c1df425d54e55eb71")
        fmt.Println("dd3:",dd3)

        /*test update
	sDAGraph_mongo.UpdatebyID(d, "test", user)
        dd3,_ := sDAGraph_mongo.FindbyID(d, "test", "5c06324c1df425d54e55eb71")
        fmt.Println("dd3:",dd3)
        */

	//in := bson.M{"name":"tet25"}
	//dd2,_ := sDAGraph_mongo.FindOne(d, "test", in)
	//fmt.Println("dd:",dd2)
	
	//in2 := bson.M{"name":"tet15","number":31}
	//sDAGraph_mongo.Update(d, "test", in, in2)
        //dd3,_ := sDAGraph_mongo.FindOne(d, "test", in2)
        //fmt.Println("dd3:",dd3)

        //aaa := newsdata.ID.Hex()
        //dd3,_ := sDAGraph_mongo.FindbyID(db, mongoCollection, aaa)

}
