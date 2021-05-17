package wecombot

import (
	"reflect"
	"testing"
)

func Test_renderText(t *testing.T) {
	type args struct {
		str     string
		options []Option
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "Test_renderText-001",
			args: args{
				str: "Hello, World.",
				options: []Option{
					WithMentionedList([]string{"L. Jiang"}),
					WithMentionAll(),
				},
			},
			want: &TextMessage{
				MsgType: "text",
				Text: struct {
					Content             string   `json:"content"`
					MentionedList       []string `json:"mentioned_list"`
					MentionedMobileList []string `json:"mentioned_mobile_list"`
				}{
					Content:             "Hello, World.",
					MentionedList:       []string{"L. Jiang", "@all"},
					MentionedMobileList: []string{},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := renderText(tt.args.str, tt.args.options...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("renderText() = %v, want %v", got, tt.want)
			}
		})
	}
}
