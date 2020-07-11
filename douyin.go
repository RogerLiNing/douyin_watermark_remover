package douyin

import "log"

func WatermarkRemover(url string)  (string, error) {
	id, err := getVideoId(url)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	videoLink, err := getVideoLink(id)

	if err != nil {
		log.Fatal(err)
		return "", err
	}

	finalLink, err := getFinalVideoLink(videoLink)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	return finalLink, nil
}

