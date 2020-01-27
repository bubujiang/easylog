package db


type Operate interface {
	Find()
	Insert()
	MkCondition()
}
