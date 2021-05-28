package main

import (
	"github.com/iancoleman/strcase"
	"github.com/line-org/protection-bot/users"
	"strings"
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

func (c *Command) ToSubTextStr() string {
	if c.SubTexts == nil || len(c.SubTexts) == 0 {
		return "[]string{}"
	}
	return "[]string{`" + strings.Join(c.SubTexts, "`,`") + "`}"
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

var (
	UserLevelStrC = map[UserLevelStr]users.UserLevel{
		PermanentBlack: users.PermanentBlack,
		GlobalBlack:    users.GlobalBlack,
		LocalBlack:     users.LocalBlack,
		None:           users.None,
		LocalWhite:     users.LocalWhite,
		GlobalWhite:    users.GlobalWhite,
		LocalBot:       users.LocalBot,
		GlobalBot:      users.GlobalBot,
		LocalAdmin:     users.LocalAdmin,
		SimpleAdmin:    users.SimpleAdmin,
		NormalAdmin:    users.NormalAdmin,
		SupremeAdmin:   users.SupremeAdmin,
		Buyer:          users.Buyer,
		Developer:      users.Developer,
	}
)

func (u UserLevelStr) ToUserLevel() users.UserLevel {
	return UserLevelStrC[u]
}
