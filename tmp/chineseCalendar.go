package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

// 大文件
var path = "/mnt/d/ed2k/cn_windows_10_business_editions_version_1909_x86_dvd_09290f8c.iso"

func download(ctx *gin.Context) {
	filename := "download"
	w := ctx.Writer
	r := ctx.Request

	file, err := os.Open(path)
	if err != nil {
		err = errors.Wrapf(err, "download openfile err")
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		err = errors.Wrapf(err, "download stat err")
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}
	md5sum, err := MD5sum(file)
	if err != nil {
		err = errors.Wrapf(err, "download md5sum err")
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}

	fmt.Println("md5sum = ", md5sum)

	w.Header().Add("Accept-Ranges", "bytes")
	w.Header().Add("Content-Disposition", "attachment; filename="+filename)
	w.Header().Add("Content-Md5", md5sum)
	var start, end int64
	if r := r.Header.Get("Range"); r != "" {
		if strings.Contains(r, "bytes=") && strings.Contains(r, "-") {
			fmt.Sscanf(r, "bytes=%d-%d", &start, &end)

			if end == 0 {
				end = info.Size() - 1
			}

			// start 从 0 开始,所以 end = info.Size() 也是有问题的，end 最大是 `info.Size() - 1`
			if start > end || start < 0 || end < 0 || end >= info.Size() {
				w.WriteHeader(http.StatusRequestedRangeNotSatisfiable)
				w.Write([]byte("参数错误...."))
				return
			}

			w.Header().Add("Content-Length", strconv.FormatInt(end-start+1, 10))
			w.Header().Add("Content-Range", fmt.Sprintf("bytes %d-%d/%d", start, end, info.Size()))
			w.Header().Set("Content-Type", "application/octet-stream")
			w.WriteHeader(http.StatusPartialContent)

		} else {
			w.WriteHeader(400)
			w.Write([]byte("header Range"))
			return
		}
	} else {
		w.Header().Add("Content-Length", strconv.FormatInt(info.Size(), 10))
		w.Header().Set("Content-Type", "application/octet-stream")
		start = 0
		end = info.Size() - 1
	}

	_, err = file.Seek(start, 0)
	if err != nil {
		err = errors.Wrapf(err, "file seek err")
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}

	n := 2048
	buf := make([]byte, n)
	for {
		if end-start+1 < int64(n) {
			n = int(end - start + 1)
		}
		_, err = file.Read(buf[:n])
		if err != nil {
			if err != io.EOF {
				err = errors.Wrapf(err, "io.Eof err")
				w.WriteHeader(500)
				w.Write([]byte(err.Error()))
				return
			}
			return
		}

		_, err = w.Write(buf[:n])
		if err != nil {
			err = errors.Wrapf(err, "Writer.Write err")
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
			return
		}

		start += int64(n)
		if start >= end+1 {
			return
		}
	}
}

func MD5sum(file *os.File) (string, error) {
	hash := md5.New()
	for buf, reader := make([]byte, 65536), bufio.NewReader(file); ; {
		n, err := reader.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			return "", err
		}
		hash.Write(buf[:n])
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}

func main() {
	route := gin.New()
	route.GET("/download", download)

	route.Run(":8080")
}
