package image

import (
	"os"
	"path/filepath"
)

func Search(dir string, srcFormat string) ([]string, error) {
	sources := make([]string, 0)

	//以下，公式ドキュメントを参考． type と　function の違いを明確に意識すること
	//WalkFunc=Walkが呼び出してくれる，無名関数の型のこと．　処理内容は自分で決める．
	//Walk=各ノードにて，WalkFuncを呼び出してくれる
	if err := filepath.Walk(dir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if filepath.Ext(path) == "."+srcFormat {
				sources = append(sources, path)
			}
			return nil
		},
	); err != nil {
		return nil, err
	}

	return sources, nil
}
