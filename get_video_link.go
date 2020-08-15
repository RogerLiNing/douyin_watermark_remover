package douyin

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type JsonData struct {
	ItemList []Item `json:"item_list"`
}

type Item struct {
	Video Video `json:"video"`
	ForwardId string `json:"forward_id"`
}

type Video struct {
	Vid string `json:"vid"`
}

func getVideoLink(id string) (string, error) {

	client := &http.Client{}
	// 通过这个接口获取视频信息，其中包括带有水印的链接
	url := "https://www.iesdouyin.com/web/api/v2/aweme/iteminfo/?item_ids=" + id

	request, err := http.NewRequest("GET", url, nil)

	request.Header.Add("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1")
	request.Header.Add("Accept", "*/*")
	request.Header.Add("Accept-Language", "zh-CN,zh;q=0.8,en-US;q=0.5,en;q=0.3")
	request.Header.Add("Connection", "keep-alive")

	response, err := client.Do(request)
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	jsonByteData := []byte(string(body))

	jsonData := JsonData{}
	err = json.Unmarshal(jsonByteData, &jsonData)
	if err != nil {
		fmt.Println(err)
	}

	videoLink := ""


	if len(jsonData.ItemList) > 0 {
		if jsonData.ItemList[0].ForwardId != "0" {
			return jsonData.ItemList[0].ForwardId, nil
		}

		vid := jsonData.ItemList[0].Video.Vid
		// 自行拼接成无水印的链接
		videoLink = "https://aweme.snssdk.com/aweme/v1/play/?video_id=" + vid + "&ratio=720p&line=0"
		return videoLink, nil

	} else {
		return videoLink, nil
	}

}
