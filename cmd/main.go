package main

import (
	"context"

	"log"

	"github.com/drkliu/zj-raya/internal/meta"
	jsoniter "github.com/json-iterator/go"

 
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connection URI
const uri = "mongodb://localhost:27017/?maxPoolSize=20&w=majority"

func main() {
	
	val := []byte(`{"ID":1,"Name":"Reds","Colors":["Crimson","Red","Ruby","Maroon"]}`)
	a:=jsoniter.Get(val, "*")
	log.Println(a)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	db := client.Database("tea")
	metaDatabase := meta.Database(*db)
	repository := meta.NewRepository(&metaDatabase)
	metaService := meta.NewService(&repository)
	productMetaTable,err:=metaService.FindMetaTableByName("products")
	if err != nil {
		log.Fatal(err)
	}
	//find all
	products, err := metaService.FindAll(productMetaTable)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("findAll:", products)
	// assert.Equal(t, 4, len(products))
	// product0,e := products[0].Get("name")
	// assert.Equal(t, true, e)
	// assert.Equal(t, "Apple iPhone 13", product0)
	log.Printf("%v",products[0])
}

 
	 