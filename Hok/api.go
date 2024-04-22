package Hok

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
	"time"
)

var (
	BaseURL     = "https://api.91m.top/hero/"
	AccessToken = ""
	OpendId     = ""
	HeroMap     = map[string]Hero{}
	CNameMap    = map[string]string{}
	LastUpadate = time.Now()
	Email       = "YourEmail"
	PassWord    = "YourPassword"
)

func init() {
	ReadCacheFile()
}

func ReadCacheFile() {
	body, err := ioutil.ReadFile("./data/heros.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	//fmt.Println(string(body))
	var Name map[string]string
	if err := json.Unmarshal(body, &Name); err != nil {
		fmt.Println(err)
	}

	for s, s2 := range Name {
		CNameMap[s] = s2
	}
}

func ApiLogin(Email, Password string) {
	req, _ := http.NewRequest(http.MethodPost, BaseURL+"/app?type=loginWebAccount&aid=1&host=91m.top&lang=zh-CN&ref=&url=https%3A%2F%2F91m.top%2Flogin&url_real=https%3A%2F%2F91m.top%2Flogin?"+"openId=&accessToken=&name=&email="+Email+"&password="+Password, nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	var Result LoginInfo
	//fmt.Println(string(body))
	if err := json.Unmarshal(body, &Result); err != nil {
		fmt.Println(err)
	}
	AccessToken = Result.Data.AccessToken
	OpendId = Result.Data.OpenId
	fmt.Println(AccessToken, OpendId)
}

func Request(EndPoint string) []byte {
	req, _ := http.NewRequest(http.MethodGet, BaseURL+EndPoint, nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(resp.StatusCode)
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	return body
	//fmt.Println(string(body))
}

func GetAllheroInfos() {
	//Re := ""
	body := Request("/app?type=getRanking&aid=0&bid=0&cid=0&did=0&id=0&openId=" + OpendId + "&accessToken=" + AccessToken)
	//fmt.Println(BaseURL + "/app?type=getRanking&aid=0&bid=0&cid=0&did=0&id=0&openId=" + OpendId + "&accessToken=" + AccessToken)
	var Result Ranking
	if err := json.Unmarshal(body, &Result); err != nil {
		fmt.Println(err)
	}
	//fmt.Println(Result.Data.Result.Rows)
	for _, data := range Result.Data.Result.Rows {
		//fmt.Println(data.Name)
		HeroMap[data.Name] = Hero{
			AllBPRate:          data.AllBPRate,
			AllBanRate:         data.AllBanRate,
			AllBrandRate:       data.AllBrandRate,
			AllMvpRate:         data.AllMvpRate,
			AllPickRate:        data.AllPickRate,
			AllScore:           data.AllScore,
			AllWinRate:         data.AllWinRate,
			AssistCnt:          data.AssistCnt,
			BehurtPerDeath:     data.BehurtPerDeath,
			DeadCnt:            data.DeadCnt,
			EquipmentMoney:     data.EquipmentMoney,
			EquipmentMoneyMin:  data.EquipmentMoneyMin,
			EvaluateGoldRate:   data.EvaluateGoldRate,
			EvaluateSilverRate: data.EvaluateSilverRate,
			Gradient:           data.Gradient,
			HurtTransRate:      data.HurtTransRate,
			ID:                 data.ID,
			JoinGamePercent:    data.JoinGamePercent,
			KillCnt:            data.KillCnt,
			LoseMvpRate:        data.LoseMvpRate,
			TotalBehurtCnt:     data.TotalBehurtCnt,
			TotalHeroHurtCnt:   data.TotalHurtCnt,
			TotalHurtCnt:       data.TotalHeroHurtCnt,
			Trend:              data.Trend,
			UpdateID:           data.UpdateID,
			WinMvpRate:         data.WinMvpRate,
			//	Tag:                data.Tag,
			Img:    data.Img,
			Change: data.Change,
			Name:   data.Name,
		}
	}
}

func GetFromCname(name string) (Name string, ONames []string) {
	for k, s := range CNameMap {
		sp := strings.Split(s, "|")
		for _, ss := range sp {
			if ss == name {
				//fmt.Println(k)
				sort.Strings(sp)
				return k, sp
			}
		}
	}
	//fmt.Println("Not Found ", name)
	return "", nil
}

func GetFromCNameMap(name string) map[string][]string {
	M := map[string][]string{}
	for k, s := range CNameMap {
		//	fmt.Println(k)
		sp := strings.Split(s, "|")
		for _, s2 := range sp {
			if strings.Contains(s2, strings.ToLower(name)) {
				M[k] = sp
			}
		}
	}
	return M
}
