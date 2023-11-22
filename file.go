package h

import (
	"bufio"
	"errors"
	"io"
	"os"
)

func ReadLineByLine(fileName string, lineHandle func(line []byte) error) error {
	return ReadLineByLine0(fileName, lineHandle, 0, 0)
}

func ReadLineByLine0(fileName string, lineHandle func(line []byte) error, maxLineSize int, delim byte) error {
	if maxLineSize == 0 {
		maxLineSize = 1024 * 1024
	}
	if delim == 0 {
		delim = '\n'
	}

	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := bufio.NewReaderSize(file, maxLineSize) // 指定缓冲区大小
	for {
		line, err := reader.ReadBytes(delim) // 按行读取文件内容并返回字节切片
		if errors.Is(err, io.EOF) {          // 结束
			break
		}
		if err != nil { // 异常
			return err
		}
		err = lineHandle(line)
		if err != nil {
			return err
		}
	}
	return nil
}
