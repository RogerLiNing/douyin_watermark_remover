# 抖音视频水印去除库
本仓库仅作为库源码使用，可以在你的代码中导入该包，并按照以下方式调用

## 使用方法
```go
import (
	douyin "github.com/RogerLiNing/douyin_watermark_remover"
)
url, _ := WatermarkRemover("https://v.douyin.com/JNhucQF/")

fmt.Println(url)

```