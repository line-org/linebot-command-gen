package template

import (
	"github.com/line-org/lineall/lineapp/service/line"
	"github.com/line-org/protection-bot/bot"
	"github.com/line-org/protection-bot/commands"
	"github.com/line-org/protection-bot/users"
)

type CommandHandler interface {
	Speed(bot *bot.ProtectionBot, msg *line.Message, cmd *commands.ParsedCmd) error
}
type CommandData struct {
	Cmds map[string]*CommandInfo
}
type CommandInfo struct {
	ID       string
	Level    users.UserLevel
	BaseText string
	SubTexts map[string]struct{}
	Help     struct {
		Ja string
		En string
	}
	Genre    commands.CmdGenre
	Function func(bot *bot.ProtectionBot, msg *line.Message, cmd *commands.ParsedCmd) error
}

func NewCommandDatas(handler CommandHandler) *CommandData {
	por := &CommandData{}
	por.Cmds["speed"] = &CommandInfo{
		ID:       "speed",
		Level:    9,
		BaseText: "speed",
		SubTexts: nil,
		Help: struct {
			Ja string
			En string
		}{},
		Genre:    "",
		Function: handler.Speed,
	}
	return por
}
