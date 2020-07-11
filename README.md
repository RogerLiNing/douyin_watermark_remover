# 抖音视频水印去除库
本仓库是我在另一个项目中使用的解析库，仅作为库源码使用，所以不提供二进制直接运行。毕竟已经有很多去水印的软件了，也不缺我这个。如果你想要在你的go项目中使用，可以在你的代码中导入该包，并按照以下方式调用。

## 使用方法
```go
import (
	douyin "github.com/RogerLiNing/douyin_watermark_remover"
)

url, _ := WatermarkRemover("https://v.douyin.com/JNhucQF/")

// url 是原始视频，不包含水印。需要注意：URL只能在user-agent 为手机的打开，iPhone、Android等移动设备
// 比如 "User-Agent":"Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1"
// 你要下载视频时，也需要使用移动设备的user-agent

fmt.Println(url)

```