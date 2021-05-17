package wecombot

import (
	"reflect"
	"testing"
)

func Test_renderImage(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "Test_renderImage-001",
			args: args{
				data: []byte("Hello"),
			},
			want: &ImageMessage{
				MsgType: "image",
				Image: struct {
					Base64 string `json:"base64"`
					MD5    string `json:"md5"`
				}{
					Base64: "SGVsbG8=",
					MD5:    "8b1a9953c4611296a827abf8c47804d7",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := renderImage(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("renderImage() = %v, want %v", got, tt.want)
			}
		})
	}
}
