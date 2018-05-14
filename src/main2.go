/**
 * Created by chaolinding on 2018/5/11.
 */

package main


import (
	bimg "gopkg.in/h2non/bimg.v1"
	"errors"
	"math"
	"fmt"
)

/*
裁剪
 */
func CropWithPath(filePath, savePath string, width, height, quality int) error {

	if len(filePath) == 0 || len(savePath) == 0 {
		return errors.New("文件路径参数有误")
	}
	if width <= 10 || height <= 10 {
		return errors.New("裁剪宽高参数有误")
	}
	if quality < 0 {
		return errors.New("quality参数有误")
	}

	buffer, err := bimg.Read(filePath)
	if err != nil {
		return err
	}

	metaData, err := bimg.Metadata(buffer)
	if err != nil {
		return err
	}

	var options bimg.Options
	if metaData.Alpha {
		options = bimg.Options{
			Height:     height,
			Width:      width,
			Background: bimg.Color{255, 255, 255},
			Crop:       true,
			Quality:    quality,
			Gravity:    bimg.GravityCentre,
		}
	} else {
		options = bimg.Options{
			Height:  height,
			Width:   width,
			Crop:    true,
			Quality: quality,
			Gravity: bimg.GravityCentre,
		}
	}

	newImage, err := bimg.NewImage(buffer).Process(options)
	if err != nil {
		return err
	}

	return bimg.Write(savePath, newImage)

}

/*
变换大小
 */
func ResizeWithPath(filePath, savePath string, width, height, quality int) error {

	if len(filePath) == 0 || len(savePath) == 0 {
		return errors.New("文件路径参数有误")
	}
	if width < 0 || height < 0 || (width == 0 && height <= 10) || (height == 0 && width <= 10) {
		return errors.New("裁剪宽高参数有误")
	}
	if quality < 0 {
		return errors.New("quality参数有误")
	}

	buffer, err := bimg.Read(filePath)
	if err != nil {
		return err
	}

	metaData, err := bimg.Metadata(buffer)
	if err != nil {
		return err
	}

	if (width > 10 && height == 0) || (width == 0 && height > 10) {
		if metaData.Size.Width <= 10 || metaData.Size.Height <= 10 {
			return errors.New("图片meta宽高有误")
		}
		if height == 0 {
			height = int(math.Floor(float64(width) * float64(metaData.Size.Height) / float64(metaData.Size.Width)))
		} else if width == 0 {
			width = int(math.Floor(float64(metaData.Size.Width) / float64(width) * float64(metaData.Size.Height)))
		}
	}

	var options bimg.Options
	if metaData.Alpha {
		options = bimg.Options{
			Height:     height,
			Width:      width,
			Background: bimg.Color{255, 255, 255},
			Quality:    quality,
		}
	} else {
		options = bimg.Options{
			Height:  height,
			Width:   width,
			Quality: quality,
		}
	}

	newImage, err := bimg.NewImage(buffer).Process(options)
	if err != nil {
		return err
	}

	return bimg.Write(savePath, newImage)
}

func FixImageWithPath(filePath, savePath string, width, height, quality int) error {
	if width <= 10 || height <= 10 {
		return errors.New("宽高参数有误")
	}

	buffer, err := bimg.Read(filePath)
	if err != nil {
		return err
	}

	options := bimg.Options{
		Height:     height,
		Width:      width,
		Background: bimg.Color{255, 255, 255},
		Quality:    quality,
		Embed:      true,
	}

	newImage, err := bimg.NewImage(buffer).Process(options)
	if err != nil {
		return err
	}

	return bimg.Write(savePath, newImage)

}


func main(){
	err := CropWithPath( "./srcImg/001.jpeg", "./destImg/crop.jpg", 400,300, 50)
	if err != nil{
		fmt.Println( "CropWithPath======", err.Error() )
	}


	err = ResizeWithPath( "./srcImg/001.jpeg", "./destImg/resize.jpg", 400,300, 50)
	if err != nil{
		fmt.Println( "ResizeWithPath======", err.Error() )
	}


	err = FixImageWithPath( "./srcImg/001.jpeg", "./destImg/fix.jpg", 400,300, 50)
	if err != nil{
		fmt.Println( "FixImageWithPath======", err.Error() )
	}




}
