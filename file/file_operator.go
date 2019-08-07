package file

import (
	"archive/zip"
	"bufio"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func WriteFile(fileName string, writeContent string) {
	file, e := os.Create(fileName)

	if e != nil {
		log.Fatal(e)
	}

	defer file.Close()

	_, e = file.WriteString(writeContent)
}

func BufferWriteFile(fileName string, writeContent string) {
	file, e := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0666)

	if e != nil {
		log.Fatal(e)
	}

	defer file.Close()

	bufferedWriter := bufio.NewWriter(file)

	_, _ = bufferedWriter.Write([]byte(writeContent))
	_ = bufferedWriter.Flush()

	file.Close()
}

func ReadFile(fileName string) {
	file, _ := os.OpenFile(fileName, os.O_RDONLY, 0666)
	defer file.Close()

	bytes, _ := ioutil.ReadAll(file)
	println(string(bytes))
}

func ScanRead(fileName string) {
	file, _ := os.OpenFile(fileName, os.O_RDONLY, 0666)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for {
		println(scanner.Text())
		scan := scanner.Scan()

		if scan == false {
			break
		}
	}
}

func ZipFile(zipFilenName string) {
	zipFile, _ := os.Create(zipFilenName)
	defer  zipFile.Close()

	zipWrite := zip.NewWriter(zipFile)
	defer  zipWrite.Close()

	writer, _ := zipWrite.Create("/Users/chenfeihao/Desktop/go1.txt")
	_, _ = writer.Write([]byte("zip zip zip"))

	_ = zipWrite.Flush()
}

func DownloadFromWeb(httpUrl string, fileName string) {
	file, _ := os.Create(fileName)
	defer  file.Close()

	resp, _ := http.Get(httpUrl)

	defer resp.Body.Close()

	_, _ = io.Copy(file, resp.Body)
}

