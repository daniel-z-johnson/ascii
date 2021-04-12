package ascii

import (
	"image"
	"image/color"
	_ "image/jpeg"
	"image/png"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func ImageToGreyScales(fileName string) ([]string, error) {
	baseName := strings.TrimSuffix(path.Base(fileName), filepath.Ext(fileName))
	dirName := path.Dir(fileName)
	f1, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer f1.Close()
	imageF1, _, err := image.Decode(f1)
	if err != nil {
		return nil, err
	}
	rect := imageF1.Bounds()
	minPoint := rect.Min
	maxPoint := rect.Max
	lumin := image.NewRGBA(rect)
	mean := image.NewRGBA(rect)
	luma := image.NewRGBA(rect)
	luster := image.NewRGBA(rect)
	for x := minPoint.X; x < maxPoint.X; x++ {
		for y := minPoint.Y; y < maxPoint.Y; y++ {
			c := imageF1.At(x, y)
			lumin.Set(x,y, Lumin(c))
			mean.Set(x, y, Mean(c))
			luma.Set(x, y, Luma(c))
			luster.Set(x, y, Luster(c))
		}
	}
	meanFile := dirName + "/" + baseName + "-mean.png"
	luminFile := dirName + "/" + baseName + "-lumin.png"
	lumaFile := dirName + "/" + baseName + "-luma.png"
	lusterFile := dirName + "/" + baseName + "-luster.png"
	fMean, err := os.Create(meanFile)
	if err != nil {
		return nil, err
	}
	fLumin, err := os.Create(luminFile)
	if err != nil {
		return nil, err
	}
	fLuma, err := os.Create(lumaFile)
	if err != nil {
		return nil, err
	}
	fLuster, err := os.Create(lusterFile)
	if err != nil {
		return nil, err
	}
	png.Encode(fMean, mean)
	png.Encode(fLumin, lumin)
	png.Encode(fLuma, luma)
	png.Encode(fLuster, luster)
	return []string{luminFile, meanFile, lumaFile, lusterFile}, nil
}


func Lumin(c color.Color) color.Color {
	r, g, b, _ := c.RGBA()
	lumin := 0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b)
	return bigGrey{Y: uint32(lumin)}
}

func Mean(c color.Color) color.Color {
	r, g, b, _ := c.RGBA()
	mean := (r+g+b)/3
	return bigGrey{Y:uint32(mean)}
}

func Luma(c color.Color) color.Color {
	r, g, b, _ := c.RGBA()
	luma := 0.2126*float64(r) + 0.7152*float64(g) + 0.0722*float64(b)
	return bigGrey{Y: uint32(luma)}
}

func Luster(c color.Color) color.Color {
	r, g, b, _ := c.RGBA()
	max, min := MaxMin([]uint32{r, g, b})
	luster := (max + min)/2
	return bigGrey{Y: uint32(luster)}
}

func Max(a []uint32) uint32 {
	answer := uint32(0)
	for _, v := range a {
		if v > answer {
			answer = v
		}
	}
	return answer
}

func MaxMin(a []uint32) (uint32, uint32) {
	return Max(a), Min(a)
}

func Min(a []uint32) uint32 {
	answer := uint32(0xffffffff)
	for _, v := range a {
		if v < answer {
			answer = v
		}
	}
	return answer
}

type bigGrey struct {
	Y uint32
}

func (g bigGrey) RGBA() (uint32, uint32, uint32, uint32) {
	return g.Y, g.Y, g.Y, 0xffffffff
}