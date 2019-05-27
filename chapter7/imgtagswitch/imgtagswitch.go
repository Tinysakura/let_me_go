package imgtagswitch

import (
	"bytes"
	"fmt"
	"image/png"
	"os"
	"runtime"
	"strconv"
)

var workers int

// 将图片文件转换成对应的<img></img>标签，打印在控制台
func ImgSwitch2H5Tag(filePaths []string) {
	workers := runtime.NumCPU()
	filePathChan := make(chan string, 2)
	done := make(chan interface{}, workers)
	resultChan := make(chan string, len(filePaths))

	go add(filePaths, filePathChan)

	for i := 0; i < workers; i++ {
		go switchs(filePathChan, resultChan, done)
	}

	go awaitCompletion(done, resultChan, workers)
	handleResult(resultChan)
}

func add(filePaths []string, filePathChan chan<- string) {
	for _, filePath := range filePaths {
		filePathChan <- filePath
	}

	close(filePathChan)
}

func switchs(filePathChan <-chan string, resultChan chan<- string, done chan interface{}) {
	for filePath := range filePathChan {
		file, e := os.Open(filePath)
		defer file.Close()
		// fmt.Println(e)
		if e == nil {
			config, e := png.DecodeConfig(file)
			// fmt.Println(e)
			if e == nil {
				buffer := bytes.NewBufferString("<img src=")
				buffer.WriteString("\"" + filePath + "\"")
				buffer.WriteString(" ")
				buffer.WriteString("width=")
				buffer.WriteString("\"" + strconv.Itoa(config.Width) + "\"")
				buffer.WriteString(" ")
				buffer.WriteString("height=")
				buffer.WriteString("\"" + strconv.Itoa(config.Height) + "\"")
				buffer.WriteString(" ")
				buffer.WriteString("/>")

				resultChan <- buffer.String()
			}
		}
	}

	done <- struct {}{}
}

// 阻塞主goroutine直到所有图片都处理完毕
func awaitCompletion(done <-chan interface{}, resultChan chan string, workers int) {
	for i := 0; i < workers; i++ {
		<- done
	}

	close(resultChan)
}

func handleResult(resultChan <-chan string) {
	for result := range resultChan {
		fmt.Println(result)
	}
}
