/*
Copyright © 2022 Shuk

*/
package cmd

import "testing"

func Test_replaceName(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "Replace prefix tags", args: args{"[1][2][3]file.ext"}, want: "file.ext"},
		{name: "Replace prefix tags with symbol", args: args{"[*][!][#]file.ext"}, want: "file.ext"},
		{name: "Replace prefix tags with empty", args: args{"[][][]file.ext"}, want: "file.ext"},
		{name: "Case 1", args: args{"[140929][メリー・ジェーン]すぽコン～SPORTS WEAR-COMPLEX～ 上巻 アスリートボディのゆ・う・わ・く[720P].mp4"}, want: "すぽコン～SPORTS WEAR-COMPLEX～ 上巻 アスリートボディのゆ・う・わ・く[720P].mp4"},

		{name: "Replace prefix tags with empty", args: args{"(18禁アニメ) 炎の孕ませおっぱいエロアプリ学園 THE ANIMATION 第1巻（1280x720 AVC_AAC）.mp4"}, want: "炎の孕ませおっぱいエロアプリ学園 THE ANIMATION 第1巻（1280x720 AVC_AAC）.mp4"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removePrefixTags(tt.args.s); got != tt.want {
				t.Errorf("removePrefixTags() = %v, want %v", got, tt.want)
			}
		})
	}
}
