package img_util

import (
	"testing"
)

type People struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestFileCache(t *testing.T) {
	//生成图片
	GenImage()
	//读取图片信息
	ReadImage()
	//裁剪图片
	CutImage()
	//压缩图片
	CompressedPicture()
	//修改图片尺寸
	ResizeImg()
	//增加水印
	AddWatermark()
}
