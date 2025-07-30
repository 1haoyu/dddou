package ocr

import (
	"log"
	"os"

	"github.com/doraemonkeys/paddleocr"
	"main.go/dataModel/SkuModel"
	"main.go/internal/types"
)

const maxConcurrentOCR = 3

func StartOCRWorkers(state *types.AppState) {
	for i := 0; i < maxConcurrentOCR; i++ {
		state.OCRWG.Add(1)
		go func() {
			defer state.OCRWG.Done()
			for job := range state.OCRQueue {
				processOCR(job, state)
			}
		}()
	}
}

func processOCR(job types.OCRJob, state *types.AppState) {
	if _, err := os.Stat(job.ImagePath); os.IsNotExist(err) {
		// ui.UpdateLog(fmt.Sprintf("图片文件不存在: %s", job.ImagePath), state)
		return
	}

	result, err := state.OCREngine.OcrFileAndParse(job.ImagePath)
	if err != nil {
		// ui.UpdateLog(fmt.Sprintf("OCR识别失败[图片]: %v", err), state)
		return
	}

	if result.Code != paddleocr.CodeSuccess {
		// ui.UpdateLog(fmt.Sprintf("OCR识别失败: %s", result.Msg), state)
		return
	}

	var recognizedText string
	for _, data := range result.Data {
		recognizedText += data.Text + "\n"
	}

	if recognizedText == "" {
		recognizedText = "未识别到文字"
	}
	log.Printf("图片 %d 识别结果:\n%s\n%s\n", job.Index, job.URL, recognizedText)
	// ui.UpdateLog(fmt.Sprintf("图片 %d 识别结果:\n%s\n%s\n", job.Index, job.URL, recognizedText), state)
	// ui.UpdateStatus(fmt.Sprintf("已完成: 图片 %d", job.Index), state)
}

func DoOcrImages(sku SkuModel.DataItem, state *types.AppState) {
	urls := make([]string, 0)
	urls = append(urls, sku.Img)
	urls = append(urls, sku.Pics...)
	productInfo := types.ProductOcr{
		ProductID: sku.ProductID,
		URLS:      urls,
	}
	log.Printf("加载图片: %s - %d", productInfo.ProductID, len(productInfo.URLS))

	for i := range productInfo.URLS {
		// 将任务加入下载队列
		state.DownloadQueue <- struct {
			URL   string
			Index int
		}{URL: productInfo.URLS[i], Index: i}
	}
}
