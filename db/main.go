package db

type DB2 struct {
	DSN string
	Database string
	Table string
}

type DB interface {
	Connect()
	Insert(interface{}) bool
	Find(map[string]interface{}) []map[string]interface{}
	Close()
}
