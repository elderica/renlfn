package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"

	"github.com/elderica/renlfn"
)

var (
	flagdir    string
	flagactual bool
	flagdepth  int
	flaglength int

	config renlfn.Config
)

func init() {
	flag.StringVar(&flagdir, "dir", "", "対象とするディレクトリ")
	flag.BoolVar(&flagactual, "actual", false, "実際にリネームする")
	flag.IntVar(&flagdepth, "depth", 2, "対象となる深さ")
	flag.IntVar(&flaglength, "length", 32, "新しい基底名の長さ(63文字まで)")
}

func processflags() {
	flag.Parse()

	if flag.NFlag() == 0 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	if flagdir == "" {
		log.Fatal("-dirを指定していません")
	}
}

func main() {
	processflags()

	abspath, err := filepath.Abs(flagdir)
	if err != nil {
		log.Fatal("abspath: ", err)
	}

	fileinfo, err := os.Stat(abspath)
	if err != nil {
		log.Fatal("stat: ", err)
	}
	if !fileinfo.IsDir() {
		log.Fatal("ディレクトリではありません:", flagdir)
	}

	if flaglength > 63 {
		log.Fatal("基底名の長さが63文字を越えています。")
	}

	config.Dir = abspath
	config.Actual = flagactual
	config.Depth = uint(flagdepth)
	config.Length = uint(flaglength)

	renlfn.RealMain(config)
}
