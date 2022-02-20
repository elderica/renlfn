package renlfn

import "testing"

func TestMakeTruncatedPath(t *testing.T) {
	from1 := "/tmp/renlfn/いろは/にほへと/ちりぬるを/ああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああ.mkv"
	to1 := "/tmp/renlfn/いろは/にほへと/ちりぬるを/0d0f797bあああああああああああああああああああああああああああああああああああああああああああああああああ.mkv"
	actual1 := TruncatePath(from1, 49)
	if to1 != actual1 {
		t.Errorf("MakeTruncatedPath(from) = \"%s\"; want \"%s\"", actual1, to1)
	}

	from2 := "にほへと/ちりぬるを/ああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああ.mkv"
	to2 := "にほへと/ちりぬるを/0d0f797bあああああああああああああああああああああああああああああああああああああああああああああああああ.mkv"
	actual2 := TruncatePath(from2, 49)
	if to1 != actual1 {
		t.Errorf("MakeTruncatedPath(from) = \"%s\"; want \"%s\"", actual2, to2)
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
