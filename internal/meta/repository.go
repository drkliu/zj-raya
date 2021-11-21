package meta

import (
	"context"
	"errors"

	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	table_name = "metas"
	timeout    = 30 * time.Second
)

type Repository interface {
	FindMetaTableById(id ID) (*MetaTable, error)
	FindMetaTableByName(tableName string) (*MetaTable, error)
	FindAllMetaTables() ([]*MetaTable, error)
	InsertMetaTable(table *MetaTable) (*ID, error)
	InsertManyMetaTables(tables []*MetaTable) ([]*ID, error)
	FindAll(table *MetaTable) ([]*DataObjectResp, error)
	FindOne(table *MetaTable, id ID) (*DataObjectResp, error)
	InsertOne(table *MetaTable, value *DataObject) (*ID, error)
	InsertMany(table *MetaTable, values []*DataObject) ([]*ID, error)
}

type repository struct {
	db *Database
}

func NewRepository(db *Database) Repository {
	return &repository{db}
}
func (r *repository) FindMetaTableById(id ID) (*MetaTable, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	db := mongo.Database(*r.db)
	coll := db.Collection(table_name)
	var table MetaTable
	err := coll.FindOne(ctx, bson.M{"_id": id.ToObjectId()}).Decode(&table)
	return &table, err
}

func (r *repository) FindMetaTableByName(tableName string) (*MetaTable, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	db := mongo.Database(*r.db)
	coll := db.Collection(table_name)
	var table MetaTable
	err := coll.FindOne(ctx, bson.M{"name": tableName}).Decode(&table)

	return &table, err
}
func (r *repository) FindAllMetaTables() ([]*MetaTable, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	db := mongo.Database(*r.db)
	coll := db.Collection(table_name)
	var tables []*MetaTable
	cursor, err := coll.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	for cursor.Next(ctx) {
		var table MetaTable
		err := cursor.Decode(&table)

		if err != nil {
			return nil, err
		}
		tables = append(tables, &table)
	}
	return tables, nil
}
func (r *repository) InsertMetaTable(table *MetaTable) (*ID, error) {
	db := mongo.Database(*r.db)
	coll := db.Collection(table_name)
	result, err := coll.InsertOne(context.TODO(), table)
	if err != nil {
		return nil, err
	}
	id := ParseID(result.InsertedID)
	return &id, nil
}
func (r *repository) InsertManyMetaTables(tables []*MetaTable) ([]*ID, error) {
	db := mongo.Database(*r.db)
	coll := db.Collection(table_name)
	//convert to bson.D
	var bsonTables []interface{}
	for _, table := range tables {
		table.CreateAt = time.Now()
		table.UpdateAt = time.Now()
		bsonTables = append(bsonTables, table)
	}
	result, err := coll.InsertMany(context.TODO(), bsonTables)
	if err != nil {
		return nil, err
	}
	var ids = make([]*ID, len(result.InsertedIDs))
	var id ID
	for i, iid := range result.InsertedIDs {
		id = ParseID(iid)
		ids[i] = &id
	}
	return ids, nil
}
func (r *repository) FindAll(table *MetaTable) ([]*DataObjectResp, error) {

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	db := mongo.Database(*r.db)
	coll := db.Collection(table.Name)
	cursor, err := coll.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	result := []*DataObjectResp{}
	for cursor.Next(ctx) {
		var dor DataObjectResp
		//var v interface{}
		//var v interface{}
		err := cursor.Decode(&dor)
		// for _, c := range table.Columns {

		// 	if _,e:=dor.Get(c.Name); e{
		// 		result=append(result, &dor)
		// 	    continue

		// 	}
		// }
		if err != nil {
			return nil, err
		}
		result = append(result, &dor)
	}
	return result, nil
}

//column value to Entry
// func columnValueToEntryValue(column *MetaColumn, value interface{}, entry interface{}) {
// 	if len(column.NestedColumns)>0{
// 		switch mapVal:=value.(type) {
// 		case map[string]interface{}:

// 			for _,nc:=range column.NestedColumns{
// 				if v,e:=mapVal[nc.Name]; e{
// 					if len(nc.NestedColumns)>0{
// 						es:=[]Entry{}
// 						*entries=append(*entries, Entry{Key:nc.Name,Value:es})
// 						columnValueToEntryValue(nc,v,&es)
// 					}else{

// 						*entries=append(*entries, Entry{Key:nc.Name,Value:v})
// 					}
// 				}
// 			}
// 		default:
// 			*entries=append(*entries, Entry{Key:column.Name,Value:value})
// 		}
// 	}
// }

