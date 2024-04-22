package main

import (
	"context"
	"fmt"
	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/api/cmdroute"
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/diamondburned/arikawa/v3/gateway"
	"github.com/diamondburned/arikawa/v3/state"
	"github.com/liuzl/gocc"
	"hokAPi/Hok"
	"hokAPi/commands"
	"hokAPi/commands/handlers"
	"hokAPi/config"
	"log"
	"strings"
	"time"
)

var (
	Token = "Your Bot Token"
)

func main() {
	Config := config.ReadConfig()
	//	fmt.Println(Config)
	if Config.Email != "" && Config.Password != "" && Config.Token != "" {
		/*
		 Login to get AccessToken
		 如果你沒有帳號, 請到 https://91m.top/login 註冊
		 ## 登入請使用Email和密碼
		*/
		//Examples:
		Hok.Email = Config.Email
		Hok.PassWord = Config.Password
		Token = Config.Token
		Md5 := Hok.Md5PassWord(Hok.PassWord)
		Hok.ApiLogin(Hok.Email, Md5)
		Hok.GetAllheroInfos()
	}

	s := state.New("Bot " + Token)
	r := cmdroute.NewRouter()
	r.Use(cmdroute.Deferrable(s, cmdroute.DeferOpts{
		Timeout: 60 * time.Second,
	}))
	s.AddInteractionHandler(r)
	s.AddIntents(gateway.IntentGuilds | gateway.IntentGuildMembers | gateway.IntentDirectMessages | gateway.IntentGuildMessages)
	s.AddHandler(func(c *gateway.ReadyEvent) {
		fmt.Println(c.User.Username + "#" + c.User.Discriminator)
		if err := overwriteCommands(s); err != nil {
			fmt.Println(err)
		}
	})
	s.AddHandler(func(e *gateway.InteractionCreateEvent) {
		var resp api.InteractionResponse
		switch d := e.Data.(type) {
		case *discord.CommandInteraction:
		case *discord.AutocompleteInteraction:
			allChoices := api.AutocompleteStringChoices{}
			switch d.Name {
			case "check":
				o := strings.ReplaceAll(d.Options.Find("name").String(), `"`, "")
				//	fmt.Println(o)
				if Name, _ := Hok.GetFromCname(o); Name != "" {
					if len(allChoices) <= 25 {
						s2t, err := gocc.New("s2t")
						if err != nil {
							fmt.Println(err)
						}

						out, err := s2t.Convert(Name)
						if err != nil {
							fmt.Println(err)
						}
						allChoices = append(allChoices, discord.StringChoice{
							Name:  fmt.Sprintf("搜尋:%v", out),
							Value: fmt.Sprintf("%v", Name),
						})
					}
					resp = api.InteractionResponse{
						Type: api.AutocompleteResult,
						Data: &api.InteractionResponseData{
							Choices: &allChoices,
						},
					}

					if err := s.RespondInteraction(e.ID, e.Token, resp); err != nil {
						log.Println("failed to send interaction callback:", err)
					}
				}
			}
		}
	})
	r.AddFunc("check", handlers.CheckHero)
	if err := s.Open(context.Background()); err != nil {
		log.Fatalln("Failed to Open:", err)
	}

	defer s.Close()

	select {}
}

func overwriteCommands(s *state.State) error {
	return cmdroute.OverwriteCommands(s, commands.Commands)
}
