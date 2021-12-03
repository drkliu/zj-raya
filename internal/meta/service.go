package meta

import (
	"bytes"
	"time"
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
	dors, err := s.FindAll(table)
	if err != nil {
		return "", err
	}
	return ToJson(dors)
}

func (s *service) InsertMetaTable(table *MetaTable) (*ID, error) {
	if len(table.ModelName) == 0 {
		table.ModelName = table.Name
	}
	setTrack(table, nil)
	return s.Repository.InsertMetaTable(table)
}

func (s *service) InsertManyMetaTables(tables []*MetaTable) ([]*ID, error) {
	for _, table := range tables {
		if len(table.ModelName) == 0 {
			table.ModelName = table.Name
		}
		setTrack(table, nil)
	}
	return s.Repository.InsertManyMetaTables(tables)
}

func (s *service) InsertMany(table *MetaTable, values []*DataObject) ([]*ID, error) {
	// for _, value := range values {
	// 	setTrack(table,nil)
	// }
	
	return s.Repository.InsertMany(table, values)
}
func setTrack(table *MetaTable, updateBy *ID) {
	table.CreatedAt = time.Now()
	table.UpdatedAt = time.Now()
	table.CreatedBy = updateBy
	table.UpdatedBy = updateBy
}

func ToJson(dors []*DataObjectResp) (string, error) {
	var buf bytes.Buffer

	buf.WriteString("[")
	for i, product := range dors {
		j, err := product.ToJson()
		if err != nil {
			return "", err
		}
		if i > 0 {
			buf.WriteString(",")
		}
		buf.WriteString(string(j))
	}
	buf.WriteString("]")
	return buf.String(), nil
}
