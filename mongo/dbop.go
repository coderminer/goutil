package main

import (
	"fmt"
	db "mongo/mongodb"
	"time"

	"github.com/globalsign/mgo/bson"
)

type Data struct {
	Id      bson.ObjectId `bson:"_id"`
	Title   string        `bson:"title"`
	Des     string        `bson:"des"`
	Content string        `bson:"content"`
	Date    time.Time     `bson:"date"`
}

const (
	database   = "Test"
	collection = "TestModel"
)

// Examples : the operation of mgo to MongoDB
func main() {
	//insert one document
	data := &Data{
		Id:      bson.NewObjectId(),
		Title:   "博客的标题 1",
		Des:     "博客描述信息 1",
		Content: "博客的具体内容 1",
		Date:    time.Now(),
	}

	err := db.Insert(database, collection, data)
	if err != nil {
		fmt.Println("insert one doc", err)
	}

	// find one with all fields
	var result Data
	err = db.FindOne(database, collection, bson.M{"_id": bson.ObjectIdHex("5b3db2334d661ff46ee14b9c")}, nil, &result)
	fmt.Println("find one with all fields", result)

	// find one without id field
	var result1 Data
	err = db.FindOne(database, collection, bson.M{"_id": bson.ObjectIdHex("5b3db2334d661ff46ee14b9c")}, bson.M{"_id": 0}, &result1)
	fmt.Println("find one without id field", result1)

	//find all documents
	var allResult []Data
	err = db.FindAll(database, collection, nil, nil, &allResult)
	fmt.Println("find all docs", allResult)

	// find all documents with query and selector
	var allResult1 []Data
	err = db.FindAll(database, collection, bson.M{"title": "博客的标题 1"}, bson.M{"_id": 0}, &allResult1)
	fmt.Println("find all docs with query and selector", allResult1)

	//find documents with page and limit
	var resultWithPage []Data
	err = db.FindPage(database, collection, 0, 4, nil, bson.M{"_id": 0}, &resultWithPage)
	fmt.Println("find docs with page and limit", resultWithPage)

	//find the cursor
	var iterAll []Data
	iter := db.FindIter(database, collection, nil)
	err = iter.All(&iterAll)
	fmt.Println("find cursor ", iterAll)

	//update one document
	err = db.Update(database, collection, bson.M{"_id": bson.ObjectIdHex("5b3db2334d661ff46ee14b9c")}, bson.M{"$set": bson.M{
		"title": "更新后的标题",
		"des":   "更新后的描述信息",
		"date":  time.Now(),
	}})

	if err != nil {
		fmt.Println("upate one error", err)
	}

	//update all docments
	/*err = db.UpdateAll(database, collection, nil, bson.M{"$set": bson.M{
		"title": "更新所有的标题",
		"date":  time.Now(),
	}})
	if err != nil {
		fmt.Println("update all docs error ", err)
	}*/

	//delete one docment
	err = db.Remove(database, collection, bson.M{"_id": bson.ObjectIdHex("5b3db2334d661ff46ee14b99")})
	if err != nil {
		fmt.Println("remove one doc error", err)
	}

	//upsert the docment
	err = db.Upsert(database, collection, bson.M{"title": "Title Upsert"}, bson.M{"$set": bson.M{
		"des":     "描述Upsert",
		"date":    time.Now(),
		"Content": "内容Upsert",
	}})
	if err != nil {
		fmt.Println("upsert docment error", err)
	}

	//bulk insert docments
	d1 := &Data{
		Id:      bson.NewObjectId(),
		Title:   "bulk title",
		Des:     "bulk Des",
		Content: "bulk content",
		Date:    time.Now(),
	}
	d2 := &Data{
		Id:      bson.NewObjectId(),
		Title:   "bulk title",
		Des:     "bulk Des",
		Content: "bulk content",
		Date:    time.Now(),
	}

	insertResult, _ := db.BulkInsert(database, collection, d1, d2)
	fmt.Println("bulk insert docs", insertResult)

	//bulk update
	up1 := bson.M{"title": "bulk update title"}
	up2 := bson.M{"$set": bson.M{"title": "bulk update title"}}

	up3 := bson.M{"_id": bson.ObjectIdHex("5b3dbd7a9d5e3e314c93d150")}
	up4 := bson.M{"$set": bson.M{"des": "bulk update des"}}

	updateResult, _ := db.BulkUpdate(database, collection, up1, up2, up3, up4)
	fmt.Println("bulk update result", updateResult)

	//bulk update all
	up5 := bson.M{"title": "bulk title"}
	up6 := bson.M{"$set": bson.M{"title": "bulk update title"}}

	updateAllResult, _ := db.BulkUpdateAll(database, collection, up5, up6)
	fmt.Println("bulk update result", updateAllResult)

}
