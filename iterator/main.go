package main

import "fmt"

// 迭代器是一种行为设计模式，让你能在不暴露复杂数据结构内部细节的情况下遍历其中所有的元素。
// 在迭代器的帮助下， 客户端可以用一个迭代器接口以相似的方式遍历不同集合中的元素。

func main() {
	user1 := &user{
		name: "a",
		age:  30,
	}

	user2 := &user{
		name: "b",
		age:  20,
	}

	userCollection := &userCollection{
		users: []*user{user1, user2},
	}

	iterator := userCollection.createIterator()

	for iterator.hasNext() {
		user := iterator.getNext()
		fmt.Println("User is %+v\n", user)
	}
}
