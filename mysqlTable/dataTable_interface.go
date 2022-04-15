package mysqlTable

type DataBaseTable interface {
	TableName() string
	ToJson() string
}
