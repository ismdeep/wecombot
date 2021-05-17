package wecombot

import (
	"reflect"
	"testing"
)

func Test_renderMarkdown(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "Test_renderMarkdown-001",
			args: args{
				str: "Hello, World.",
			},
			want: &MarkdownMessage{
				MsgType: "markdown",
				Markdown: struct {
					Content string `json:"content"`
				}{
					Content: "Hello, World.",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := renderMarkdown(tt.args.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("renderMarkdown() = %v, want %v", got, tt.want)
			}
		})
	}
}
