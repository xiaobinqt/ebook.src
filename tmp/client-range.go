package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"strconv"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func DownloadDownloadArtifact(downloadPath, surl string) (err error) {
	dfn := downloadPath
	var (
		file *os.File
		size int64
		headerMd5sum,
		downloadMd5sum string
	)

	file, err = os.OpenFile(dfn, os.O_RDWR|os.O_CREATE, 0644)
	defer file.Close()
	if err != nil {
		err = errors.Wrapf(err, "download openfile err")
		return err
	}
	stat, _ := file.Stat()
	size = stat.Size()
	sk, err := file.Seek(size, 0)
	if err != nil {
		err = errors.Wrapf(err, "seek err")
		return err
	}

	if sk != size {
		err = fmt.Errorf("seek length not equal file size,seek=%d,size=%d", sk, size)
		logrus.Error(err.Error())
		return err
	}

	request := http.Request{}
	request.Method = http.MethodGet
	if size != 0 {
		header := http.Header{}
		header.Set("Range", "bytes="+strconv.FormatInt(size, 10)+"-")
		request.Header = header
	}
	parse, _ := url.Parse(surl)
	request.URL = parse
	resp, err := http.DefaultClient.Do(&request)
	//resp, err := http.DefaultClient.Do(&request)
	defer resp.Body.Close()
	if err != nil {
		err = errors.Wrapf(err, "client do err")
		logrus.Error(err.Error())
		return err
	}

	headerMd5sum = resp.Header.Get("Content-Md5")
	if headerMd5sum == "" {
		return fmt.Errorf("resp header md5sum empty")
	}

	body := resp.Body
	writer := bufio.NewWriter(file)
	bs := make([]byte, 1024*1024)
	for {
		var read int
		read, err = body.Read(bs)
		if err != nil {
			fmt.Println("for loop = ", read, err)
			if err != io.EOF {
				err = errors.Wrapf(err, "body read not io eof")
				logrus.Error(err.Error())
				return err
			}

			if err == io.EOF && resp.StatusCode != http.StatusOK {
				err = nil
				return
			}

			if read != 0 {
				_, err = writer.Write(bs[:read])
				if err != nil {
					err = errors.Wrapf(err, "writer write err")
					return err
				}
			}

			err = nil
			break
		}
		_, err = writer.Write(bs[:read])
		if err != nil {
			err = errors.Wrapf(err, "writer write err")
			return err
		}
	}

	if err != nil {
		return err
	}

	err = writer.Flush()
	if err != nil {
		err = errors.Wrapf(err, "writer.Flush err")
		return err
	}

	// 比对 md5 是否一致
	downloadMd5sum, err = md5sum(downloadPath)
	if err != nil {
		err = errors.Wrapf(err, "get download md5dum err")
		logrus.Error(err.Error())
		// md5 不一致直接删除
		//os.Remove(downloadPath)
		return err
	}
	logrus.Debugf("downloadMd5sum: %s,headerMd5sum:%s ", downloadMd5sum, headerMd5sum)

	if downloadMd5sum == headerMd5sum {
		return nil
	}

	// 错误了删除 tar 包
	//os.Remove(downloadPath)
	return fmt.Errorf("download md5sum not equal header md5dum")
}

func md5sum(downloadPath string) (string, error) {
	cmdStr := fmt.Sprintf("printf $(md5sum %s)", downloadPath)
	cmdOutput, err := exec.Command("/bin/sh", "-c", cmdStr).CombinedOutput()
	logrus.Debugf("md5sum: %s ", cmdStr)
	if err != nil {
		err = errors.Wrapf(err, "md5sum [%s] exec.Command err", cmdStr)
		logrus.Error(err.Error())
		return "", err
	}
	return string(cmdOutput), nil
}

func main() {
	err := DownloadDownloadArtifact("/mnt/d/tmp/xxx.111.test", "http://127.0.0.1:8080/download")
	if err != nil {
		fmt.Println("download err", err.Error())
		return
	}
	fmt.Println("success..........")
}
