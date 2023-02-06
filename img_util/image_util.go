package img_util

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"os"

	"github.com/nfnt/resize"
)

func GenImage() {
	m := image.NewRGBA(image.Rect(0, 0, 640, 480))

	draw.Draw(m, m.Bounds(), &image.Uniform{C: color.White}, image.Point{}, draw.Src)

	f, err := os.Create("./result_pic/demo.jpeg")
	if err != nil {
		panic(err)
	}
	err = jpeg.Encode(f, m, nil)
	if err != nil {
		panic(err)
	}
}

// 读取图片信息
func ReadImage() {
	f, err := os.Open("ubuntu.png")
	if err != nil {
		panic(err)
	}
	// decode图片
	m, err := png.Decode(f)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", m.Bounds())     // 图片长宽
	fmt.Printf("%v\n", m.ColorModel()) // 图片颜色模型
	fmt.Printf("%v\n", m.At(100, 100)) // 该像素点的颜色
}

//裁剪图片
func CutImage() {
	f, err := os.Open("ubuntu.png")
	if err != nil {
		panic(err)
	}
	// decode图片
	m, err := png.Decode(f)
	if err != nil {
		panic(err)
	}
	rgba := m.(*image.RGBA)
	subImage := rgba.SubImage(image.Rect(0, 0, 266, 133)).(*image.RGBA)

	// 保存图片
	create, _ := os.Create("./result_pic/new.png")
	err = png.Encode(create, subImage)
	if err != nil {
		panic(err)
	}
}

//压缩图片
func CompressedPicture() {
	file, err := os.Open("duola.jpg")
	defer file.Close()
	if err != nil {
		panic(err)
	}
	// decode jpeg into image.Image
	img, err := jpeg.Decode(file)
	// 处理png
	//img, err := png.Decode(file)

	if err != nil {
		panic(err)
	}
	file.Close()

	// resize to width 1000 using Lanczos resampling
	// and preserve aspect ratio
	m := resize.Resize(1000, 0, img, resize.Lanczos3)

	out, err := os.Create("./result_pic/test_compressed.jpg")
	if err != nil {
		panic(err)
	}
	defer out.Close()

	// write new image to file
	jpeg.Encode(out, m, nil)

}

// 修改图片尺寸
func ResizeImg() {
	file, _ := os.Open("duola.jpg")
	img, _ := jpeg.Decode(file)
	defer file.Close()

	b := img.Bounds()
	imgW := b.Size().X
	imgH := b.Size().Y
	fmt.Println("imgW", imgW, "imgH", imgH)

	newImg := ImageResize(img, imgW/2, imgH/2)

	out, err := os.Create("./result_pic/resize.jpg")
	if err != nil {
		panic(err)
	}
	defer out.Close()
	// write new image to file
	jpeg.Encode(out, newImg, &jpeg.Options{jpeg.DefaultQuality})
}

// 增加水印
func AddWatermark() {
	//加载图片
	file, _ := os.Open("duola.jpg")
	img, _ := jpeg.Decode(file)
	defer file.Close()

	//加载水印
	watermark, err := os.Open("water_test.png")
	if err != nil {
		panic(err)
	}
	defer watermark.Close()
	imgwatermark, err := png.Decode(watermark)
	if err != nil {
		panic(err)
	}

	//img.Bounds().Dx()指的是原图的宽度，img.Bounds().Dy()指的是原图的高度
	offset := image.Pt(img.Bounds().Dx()-imgwatermark.Bounds().Dx()-19, img.Bounds().Dy()-imgwatermark.Bounds().Dy()-19)
	b := img.Bounds()
	m := image.NewNRGBA(b) //按原图生成新图

	//新图写入原图和背景图

	//image.ZP代表Point结构体，目标的源点，即(0,0)
	//draw.Src源图像透过遮罩后，替换掉目标图像
	//draw.Over源图像透过遮罩后，覆盖在目标图像上（类似图层）
	draw.Draw(m, b, img, image.ZP, draw.Src)
	draw.Draw(m, imgwatermark.Bounds().Add(offset), imgwatermark, image.ZP, draw.Over)

	//输出图像
	imgw, _ := os.Create("./result_pic/new-water.jpg")
	jpeg.Encode(imgw, m, &jpeg.Options{100})

}

// 图片大小调整
func ImageResize(src image.Image, w, h int) image.Image {
	return resize.Resize(uint(w), uint(h), src, resize.Lanczos3)
}
