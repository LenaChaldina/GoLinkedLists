package main

import (
	"fmt"
)

type list struct {
	a    int
	next *list
}

func find(l **list, elt int) **list {
	for *l != nil {
		if (*l).a == elt {
			return l
		}
		l = &(*l).next
	}
	return nil
}

func f(l **list) int {
	k := 0
	for *l != nil {
		k += 1
		fmt.Println((*l).a)
		l = &((*l).next)
	}
	return k
}

func remove(l **list) {
	*l = (*l).next
}

func append(l **list, elt int) {
	for *l != nil {
		l = &(*l).next

	}
	l4 := new(list)
	l4.a = elt
	l4.next = nil
	*l = (l4)
}

func prepend(l **list, elt int) {
	l5 := new(list)
	l5.a = elt
	l5.next = *l
	*l = l5

}

func extend(l1 **list, l2 **list) {
	for *l1 != nil {
		l1 = &((*l1).next)
	}
	*l1 = *l2

}

func main() {
	var l1 *list
	append(&l1, 1)
	append(&l1, 5)
	append(&l1, 3)

	var l2 *list
	append(&l2, 11)
	append(&l2, 12)
	append(&l2, 13)

	extend(&l1, &l2)

	//fmt.Println(*find(&l1, 5))

	//remove(find(&l1, 5))
	//append(&l1, 4)
	//prepend(&l1, 0)
	fmt.Println(f(&l1))

}
