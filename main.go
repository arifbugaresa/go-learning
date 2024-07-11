package main

import (
	"fmt"
	"go-learning/model"
	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	var user1 = &model.User{
		Id:       "u001",
		Name:     "Sylvana Windrunner",
		Password: "Password123",
		Gender:   model.UserGender_FEMALE,
	}

	var userList = &model.UserList{
		List: []*model.User{
			user1,
		},
	}

	var garage1 = &model.Garage{
		Id:   "g001",
		Name: "Kalimdor",
		Coordinate: &model.GarageCoordinate{
			Latitude:  23.2222222,
			Longitude: 24.33333,
		},
	}

	var garageList = &model.GarageList{
		List: []*model.Garage{
			garage1,
		},
	}

	var garageListByUser = &model.GarageListByUser{
		List: map[string]*model.GarageList{
			user1.Id: garageList,
		},
	}

	fmt.Println(user1)
	fmt.Println(user1.String())

	fmt.Println(garageListByUser)
	fmt.Println(userList)

	jsonb, _ := protojson.Marshal(garageList)

	fmt.Println(string(jsonb))
}
