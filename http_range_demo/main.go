package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
)

type FilePart struct {
	index int
	from  int
	to    int
	size  int
	data  []byte
}
type FileDownloader struct {
	url      string
	fileSize int
	// fileParts   []FilePart
	fileName string
	partNum  int
	partSize int
}

func NewFileDownloader(url string, partSize int, fileName string) (filedownloader *FileDownloader, err error) {
	filedownloader = &FileDownloader{
		url: url,
		// filePartNum: filePartNum,
		fileName: fileName,
		partSize: partSize,
	}
	return filedownloader, nil
}
func (f *FileDownloader) DownLoadFile() error {
	fileSize, err := f.getFileSize()
	if err != nil {
		return err
	}
	f.fileSize = fileSize
	f.partNum = f.fileSize/f.partSize + 1
	filePartList := make([]*FilePart, f.partNum)

	log.Println("download file begin", f)
	for i := 0; i < len(filePartList); i++ {
		filePart := &FilePart{
			index: i,
			from:  i * f.partSize,
			to:    (i+1)*f.partSize - 1,
			size:  f.partSize,
		}
		if i == len(filePartList)-1 {
			filePart.to = filePart.from + f.fileSize - (len(filePartList)-1)*f.partSize - 1
			filePart.size = f.fileSize - (len(filePartList)-1)*f.partSize
		}
		filePartList[i] = filePart
	}
	for _, part := range filePartList {
		log.Println(part)
	}

	var wg sync.WaitGroup
	for i := 0; i < len(filePartList); i++ {
		// for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(i int) {
			// log.Println("part downloading", i)
			defer wg.Done()
			filePart := filePartList[i]
			err := f.downloadPart(filePart)
			if err != nil {
				log.Println("downloadPart err", err, i)
				return

			}
			log.Println("part download succ", i, len(filePart.data))
		}(i)
	}
	wg.Wait()
	f.mergePart(filePartList)
	log.Println("downlaod file succ", f)
	return nil
}
func (f *FileDownloader) mergePart(filePartList []*FilePart) error {
	file, err := os.OpenFile(f.fileName, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}
	totalSize := 0
	for _, filePart := range filePartList {
		file.Write(filePart.data)
		totalSize += len(filePart.data)
	}
	if totalSize != f.fileSize {
		return errors.New("文件不完整")
	}
	defer file.Close()
	return nil
}
func (f *FileDownloader) downloadPart(p *FilePart) error {
	req, err := http.NewRequest("GET", f.url, nil)
	if err != nil {
		return err
	}
	// req.Header.Set("User-Agent", "mojocn")
	req.Header.Set("Range", fmt.Sprintf("bytes=%v-%v", p.from, p.to))
	// req.Header.Set("Range", fmt.Sprintf("bytes=%v-%v", 1, 250))

	client := http.DefaultClient
	resp, err := client.Do(req)
	// resp, err := http.Get(f.url)
	if err != nil {
		return err
	}
	// log.Println("show req", p.index, req, resp.Status, resp.Header)
	if resp.StatusCode > 299 {
		return errors.New("http status err:" + resp.Status)
	}

	defer resp.Body.Close()
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if len(buf) != p.size {
		return errors.New("下载文件分片长度错误,buf=" + strconv.Itoa(len(buf)) + ",p.size=" + strconv.Itoa(p.size))
	}
	p.data = buf
	// copy(p.data, buf)
	// log.Println("downloadPart", p.size, len(buf))
	return nil

}
func (f *FileDownloader) getFileSize() (int, error) {
	// http.Header.Set()
	resp, err := http.Head(f.url)
	if err != nil {
		// fmt.Println("head err", f, err)
		return 0, err
	}
	if resp.Header.Get("Accept-Ranges") != "bytes" {
		return 0, errors.New("不支持断点下载")
	}
	//resp.Header["Content-Length"]
	log.Println("show head", resp.Header, resp.Status)
	filesize, err := strconv.Atoi(resp.Header.Get("Content-Length"))
	return filesize, err

}
func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}
func main() {
	// url := "https://download.jetbrains.com/go/goland-2020.2.2.dmg"
	url := "https://download-cdn.jetbrains.com/go/goland-2020.2.2.dmg"
	partSize := 40000000
	f, _ := NewFileDownloader(url, partSize, "a.txt")
	err := f.DownLoadFile()
	if err != nil {
		log.Println("DownLoadFile err", err)
	}
	return
}
