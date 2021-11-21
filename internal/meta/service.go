package meta

import (
	"bytes"
 
)

type MetaService interface {
	Repository
	FindAllToJson(table *MetaTable) (string, error)
}
type service struct {
	Repository
}

func NewService(repository *Repository) MetaService {
	return &service{
		Repository: *repository,
	}
}
func (s *service) FindAllToJson(table *MetaTable) (string, error) {
	dors,err:=s.FindAll(table)
	if err!=nil{
		return "",err
	}
	return ToJson(dors)
}
func ToJson(dors []*DataObjectResp) (string,error) {
	var buf bytes.Buffer
	
	buf.WriteString("[")
	for i, product := range dors {
		j,err:=product.ToJson()
		if err != nil {
			return "",err
		}
		if i>0{
			buf.WriteString(",")
		}
		buf.WriteString(string(j))
	}
	buf.WriteString("]")
	return buf.String(),nil
}