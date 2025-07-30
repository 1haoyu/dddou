package downloader

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"main.go/internal/types"
)

const (
	maxConcurrentDownloads = 5
	tempDir                = "ocr_temp"
)

func StartDownloadWorkers(state *types.AppState) {
	for i := 0; i < maxConcurrentDownloads; i++ {
		state.DownloadWG.Add(1)
		go func() {
			defer state.DownloadWG.Done()
			for item := range state.DownloadQueue {
				downloadImage(item.URL, item.Index, state)
			}
		}()
	}
}

func downloadImage(url string, index int, state *types.AppState) {
	fmt.Printf("访问索引: %d (切片长度: %d)\n", index, len(state.ImageObjects))

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
		Timeout: 30 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		// ui.UpdateStatus("下载失败: "+url, state)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		// ui.UpdateStatus(fmt.Sprintf("下载失败: %s (状态码: %d)", url, resp.StatusCode), state)
		return
	}

	absTempDir, _ := filepath.Abs(tempDir)
	if _, err := os.Stat(absTempDir); os.IsNotExist(err) {
		os.MkdirAll(absTempDir, 0755)
	}

	filePath := filepath.Join(absTempDir, "image_"+strconv.FormatInt(time.Now().UnixNano(), 10)+".png")
	outFile, err := os.Create(filePath)
	if err != nil {
		// ui.UpdateStatus("创建文件失败: "+filePath, state)
		return
	}
	defer outFile.Close()

	if _, err := io.Copy(outFile, resp.Body); err != nil {
		// ui.UpdateStatus("保存图片失败: "+url, state)
		return
	}

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		// ui.UpdateStatus("文件保存后消失: "+filePath, state)
		return
	}

	// 添加到OCR队列
	state.OCRQueue <- types.OCRJob{
		URL:       url,
		Index:     index,
		ImagePath: filePath,
	}

	// ui.UpdateStatus("已下载: "+filepath.Base(filePath), state)
}
