package meta

import (
	"time"

	//"github.com/emirpasic/gods/maps/linkedhashmap"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type 
(
	DataType int8
	ID primitive.ObjectID
	Database mongo.Database
	InputType int8
	RelationShipType int8
	IdGeneratorType int8
	DataObject map[string]interface{}
	DataObjectResp bson.D
)

type Entry struct {
	Key string
	Value interface{}
}
 
const (
	DataTypeUnknown DataType = iota
	DataTypeString
	DataTypeInt
	DataTypeLong
	DataTypeFloat
	DataTypeDouble
	DataTypeDecimal
	DataTypeBool
	DataTypeDateTime
	DataTypeTime
	DataTypeTimestamp
	DataTypeObjectId
	DataTypeJson
	DataTypeObject
	DataTypeUrl
)

func (d DataType) String() string {
	switch d {
	case DataTypeUnknown:
		return "unknown"
	case DataTypeString:
		return "string"
	case DataTypeInt:
		return "int"
	case DataTypeLong:
		return "long"
	case DataTypeFloat:
		return "float"
	case DataTypeDouble:
		return "double"
	case DataTypeDecimal:
		return "decimal"
	case DataTypeBool:
		return "bool"
	case DataTypeDateTime:
		return "dateTime"
	case DataTypeTime:
		return "time"
	case DataTypeTimestamp:
		return "timestamp"
	case DataTypeObjectId:
		return "objectId"
	case DataTypeJson:
		return "json"
	case DataTypeObject:
		return "object"
	case DataTypeUrl:
		return "url"
	default:
		return "unknown"
	}
}
func ParseDataType(i int8) DataType {
	switch i {
	case 0:
		return DataTypeUnknown
	case 1:
		return DataTypeString
	case 2:
		return DataTypeInt
	case 3:
		return DataTypeLong
	case 4:
		return DataTypeFloat
	case 5:
		return DataTypeDouble
	case 6:
		return DataTypeDecimal
	case 7:
		return DataTypeBool
	case 8:
		return DataTypeDateTime
	case 9:
		return DataTypeTime
	case 10:
		return DataTypeTimestamp
	case 11:
		return DataTypeObjectId
	case 12:
		return DataTypeJson
	case 13:
		return DataTypeObject
	case 14:
		return DataTypeUrl
	default:
		return DataTypeUnknown
	}
}


const (
	InputTypeUnknown InputType = iota
	InputTypeText
	InputTypeTextArea
	InputTypeNumber
	InputTypeDate
	InputTypeRadio
	InputTypeCheckbox
	InputTypeSelect
	InputTypeFile
	InputTypeImage
	InputTypeHidden
	InputTypePassword
)

func (i InputType) String() string {
	switch i {
	case InputTypeText:
		return "text"
	case InputTypeTextArea:
		return "textArea"
	case InputTypeNumber:
		return "number"
	case InputTypeDate:
		return "date"
	case InputTypeRadio:
		return "radio"
	case InputTypeCheckbox:
		return "checkbox"
	case InputTypeSelect:
		return "select"
	case InputTypeFile:
		return "file"
	case InputTypeImage:
		return "image"
	case InputTypeHidden:
		return "hidden"
	case InputTypePassword:
		return "password"
	default:
		return "unknown"
	}
}
func ParseInputType(i int8) InputType {
	switch i {
	case 0:
		return InputTypeUnknown
	case 1:
		return InputTypeText
	case 2:
		return InputTypeTextArea
	case 3:
		return InputTypeNumber
	case 4:
		return InputTypeDate
	case 5:
		return InputTypeRadio
	case 6:
		return InputTypeCheckbox
	case 7:
		return InputTypeSelect
	case 8:
		return InputTypeFile
	case 9:
		return InputTypeImage
	case 10:
		return InputTypeHidden
	case 11:
		return InputTypePassword
	default:
		return InputTypeUnknown
	}
}
const (
	IdGeneratorTypeUnknown IdGeneratorType = iota
	IdGeneratorTypeUUID
	IdGeneratorTypeObjectId
	IdGeneratorTypeSequence
	//auto increment
	IdGeneratorTypeAutoIncrement
)
//IdGeneratorType string
func (id IdGeneratorType) String() string {
	switch id {
	case IdGeneratorTypeUnknown:
		return "unknown"
	case IdGeneratorTypeUUID:
		return "uuid"
	case IdGeneratorTypeObjectId:
		return "objectId"
	case IdGeneratorTypeSequence:
		return "sequence"
	case IdGeneratorTypeAutoIncrement:
		return "autoIncrement"
	default:
		return "unknown"
	}
}
//IdGeneratorType parse
func ParseIdGeneratorType(i int8) IdGeneratorType {
	switch i {
	case 0:
		return IdGeneratorTypeUnknown
	case 1:
		return IdGeneratorTypeUUID
	case 2:
		return IdGeneratorTypeObjectId
	case 3:
		return IdGeneratorTypeSequence
	case 4:
		return IdGeneratorTypeAutoIncrement
	default:
		return IdGeneratorTypeUnknown
	}
}
//DataObject
func (do *DataObject) Put(key string, value interface{})  {
	v:=map[string]interface{}(*do)
	v[key]=value
}
func (do *DataObject) Get(key string) (interface{},bool) {
	v:=map[string]interface{}(*do)
	val,ok:=v[key]
	return val,ok
}
//dataObjectResp get
func (do *DataObjectResp) Get(key string) (interface{},bool) {
	v:=bson.D(*do)
	
	value,ok:=v.Map()[key]
	return value,ok
}
	
	 
//map to DataObject
func MapToDataObject(m *map[string]interface{}) *DataObject {
	var do DataObject
	for k,v := range *m {
		do.Put(k,v)
	}
	return &do
}

type MetaTable struct {
	Id            primitive.ObjectID `bson:"_id,omitempty"`
	Name          string
	Columns       []*MetaColumn
	PrimaryKey    *PrimaryKey
	RelationShips []*RelationShip
	Indexes       []*MetaIndex
	//createAt
	CreateAt time.Time
	//updateAt
	UpdateAt time.Time
	//createBy
	CreateBy *primitive.ObjectID
	//updateBy
	UpdateBy *primitive.ObjectID
	
}
type MetaColumn struct {
	Name         string
	DataType     DataType
	Length       int
	Precision    int
	Scale        int
	IsNullable   bool
	Validators   []string
	IsNestable   bool
	IsArray	  	 bool
	IsSerachable bool
	InputType    InputType
	NestedColumns []*MetaColumn
}
type PrimaryKey struct {
	Name string
	ColumnNames []string
	IdGeneratorType IdGeneratorType
}
type IdGenerator struct {
	Type IdGeneratorType
	Config string
}
//Primarykey ObjectId
func NewObjectIdPrimaryKey(name string,columnNames []string) *PrimaryKey {
	return &PrimaryKey{
		Name: name,
		ColumnNames: columnNames,
		IdGeneratorType: IdGeneratorTypeObjectId,
	}
}

type MetaIndex struct {
}
type RelationShip struct {
	Name         string
	Type         RelationShipType
	Column  string
	RefTable     string
	RefColumn 	string
}


const (
	RelationShipTypeUnknown RelationShipType = iota
	RelationShipTypeOneToOne
	RelationShipTypeOneToMany
	RelationShipTypeManyToOne
	RelationShipTypeManyToMany
)

func (r RelationShipType) String() string {
	switch r {
	case RelationShipTypeOneToOne:
		return "oneToOne"
	case RelationShipTypeOneToMany:
		return "oneToMany"
	case RelationShipTypeManyToOne:
		return "manyToOne"
	case RelationShipTypeManyToMany:
		return "manyToMany"
	default:
		return "unknown"
	}
}
func ParseRelationShipType(i int8) RelationShipType {
	switch i {
	case 0:
		return RelationShipTypeUnknown
	case 1:
		return RelationShipTypeOneToOne
	case 2:
		return RelationShipTypeOneToMany
	case 3:
		return RelationShipTypeManyToOne
	case 4:
		return RelationShipTypeManyToMany
	default:
		return RelationShipTypeUnknown
	}
}

 
func NilObjectID() ID {
	return ID(primitive.NilObjectID)
}
func (id ID) String() string {
	
	return primitive.ObjectID(id).Hex()
}
func (id ID) ToObjectId() primitive.ObjectID {
	return primitive.ObjectID(id)
}
//toID
func ParseID(id interface{}) ID {
	switch id:=id.(type) {
	case string:
		objID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			panic(err)
		}
		return ID(objID)
	case primitive.ObjectID:
		return ID(id)
	default:
		return NilObjectID()
	}
}

