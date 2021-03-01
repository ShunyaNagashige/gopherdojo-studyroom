package image_test

import (
	"testing"

	"github.com/ShunyaNagashige/imgconv/image"
)

type ConvertInput struct {
	src       string
	srcFormat string
	dstFormat string
}

func TestConvert(t *testing.T) {
	cases := []struct {
		name  string
		input ConvertInput
		//エラーが存在するかどうか
		isErr bool
	}{
		{
			name:  "png:jpg",
			input: ConvertInput{src: "testdata/dog_hachi_sasareta.png", srcFormat: "png", dstFormat: "jpg"},
			isErr: false,
		},
		{
			name:  "png:gif",
			input: ConvertInput{src: "testdata/dog_hachi_sasareta.png", srcFormat: "png", dstFormat: "gif"},
			isErr: false,
		},
		{
			name:  "jpg:png",
			input: ConvertInput{src: "testdata/itu.jpg", srcFormat: "jpg", dstFormat: "png"},
			isErr: false,
		},
		{
			name:  "jpg:gif",
			input: ConvertInput{src: "testdata/itu.jpg", srcFormat: "jpg", dstFormat: "gif"},
			isErr: false,
		},
		{
			name:  "gif:png",
			input: ConvertInput{src: "testdata/pop.gif", srcFormat: "gif", dstFormat: "png"},
			isErr: false,
		},
		{
			name:  "gif:jpg",
			input: ConvertInput{src: "testdata/pop.gif", srcFormat: "gif", dstFormat: "jpg"},
			isErr: false,
		},
		{
			name:  "source file open error",
			input: ConvertInput{src: "testdata/nothing.png", srcFormat: "gif", dstFormat: "jpg"},
			isErr: true,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			err := image.Convert(c.input.src, c.input.srcFormat, c.input.dstFormat)
			if !CheckErr(t, c.isErr, err){
				t.Errorf("case[%s]: expected isErr = %t ,but actual isErr = %t, err = %s",
					c.name, c.isErr, !(c.isErr), err)
			}
		})
	}
}
