#!/bin/bash
set -euo pipefail

renlfn="go run github.com/elderica/renlfn/bin/renlfn"
bindir="$(pwd)"
basedir="/tmp/renlfn"
rm -r -f "$basedir"

rep() {
	s="${1:-あ}"
	r="${2:-10}"
	printf "${s}%.0s" $(seq 1 "${r}")
}
function mkcd {
	last=$(eval "echo \$$#")
	if [ -z "$last" ]; then
		echo "Enter a directory name"
	elif [ -d "$last" ]; then
		echo "\`$last' already exists"
	else
		mkdir -p "$@" && cd "$last"
	fi
}

mkcd "$basedir"
	mkcd "いろは"
		mkcd "にほへと"
			mkcd "ちりぬるを"
				:> "私がその掛茶屋で先生を見た時は、先生がちょうど着物を脱いでこれから海へ入ろうとするところであった。私はその時反対に濡れた身体を風に吹かして水から上がって来た。.mkv"
				:> "その西洋人の優れて白い皮膚の色が、掛茶屋へ入るや否や、すぐ私の注意を惹いた。純粋の日本の浴衣を着ていた彼は、それを床几の上にすぽりと放り出したまま.mp4"
				cd ..
			:> "彼は我々の穿く猿股一つの外何物も肌に着けていなかった。.avi"
			mkcd "わかをたれそ"
				:> "私にはそれが第一不思議だった。私はその二日前に由井が浜まで行って、砂の上にしゃがみながら、長い間西洋人の海へ入る様子を眺めていた。.ogv"
				cd ..
			cd ..
		mkcd "つねならむ"
			:> "彼はやがて自分の傍を顧みて、そこにこごんでいる日本人に、一言二言何かいった。.ogg"
			cd ..
		mkcd "うゐのおくやま"
			:> "女は殊更肉を隠しがちであった。.html"
			cd ..
		cd ..
	mkcd "けふこえて"
		mkcd "あさきゆめみし"
			:> "大抵は頭に護謨製の頭巾を被って、海老茶や紺や藍の色を波間に浮かしていた。.mov"
			cd ..
		mkcd "ゑひもせす"
			cd ..
		:> "私は単に好奇心のために、並んで浜辺を下りて行く二人の後姿を見守っていた。.mp3"

cd "$bindir"
find "$basedir"
echo --------------------------------------------------------------------
$renlfn -dir "$basedir" -depth 8 -length 10 -leavedirs
echo --------------------------------------------------------------------
$renlfn -dir "$basedir" -actual -depth 8 -length 10 -leavedirs
echo --------------------------------------------------------------------
find "$basedir"
