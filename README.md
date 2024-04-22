# Go-HokDiscord

# 資料來源: 蘇蘇的榮耀助手 https://91m.top/

# 這是一個 Discord Bot，用來查詢王者榮耀的資訊

# 使用方法
- 首次運行: go run main.go 
- 到config.json 裡面填入你在上述網站中註冊的帳號和密碼 以及Discord Bot Token
- 如未有帳號, 請到 https://91m.top/login 註冊
- 開始使用 Discord Bot

# 指令
- /check: {name} 輸入英雄名字或暱稱，查詢英雄資訊
* 例如: /check 阿轲

# 關鍵字
- 存於/data/heros.json中,可以根據自身需求並按照對應格式新增英雄資訊
- 例如:
```json
{
  "云缨": "云嘤|yy"
}
```
- 這樣就可以用云缨或云嘤或yy查詢英雄資訊
```json
{
  "云缨": "云嘤|yy|你老婆"
}
```
- 關鍵字因設計問題只限輸入簡體中文(日後會改善),但搜尋時可以用繁體中文或英文