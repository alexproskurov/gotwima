package main

import (
	"bytes"
	_ "embed"
	"io"
	"net/http"
	"reflect"
	"testing"

	"github.com/soranoba/googp"
)

//go:embed image.html
var imageHTML []byte

//go:embed video.html
var videoHTML []byte

func TestParse(t *testing.T) {
	type args struct {
		res  *http.Response
		opts []googp.ParserOpts
	}
	tests := []struct {
		name    string
		args    args
		want    CustomOGP
		wantErr bool
	}{
		{
			name: "test_image",
			args: args{
				res: &http.Response{
					StatusCode: 200,
					Body:       io.NopCloser(bytes.NewReader(imageHTML)),
				},
				opts: nil,
			},
			want: CustomOGP{
				Title:       "chosen undead (@chosenundeadone)",
				Description: "",
				Image:       "https://pbs.twimg.com/media/GZnEcwZXUAckOXs.jpg",
			},
			wantErr: false,
		},
		{
			name: "test_video",
			args: args{
				res: &http.Response{
					StatusCode: 200,
					Body:       io.NopCloser(bytes.NewReader(videoHTML)),
				},
				opts: nil,
			},
			want: CustomOGP{
				Title:       "Nature is Amazing ☘️ (@AMAZlNGNATURE)",
				Description: "Moo Deng got in the bowl so momma would stop eating and pay attention to her 😂",
				Image:       "https://pbs.twimg.com/ext_tw_video_thumb/1850633711596793856/pu/img/I2bnm3FGu2sImvUq.jpg",
				Video:       "https://api.fxtwitter.com/2/go?url=https%3A%2F%2Fvideo.twimg.com%2Fext_tw_video%2F1850633711596793856%2Fpu%2Fvid%2Favc1%2F576x1024%2FZLOKNCguRtMigowx.mp4%3Ftag%3D12",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.args.res, tt.args.opts...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}
