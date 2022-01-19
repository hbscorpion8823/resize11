# resize11
レトロフリークで出力された画像が微妙な縦横比なので1:1にするツールを作成しました  

https://github.com/golang-standards/project-layout  
を参考にディレクトリ切ってはみましたが、あまりよくわかっていないかもしれません  

Makefileのinstallはgo installを使ってはいませんが、Path上にバイナリを配置するようにはしています  

[前提条件]  

以下コマンドを実行してlibmagickwand-devを導入していること  
```  
$ sudo apt install -y libmagickwand-dev
```  
その後、以下コマンドを実行してImageMagickのgoモジュールを取得していること  
```  
$ go get gopkg.in/gographics/imagick.v2/imagick
```  
