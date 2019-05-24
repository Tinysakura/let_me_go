package chapter5

import (
	"archive/tar"
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

var addFunctionForSuffix map[string] func(string) string
var fileListFunctionForSuffix map[string] func(string) ([]string, error)

// 创建为文件添加后缀的方法的通用函数
func makeAddSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasPrefix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

// 创建解压特定后缀文件的方法的通用函数
func makeListCompressFileFunc(suffix string) func(string) ([]string, error) {
	if strings.HasSuffix(suffix, ".tar.gz") {
		return func(s string) (i []string, e error) {
			// file read
			fr, err := os.Open(s)
			if err != nil {
				panic(err)
			}
			defer fr.Close()

			// gzip read
			gr, err := gzip.NewReader(fr)
			if err != nil {
				panic(err)
			}
			defer gr.Close()

			// tar read
			tr := tar.NewReader(gr)

			files := make([]string, 10)

			// 读取文件
			for {
				h, err := tr.Next()

				if err == io.EOF {
					break
				}

				if err != nil {
					panic(err)
				}

				files = append(files, h.Name)
			}

			return files, err
		}
	}

	return nil
}

func init() {
	tarFileList := makeListCompressFileFunc(".tar.gz")

	fileListFunctionForSuffix = map[string]func(string) ([]string, error){".tar.gz":tarFileList}
}

func ArchiveFileListMap(filePath string)([]string, error) {
	if function, exist := fileListFunctionForSuffix[fileSuffix(filePath)]; exist{
		return function(filePath)
	}

	return nil, errors.New("nonsupport compress file format")
}

func fileSuffix(filePath string) string {
	suffix := filePath[strings.Index(filePath, "."):]
	fmt.Println(suffix)
	return suffix
}

// 返回一个切片中所有字符串的共同前缀
func CommonPrefix(slice []string) string {
	if len(slice) == 0 {
		return ""
	}

	if len(slice) == 1 {
		return slice[0]
	}

	s1 := slice[0]
	s2 := slice[1]

	var commonPrefix []rune

	for index, charOutside := range []rune(s1) {
		if charOutside == []rune(s2)[index] {
			commonPrefix = append(commonPrefix, rune(charOutside))
		} else {
			break
		}
	}

	for _, s := range slice[2:] {
		for i := len(commonPrefix); i > 0 ; i -- {
			if strings.HasPrefix(s, string(commonPrefix[0:i])) {
				commonPrefix = commonPrefix[0:i]
				break
			}
		}
	}

	fmt.Println(fmt.Sprintf("commonPrefix:%s", string(commonPrefix)))
	return string(commonPrefix)
}

