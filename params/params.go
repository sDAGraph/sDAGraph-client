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
}

type Env struct {
        Sue map[string] *EnvData
}

func Chain()(*ChainConfigStructure){
	d := &EnvData{
		Fee : big.NewInt(10),
		FeeAddress: "183344f5ae82fb707de5927120ab05398fd89517",
		FeeToken: "def",
		BlockSpeed: 1,
	}
	v := &Env{
		Sue : make(map[string]*EnvData),
	}
	v.Sue["dev"] = d

	s := &ChainConfigStructure{
		Hash: "fea4910f5d3e2d3af187cec5b8d8b1cfe99a9f5545ba50495bd42f4bae234b3a",
		Id: 101,
		V: "Kaman",
		Version: v,
	}
	return s
}
