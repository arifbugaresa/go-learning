package main

import (
	"fmt"
	"go-learning/utils/commonFuntion"
)

func main() {
	err := commonFuntion.GetDBConnection()
	if err != nil {
		fmt.Println("Error:", err)
	}

	err = commonFuntion.GetElasticConnection()
	if err != nil {
		fmt.Println("Error:", err)
	}
}
