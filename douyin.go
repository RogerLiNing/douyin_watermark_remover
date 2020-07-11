package douyin

import "log"

func WatermarkRemover(url string) (string, error) {
	id, err := getVideoId(url)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	if len(id) == 0 {
		return "", nil
	}

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
