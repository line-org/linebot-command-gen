package main

import (
	"github.com/iancoleman/strcase"
	"github.com/line-org/protection-bot/users"
)

type Commands struct {
	Cmds map[string]Command
}

type Command struct {
	ID       string       `yaml:"id"`
	Level    UserLevelStr `yaml:"level"`
	BaseText string       `yaml:"base_text"`
	SubTexts []string     `yaml:"sub_texts"`
	Help     struct {
		Ja string `yaml:"ja"`
		En string `yaml:"en"`
	} `yaml:"help"`
	Genre CmdGenre `yaml:"genre"`
}

func (c *Command) GetUpperId() string {
	return strcase.ToCamel(c.ID)
}

type CmdGenre string

const (
	ManageCmd CmdGenre = "manage"
)

type UserLevelStr string

const (
	PermanentBlack UserLevelStr = "PermanentBlack"
	GlobalBlack    UserLevelStr = "GlobalBlack"
	LocalBlack     UserLevelStr = "LocalBlack"
	None           UserLevelStr = "None"
	LocalWhite     UserLevelStr = "LocalWhite"
	GlobalWhite    UserLevelStr = "GlobalWhite"
	LocalBot       UserLevelStr = "LocalBot"
	GlobalBot      UserLevelStr = "GlobalBot"
	LocalAdmin     UserLevelStr = "LocalAdmin"
	SimpleAdmin    UserLevelStr = "SimpleAdmin"
	NormalAdmin    UserLevelStr = "NormalAdmin"
	SupremeAdmin   UserLevelStr = "SupremeAdmin"
	Buyer          UserLevelStr = "Buyer"
	Developer      UserLevelStr = "Developer"
)

func (u UserLevelStr) ToUserLevel() users.UserLevel {
	switch u {
	case PermanentBlack:
		return users.PermanentBlack
	case GlobalBlack:
		return users.GlobalBlack
	case LocalBlack:
		return users.LocalBlack
	case None:
		return users.None
	case LocalWhite:
		return users.LocalWhite
	case GlobalWhite:
		return users.GlobalWhite
	case LocalBot:
		return users.LocalBot
	case GlobalBot:
		return users.GlobalBot
	case LocalAdmin:
		return users.LocalAdmin
	case SimpleAdmin:
		return users.SimpleAdmin
	case NormalAdmin:
		return users.NormalAdmin
	case SupremeAdmin:
		return users.SupremeAdmin
	case Buyer:
		return users.Buyer
	case Developer:
		return users.Developer
	}
	return users.None
}
