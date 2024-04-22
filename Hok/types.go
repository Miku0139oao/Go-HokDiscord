package Hok

type LoginInfo struct {
	Data struct {
		OpenId        string `json:"openId"`
		AccessToken   string `json:"accessToken"`
		Certification struct {
			Text  string `json:"text"`
			Color string `json:"color"`
		} `json:"certification"`
		Description string        `json:"description"`
		AreaType    int           `json:"areaType"`
		FriendsType int           `json:"friendsType"`
		HeroList    []interface{} `json:"heroList"`
		Img         string        `json:"img"`
		Tips        struct {
			Text     string      `json:"text"`
			CopyData interface{} `json:"copyData"`
		} `json:"tips"`
		Name string `json:"name"`
		Rank struct {
			StarType int    `json:"starType"`
			StarIcon string `json:"starIcon"`
			Score    int    `json:"score"`
		} `json:"rank"`
		Statistics struct {
			Team  int `json:"team"`
			Label int `json:"label"`
		} `json:"statistics"`
	} `json:"data"`
	Status struct {
		Code      int         `json:"code"`
		Msg       string      `json:"msg"`
		Usedtime  float64     `json:"usedtime"`
		Cache     bool        `json:"cache"`
		RequestId string      `json:"requestId"`
		LoginId   int         `json:"loginId"`
		Icp       interface{} `json:"icp"`
	} `json:"status"`
	Hello struct {
		Tips   string `json:"_tips"`
		Afdian string `json:"afdian"`
		Docs   string `json:"docs"`
	} `json:"hello"`
}

/*
allBPRate:89.74 allBanRate:72.65 allBrandRate:24.82 allMvpRate:24.29 allPickRate:17.09 allScore:152 allWinRate:53.66 assistCnt:5.43 behurtPerDeath:26314 change:map[updateType:0 updateValue:0] deadCnt:2.46 equipmentMoney:10378 equipmentMoneyMin:730 evaluateGoldRate:8.96 evaluateSilverRate:15.85 gradient:0 hurtTransRate:1.16 id:519 img:https://game.gtimg.cn/images/yxzj/img201606/heroimg/519/519.jpg joinGamePercent:56.25 killCnt:4.25 loseMvpRate:11.79 name:敖隐 skill:[map[id:80115 img:https://game.gtimg.cn/images/yxzj/img201606/summoner/80115.jpg pickRate:98] map[id:1422 img:https://game.gtimg.cn/images/yxzj/img201606/itemimg/1422.jpg pickRate:29]] tag:map[color:red text:有
更新] totalBehurtCnt:65559 totalHeroHurtCnt:85312 totalHurtCnt:330941 trend:0 type:[3 0 0] updateId:12345 usedtime:14.31 winMvpRate:12.49
*/
type Hero struct {
	AllBPRate      float64 `json:"allBPRate"`      //BP率
	AllBanRate     float64 `json:"allBanRate"`     //Ban率
	AllBrandRate   float64 `json:"allBrandRate"`   //熱度
	AllMvpRate     float64 `json:"allMvpRate"`     //Mvp率
	AllPickRate    float64 `json:"allPickRate"`    //出場率
	AllScore       float64 `json:"allScore"`       //最高戰力
	AllWinRate     float64 `json:"allWinRate"`     //勝率
	AssistCnt      float64 `json:"assistCnt"`      //場均助攻
	BehurtPerDeath float64 `json:"behurtPerDeath"` //場均每死承傷
	Change         struct {
		UpdateType  int `json:"updateType"`
		UpdateValue int `json:"updateValue"`
	} `json:"change"`
	DeadCnt            float64 `json:"deadCnt"`        //場均死亡
	EquipmentMoney     float64 `json:"equipmentMoney"` //場均經濟
	EquipmentMoneyMin  float64 `json:"equipmentMoneyMin"`
	EvaluateGoldRate   float64 `json:"evaluateGoldRate"`   //場均金牌率
	EvaluateSilverRate float64 `json:"evaluateSilverRate"` //場均銀牌率
	Gradient           float64 `json:"gradient"`
	HurtTransRate      float64 `json:"hurtTransRate"` //場均傷害轉化比
	ID                 float64 `json:"id"`
	Img                string  `json:"img"`
	JoinGamePercent    float64 `json:"joinGamePercent"` //場均參團
	KillCnt            float64 `json:"killCnt"`         //場均擊殺
	LoseMvpRate        float64 `json:"loseMvpRate"`     //場均敗方Mvp
	Name               string  `json:"name"`            //英雄名稱
	Skill              []struct {
		ID       float64 `json:"id"`
		Img      string  `json:"img"`
		PickRate float64 `json:"pickRate"`
	} `json:"skill"`
	//Tag []struct {
	//	Color string `json:"color"`
	//	Text  string `json:"text"`
	//} `json:"tag"`
	TotalBehurtCnt   float64   `json:"totalBehurtCnt"`   //場均承傷
	TotalHeroHurtCnt float64   `json:"totalHeroHurtCnt"` //場均英雄傷害
	TotalHurtCnt     float64   `json:"totalHurtCnt"`     //場均總傷害
	Trend            float64   `json:"trend"`            //熱度
	Type             []float64 `json:"type"`
	UpdateID         float64   `json:"updateId"`
	Usedtime         float64   `json:"usedtime"`
	WinMvpRate       float64   `json:"winMvpRate"` //場均勝方Mvp
}

