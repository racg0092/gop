package gop

import (
	"github.com/racg0092/gop/rdb"
	_ "modernc.org/sqlite"

	// _ "github.com/tursodatabase/go-libsql"
	"os"
	"testing"
)

func TestSqlUserCreateTable(t *testing.T) {
	orm, err := rdb.Open("libsql", "file:./rdb/test.db")
	if err != nil {
		t.Error(err)
	}

	u := User{}

	err = orm.CreateTable(u, "")
	if err != nil {
		t.Error(err)
	}
}

func TestMongoUser(t *testing.T) {
	d, e := NewDriver(MONGO, DriverConfig{Conn: os.Getenv("mdb"), Database: "ttt", Collection: "users"})
	if e != nil {
		t.Error(e)
	}

	c := GetConfig()
	c.IfUserNameBlankUseEmail().CheckIfPawnedPassword().CheckIfBadPassword()

	u := User{
		Email:     "richard@test.email",
		FirstName: "Richard",
		LastName:  "Chapman",
		Phone:     "150f25440094",
		Password:  "password123#foobarBarbar",
		Password2: "password123#foobarBarbar",
	}

	id, e := u.Save(d)
	if e != nil {
		t.Error(e)
	}
	t.Log("user id", id)
	userdata, e := d.ReadNonCritical(id, false)
	if e != nil {
		t.Error(e)
	}
	t.Log(userdata)
}
