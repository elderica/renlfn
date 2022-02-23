// renlfnはディレクトリを再帰的に探索して、ディレクトリとファイルの名前を切り詰めるためのルーチンを実装します。
//
// 短縮されたファイル名は、ファイルパス全体のCRC32値、拡張子のない基底名(basename)、ドットを含む拡張子からなります。
package renlfn

import (
	"fmt"
	"hash/crc32"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/rivo/uniseg"
)

func ChecksumAscii(s string) string {
	cksum := crc32.ChecksumIEEE([]byte(s))
	return fmt.Sprintf("%08x", cksum)
}

// パスをとり、短縮されたファイル名を持つパスを作成します。
func TruncatePath(path string, newbasenamelen uint) string {
	cksum := ChecksumAscii(path)
	base := filepath.Base(path)
	ext := filepath.Ext(path)
	basename := base[:len(base)-len(ext)]
	truncatedBasename := TruncateString(basename, newbasenamelen)
	trimedTruncatedBasename := strings.TrimSpace(truncatedBasename)
	dir := filepath.Dir(path)
	return filepath.Join(dir, trimedTruncatedBasename+cksum+ext)
}

func TruncateString(s string, len uint) string {
	g := uniseg.NewGraphemes(s)
	c := uint(0)
	t := make([]string, 0)
	for g.Next() && c < len {
		t = append(t, g.Str())
		c++
	}
	r := strings.Join(t, "")
	return r
}

func RenameRec(config Config, dir string, depth uint) {
	if depth <= 0 {
		return
	}
	dentries, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal("ディレクトリを読み取れません:", err)
	}
	for _, dentry := range dentries {
		path := filepath.Join(dir, dentry.Name())
		if dentry.IsDir() {
			RenameRec(config, path, depth-1)
		}
		if config.LeaveDirs && dentry.Type().IsDir() {
			log.Printf("ディレクトリであるため変更しない:%s", path)
			continue
		}
		newpath := TruncatePath(path, config.Length)
		log.Printf("旧:%s 新:%s", path, newpath)
		if config.Actual {
			if os.Rename(path, newpath) != nil {
				log.Fatal("リネームに失敗しました:", err, path, newpath)
			}
		}
	}
}

type Config struct {
	Dir       string
	Actual    bool
	Depth     uint
	Length    uint
	LeaveDirs bool
}

func RealMain(config Config) {
	RenameRec(config, config.Dir, config.Depth)
}
