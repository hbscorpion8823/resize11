package main

import (
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"

	"math"
	"path/filepath"
	"time"

	"gocv.io/x/gocv"
)

func main() {
	// 引数チェック
	flag.Parse()
	if flag.NArg() != 1 {
		// 引数が1でない場合はエラー
		// arg1: io.Writer, arg2: interface{}
		fmt.Fprintln(os.Stderr, os.ErrInvalid)
		return
	}

	// 入力ファイルからイメージを取得
	srcPath := flag.Arg(0)
	inCvImg := gocv.IMRead(srcPath, gocv.IMReadColor)
	// 画像の型を特定するための処理
	srcfile, err := os.Open(srcPath) //maybe file path
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	defer srcfile.Close() // file close (after operation end)
	_, t, err := image.Decode(srcfile)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	//rectange info of image
	rowSize := inCvImg.Rows()
	colSize := inCvImg.Cols()
	minNum := int(math.Min(float64(rowSize), float64(colSize)))

	//scale 1:1
	originPoint := image.Point{colSize - minNum, rowSize - minNum}
	outputImg := gocv.NewMatWithSize(minNum, minNum, gocv.MatTypeCV8U)
	gocv.Resize(inCvImg, &outputImg, originPoint, 1.0, 1.0, gocv.InterpolationMax)
	outputImg = outputImg.Region(image.Rectangle{image.Point{0, 0}, image.Point{minNum, minNum}})
	dstimg, _ := outputImg.ToImage()

	fmt.Printf("dstimg.Bounds.Dx: %d\n", dstimg.Bounds().Dx())
	fmt.Printf("dstimg.Bounds.Dy: %d\n", dstimg.Bounds().Dy())

	fmt.Printf("Width: %d --> %d \n", colSize, minNum)
	fmt.Printf("Height: %d --> %d \n", rowSize, minNum)

	// dst dir
	dstDir := filepath.Dir(srcPath)
	// dst extension
	dstExt := filepath.Ext(srcPath)

	// dst file name
	nowTime := time.Now()
	dstFilename := nowTime.Format("2006-01-02-150405")

	dstPath := filepath.Join(dstDir, dstFilename+dstExt)
	dstfile, err := os.Create(dstPath)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	defer dstfile.Close()

	// encode resized image
	// 画像フォーマットごとに出力方式を変える
	switch t {
	case "jpeg":
		if err := jpeg.Encode(dstfile, dstimg, &jpeg.Options{Quality: 100}); err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
	case "gif":
		if err := gif.Encode(dstfile, dstimg, nil); err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
	case "png":
		if err := png.Encode(dstfile, dstimg); err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
	default:
		fmt.Fprintln(os.Stderr, "format error")
	}

	// 終了メッセージ出力
	fmt.Println(dstPath, " output complete!!")
}
