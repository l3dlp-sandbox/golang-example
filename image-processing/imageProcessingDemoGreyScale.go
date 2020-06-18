package image_processing

import (
	"fmt"
	"image/color"
	"sync"
)

func greyScale(pixels *[][]color.Color)  {
	ppixels := *pixels
	xLen:=len(ppixels)
	yLen := len(ppixels[0])
	//create new image
	newImage:=make([][]color.Color, xLen)
	for i:=0;i<len(newImage);i++{
		newImage[i] = make([]color.Color,yLen)
	}
	//idea is processing pixels in parallel
	wg := sync.WaitGroup{}
	for x:=0;x<xLen;x++{
		for y:=0;y<yLen; y++{
			wg.Add(1)
			go func(x,y int) {
				pixel :=ppixels[x][y]
				originalColor,ok := color.RGBAModel.Convert(pixel).(color.RGBA)
				if !ok{
					fmt.Println("type conversion went wrong")
				}
				grey := uint8(float64(originalColor.R)*0.21 + float64(originalColor.G)*0.72 + float64(originalColor.B)*0.07)
				col :=color.RGBA{
					grey,
					grey,
					grey,
					originalColor.A,
				}
				newImage[x][y] = col
				wg.Done()
			}(x,y)

		}
	}
	wg.Wait()
	*pixels = newImage
}



