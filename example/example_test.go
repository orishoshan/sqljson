package example

import (
	"github.com/orishoshan/sqljson"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

type UserData struct {
	Name string `json:"name"`
}

type Model struct {
	ID       uint `gorm:"primarykey"`
	UserData sqljson.JSON[UserData]
}

func TestAutoScanAndValue(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = db.Debug()
	err = db.AutoMigrate(Model{})
	if err != nil {
		panic(err)
	}
	user := UserData{Name: "Beautiful Name"}
	err = db.Save(&Model{UserData: sqljson.From(user)}).Error
	if err != nil {
		panic(err)
	}

	m := Model{}
	err = db.First(&m).Error
	if err != nil {
		panic(err)
	}

	if m.UserData.Item.Name != user.Name {
		t.Errorf("expected name %s, got %s", user.Name, m.UserData.Item.Name)
	}

	println(m.UserData.Item.Name)

}
