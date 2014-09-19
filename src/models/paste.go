// paste
package models

import (
	"labix.org/v2/mgo/bson"
	"log"
	"strings"
	"time"
)

type Paste struct {
	Id        bson.ObjectId `bson:"_id"`
	Title     string
	Content   string
	CreatedOn time.Time
}

const DATABASE = "mgopastebin"
const PASTES = "pastes"

func (this Paste) Add() (id bson.ObjectId, err error, createdOn time.Time) {
	session, err1 := GetDB()
	if err1 != nil {
		log.Fatal("Error on database start - Add:", err1)
	}
	collection := session.DB(DATABASE).C(PASTES)
	id = bson.NewObjectId()
	createdOn = time.Now()
	this.Id = id
	this.CreatedOn = createdOn
	err = collection.Insert(&this)
	return
}
func ToObjectId(rawId string) bson.ObjectId {
	rawId = strings.TrimLeft(rawId, "ObjectIdHex(")
	rawId = strings.TrimRight(rawId, ")")
	rawId = strings.Trim(rawId, "\"")
	id := bson.ObjectIdHex(rawId)
	return id
}

func GetPaste(id bson.ObjectId) Paste {
	session, err1 := GetDB()
	if err1 != nil {
		log.Fatal("Error on database start - GetPaste():", err1)
	}
	collection := session.DB(DATABASE).C(PASTES)
	var paste Paste
	err := collection.FindId(id).One(&paste)
	if err != nil {
		log.Fatal("Error on database get - GetPaste():", err)
	}
	return paste
}

/*
func main() {
	fmt.Println("Hello World!")
}
*/
