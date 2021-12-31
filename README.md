# resize11
レトロフリークで出力された画像が微妙な縦横比なので1:1にするツールを作成しました  

https://github.com/golang-standards/project-layout  
を参考にディレクトリ切ってはみましたが、あまりよくわかっていないかもしれません  

Makefileのinstallはgo installを使ってはいませんが、Path上にバイナリを配置するようにはしています  

[前提条件]  

以下URLを参考に、最新版のopencvを導入していること  
Ubuntu20.04/go1.16.10環境の場合は、go getした際に取得したzipファイルを解凍して生成されたフォルダでmake installすることで最新のopencvがインストールされます  

https://gocv.io/getting-started/  
