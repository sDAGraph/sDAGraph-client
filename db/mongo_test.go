package sDAGraph_mongo

import (
	"fmt"
	"testing"
	"gopkg.in/mgo.v2/bson"
)

type newput struct{
	tmp string
}

func TestMongo(t *testing.T) {
	//連線
	db, session := GetDB("mongodb://192.168.51.202:27017", "sDAG")
	//範例
	in := bson.M{"a": 1, "b": true}
	//插入
	result := Insert(db,"test",in)
	fmt.Println(result)
	//尋找一個
	result2 := FindOne(db,"test",in)
	fmt.Println("final:",result2)
	//尋找多個
	result3 := FindAll(db,"test",in)
	fmt.Println(result3[0])
	//資料庫關閉
	session.Close()
}
