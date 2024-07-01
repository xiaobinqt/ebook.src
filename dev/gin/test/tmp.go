package main

import (
	"bytes"
	"fmt"
	"golang.org/x/net/html"
	"strings"
)

func main() {
	htmlContent := `<h2>1.1         列表</h2><figure class="table"><table style=";"><tbody><tr><td style="border:1.0pt solid windowtext;padding:0cm 5.4pt;vertical-align:top;width:77.75pt;"><strong>操作用户</strong></td><td style="border-bottom-style:solid;border-color:windowtext;border-left-style:none;border-right-style:solid;border-top-style:solid;border-width:1.0pt;padding:0cm 5.4pt;vertical-align:top;width:397.8pt;"><span style="color:red;">系统管理员</span></td></tr><tr><td style="border-bottom-style:solid;border-color:windowtext;border-left-style:solid;border-right-style:solid;border-top-style:none;border-width:1.0pt;padding:0cm 5.4pt;vertical-align:top;width:77.75pt;"><strong>预置条件</strong></td><td style="border-bottom:1.0pt solid windowtext;border-left-style:none;border-right:1.0pt solid windowtext;border-top-style:none;padding:0cm 5.4pt;vertical-align:top;width:397.8pt;"><p style="margin-left:18.0pt;">1.        系统Root用户(admin)，已登录</p></td></tr><tr><td style="border-bottom-style:solid;border-color:windowtext;border-left-style:solid;border-right-style:solid;border-top-style:none;border-width:1.0pt;padding:0cm 5.4pt;vertical-align:top;width:77.75pt;"><strong>测试步骤</strong></td><td style="border-bottom:1.0pt solid windowtext;border-left-style:none;border-right:1.0pt solid windowtext;border-top-style:none;padding:0cm 5.4pt;vertical-align:top;width:397.8pt;"><p style="margin-left:18.0pt;">1.        登录容器云平台</p><p style="margin-left:18.0pt;">2.        点击左侧权限树中组织管理功能</p></td></tr><tr><td style="border-bottom-style:solid;border-color:windowtext;border-left-style:solid;border-right-style:solid;border-top-style:none;border-width:1.0pt;padding:0cm 5.4pt;vertical-align:top;width:77.75pt;"><strong>预期结果</strong></td><td style="border-bottom:1.0pt solid windowtext;border-left-style:none;border-right:1.0pt solid windowtext;border-top-style:none;padding:0cm 5.4pt;vertical-align:top;width:397.8pt;"><p style="margin-left:18.0pt;">1.        可以看到当前系统中存在的全部组织</p></td></tr><tr><td style="border-bottom-style:solid;border-color:windowtext;border-left-style:solid;border-right-style:solid;border-top-style:none;border-width:1.0pt;padding:0cm 5.4pt;width:77.75pt;"><strong>预览图</strong></td><td style="border-bottom:1.0pt solid windowtext;border-left-style:none;border-right:1.0pt solid windowtext;border-top-style:none;padding:0cm 5.4pt;vertical-align:top;width:397.8pt;"><figure class="image"><img src="/xae-api/xae-images/2023-10-13/be8a6f5d-8359-41d4-925a-226aa007138a.png"></figure></td></tr><tr><td style="border-bottom-style:solid;border-color:windowtext;border-left-style:solid;border-right-style:solid;border-top-style:none;border-width:1.0pt;padding:0cm 5.4pt;vertical-align:top;width:77.75pt;"><strong>备注</strong></td><td style="border-bottom:1.0pt solid windowtext;border-left-style:none;border-right:1.0pt solid windowtext;border-top-style:none;padding:0cm 5.4pt;vertical-align:top;width:397.8pt;"><span style="color:#ED7D31;"><strong>只有admin用户可查看全部组织</strong></span></td></tr></tbody></table></figure><p> </p>`
	ip := "192.168.14.72"

	// 使用HTML解析器解析HTML内容
	doc, err := html.Parse(strings.NewReader(htmlContent))
	if err != nil {
		fmt.Println(err)
		return
	}

	// 定义一个函数，用于递归遍历HTML节点
	var modifySrc func(*html.Node)
	modifySrc = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "img" {
			// 查找并修改img标签的src属性
			for i, attr := range n.Attr {
				if attr.Key == "src" {
					n.Attr[i].Val = fmt.Sprintf("http:%s:9527", ip) + n.Attr[i].Val
				}
			}
		}

		// 递归遍历子节点
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			modifySrc(c)
		}
	}

	// 调用函数开始遍历和修改HTML节点
	modifySrc(doc)

	// 生成修改后的HTML文本
	var result bytes.Buffer
	html.Render(&result, doc)
	modifiedHTML := result.String()

	// 打印修改后的HTML内容
	fmt.Println(modifiedHTML)
}
