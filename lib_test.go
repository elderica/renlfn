package renlfn

import "testing"

func TestTruncatePath(t *testing.T) {
	from1 := "/tmp/renlfn/いろは/にほへと/ちりぬるを/ああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああ.mkv"
	to1 := "/tmp/renlfn/いろは/にほへと/ちりぬるを/あああああああああああああああああああああああああああああああああああああああああああああああああ0d0f797b.mkv"
	actual1 := TruncatePath(from1, 49)
	if to1 != actual1 {
		t.Errorf("MakeTruncatedPath(from) = \"%s\"; want \"%s\"", actual1, to1)
	}

	from2 := "にほへと/ちりぬるを/ああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああ.mkv"
	to2 := "にほへと/ちりぬるを/あああああああああああああああああああああああああああああああああああああああああああああああああ0d0f797b.mkv"
	actual2 := TruncatePath(from2, 49)
	if to1 != actual1 {
		t.Errorf("MakeTruncatedPath(from) = \"%s\"; want \"%s\"", actual2, to2)
	}
}

func TestWhitespaceIncluded(t *testing.T) {
	from1 := "/tmp/renlfn/いろは/にほへと/ちりぬるを/　あいうえお　かきくけこ.mkv"
	to1 := "/tmp/renlfn/いろは/にほへと/ちりぬるを/あいうえお29f9cf39.mkv"
	actual1 := TruncatePath(from1, 6)
	if to1 != actual1 {
		t.Errorf("MakeTruncatedPath(from) = \"%s\"; want \"%s\"", actual1, to1)
	}
}

func TestTruncatedString(t *testing.T) {
	from := "【重要】💕私はその人を常に先生と呼んでいた🥰"
	to := "【重要】💕私はその人を常に先生と"
	actual := TruncateString(from, 16)
	if actual != to {
		t.Errorf("MakeTruncatedString(\"%s\") = \"%s\"; want \"%s\"", from, actual, to)
	}
}

func TestBasenameLength(t *testing.T) {
	path := "/tmp/renlfn/いろは/にほへと/ちりぬるを/🗻で撮られた衝撃の映像😮.mkv"
	expected := uint(12)
	actual := BasenameLength(path)
	if actual != expected {
		t.Errorf("BasenameLength(from) = %d; want %d", actual, expected)
	}
}
