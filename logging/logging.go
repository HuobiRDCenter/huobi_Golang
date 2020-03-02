package logging

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type PerformanceLogging struct {
	logger *log.Logger
	file   *os.File
	index int
	start time.Time
}

func (p *PerformanceLogging) Init() *PerformanceLogging {
	var err error
	fileName := time.Now().Format("20060102_150405.txt")
	p.file, err = os.OpenFile(fileName, os.O_CREATE | os.O_WRONLY | os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open file: ", fileName)
	}
	p.logger = log.New(p.file, "", 0)
	p.index = 1
	p.logger.Println("Index, Duration(ms), URL")
	return p
}

func (p *PerformanceLogging) Start() {
	p.start = time.Now()
}

func (p *PerformanceLogging) StopAndLog(method string, url string) {
	duration := time.Since(p.start).Milliseconds()

	i := strings.IndexByte(url, '?')
	var path string
	if i > 0 {
		path = url[0:i]
	} else {
		path = url
	}

	p.logger.Println(fmt.Sprintf("%d, %d, %s %s", p.index, duration, method, path))
	p.index++
}