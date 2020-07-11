package douyin

import (
	"strings"
	"testing"
)

func TestGetAvailableDouyinLink(t *testing.T) {
	t.Log("测试正常的抖音视频链接")
	url := "https://v.douyin.com/JNhucQF/"

	videoLink, err := WatermarkRemover(url)

	if err != nil {
		t.Fail()
		t.Log("测试通过")
	}

	if strings.Contains(videoLink, ".ixigua.com") == false {
		t.Fail()
		t.Log("测试失败")
	} else {
		t.Log("视频链接为：", videoLink)

	}
}

func TestGetUnAvailableDouyinLink(t *testing.T) {
	t.Log("测试视频不存在的抖音链接")
	url := "https://v.douyin.com/JNSy3pH/"
	videoLink, err := WatermarkRemover(url)

	if err != nil {
		t.Fail()
	}

	if strings.Contains(videoLink, ".ixigua.com") {
		t.Fail()
		t.Log("测试通过")
	}
}
