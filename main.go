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

	"github.com/rivo/uniseg"
)

const (
	OriginalFilenameListFileExt = ".orfn"
	ShortNameBytesLength        = 204
	NAMEMAX                     = 255
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
	// StatもしくはReadDirに失敗するようなら飛ばす
	if err != nil {
		return fs.SkipDir
	}

	// 対象ディレクトリそのものはリネームしない
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

	base := filepath.Base(path)
	ext := filepath.Ext(path)
	basewithoutext := base[:len(base)-len(ext)]
	newbasewithoutext := TruncateBase(basewithoutext)

	newbase := fmt.Sprintf("%s_%x%s", newbasewithoutext, hv, ext)
	dir := filepath.Dir(path)
	newpath := filepath.Join(dir, newbase)

	return newpath
}

func TruncateBase(a string) string {
	grs := uniseg.NewGraphemes(a)
	newa := make([]byte, 0, NAMEMAX)

	for grs.Next() {
		bs := grs.Bytes()
		newlen := len(newa) + len(bs)
		if newlen > ShortNameBytesLength {
			break
		}
		newa = append(newa, bs...)
	}

	return string(newa)
}
