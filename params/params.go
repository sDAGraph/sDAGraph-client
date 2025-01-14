package params

import (
	//"fmt"
	"math/big"
	"gopkg.in/mgo.v2/bson"
) 

const UPLOAD_PATH string = "/home/administrator/upLoad/"

type Img struct {
    ImgName string `bson:"_id"`
    ImgUrl  string `bson:"imgUrl"`
}

type NewsData struct {
    ID         bson.ObjectId `bson:"_id" json:"id"`
    Name       string `bson:"name" json:"name"`
    Intro      string `bson:"intro" json:"intro"`
    Title      string `bson:"title" json:"title"`
    Article    string `bson:"article" json:"article"`
}

type NewsFile struct{
    Abspath string `bson:"abspath" json:"abspath"`
    Name string
}

type ChainConfigStructure struct {
	Id int64
	Hash string
	V string
	Version *Env
}

type EnvData struct {
        Fee *big.Int
        FeeAddress string
        FeeToken string
        BlockSpeed int
        BlockTransaction int
        Consensus string
	MongoIp string
	MongoName string
	MongoCollection string
}

type Env struct {
        Sue map[string] *EnvData
	Eleve map[string] *EnvData
}

func Chain()(*ChainConfigStructure){
	dev := &EnvData{
		Fee : big.NewInt(10),
		FeeAddress: "183344f5ae82fb707de5927120ab05398fd89517",
		FeeToken: "def",
		BlockSpeed: 1,
		MongoIp:"mongodb://192.168.51.202:27017",
		MongoName:"sDAG",
		MongoCollection:"test",
	}
	prod := &EnvData{
        Fee : big.NewInt(10),
        FeeAddress: "183344f5ae82fb707de5927120ab05398fd89517",
        FeeToken: "def",
        BlockSpeed: 1,
        MongoIp:"mongodb://192.168.51.202:27017",
		MongoName:"sDAG",
        MongoCollection:"test",
    }
    dev203 := &EnvData{
        Fee : big.NewInt(10),
        FeeAddress: "183344f5ae82fb707de5927120ab05398fd89517",
        FeeToken: "def",
        BlockSpeed: 1,
        MongoIp:"mongodb://192.168.51.203:27017",
        MongoName:"sDAG",
        MongoCollection:"test",
    }
	v := &Env{
		Sue : make(map[string]*EnvData),
		Eleve : make(map[string]*EnvData),
	}
	v.Sue["dev"] = dev
	v.Sue["prod"] = prod
	v.Sue["dev203"] = dev203
	s := &ChainConfigStructure{
		Hash: "fea4910f5d3e2d3af187cec5b8d8b1cfe99a9f5545ba50495bd42f4bae234b3a",
		Id: 101,
		V: "Kaman",
		Version: v,
	}
	return s
}
