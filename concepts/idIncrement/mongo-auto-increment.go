package main

import (
"fmt"
"sync"

"gopkg.in/mgo.v2"
"gopkg.in/mgo.v2/bson"
)

type idGenerator struct {
db  *mgo.Database
N   int `bson:"n"`
Key string  // Don't forget to add an unique index to this field.
}

func newIDGenerator(db *mgo.Database) *idGenerator {
return &idGenerator{
db: db,
}
}

// Generate a auto increment version ID for the given key
func (idg *idGenerator) Next(key string) (int, error) {
change := mgo.Change{
Update:    bson.M{"$inc": bson.M{"n": 1}},
ReturnNew: true,
}
r := &idGenerator{}
_, err := idg.db.C("idgen").Find(bson.M{"key": key}).Apply(change, &r)

if err == mgo.ErrNotFound {
err := idg.db.C("idgen").Insert(bson.M{"key": key, "n": 1})
if err != nil {
return -1, err
}
return 1, nil
} else if err != nil {
return -1, err
}

return r.N, nil
}

func main() {
mongo, err := mgo.Dial("mongodb://127.0.0.1:27017")
if err != nil {
panic(err)
}

db := mongo.DB("test")
idg := newIDGenerator(db)

wg := new(sync.WaitGroup)
wg.Add(100)
for i := 0; i < 100; i++ {
go func() {
n, err := idg.Next("my-document")
if err != nil {
panic(err)
}
fmt.Println(n)
wg.Done()
}()
}

wg.Wait()
}
