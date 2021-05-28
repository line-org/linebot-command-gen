package main

type Commands struct {
	Cmds map[string]Command
}

type Command struct {
	ID       string       `yaml:"id"`
	IDUpper  string       `yaml:"-"`
	Level    UserLevelStr `yaml:"level"`
	BaseText string       `yaml:"base_text"`
	SubTexts []string     `yaml:"sub_texts"`
	Help     struct {
		Ja string `yaml:"ja"`
		En string `yaml:"en"`
	} `yaml:"help"`
	Genre CmdGenre `yaml:"genre"`
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

func (u UserLevelStr) ToUserLevel() UserLevel {
	switch u {
	case PermanentBlack:
		return PermanentBlackInt
	case GlobalBlack:
		return GlobalBlackInt
	case LocalBlack:
		return LocalBlackInt
	case None:
		return NoneInt
	case LocalWhite:
		return LocalWhiteInt
	case GlobalWhite:
		return GlobalWhiteInt
	case LocalBot:
		return LocalBotInt
	case GlobalBot:
		return GlobalBotInt
	case LocalAdmin:
		return LocalAdminInt
	case SimpleAdmin:
		return SimpleAdminInt
	case NormalAdmin:
		return NormalAdminInt
	case SupremeAdmin:
		return SupremeAdminInt
	case Buyer:
		return BuyerInt
	case Developer:
		return DeveloperInt
	}
	return NoneInt
}

type UserLevel int

const (
	PermanentBlackInt UserLevel = -3
	GlobalBlackInt    UserLevel = -2
	LocalBlackInt     UserLevel = -1

	NoneInt UserLevel = 0

	LocalWhiteInt  UserLevel = 1
	GlobalWhiteInt UserLevel = 2

	LocalBotInt  UserLevel = 5
	GlobalBotInt UserLevel = 6

	LocalAdminInt   UserLevel = 8
	SimpleAdminInt  UserLevel = 9
	NormalAdminInt  UserLevel = 10
	SupremeAdminInt UserLevel = 11

	BuyerInt     UserLevel = 20
	DeveloperInt UserLevel = 21
)
