package main

import (
	"context"

	"log"

	"github.com/drkliu/zj-raya/internal/meta"
 

 
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connection URI
const uri = "mongodb://localhost:27017/?maxPoolSize=20&w=majority"

func main() {
	
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
	json, err := metaService.FindAllToJson(productMetaTable)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s",json)
}

 
	 