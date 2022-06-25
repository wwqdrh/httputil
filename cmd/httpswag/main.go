package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"path"
	"strings"

	"github.com/wwqdrh/httputil/gen"
	"github.com/wwqdrh/logger"
)

var (
	api = flag.String("api", "", "指定api列表, main.go")
	dst = flag.String("dst", "docs", "指定生成的目录")
)

func main() {
	flag.Parse()

	if *api == "" {
		logger.DefaultLogger.Panic("目标文件为空")
	}

	fi, err := os.Open(*api)
	if err != nil {
		logger.DefaultLogger.Error(err.Error())
		return
	}
	defer fi.Close()

	ans := []string{}
	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		line := strings.TrimSpace(string(a))
		if strings.Index(line, `{"GET`) == 0 ||
			strings.Index(line, `{"POST`) == 0 ||
			strings.Index(line, `{"PUT`) == 0 ||
			strings.Index(line, `{"DELETE`) == 0 {
			ans = append(ans, line)
		}
	}

	if err := os.MkdirAll(*dst, 0o777); err != nil {
		logger.DefaultLogger.Error(err.Error())
		return
	}
	f, err := os.Create(path.Join(*dst, "main.go"))
	if err != nil {
		logger.DefaultLogger.Error(err.Error())
		return
	}
	defer f.Close()
	_, err = f.WriteString(fmt.Sprintf("package %s\n", *dst))
	if err != nil {
		logger.DefaultLogger.Error(err.Error())
		return
	}
	for _, item := range gen.NewSwagInfoList(ans) {
		_, err := f.WriteString(item.String())
		if err != nil {
			logger.DefaultLogger.Warn(err.Error())
		}
	}
}
