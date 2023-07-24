// Package links provides a link-extraction function.
package links

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
)

// Extract makes an HTTP GET request to the specified URL, parses
// the response as HTML, and returns the links in the HTML document. 返回一个url页面中的所有链接
func Extract(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}
	//解析html页面 doc类型是*html.Node
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	var links []string
	//把匿名函数赋值给visitNode变量，因为在后面的forEachNode函数里面要不断调用这个匿名函数来判断每一个节点里面是否包含链接节点href
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				//把链接解析成绝对路径
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue // ignore bad URLs
				}
				//当a.Key是href时，把这个链接加入结果集
				links = append(links, link.String())
			}
		}
	}
	forEachNode(doc, visitNode, nil)
	return links, nil
}

// forEachNode 遍历html页面 针对每个结点x，都会调用pre(x)和post(x)。pre和post都是函数参数
// pre和post都是可选的。
// 遍历孩子结点之前，pre被调用
// 遍历孩子结点之后，post被调用
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

func main()  {
	links, _ := Extract("https://google.com")
	for _, link := range links {
		fmt.Println(link)
	}
}
