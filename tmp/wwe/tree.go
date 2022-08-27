package main

import (
	"os"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

//func main() {
//	t, err := DirTreeJSON("d:/tmp/package")
//	if err != nil {
//		fmt.Println(err)
//	}
//
//	b, _ := json.Marshal(t.Children)
//	fmt.Println(string(b))
//}

type Tree struct {
	Title    string `json:"title"`
	Children []Tree `json:"children"`
	ID       string `json:"id"`
}

func DirTreeJSON(dstPath string) (tree Tree, err error) {
	dstF, err := os.Open(dstPath)
	if err != nil {
		err = errors.Wrapf(err, "打开目录失败，目录：%s", dstPath)
		return tree, err
	}

	defer dstF.Close()
	fileInfo, err := dstF.Stat()
	if err != nil {
		return tree, err
	}

	if fileInfo.IsDir() == false { //如果是文件
		tree.Title = fileInfo.Name()
		tree.ID = uuid.NewV4().String()
		return tree, nil
	} else { //如果是文件夹

		dir, err := dstF.Readdir(0) //获取文件夹下各个文件或文件夹的fileInfo
		if err != nil {
			return tree, err
		}

		for _, fileInfo = range dir {
			x, err := DirTreeJSON(dstPath + "/" + fileInfo.Name())
			if err != nil {
				return tree, err
			}
			x.Title = fileInfo.Name()
			tree.ID = uuid.NewV4().String()
			tree.Children = append(tree.Children, x)
		}

		return tree, nil
	}
}
