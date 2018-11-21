package main

import (
	"context"
        "fmt"
	//"log"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/common"
	"gopkg.in/mgo.v2"
        "gopkg.in/mgo.v2/bson"
	"os"
	"strconv"
	//"time"
)

type Person struct {
        Adress string
        Phone string
}

type Address struct {
        Address string "bson:`address`"
	Balance string "bson:`balance`"
}

func main() {
	db := getDB()
	data := findMuit(db,"topcollection")
	client, err := ethclient.Dial("http://192.168.51.203:19999")//"https://mainnet.infura.io")
	client2, err2 := ethclient.Dial("https://mainnet.infura.io")
	client3, err3 := ethclient.Dial("http://192.168.51.203:9999")
	
	if err != nil {
		fmt.Println(err)
	}
        if err2 != nil {
                fmt.Println(err)
        }
        if err3 != nil {
                fmt.Println(err)
        }

	//go process(client, db, u)
	for index, element := range data {
		fmt.Println(index)
		//time.Sleep(2000 * time.Microsecond)
		if (index%3==0){
			//time.Sleep(3000 * time.Microsecond)
			process(client, db, element)
		}
		if (index%3==1){
			//time.Sleep(1000 * time.Microsecond)
                        go process(client2, db, element)
                }
                if (index%3==2){
			//time.Sleep(10000 * time.Microsecond)
                        process(client3, db, element)
                }
	}
	select {}
	/*
	for index, element := range u {
		fmt.Println(index)
		fmt.Println(element.Address)
		fmt.Println(getBalance(client,element.Address))
		bal, status := getBalance(client,element.Address)
		if(status == true){
			account := Address{Address:element.Address,Balance:bal}
			insert(db,"account",account)
		}
	}*/
}

func process(client *ethclient.Client, db *mgo.Database, element Address){
        //for index, element := range data {
                //fmt.Println(index)
                fmt.Println(element.Address)
                fmt.Println(getBalance(client,element.Address))
                bal, status := getBalance(client,element.Address)
                if(status == true){
                        account := Address{Address:element.Address,Balance:bal}
			insert(db,"account",account)
                }
        //}
}

func getBalance(client *ethclient.Client, address string)(string, bool){
        header, err := client.BalanceAt(context.Background(),common.HexToAddress(address), nil)
        if err != nil {
                fmt.Println(err)
		return getBalance(client, address)//header.String(), false
        }
        return header.String(), true
}

func findOne() {
    db := getDB()
    c := db.C("topcollection")
    type Address struct {
        Address string "bson:`address`"
    }
    user := Address{}
    err := c.Find(bson.M{}).One(&user)
    if err != nil {
        panic(err)
    }
    fmt.Println(user)
    result := bson.M{}
    err = c.Find(nil).One(&result)
    if err != nil {
        panic(err)
    }
    fmt.Println(result)
}


func findMuit(db *mgo.Database, collection string) []Address{
    c := db.C(collection)
    var users []Address
    i, _ := strconv.Atoi(os.Args[1])
    f, _ := strconv.Atoi(os.Args[2])
    err := c.Find(nil).Skip(i).Limit(f).All(&users)

    if err != nil {
        panic(err)
    }
    
    //fmt.Println(users)
/*
    for index, element := range users {
	fmt.Println(index)
	fmt.Println(element.Address)
    }*/
	return users
/*
    var user Address
    iter := c.Find(nil).Iter()
    for iter.Next(&user) {
        fmt.Println(user)
    }
*/
}

func getDB() *mgo.Database {
    session, err := mgo.Dial("mongodb://192.168.51.202:27017")
    if err != nil {
        panic(err)
    }

    session.SetMode(mgo.Monotonic, true)
    db := session.DB("testdatabase")
    return db
}

func insert(db *mgo.Database, collection string, add Address) {
    c := db.C(collection)
    err := c.Insert(&add)
    if err != nil {
        fmt.Println(err)
    }
}
