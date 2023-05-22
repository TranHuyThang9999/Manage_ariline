package resources

import (
	entities "btl/core/entity"
	"encoding/json"
)

func toJson(val []byte) entities.User {
	user := entities.User{}
	err := json.Unmarshal(val, &user)
	if err != nil {
		panic(err)
	}
	return user
}
