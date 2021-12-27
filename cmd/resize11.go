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

	"golang.org/x/image/draw"
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

	//open image file
	srcfile, err := os.Open(flag.Arg(0)) //maybe file path
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	defer srcfile.Close() // file close (after operation end)

	//decode image
	srcimg, t, err := image.Decode(srcfile)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	fmt.Println("Type of image:", t)

	//rectange info of image
	srcrct := srcimg.Bounds()
	minNum := int(math.Min(float64(srcrct.Dx()), float64(srcrct.Dy())))

	//scale 1:1
	dstimg := image.NewRGBA(image.Rect(0, 0, minNum, minNum))
	draw.CatmullRom.Scale(dstimg, dstimg.Bounds(), srcimg, srcrct, draw.Over, nil)

	fmt.Printf("Width: %d --> %d \n", srcrct.Dx(), dstimg.Bounds().Dx())
	fmt.Printf("Height: %d --> %d \n", srcrct.Dy(), dstimg.Bounds().Dy())

	// dst dir
	dstDir := filepath.Dir(flag.Arg(0))
	// dst extension
	dstExt := filepath.Ext(flag.Arg(0))

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
