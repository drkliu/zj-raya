package meta_test

import (
 
	"context"
	"encoding/json"
	"log"
	"testing"

	"github.com/drkliu/zj-raya/internal/meta"

	"github.com/stretchr/testify/assert"

	 
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const uri = "mongodb://localhost:27017/?maxPoolSize=20&w=majority"

var productMetaTable = meta.MetaTable{
	Name: "products",
	RelationShips: []*meta.RelationShip{
		{
			Name:      "brandRelation",
			Type:      meta.RelationShipTypeManyToOne,
			Column:    "brand._id",
			RefTable:  "brands",
			RefColumn: "_id",
		},
	},
	PrimaryKey: meta.NewObjectIdPrimaryKey("products_pk_id", []string{"_id"}),
	Columns: []*meta.MetaColumn{
		{
			Name:     "_id",
			DataType: meta.DataTypeObjectId,
		},
		{
			Name:     "name",
			DataType: meta.DataTypeString,
		},
		{
			Name:     "shortDescription",
			DataType: meta.DataTypeString,
		},
		{
			Name:     "longDescription",
			DataType: meta.DataTypeString,
		},
		{
			Name:     "brand",
			DataType: meta.DataTypeJson,
			NestedColumns: []*meta.MetaColumn{
				{
					Name:     "_id",
					DataType: meta.DataTypeObjectId,
				},
				{
					Name:     "name",
					DataType: meta.DataTypeString,
				},
				{
					Name:     "logo",
					DataType: meta.DataTypeUrl,
				},
				{
					Name:     "media",
					DataType: meta.DataTypeJson,
					NestedColumns: []*meta.MetaColumn{
						{
							Name:     "_id",
							DataType: meta.DataTypeObjectId,
						},
						{
							Name:     "name",
							DataType: meta.DataTypeString,
						},
						{
							Name:     "url",
							DataType: meta.DataTypeUrl,
						},
					},
				},
			},
		},
		{
			Name:     "medias",
			DataType: meta.DataTypeJson,
			IsArray:  true,
			NestedColumns: []*meta.MetaColumn{
				{
					Name:     "_id",
					DataType: meta.DataTypeObjectId,
				},
				{
					Name:     "name",
					DataType: meta.DataTypeString,
				},
				{
					Name:     "url",
					DataType: meta.DataTypeUrl,
				},
			},
		},
		{
			Name:     "thumbnails",
			DataType: meta.DataTypeJson,
			IsArray:  true,
			NestedColumns: []*meta.MetaColumn{
				{
					Name:     "_id",
					DataType: meta.DataTypeObjectId,
				},
				{
					Name:     "name",
					DataType: meta.DataTypeString,
				},
				{
					Name:     "url",
					DataType: meta.DataTypeUrl,
				},
			},
		},
		{
			Name:     "galleries",
			DataType: meta.DataTypeJson,
			IsArray:  true,
			NestedColumns: []*meta.MetaColumn{
				{
					Name:     "_id",
					DataType: meta.DataTypeObjectId,
				},
				{
					Name:     "name",
					DataType: meta.DataTypeString,
				},
				{
					Name:     "url",
					DataType: meta.DataTypeUrl,
				},
			},
		},
		{
			Name:       "attributeSets",
			DataType:   meta.DataTypeJson,
			IsArray:    true,
			IsNestable: true,
			NestedColumns: []*meta.MetaColumn{
				{
					Name:     "name",
					DataType: meta.DataTypeString,
				},
				{
					Name:       "attributes",
					DataType:   meta.DataTypeJson,
					IsArray:    true,
					IsNestable: true,
					NestedColumns: []*meta.MetaColumn{
						{
							Name:     "name",
							DataType: meta.DataTypeString,
						},
						{
							Name:     "value",
							DataType: meta.DataTypeObject,
						},
						{
							Name:     "icon",
							DataType: meta.DataTypeUrl,
						},
					},
				},
			},
		},
		{
			Name:       "specifications",
			DataType:   meta.DataTypeJson,
			IsArray:    true,
			IsNestable: true,
			NestedColumns: []*meta.MetaColumn{
				{
					Name:     "name",
					DataType: meta.DataTypeString,
				},
				{
					Name:     "value",
					DataType: meta.DataTypeString,
				},
			},
		},
		{
			Name:       "packageLists",
			DataType:   meta.DataTypeJson,
			IsArray:    true,
			IsNestable: true,
			NestedColumns: []*meta.MetaColumn{
				{
					Name:     "name",
					DataType: meta.DataTypeString,
				},
				{
					Name:     "value",
					DataType: meta.DataTypeObject,
				},
			},
		},
		{
			Name:       "price",
			DataType:   meta.DataTypeJson,
			IsNestable: true,
			NestedColumns: []*meta.MetaColumn{
				{
					Name:     "currency",
					DataType: meta.DataTypeString,
				},
				{
					Name:      "amount",
					DataType:  meta.DataTypeDecimal,
					Length:    19,
					Precision: 2,
				},
			},
		},
		{
			Name:     "deleted",
			DataType: meta.DataTypeBool,
		},
		{
			Name:     "createAt",
			DataType: meta.DataTypeDateTime,
		},
		{
			Name:     "updateAt",
			DataType: meta.DataTypeDateTime,
		},
		{
			Name:     "deleteAt",
			DataType: meta.DataTypeDateTime,
		},
		{
			Name:     "createBy",
			DataType: meta.DataTypeObjectId,
		},
		{
			Name:     "updateBy",
			DataType: meta.DataTypeObjectId,
		},
	},
}
var brandsMetaTable = meta.MetaTable{
	Name:       "brands",
	PrimaryKey: meta.NewObjectIdPrimaryKey("products_pk_id", []string{"_id"}),
	Columns: []*meta.MetaColumn{
		{
			Name:     "_id",
			DataType: meta.DataTypeObjectId,
		},
		{
			Name:     "name",
			DataType: meta.DataTypeString,
		},
		{
			Name:     "logo",
			DataType: meta.DataTypeUrl,
		},
		{
			Name:     "description",
			DataType: meta.DataTypeString,
		},
		{
			Name:     "deleted",
			DataType: meta.DataTypeBool,
		},
		{
			Name:     "createAt",
			DataType: meta.DataTypeDateTime,
		},
		{
			Name:     "updateAt",
			DataType: meta.DataTypeDateTime,
		},
		{
			Name:     "deleteAt",
			DataType: meta.DataTypeDateTime,
		},
		{
			Name:     "createBy",
			DataType: meta.DataTypeObjectId,
		},
		{
			Name:     "updateBy",
			DataType: meta.DataTypeObjectId,
		},
	},
}
var cartsMetaTable = meta.MetaTable{
	Name:       "carts",
	PrimaryKey: meta.NewObjectIdPrimaryKey("carts_pk_id", []string{"_id"}),
	Columns: []*meta.MetaColumn{
		{
			Name:     "_id",
			DataType: meta.DataTypeObjectId,
		},
		{
			Name:     "userId",
			DataType: meta.DataTypeObjectId,
		},
		{
			Name:     "cartItems",
			DataType: meta.DataTypeJson,
			IsArray:  true,
			NestedColumns: []*meta.MetaColumn{
				{
					Name:     "productId",
					DataType: meta.DataTypeObjectId,
				},
				{
					Name:     "quantity",
					DataType: meta.DataTypeInt,
				},
				{
					Name:     "price",
					DataType: meta.DataTypeJson,
					NestedColumns: []*meta.MetaColumn{
						{
							Name:     "currency",
							DataType: meta.DataTypeString,
						},
						{
							Name:      "amount",
							DataType:  meta.DataTypeDecimal,
							Length:    19,
							Precision: 2,
						},
					},
				},
				{
					Name:     "promotions",
					DataType: meta.DataTypeJson,
					IsArray:  true,
					NestedColumns: []*meta.MetaColumn{
						{
							Name:     "promotionId",
							DataType: meta.DataTypeObjectId,
						},
						{
							Name:     "quantity",
							DataType: meta.DataTypeInt,
						},
						{
							Name:     "price",
							DataType: meta.DataTypeJson,
							NestedColumns: []*meta.MetaColumn{
								{
									Name:     "currency",
									DataType: meta.DataTypeString,
								},
								{
									Name:      "amount",
									DataType:  meta.DataTypeDecimal,
									Length:    19,
									Precision: 2,
								},
							},
						},
					},
				},
			},
		},
	},
}

func TestInsertMetaTable(t *testing.T) {

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	//insert table
	db := client.Database("tea")
	metaDatabase := meta.Database(*db)
	repository := meta.NewRepository(&metaDatabase)
	metaService := meta.NewService(&repository)
	id, err := metaService.InsertMetaTable(&productMetaTable)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("insertId:", id)
	assert.Equal(t, 1, 1)

}

func TestInsertManyMetaTables(t *testing.T) {

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	//insert table
	db := client.Database("tea")
	metaDatabase := meta.Database(*db)
	repository := meta.NewRepository(&metaDatabase)
	metaService := meta.NewService(&repository)
	ids, err := metaService.InsertManyMetaTables([]*meta.MetaTable{&productMetaTable, &brandsMetaTable,&cartsMetaTable})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("insertId:", ids)
	assert.Equal(t, 3, len(ids))

}

//test FindAllMetaTables
func TestFindAllMetaTables(t *testing.T) {
	
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	//insert table
	db := client.Database("tea")
	metaDatabase := meta.Database(*db)
	repository := meta.NewRepository(&metaDatabase)
	metaService := meta.NewService(&repository)
	//find all tables
	tables, err := metaService.FindAllMetaTables()
	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, 2, len(tables))
	//loop tables
	for _, table := range tables {
		t.Log(table.Id)
		log.Println(table.Name,table.Id)
	}
}

