# goutil
对 [mgo]()关于MongoDB的一些基础操作做一下封装，便于直接调用  

* 插入一个或多个`document`

```
func Insert(db, collection string, docs ...interface{}) error
```

* 查询一个满足条件的`document`  

**query**查询条件 `bson.M{"_id":"id"}`
**selector**相当于MongoDB中的projection,查询结果中是否显示某个字段 `bson.M{"_id":0}`

```
func FindOne(db, collection string, query, selector, result interface{}) error
```

* 查询满足条件的所有结果  

```
func FindAll(db, collection string, query, selector, result interface{}) error
```

* 对查询结果分页处理

```
func FindPage(db, collection string, page, limit int, query, selector, result interface{}) error
```

* 查询满足条件的 `cursor`

```
func FindIter(db, collection string, query interface{}) *mgo.Iter
```

* 更新满足条件的一个`document`

```
func Update(db, collection string, selector, update interface{}) error
```

* 更新满足条件的所有的`docuement`

```
func UpdateAll(db, collection string, selector, update interface{}) error
```

* 更新，如果不存在就插入一个新的`document`

```
func Upsert(db, collection string, selector, update interface{}) error
```

* 删除满足条件的一个 `document`

```
func Remove(db, collection string, selector interface{}) error
```

* 删除满足条件的所有的`document`

```
func RemoveAll(db, collection string, selector interface{}) error
```

* 批量的插入

```
func BulkInsert(db, collection string, docs ...interface{}) (*mgo.BulkResult, error)
```

* 批量更新

```
func BulkUpdate(db, collection string, pairs ...interface{}) (*mgo.BulkResult, error)
```

* 批量更新所有

```
func BulkUpdateAll(db, collection string, pairs ...interface{}) (*mgo.BulkResult, error)
```

* 批量删除

```
func BulkRemove(db, collection string, selector ...interface{}) (*mgo.BulkResult, error)
```

* 批量删除所有  

```
func BulkRemoveAll(db, collection string, selector ...interface{}) (*mgo.BulkResult, error)
```

* 聚合操作查找所有的

```
func PipeAll(db, collection string, pipeline, result interface{}, allowDiskUse bool) error
```

* 聚合操作查找一个

```
func PipeOne(db, collection string, pipeline, result interface{}, allowDiskUse bool) error
```

* 聚合操作查找`cursor`

```
func PipeIter(db, collection string, pipeline interface{}, allowDiskUse bool) *mgo.Iter
```
