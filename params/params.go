package params

import (
	//"fmt"
	"math/big"
) 

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
	MongoSession string
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
		MongoSession:"test",
	}
	prod := &EnvData{
                Fee : big.NewInt(10),
                FeeAddress: "183344f5ae82fb707de5927120ab05398fd89517",
                FeeToken: "def",
                BlockSpeed: 1,
                MongoIp:"mongodb://192.168.51.202:27017",
		MongoName:"sDAG",
                MongoSession:"test",
        }
	v := &Env{
		Sue : make(map[string]*EnvData),
		Eleve : make(map[string]*EnvData),
	}
	v.Sue["dev"] = dev
	v.Sue["prod"] = prod
	s := &ChainConfigStructure{
		Hash: "fea4910f5d3e2d3af187cec5b8d8b1cfe99a9f5545ba50495bd42f4bae234b3a",
		Id: 101,
		V: "Kaman",
		Version: v,
	}
	return s
}
