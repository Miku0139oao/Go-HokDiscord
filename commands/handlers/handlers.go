package handlers

import (
	"context"
	"fmt"
	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/api/cmdroute"
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/diamondburned/arikawa/v3/utils/json/option"
	"github.com/liuzl/gocc"
	"hokAPi/Hok"
	"strings"
	"time"
)

func CheckHero(context context.Context, d cmdroute.CommandData) *api.InteractionResponseData {
	Hok.ReadCacheFile()
	Name := d.Options.Find("name").String()
	//fmt.Println(Name)
	if Name == "" {
		return &api.InteractionResponseData{Flags: discord.EphemeralMessage, Content: option.NewNullableString(fmt.Sprintf("請輸入正確英雄名字或暱稱"))}
	}
	t2s, err := gocc.New("t2s")
	if err != nil {
		fmt.Println(err)
	}

	out, err := t2s.Convert(strings.ToLower(Name))
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(out)

	if info, Onames := Hok.GetFromCname(out); info != "" {
		//	fmt.Println(info)
		if Data, embed := Hok.GenerateHeroData(info); Data != "" {
			if Hinfo := Hok.GetHeroInfo(info); Hinfo.ID != 0 {
				//fmt.Println(Hinfo.Img)
				Names := ""
				for i, oname := range Onames {
					s2t, err := gocc.New("s2t")
					if err != nil {
						fmt.Println(err)
					}

					out2, err := s2t.Convert(oname)
					if err != nil {
						fmt.Println(err)
					}
					if i != len(Onames)-1 {
						Names += fmt.Sprintf("%v,", out2)
					} else {
						Names += fmt.Sprintf("%v", out2)
					}
				}
				if embed != nil {
					embed = append(embed, discord.EmbedField{
						Name:   "相關關鍵字(自動生成)",
						Value:  fmt.Sprintf("```%v```", Names),
						Inline: false,
					})
					return &api.InteractionResponseData{
						//Content: option.NewNullableString(fmt.Sprintf("%v", Data)),
						Embeds: &[]discord.Embed{
							{
								Title: fmt.Sprintf("%v (搜尋: %v)\n", Hinfo.Name, Name),
								//Description: fmt.Sprintf("相關關鍵字(自動生成):\n```%v```", Names),
								Timestamp: discord.NowTimestamp(),
								Color:     5763719,
								Thumbnail: &discord.EmbedThumbnail{
									URL:    fmt.Sprintf("https://game.gtimg.cn/images/yxzj/img201606/heroimg/%v/%v.jpg", Hinfo.ID, Hinfo.ID),
									Width:  128,
									Height: 128,
								},
								Fields: embed,
								Footer: &discord.EmbedFooter{
									Text: fmt.Sprintf("資料最後更新: %v 由 %v 請求", Hok.LastUpadate.Format(time.RFC822), d.Event.Sender().Tag()),
									Icon: d.Event.Sender().AvatarURL(),
									//ProxyIcon: "",
								},
							},
						},
					}
				}
			} else {
				return &api.InteractionResponseData{Content: option.NewNullableString(fmt.Sprintf("%v", Data))}
			}
		}
	}

	if Data, embed := Hok.GenerateHeroData(out); Data != "" {
		//fmt.Println(out)
		if Hinfo := Hok.GetHeroInfo(out); Hinfo.Name != "" {
			return &api.InteractionResponseData{
				//Content: option.NewNullableString(fmt.Sprintf("%v", Data)),
				Embeds: &[]discord.Embed{
					{
						Title:     fmt.Sprintf("%v (搜尋: %v)", Hinfo.Name, Name),
						Fields:    embed,
						Timestamp: discord.Timestamp{},
						Color:     5763719,
						Thumbnail: &discord.EmbedThumbnail{
							URL:    fmt.Sprintf("https://game.gtimg.cn/images/yxzj/img201606/heroimg/%v/%v.jpg", Hinfo.ID, Hinfo.ID),
							Width:  128,
							Height: 128,
						},
					},
				},
			}
		}
		return &api.InteractionResponseData{Content: option.NewNullableString(fmt.Sprintf("%v", Data))}
	}

	return &api.InteractionResponseData{Content: option.NewNullableString(fmt.Sprintf("查詢英雄: ```%v``` 失敗, 沒有找到相關英雄或暱稱", Name))}
}