type Ranking struct {
	Data struct {
		Result struct {
			Color struct {
				Bp   int `json:"bp"`
				Ban  int `json:"ban"`
				Pick int `json:"pick"`
				Win  int `json:"win"`
			} `json:"color"`
			Rows []Hero `json:"rows"`
		} `json:"result"`
	} `json:"data"`
	Status struct {
		Code      int         `json:"code"`
		Msg       string      `json:"msg"`
		Usedtime  float64     `json:"usedtime"`
		Cache     bool        `json:"cache"`
		RequestId interface{} `json:"requestId"`
		LoginId   interface{} `json:"loginId"`
	} `json:"status"`
	Hello struct {
		Tips   string `json:"_tips"`
		Afdian string `json:"afdian"`
		Docs   string `json:"docs"`
	} `json:"hello"`
}

type HeroCache struct {
	Hero
}

type MoreHereInfo struct {
	Data struct {
		Type       string `json:"type"`
		CircleInfo struct {
			Text string `json:"text"`
			Tips string `json:"tips"`
			Vote []struct {
				Img  string `json:"img"`
				Text string `json:"text"`
			} `json:"vote"`
		} `json:"circleInfo"`
		HeroInfo struct {
			BanRate  []string `json:"banRate"`
			PickRate []string `json:"pickRate"`
			BpRate   []string `json:"bpRate"`
			WinRate  []string `json:"winRate"`
			Score    []string `json:"score"`
			Behavior []struct {
				Id    string `json:"id"`
				Title string `json:"title"`
				Img   string `json:"img"`
			} `json:"behavior"`
			HighLight []struct {
				Title string `json:"title"`
			} `json:"highLight"`
			Feature struct {
				BeginnerRate float64     `json:"beginnerRate"`
				SupportRate  float64     `json:"supportRate"`
				RedBlue      interface{} `json:"redBlue"`
				Interference interface{} `json:"interference"`
				Blue         struct {
					Blue struct {
						PickRate float64 `json:"pickRate"`
						WinRate  float64 `json:"winRate"`
					} `json:"blue"`
					Red struct {
						PickRate float64 `json:"pickRate"`
						WinRate  float64 `json:"winRate"`
					} `json:"red"`
				} `json:"blue"`
				Red struct {
					Blue struct {
						PickRate float64 `json:"pickRate"`
						WinRate  float64 `json:"winRate"`
					} `json:"blue"`
					Red struct {
						PickRate float64 `json:"pickRate"`
						WinRate  float64 `json:"winRate"`
					} `json:"red"`
				} `json:"red"`
				KillTime []float64 `json:"killTime"`
			} `json:"feature"`
			More []struct {
				Icon   interface{} `json:"icon"`
				Title  string      `json:"title"`
				Label  string      `json:"label"`
				Value  string      `json:"value"`
				IsLink bool        `json:"isLink"`
				To     interface{} `json:"to"`
				Url    string      `json:"url"`
			} `json:"more"`
			New   []interface{} `json:"new"`
			Skill []struct {
				Id       float64 `json:"id"`
				Img      string  `json:"img"`
				PickRate float64 `json:"pickRate"`
			} `json:"skill"`
			Skin              []string    `json:"skin"`
			Type              []float64   `json:"type"`
			DominanceRate     float64     `json:"dominanceRate"`
			EquipmentMoney    float64     `json:"equipmentMoney"`
			EquipmentMoneyMin float64     `json:"equipmentMoneyMin"`
			Id                float64     `json:"id"`
			Img               string      `json:"img"`
			LikeStatus        float64     `json:"likeStatus"`
			Name              string      `json:"name"`
			CName             string      `json:"cName"`
			Gradient          float64     `json:"gradient"`
			Title             string      `json:"title"`
			Label             string      `json:"label"`
			Value             string      `json:"value"`
			To                string      `json:"to"`
			Url               interface{} `json:"url"`
			IsNew             bool        `json:"isNew"`
			IsLink            bool        `json:"isLink"`
			UpdateId          float64     `json:"updateId"`
			WikiId            float64     `json:"wikiId"`
			Usedtime          float64     `json:"usedtime"`
			AdjustmentTime    string      `json:"adjustmentTime"`
			UpdateTime        string      `json:"updateTime"`
			Change            struct {
				TrendType   float64 `json:"trendType"`
				TrendIcon   string  `json:"trendIcon"`
				ScoreIcon   string  `json:"scoreIcon"`
				UpdateType  float64 `json:"updateType"`
				UpdateValue float64 `json:"updateValue"`
			} `json:"change"`
		} `json:"heroInfo"`
	} `json:"data"`
	Status struct {
		Code      int     `json:"code"`
		Msg       string  `json:"msg"`
		Usedtime  float64 `json:"usedtime"`
		Cache     bool    `json:"cache"`
		RequestId string  `json:"requestId"`
		LoginId   string  `json:"loginId"`
	} `json:"status"`
	Hello struct {
		Tips   string `json:"_tips"`
		Afdian string `json:"afdian"`
		Docs   string `json:"docs"`
	} `json:"hello"`
}

type CCToHero struct {
	Name string `json:"name"`
}
