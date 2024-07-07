package modules

import (
	"encoding/json"
	"fmt"
)

type User struct {
	FullName string `json:"full_name"`
	Age      int64  `json:"age"`
}

func LearningJSONData() {
	UnMarshalJSONToStruct()
	UnMarshalJSONToMap()
	UnMarshalJSONToInterface()
	MarshalStringToObject()
}

func UnMarshalJSONToStruct() {
	var jsonString = `{"full_name": "john doe", "age": 27}`

	var jsonData = []byte(jsonString)

	var student1 User

	var err = json.Unmarshal(jsonData, &student1)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("user :", student1.FullName)
	fmt.Println("age  :", student1.Age)
}

func UnMarshalJSONToMap() {
	var jsonString = `{"full_name": "john doe", "age": 27}`

	var jsonData = []byte(jsonString)

	var student2 map[string]interface{}

	json.Unmarshal(jsonData, &student2)

	fmt.Println("user :", student2["full_name"])
	fmt.Println("age  :", student2["age"])
}

func UnMarshalJSONToInterface() {
	var jsonString = `{"full_name": "john doe", "age": 27}`

	var jsonData = []byte(jsonString)

	var student3 interface{}

	json.Unmarshal(jsonData, &student3)

	var decodedData = student3.(map[string]interface{})

	fmt.Println("user :", decodedData["full_name"])
	fmt.Println("age  :", decodedData["age"])
}

func MarshalStringToObject() {
	var object = []User{
		{
			"john doe",
			27,
		},
		{
			"doe john",
			32,
		},
	}

	var jsonData, err = json.Marshal(object)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var jsonString = string(jsonData)

	fmt.Println(jsonString)
}
