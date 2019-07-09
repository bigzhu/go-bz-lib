package modelsbz

import (
	"log"
	"testing"
)

// Base 基本模型的定义
type Son struct {
	Base
	SonName  string
	LastName string
}

// GetInstance 获取实例
func (son *Son) GetInstance() (result interface{}) {
	return son
}

func TestBase_GetFirst(t *testing.T) {
	Connect()
	DB.AutoMigrate(
		&Base{},
		&Son{},
	)
	GetFirst(&Base{ID: 1}, &Base{})
	son := Son{SonName: "kakazhu"}
	err := GetFirst(&son, &son)
	if err != nil {
		t.Fatal(err)
	}
	log.Println(son)
}

func TestSelf_GetFirst(t *testing.T) {
	Connect()
	base := Base{ID: 1}
	base.Get(&base)
	log.Println(base)
	son := Son{SonName: "kakazhu"}
	son.Get(&son)
	log.Println(son)
}
