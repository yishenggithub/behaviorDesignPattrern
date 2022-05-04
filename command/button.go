package main

type buttom struct {
	command command
}

func (b *buttom) press() {
	b.command.execute()
}
