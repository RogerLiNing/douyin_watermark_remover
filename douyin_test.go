package douyin

import (
	"fmt"
	"strings"
	"testing"
)

func TestGetAvailableDouyinLink(t *testing.T) {
	t.Log("测试正常的抖音视频链接")
	url := "https://v.douyin.com/JNhucQF/" // https://v.douyin.com/JMyLw4U/

	videoLink, err := WatermarkRemover(url)

	if err != nil {
		t.Fail()
		t.Log("pass")
	}

	fmt.Println(videoLink)

	if strings.Contains(videoLink, ".ixigua.com") == false {
		t.Fail()
	} else {
		t.Log("link：", videoLink)
	}
}

func TestGetUnAvailableDouyinLink(t *testing.T) {
	t.Log("测试视频不存在的抖音链接")
	url := "https://v.douyin.com/JNhu000"
	videoLink, err := WatermarkRemover(url)

	if err != nil {
		t.Fail()
	}

	if strings.Contains(videoLink, ".ixigua.com") {
		t.Fail()
		t.Log("测试通过")
	}
}
