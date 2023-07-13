/*
デフォルトは打ち込まれた日の1日分の京都産業大学（緯度：35.07, 軽度：135.75）の気温を表示する.
*/
package sleet

import (
	//"encoding/json"

	"encoding/json"
	//"fmt"
	"io/ioutil"
	"net/http"
	//"strings"
)

type Result struct {
	Day *Daily `json:"daily"`
}

type Daily struct {
	Whetercode []int `json:"weathercode"`
	Time       []string `json:"time"`
}

/* urlをopen-meteoにリクエストしてjsonファイルに書き込み, ウェザーコードと日にちを返す（parse関数も関わりあり） */
func Meteo(url string) (*Result, error) {
	method := "GET"
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	//fmt.Println(q)
	req.URL.RawQuery = q.Encode()

	client := &http.Client{}
	//fmt.Println(client)
	resp, err := client.Do(req)
	//fmt.Println(resp)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(body))
	if err != nil {
		return nil, err
	}

	r, err := parse(body)
	return r, err
}

// jsonファイルにウェザーコードと時間を追加
func parse(body []byte) (*Result, error) {
	result := Result{}
	err := json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	//fmt.Println(result.Day)
	return &result, nil
}
