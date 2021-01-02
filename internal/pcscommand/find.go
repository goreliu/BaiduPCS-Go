package pcscommand

import (
	"fmt"
	"github.com/qjfoidnh/BaiduPCS-Go/baidupcs"
)

func getFind(pcspath string, depth int) {
	var (
		err   error
		files baidupcs.FileDirectoryList
	)
	if depth == 0 {
		err := matchPathByShellPatternOnce(&pcspath)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	files, err = GetBaiduPCS().FilesDirectoriesList(pcspath, baidupcs.DefaultOrderOptions)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, file := range files {
		if file.Isdir {
			getFind(file.Path, depth+1)
			continue
		}

		fmt.Printf("%v\n", file.Path)
	}

	return
}

// RunFind 列出文件路径
func RunFind(path string) {
	getFind(path, 0)
}
