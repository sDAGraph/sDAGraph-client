package route

import "testing"
import "fmt"
import "sDAGraph-client/db"
import "gopkg.in/mgo.v2/bson"

func TestSignbatch(t *testing.T) {
        // write some test
	d, _ := sDAGraph_mongo.GetDB("mongodb://192.168.51.202:27017","sDAG")
	dd, _ := sDAGraph_mongo.FindbyID(d, "test", "5c06324c1df425d54e55eb71")
	fmt.Println("dd:",dd)
	
	in := bson.M{"name":"tet15"}
	dd2 := sDAGraph_mongo.FindOne(d, "test", in)
	fmt.Println("dd:",dd2)
}
