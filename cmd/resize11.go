package main

import (
	"flag"
	"fmt"
	"math"
	"os"

	"path/filepath"
	"time"

	"gopkg.in/gographics/imagick.v2/imagick"
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
	// ImageMagickモジュール初期化（終わったら停止）
	imagick.Initialize()
	defer imagick.Terminate()

	// 入力ファイルからイメージを取得
	srcPath := flag.Arg(0)
	mw1 := imagick.NewMagickWand()
	defer mw1.Destroy()
	err := mw1.ReadImage(srcPath)
	if err != nil {
		panic(err)
	}

	// scale 1:1
	width := mw1.GetImageWidth()
	height := mw1.GetImageHeight()
	maxNum := int(math.Max(float64(width), float64(height)))
	fmt.Printf("Width: %d --> %d \n", width, maxNum)
	fmt.Printf("Height: %d --> %d \n", height, maxNum)

	// resize canvas
	err = mw1.ExtentImage(uint(maxNum), uint(maxNum), 0, 0)
	if err != nil {
		panic(err)
	}

	// dst dir
	dstDir := filepath.Dir(srcPath)
	// dst extension
	dstExt := filepath.Ext(srcPath)

	// dst file name
	nowTime := time.Now()
	dstFilename := nowTime.Format("2006-01-02-150405")

	dstPath := filepath.Join(dstDir, dstFilename+dstExt)

	// ファイル出力
	mw1.WriteImage(dstPath)

	// 終了メッセージ出力
	fmt.Println(dstPath, " output complete!!")
}