//insertOne
func TestInsertOne(t *testing.T) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	//insert table
	db := client.Database("tea")
	metaDatabase := meta.Database(*db)
	repository := meta.NewRepository(&metaDatabase)
	metaService := meta.NewService(&repository)
	productMetaTable,err:=metaService.FindMetaTableByName("products")
	if err != nil {
		log.Fatal(err)
	}
	
	product:=meta.DataObject{}
 
	err=json.Unmarshal([]byte(`{ 
		"name":"Apple iPhone 17",
		"shortDescription":"Apple iPhone 13 (A2634) 128GB 午夜色 支持移动联通电信5G 双卡双待手机",
		"longDescription": "品牌： Apple商品名称：AppleiPhone 13商品编号：100026667880商品毛重：320.00g商品产地：中国大陆运行内存：其他机身存储：128GB摄像头数量：后置双摄分辨率：其他屏幕比例：其他屏幕前摄组合：其他充电器：其他扬声器：立体声系统：iOS功能：超大字体，SOS功能，语音命令，语音识别(文字语音互转)支持IPv6：支持IPv6",
		"brand":{
			
		
			"name":"Apple",
			"logo":"https://img10.360buyimg.com/n9/s40x40_jfs/t1/90633/7/18242/187090/614be619E71982212/ee282423ecdc028c.jpg",
			"aaa":"ddddd",
			"media":{
				"name":"Apple",
				"url":"https://img10.360buyimg.com/n9/s40x40_jfs/t1/90633/7/18242/187090/614be619E71982212/ee282423ecdc028c.jpg"
			}
		},
		"medias":[{
			"name":"Apple iPhone 13",
			"url":"https://img10.360buyimg.com/n1/s450x450_jfs/t1/85203/17/16688/176934/614be5bfE28ca6187/6699eda01a9871e3.jpg"
		},
		{
			"url":"https://img10.360buyimg.com/n1/s450x450_jfs/t1/85203/17/16688/176934/614be5bfE28ca6187/6699eda01a9871e3.jpg",
			"name":"Apple iPhone 13 h",
			"xxx":"bbbb"
		}
		],
		"thumbnails":[{
			"name":"Apple iPhone 13",
			"url":"https://img10.360buyimg.com/n1/s450x450_jfs/t1/85203/17/16688/176934/614be5bfE28ca6187/6699eda01a9871e3.jpg"
		},
		{
			"url":"https://img10.360buyimg.com/n1/s450x450_jfs/t1/85203/17/16688/176934/614be5bfE28ca6187/6699eda01a9871e3.jpg",
			"name":"Apple iPhone 13 h",
			"xxx":"bbbb"
		}
		],
		"price":{
			"currency":"CNY",
			"amount":129900.345
		},
		"attributeSets":[
			{
			"name":"Apple iPhone 13",
			"attributes":[{
				"name":"Apple iPhone 13",
				"value":"Apple iPhone 13"
			},{
				"name":"color",
				"value":"red",
				"xxx":"bbbb",
				"fff":"gggg"
			}]
			}
		],
		"deleted":false,
		"createAt":"2019-01-01T00:00:00Z",
		"updateAt":"2019-01-01T00:00:00Z",
		"deleteAt":"2019-01-01T00:00:00Z",
		"createBy":"5c3c8f8f9f8f8e2c6a0a0a0a",
		"updateBy":"5c3c8f8f9f8f8e2c6a0a0a0a"}`),&product)
	 
	if err != nil {
		log.Fatal(err)
	}
	 
	 
	id, err := metaService.InsertOne(productMetaTable,  &product)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("insertIdd:", id)
	assert.Equal(t, 1, 1)
}
//insert carts
func TestInsertCarts(t *testing.T) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	//insert table
	db := client.Database("tea")
	metaDatabase := meta.Database(*db)
	repository := meta.NewRepository(&metaDatabase)
	metaService := meta.NewService(&repository)
	cartMetaTable,err:=metaService.FindMetaTableByName("carts")
	if err != nil {
		log.Fatal(err)
	}  
	
	cart:=meta.DataObject{}
	err=json.Unmarshal([]byte(`{
		"userId":"5c3c8f8f9f8f8e2c6a0a0a01",
		"cartItems":[
			{
				"productId":"5c3c8f8f9f8f8e2c6a0a0a0a",
				"quantity":1,
				"price":{
					"currency":"CNY",
					"amount":1299.00
				},
				"promotions":[
					{
						"promotionId":"5c3c8f8f9f8f8e2c6a0a0a0a",
						"promotionName":"Apple iPhone 13",
						"quantity":1,
						"price":{
							"currency":"CNY",
							"amount":1299.00
						}
					},
					{
						"promotionId":"5c3c8f8f9f8f8e2c6a0a0a0b",
						"promotionName":"Apple iPhone 131",
						"quantity":2,
						"price":{
							"currency":"CNY",
							"amount":6299.00,
							"xxx":"ffff"
						}
					}
				]
			},
			{
				"productId":"5c3c8f8f9f8f8e2c6a0a0a0b",
				"quantity":11,
				"price":{
					"currency":"CNY",
					"amount":1299.00
				},
				"promotions":[
					{
						"promotionId":"5c3c8f8f9f8f8e2c6a0a0a0a",
						"promotionName":"Apple iPhone 13",
						"quantity":1,
						"price":{
							"currency":"CNY",
							"amount":5299.00
						}
					},
					{
						"promotionId":"5c3c8f8f9f8f8e2c6a0a0a0b",
						"promotionName":"Apple iPhone 131",
						"quantity":2,
						"price":{
							"currency":"CNY",
							"amount":6299.07
						}
					}
				]
			}
		]
	}`),&cart)
	if err != nil {
		panic(err)
	}
	id, err := metaService.InsertOne(cartMetaTable,  &cart)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("insertIdd:", id)
	assert.Equal(t, 1, 1)
}


//findAll
func TestMetaServiceFindAll(t *testing.T) {

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
	t.Errorf("%s",json)
}