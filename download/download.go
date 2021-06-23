package download

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type Bar struct {
	percent int64  //百分比
	cur     int64  //当前进度位置
	total   int64  //总进度
	rate    string //进度条
	graph   string //显示符号
}


func (bar *Bar) NewOption(start, total int64) {
	bar.cur = start
	bar.total = total
	if bar.graph == "" {
		bar.graph = "█"
	}
	bar.percent = bar.getPercent()
	for i := 0; i < int(bar.percent); i += 2 {
		bar.rate += bar.graph //初始化进度条位置
	}
}
func (bar *Bar) getPercent() int64 {
	return int64(float32(bar.cur) / float32(bar.total) * 100)
}

func (bar *Bar) NewOptionWithGraph(start, total int64, graph string) {
	bar.graph = graph
	bar.NewOption(start, total)
}

func (bar *Bar) Play(cur int64) {
	bar.cur = cur
	last := bar.percent
	bar.percent = bar.getPercent()
	if bar.percent != last && bar.percent%2 == 0 {
		bar.rate += bar.graph
	}
	fmt.Printf("\r[%-50s]%3d%%  %8d/%d", bar.rate, bar.percent, bar.cur, bar.total)
}
func (bar *Bar) Finish(){
	fmt.Println()
}


type WriteCounter struct {
	Total int64
	All int64
	Procbar Bar
}

func (wc *WriteCounter) Write(p []byte) (int, error) {
	n := len(p)
	wc.Total += int64(n)
	wc.Procbar.Play(wc.Total)
	return n, nil
}


func DownloadFile(Url string, DownPath string) (error) {
	out, err := os.Create(DownPath + ".tmp")
	if err != nil {
		return err
	}
	resp, err := http.Get(Url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var bar Bar
	bar.NewOption(0,resp.ContentLength)
	counter := &WriteCounter{0,resp.ContentLength,bar}
	if _, err = io.Copy(out, io.TeeReader(resp.Body, counter)); err != nil {
		out.Close()
		return err
	}
	out.Close()
	bar.Finish()

	err = os.Rename(DownPath+".tmp", DownPath)
	if err != nil {
		return err
	}
	return nil
}

