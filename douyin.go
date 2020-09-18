package douyin

import (
	"log"
)

func WatermarkRemover(url string) (string, error) {
	// 先获取视频的ID
	id, err := getVideoId(url)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	// 如果获取不到就返回
	if len(id) == 0 {
		return "", nil
	}

	// 根据ID获取视频信息
	videoLink, err := getVideoLink(id)

	if err != nil {
		log.Fatal(err)
		return "", err
	}

	if len(videoLink) == 0 {
		return "", nil
	}



	finalLink, err := getFinalVideoLink(videoLink)

	if err != nil {
		log.Fatal(err)
		return "", err
	}

	if len(finalLink) == 0 {
		return "", nil
	}



	return finalLink, nil
}
