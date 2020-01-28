package db


type Client interface {
	Connect() error
	Close() error
	Validate() bool
	Find(map[string]interface{}) []map[string]interface{}
	Total(map[string]interface{}) int64
	Insert()
}
