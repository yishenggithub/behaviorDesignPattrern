package main

import "fmt"

type item struct {
	observerList []observer
	name         string
	inStock      bool
}

func newItem(name string) *item {
	return &item{
		name: name,
	}
}

func (i *item) updateAvailability() {
	fmt.Println("Item %s is now in stock\n", i.name)
	i.inStock = true
	i.notifyAll()
}

func (i *item) register(o observer) {
	i.observerList = append(i.observerList, o)
}

func (i *item) deregister(o observer) {
	i.observerList = removeFromSlice(i.observerList, o)
}

func (i *item) notifyAll() {
	for _, observer := range i.observerList {
		observer.update(i.name)
	}
}

func removeFromSlice(observerList []observer, observerToRemove observer) []observer {
	observerListLength := len(observerList)
	for i, observer := range observerList {
		if observerToRemove.getID() == observer.getID() { // 交换移除法？
			observerList[observerListLength-1], observerList[i] = observerList[i], observerList[observerListLength-1]
			return observerList[:observerListLength-1]
		}
	}
	return observerList
}
