package database

type DataBase struct {
	ApiKey string
}

func NewDataBase(ApiKey string) *DataBase {
	return &DataBase{ApiKey: ApiKey}
}
