package image_test

import (
	"fmt"
	"testing"

	"github.com/ShunyaNagashige/imgconv/image"
)

type SearchInput struct {
	dir    string
	format string
}

//エラーの存在が期待された通りになっているかチェック
func CheckErr(t *testing.T, isErrExpected bool, err error) bool {
	t.Helper()

	//エラーが期待されない(nilが期待される)場合
	if !isErrExpected {
		if err != nil {
			//t.Errorf("case[%s] err should be nil,but got:%s", t.Name(), err)
			return false
		}

		return true
	}

	//エラーが期待されているがnilであった場合
	if err == nil {
		//t.Errorf("case[%s] err should not be nil,but got:%s", t.Name(), err)
		return false
	}

	return true
}

func TestSearch(t *testing.T) {

	cases := []struct {
		name  string
		input SearchInput
		//エラーが存在するかどうか
		isErr    bool
		expected []string
	}{
		{
			name:     "png",
			input:    SearchInput{dir: ".", format: "png"},
			isErr:    false,
			expected: []string{"testdata/dog_hachi_sasareta.png"},
		},
		{
			name:     "jpg",
			input:    SearchInput{dir: ".", format: "jpg"},
			isErr:    false,
			expected: []string{"testdata/itu.jpg"},
		},
		{
			name:     "gif",
			input:    SearchInput{dir: ".", format: "gif"},
			isErr:    false,
			expected: []string{"testdata/pop.gif"},
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			actual, err := image.Search(c.input.dir, c.input.format)

			if !CheckErr(t, c.isErr, err) {
				t.Errorf("case[%s]: expected isErr = %t ,but actual isErr = %t, err = %s",
					c.name, c.isErr, !(c.isErr), err)
			}

			if len(actual) != len(c.expected) {
				t.Errorf("実際の出力と期待値の,lengthが異なります: want Search(%s,%s) = %s, got %s",
					c.input.dir, c.input.format, c.expected, actual)
			}

			for i := 0; i < len(c.expected); i++ {
				if actual[i] != c.expected[i] {
					t.Errorf("actual[%d]=%s と c.expected[%d]=%sの値が一致しません: want Search(%s,%s) = %s, got %s",
						i, actual[i], i, c.expected[i], c.input.dir, c.input.format, c.expected, actual)
				}
			}
		})
	}
}

func ExampleSearch() {
	fmt.Println(image.Search(".", "png"))
	// output: [testdata/dog_hachi_sasareta.png] <nil>
}
