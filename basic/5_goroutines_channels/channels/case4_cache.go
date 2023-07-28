package main

func mirroredQuery() string {
	//创建一个带缓存的channel，容量是3
	//向缓存Channel的发送操作就是向内部缓存队列的尾部插入元素，接收操作则是从队列的头部删除元素。
	//如果内部缓存队列是满的，那么发送操作将阻塞直到因另一个goroutine执行接收操作而释放了新的队列空间。相反，如果channel是空的，接收操作将阻塞直到有另一个goroutine执行发送操作而向队列插入元素。
	responses := make(chan string, 3)
	go func() { responses <- request("asia.gopl.io") }()
	go func() { responses <- request("europe.gopl.io") }()
	go func() { responses <- request("americas.gopl.io") }()
	//获取最快响应的镜像
	return <-responses // return the quickest response
}

func request(hostname string) (response string) {
	/* ... */
	response = ""
	return
}

