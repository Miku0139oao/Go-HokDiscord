package main

import (
	"context"
	"fmt"
	"github.com/diamondburned/arikawa/v3/api/cmdroute"
	"github.com/diamondburned/arikawa/v3/gateway"
	"github.com/diamondburned/arikawa/v3/state"
	"hokAPi/Hok"
	"hokAPi/commands"
	"hokAPi/commands/handlers"
	"log"
	"time"
)

var (
	Token = "Your Bot Token"
)

func main() {
	Hok.Email = "Your Email"
	Hok.PassWord = "Your Password"
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
