package wecombot

import (
	"reflect"
	"testing"
)

func Test_renderImageText(t *testing.T) {
	type args struct {
		options []Option
	}
	tests := []struct {
		name string
		args args
		want *ImageTextMessage
	}{
		{
			name: "Test_renderImageText1-001",
			args: args{
				options: []Option{
					WithAddArticle(&Article{
						Title:       "Hello, World.",
						Description: "This is a hello world description.",
						URL:         "https://ismdeep.com",
						PicURL:      "https://ismdeep.com/favicon.ico",
					}),
				},
			},
			want: &ImageTextMessage{
				MsgType: "news",
				News: struct {
					Articles []*Article `json:"articles"`
				}{
					[]*Article{
						&Article{
							Title:       "Hello, World.",
							Description: "This is a hello world description.",
							URL:         "https://ismdeep.com",
							PicURL:      "https://ismdeep.com/favicon.ico",
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := renderImageText(tt.args.options...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("renderImageText() = %v, want %v", got, tt.want)
			}
		})
	}
}
