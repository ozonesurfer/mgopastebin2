// repository
package models

import (
	//	"fmt"
	"labix.org/v2/mgo"
	"log"
	"sort"
)

/*
func main() {
	fmt.Println("Hello World!")
}
*/

const DB_HOST = "127.0.0.1"

// needed for slice sorting - start
type ByCreated []Paste

func (this ByCreated) Len() int           { return len(this) }
func (this ByCreated) Less(i, j int) bool { return this[i].CreatedOn.After(this[j].CreatedOn) } // desc
func (this ByCreated) Swap(i, j int)      { this[i], this[j] = this[j], this[i] }

// end

func init() {
	session, err := GetDB()
	defer session.Close()
	if err == nil {
		log.Println("Database connection verified")
	} else {
		log.Fatalln("models.init failed", err)
	}

}
func GetDB() (session *mgo.Session, err error) {
	session, err = mgo.Dial(DB_HOST)
	return
}

func GetAll() []Paste {
	session, err1 := GetDB()
	var pastes []Paste
	defer session.Close()
	if err1 != nil {
		log.Fatal("Error on database start - Add:", err1)
	}
	collection := session.DB(DATABASE).C(PASTES)
	collection.Find(nil).All(&pastes)
	sort.Sort(ByCreated(pastes))
	var limited []Paste
	if len(pastes) > 5 {
		limited = pastes[0:5]
	} else {
		limited = pastes
	}
	return limited
}
