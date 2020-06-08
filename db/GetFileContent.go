package db

import (
	// "fmt"
	"io/ioutil"
	"strings"

	"github.com/google/logger"
)

func GetFileContentAsStringLines(filePath string) []string {
	logger.Infof("get file content as lines: %v", filePath)
	result := []string{}
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		logger.Errorf("read file: %v error: %v", filePath, err)
		return result
	}
	s := string(b)
	for _, lineStr := range strings.Split(s, "\n") {
		lineStr = strings.TrimSpace(lineStr)
		if lineStr == "" {
			continue
		}
		result = append(result, lineStr)
	}
	logger.Infof("get file content as lines: %v, size: %v", filePath, len(result))
	return result
}
