## Gorm-Scope-Utils

gorm-scope-utils is a set of tool kit help you build SQL filter.
You can easily build you scope filter with this tool set. 
Go generics supported. You need to upgrade you Go version to at least 1.18.   




### Get Started

```shell
go get github.com/fenghaojiang/gorm-scope-utils
``` 


### Examples 

```go
func main() {
	dbConns, err := gorm.Open(
		mysql.Open("root:fenghaojiang@tcp(127.0.0.1:4000)/dbname?charset=utf8mb4&parseTime=True&loc=Local"),
		&gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic(err.Error())
	}
	equalValues := []value.ValueEqual[string]{
		{
			Field:        "col1",
			Value:        "value1",
			IncludeEmpty: false,
		},
	}
	rangeValues := []value.ValueRange[int64]{
		{
			Field:        "col2",
			From:         -1,
			To:           2,
			IncludeEmpty: true,
		},
	}

	inValues := []value.ValueIn[string]{
		{
			Field:  "col3",
			Values: []string{"123", "234", "345"},
		},
	}

	var res []any
	err = dbConns.Table("example_table").Select("*").
		Scopes(scope.ScopeEqual[string](equalValues...)).
		Scopes(scope.ScopeRange[int64](rangeValues...)).
		Scopes(scope.ScopeIn[string](inValues...)).
		Scan(&res).Error
	if err != nil {
		fmt.Println(err.Error())
		panic(err.Error())
	}
	for i := range res {
		fmt.Printf("result values: %+v \n", res[i])
	}
}
```

