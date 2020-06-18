package image_processing

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	_"image/png"
	"math"
	"os"
)

func ImageProcessingImplementation(){
	img,err:=openImage("/home/damitha/go/src/test/test/image-processing/portraitofaladyonfire_02.jpg")
	if err!=nil{
		return
	}
	size:= img.Bounds().Size()
	var pixels [][]color.Color
	//put pixels into two three two dimensional array
	for i:=0; i<size.X;i++{
		var y []color.Color
		for j:=0; j<size.Y;j++{
			y = append(y,img.At(i,j))
		}
		pixels = append(pixels,y)
	}
	//pixel manipulation 1) upside down image
	//This can be achieved just by reverting y array
		upsideDown(pixels)
	//pixel manipulation 2) rotate image from 60 angle
	//	radius :=int(math.Sqrt(float64(size.X*size.X) +float64(size.Y*size.Y)))
	//	rotateByAngle(math.Pi/10,&pixels,radius)
	//pixel manipulation 3 rotate by three shear matrix
		//rotateByAngleWithoutAliasing(math.Pi/10,&pixels,radius)
	//creating image back again
	rect := image.Rect(0,0,len(pixels),len(pixels[0]))
	nImg := image.NewRGBA(rect)

	for x:=0; x<len(pixels);x++{
		for y:=0; y<len(pixels[0]);y++ {
			q:=pixels[x]
			if q==nil{
				continue
			}
			p := pixels[x][y]
			if p==nil{
				continue
			}
			original,ok := color.RGBAModel.Convert(p).(color.RGBA)
			if ok{
				nImg.Set(x,y,original)
			}
		}
	}

	fg,err:= os.Create("/home/damitha/go/src/test/test/image-processing/upside_down.jpg")
	if err!=nil{
		fmt.Println("Creating file:",err)
	}
	defer  fg.Close()
	err = jpeg.Encode(fg, nImg,nil)
	if err !=nil{
		fmt.Println("Encoding error",err)
	}
}

func openImage(path string)(image.Image,error){
	f, err := os.Open("/home/damitha/go/src/test/test/image-processing/portraitofaladyonfire_02.jpg")
	if err !=nil{
		fmt.Println(err)
		return nil,err
	}
	fi,_:=f.Stat()
	fmt.Println(fi.Name())
	//defer f.Close()sss
	img,format,err:=image.Decode(f)
	if err!=nil{
		fmt.Println("Decoding error:",err.Error())
		return nil,err
	}
	if format != "jpeg"{
		fmt.Println("image format is not jpeg")
		return nil,errors.New("")
	}
	return img,nil
}



func rotateByAngle(angle float64, pixels *[][]color.Color, radius int){
	ppixels := *pixels
	cos:=math.Cos(angle)
	sin:=math.Sin(angle)

	var newImage  [][]color.Color

	newImage=make([][]color.Color, 2*radius+1)
	for i:=0;i<len(newImage);i++{
		newImage[i] = make([]color.Color,2*radius+1)
	}

	fmt.Println("x:",len(newImage),"y:",len(newImage[0]),"radius:",radius)

	//rotating image pixels
	//In this approach xNew and yNew values are rounded there for it's possible to be assigned
	//different original pixels values to one point several times which is resulting black dots in the
	//rotated image this problem is called aliasing
	for i:=0; i<len(ppixels);i++{
		for j:=0; j<len(ppixels[i]);j++{
			xNew := int(float64(i)*cos - float64(j)*sin) +radius
			yNew := int(float64(i)*sin+float64(j)*cos) +radius

			newImage[xNew][yNew] = ppixels[i][j]
		}
	}
	//
	*pixels = newImage
}


