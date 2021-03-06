package models

import (
	"2019_2_Shtoby_shto/src/customType"
	"2019_2_Shtoby_shto/src/dicts"
)

const boardTableName = "boards"

//easyjson:json
type Board struct {
	dicts.BaseInfo
	Name         string                `json:"name"`
	Users        []string              `json:"users" sql:"-"`
	BoardUsersID customType.StringUUID `json:"-"`
	CardGroups   []CardGroup           `json:"card_groups" sql:"-"`
}

func (b Board) GetTableName() string {
	return boardTableName
}

func (b Board) IsValid() bool {
	return b.Name != ""
}
