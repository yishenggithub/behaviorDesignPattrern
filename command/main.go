package main

// 命令是一种行为设计模式， 它可将请求或简单操作转换为一个对象。
// 此类转换让你能够延迟进行或远程执行请求， 还可将其放入队列中。

func main() {
	tv := &tv{}

	onCommand := &onCommand{
		device: tv,
	}

	offCommand := &offCommand{
		device: tv,
	}

	onButton := &buttom{
		command: onCommand,
	}
	onButton.press()

	offButton := &buttom{
		command: offCommand,
	}
	offButton.press()
}
