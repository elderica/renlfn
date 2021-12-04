package main

import (
	"flag"
	"fmt"
	"hash/fnv"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

const (
	OriginalFilenameListFileExt = ".orfn"
)

var (
	flagdir string

	orfnfilepath string
	orfnfile     *os.File
)

func init() {
	flag.StringVar(&flagdir, "dir", "", "対象とするディレクトリ")
}

func processflagdir() {
	if flagdir == "" {
		log.Fatalln("-dirを指定していません")
	}

	abspath, err := filepath.Abs(flagdir)
	if err != nil {
		log.Fatalln("abspath: ", err)
	}
	flagdir = abspath

	fileinfo, err := os.Stat(flagdir)
	if err != nil {
		log.Fatalln("processflagdir.stat: ", err)
	}
	if !fileinfo.IsDir() {
		log.Fatalln("ディレクトリではありません:", flagdir)
	}
	orfnfilepath = flagdir + OriginalFilenameListFileExt

	log.Println("オリジナルファイル名をここに格納する:", orfnfilepath)
}

func processflags() {
	flag.Parse()

	processflagdir()
}

func main() {
	processflags()

	startwalk()
}

// startwalk は指定されたディレクトリを見にいく。
func startwalk() {
	f, err := os.Create(orfnfilepath)
	if err != nil {
		log.Fatalln("startwalk:", err)
	}
	orfnfile = f
	defer orfnfile.Close()

	filepath.WalkDir(flagdir, dirwalker)
}

// dirwalker はリネーム処理を行なう。
// dirwalker: io/fs.WalkDirFunc
func dirwalker(path string, d fs.DirEntry, err error) error {
	// 対象ディレクトリそのものは飛ばす
	if path == flagdir {
		return nil
	}

	newpath := makenewpath(path)
	fmt.Fprintln(orfnfile, path, newpath)

	return nil
}

// makenewpath はファイルやディレクトリの新しい名前を生成する。
func makenewpath(path string) string {
	hs := fnv.New64()
	io.WriteString(hs, path)
	hv := hs.Sum(nil)

	ext := filepath.Ext(path)
	newname := fmt.Sprintf("%x%s", hv, ext)
	dir := filepath.Dir(path)
	newpath := filepath.Join(dir, newname)

	return newpath
}