func (r *repository) FindOne(table *MetaTable, id ID) (*DataObjectResp, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	db := mongo.Database(*r.db)
	coll := db.Collection(table.Name)
	var value DataObjectResp
	err := coll.FindOne(ctx, bson.M{"_id": id.ToObjectId()}).Decode(&value)
	return &value, err
}
func (r *repository) InsertOne(table *MetaTable, do *DataObject) (*ID, error) {
	db := mongo.Database(*r.db)
	coll := db.Collection(table.Name)
	insertDocument := bson.D{}
	for _, c := range table.Columns {
		if e, exist := do.Get(c.Name); exist {
			v, err := assemblyNestedColumns(c, e)
			if err != nil {
				return nil, err
			}
			insertDocument = append(insertDocument, bson.E{Key: c.Name, Value: v})
		}
	}
	result, err := coll.InsertOne(context.TODO(), insertDocument)
	if err != nil {
		return nil, err
	}
	id := ParseID(result.InsertedID)
	return &id, nil
}

func assemblyNestedColumns(c *MetaColumn, val interface{}) (interface{}, error) {
	switch c.DataType {
	case DataTypeJson:
		switch c.IsArray {
		case false:
			if val, ok := val.(map[string]interface{}); ok {
				b := bson.D{}
				for _, nc := range c.NestedColumns {
					if v, e := val[nc.Name]; e {
						if len(nc.NestedColumns) > 0 {
							ncv, err := assemblyNestedColumns(nc, v)
							if err != nil {
								return nil, err
							}
							b = append(b, bson.E{Key: nc.Name, Value: ncv})
						} else {
							b = append(b, bson.E{Key: nc.Name, Value: v})
						}
					}
				}
				return b, nil
			}
			return nil, errors.New("column:" + c.Name + ",value is not map[string]interface{}")
		default:
			var ds []bson.D
			if val, ok := val.([]interface{}); ok {
				for _, vi := range val {
					d := bson.D{}
					for _, nc := range c.NestedColumns {
						if vm, e := vi.(map[string]interface{}); e {
							if v, e := vm[nc.Name]; e {
								if len(nc.NestedColumns) > 0 {
									ncv, err := assemblyNestedColumns(nc, v)
									if err != nil {
										return nil, err
									}
									d = append(d, bson.E{Key: nc.Name, Value: ncv})
								} else {
									d = append(d, bson.E{Key: nc.Name, Value: v})
								}
							}
						}
					}
					if len(d) > 0 {
						ds = append(ds, d)
					}
				}
				return ds, nil
			}
			return nil, errors.New("column:" + c.Name + ",value is not []interface{}")
		}
	default:
		switch v := val.(type) {
		case map[string]interface{}, []interface{}:
			return v, errors.New("column:" + c.Name + ",value should not be a map[string]interface{} or []interface{}")
		default:
			return v, nil
		}
	}

	// switch val:=val.(type) {
	// case map[string]interface{}:
	// 	b:=bson.D{}
	// 	for _,nc:=range c.NestedColumns{
	// 		if v,e:=val[nc.Name]; e{
	// 			if len(nc.NestedColumns)>0{
	// 				b=append(b,bson.E{Key: nc.Name,Value: assemblyNestedColumns(nc,v)})
	// 			}else{
	// 				b=append(b, bson.E{Key:nc.Name, Value:v})
	// 			}
	// 		}
	// 	}
	// 	return b
	// case []interface{}:
	// 	var ds []bson.D
	// 	for _,vi:=range val{
	// 		d:=bson.D{}
	// 		for _,nc:=range c.NestedColumns{
	// 			if vm,e:=vi.(map[string]interface{}); e{
	// 				if v,e:=vm[nc.Name]; e{
	// 					if len(nc.NestedColumns)>0{
	// 						d=append(d,bson.E{Key: nc.Name,Value: assemblyNestedColumns(nc,v)})
	// 					}else{
	// 						d=append(d, bson.E{Key:nc.Name, Value:v})
	// 					}
	// 				}
	// 			}
	// 		}
	// 		if len(d)>0{
	// 			ds=append(ds, d)
	// 		}
	// 	}
	// 	return ds
	// default:
	// 	return val
	// }
}
func (r *repository) InsertMany(table *MetaTable, values []*DataObject) ([]*ID, error) {
	db := mongo.Database(*r.db)
	coll := db.Collection(table.Name)
	//convert to bson.D
	var bsonValues []interface{}
	for _, value := range values {
		(*value).Put("createAt", time.Now())
		(*value).Put("updateAt", time.Now())
		bsonValues = append(bsonValues, value)
	}
	result, err := coll.InsertMany(context.TODO(), bsonValues)
	if err != nil {
		return nil, err
	}
	var ids = make([]*ID, len(result.InsertedIDs))
	var id ID
	for i, iid := range result.InsertedIDs {
		id = ParseID(iid)
		ids[i] = &id
	}
	return ids, nil
}
