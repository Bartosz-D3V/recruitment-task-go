package helpers

import (
	"bufio"
	"fmt"
	"go.uber.org/zap"
	"os"
	"strconv"
)

func ReadFileIntoSlice(log *zap.SugaredLogger, path string) (*[]int, error) {
	log.Infof("Opening file at path %s", path)
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic("Failed to close file")
		}
	}(file)
	var lines []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		atoi, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, fmt.Errorf("file contains line with invalid number: %s", scanner.Text())
		}
		lines = append(lines, atoi)
	}
	return &lines, scanner.Err()
}
