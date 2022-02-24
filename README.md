# renlfn [![License: Unlicense](https://img.shields.io/badge/license-Unlicense-blue.svg)](http://unlicense.org/)

## これは何？
長い名前がついたファイルやディレクトリが大量にあるが、それをファイルサーバにアップロードしようとするとエラーになってしまう!  
こんなに沢山あると、リネームするのめんどくさすぎる!

という問題に対処するための、次善策的バッチリネームツールです。

## 使いかた
### 概念
このツールでは **基底名** を基準にファイルやディレクトリの長さを処理します。基底名とは、ファイルやディレクトリから拡張子を取り除いたものです。

実行すると、指定したディレクトリを1階層目として、再帰的にリネーム処理を行います。  
リネーム後のファイルやディレクトリの名前は、「切り詰めた基底名」と「リネーム元の絶対パスのCRC32チェックサム」と「拡張子」を連結したものに変更します。

### 最も基本的な使いかた
```bat
> renlfn -actual -dir D:\path\to\dir
```

### Usage
```
Usage of renlfn:
  -actual
        実際にリネームする
  -depth int
        対象となる深さ (default 4)
  -dir string
        対象とするディレクトリ
  -leavedirs
        ディレクトリをリネームしない
  -leavelength int
        これ以下の長さの基底名を対象としない (default 30)
  -length int
        新しい基底名の長さ(日本語で63文字まで) (default 50)
```

## 作者
elderica