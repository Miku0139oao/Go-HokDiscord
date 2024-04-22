package commands

import (
	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/discord"
)

var Commands = []api.CreateCommandData{
	{
		Name:        "check",
		Description: "查詢王者榮耀英雄資料",
		Options: []discord.CommandOption{
			&discord.StringOption{
				OptionName:   "name",
				Description:  "英雄名稱",
				Autocomplete: true,
				Required:     true,
			},
		},
	},
}
