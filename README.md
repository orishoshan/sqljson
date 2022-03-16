# sqljson
Wraps a Go type for automatic marshalling and unmarshalling with SQL, so you don't have to implement Scan() and Value().

You have a struct that is marshalled/unmarshalled to JSON data, and you want to store it in a SQL database, and have the struct type automatically marshal/unmarshal when you save/load it, without having to marshal to JSON manually. With generics in Go 1.18, this is now possible! 🥳

# Usage
```go
// You have a struct that you marshal/unmarshal from JSON and you want to automatically unmarshal.
type UserData struct {
	Name string `json:"name"`
}

// Your SQL model, in this example used with GORM.
type Model struct {
	ID       uint                    `gorm:"primarykey"`
	UserData autojson.JSON[UserData] // Use the JSON type with UserData, your data struct, as a type parameter.
}


func main() {
  db, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
  user := UserData{Name: "Beautiful Name"}
  // INSERT INTO `models` (`user_data`) VALUES ("{\"name\":\"user\"}") RETURNING `id`
  db.Save(&Model{UserData: autojson.From(user)})
  
  modelFromDb := Model{}
  db.First(&modelFromDb)
  println(m.UserData.Item.Name) // "Beautiful Name"
}
```
