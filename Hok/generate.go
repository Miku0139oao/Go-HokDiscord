package Hok

import (
	"fmt"
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/liuzl/gocc"
)

func GenerateHeroData(key string) (Result string, E []discord.EmbedField) {
	s2t, err := gocc.New("s2t")
	if err != nil {
		fmt.Println(err)
	}

	out, err := s2t.Convert(key)
	if err != nil {
		fmt.Println(err)
	}
	if data, ok := HeroMap[key]; ok {
		Result = fmt.Sprintf("```%v:\n出場率: %v%% | Ban: %v%% | 勝率:%v%% | 禁選(B+P): %v%% \nMVP: %v%% | 勝方MVP: %v%% | 敗方MVP: %v%%\n金牌: %v%% | 銀牌: %v%%\n分均經濟: %v | 場均傷害轉化比: %v | 場均參團: %v%%\n場均死亡: %v | 場均承傷: %v | 場均每死承傷: %v```", out, data.AllPickRate, data.AllBanRate, data.AllWinRate, data.AllBPRate, data.AllMvpRate, data.WinMvpRate, data.LoseMvpRate, data.EvaluateGoldRate, data.EvaluateSilverRate, data.EquipmentMoneyMin, data.HurtTransRate, data.JoinGamePercent, data.DeadCnt, data.TotalBehurtCnt, data.BehurtPerDeath)
		E = []discord.EmbedField{
			{
				Name:   "出場率",
				Value:  fmt.Sprintf("%v%%", data.AllPickRate),
				Inline: true,
			},
			{
				Name:   "Ban率",
				Value:  fmt.Sprintf("%v%%", data.AllBanRate),
				Inline: true,
			},
			{
				Name:   "禁選(B+P)",
				Value:  fmt.Sprintf("%v%%", data.AllBPRate),
				Inline: true,
			},
			{
				Name:   "勝率",
				Value:  fmt.Sprintf("%v%%", data.AllWinRate),
				Inline: true,
			},
			{
				Name:   "MVP率",
				Value:  fmt.Sprintf("%v%%", data.AllMvpRate),
				Inline: true,
			},
			{
				Name:   "勝方MVP率",
				Value:  fmt.Sprintf("%v%%", data.WinMvpRate),
				Inline: true,
			},
			{
				Name:   "敗方MVP率",
				Value:  fmt.Sprintf("%v%%", data.LoseMvpRate),
				Inline: true,
			},
			{
				Name:   "金牌率",
				Value:  fmt.Sprintf("%v%%", data.EvaluateGoldRate),
				Inline: true,
			},
			{
				Name:   "銀牌率",
				Value:  fmt.Sprintf("%v%%", data.EvaluateSilverRate),
				Inline: true,
			},
			{
				Name:   "分均經濟",
				Value:  fmt.Sprintf("%v", data.EquipmentMoneyMin),
				Inline: true,
			},
			{
				Name:   "場均傷害轉化比",
				Value:  fmt.Sprintf("%v", data.HurtTransRate),
				Inline: true,
			},
			{
				Name:   "場均參團",
				Value:  fmt.Sprintf("%v%%", data.JoinGamePercent),
				Inline: true,
			},
			{
				Name:   "場均死亡",
				Value:  fmt.Sprintf("%v", data.DeadCnt),
				Inline: true,
			},
			{
				Name:   "場均承傷",
				Value:  fmt.Sprintf("%v", data.TotalBehurtCnt),
				Inline: true,
			},
			{
				Name:   "場均每死承傷",
				Value:  fmt.Sprintf("%v", data.BehurtPerDeath),
				Inline: true,
			},
		}
	} else {
		return "", []discord.EmbedField{}
	}

	return Result, E
}

func GetHeroInfo(key string) Hero {
	t2s, err := gocc.New("t2s")
	if err != nil {
		fmt.Println(err)
	}

	out, err := t2s.Convert(key)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(out)
	if data, ok := HeroMap[out]; ok {
		//		fmt.Println(data)
		//fmt.Println(data.Name)
		s2t, err := gocc.New("s2t")
		if err != nil {
			fmt.Println(err)
		}

		out, err := s2t.Convert(data.Name)
		if err != nil {
			fmt.Println(err)
		}
		data.Name = out
		return data
	} else {
		return Hero{}
	}
}
