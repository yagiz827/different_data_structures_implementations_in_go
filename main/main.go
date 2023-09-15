package main

import (
	"fmt"
)

type Fields struct {
	Id       int
	Name     string
	Includes []User
}

type User struct {
	Id   int
	Name string
}

type shape interface {
	area() float32
	perimete() float32
}
type square struct {
	x   float32
	rrt string
}
type rectangle struct {
	x float32
	y float32
}

func (s square) area() float32 {
	return s.x * s.x
}
func (s square) perimete() float32 {
	return s.x + s.x
}
func (s square) asd() string {
	return "sdsddd"
}
func (s rectangle) perimete() float32 {
	return s.x + s.x
}
func (s rectangle) area() float32 {
	return s.x * s.y
}
func abcc(s shape) {
	s.area()
	switch a := s.(type) {
	case square:
		fmt.Println(a.rrt)
		fmt.Println(a.asd())
	case rectangle:
		fmt.Println(a.y)
	}
}

var x int = 0

func (user *User) createUser(name string) User {
	user.Name = name
	user.Id = x
	x++
	return *user
}

type node struct {
	value int
	left  *node
	right *node
}

func (Node *node) insertIntoTree(n int) {
	var newn = node{
		value: n,
		left:  nil,
		right: nil,
	}
	if Node != nil {
		if n > Node.value {
			if Node.right == nil {
				Node.right = &newn
			} else {
				Node.right.insertIntoTree(n)
			}
		} else {
			if Node.left == nil {
				Node.left = &newn
			} else {
				Node.left.insertIntoTree(n)
			}
		}
	}
}

func (Node *node) PreOrder() {
	if Node != nil {
		fmt.Println(Node.value)
		Node.left.PreOrder()
		Node.right.PreOrder()
	}
}

type linked struct {
	value int
	next  *linked
}

func (ty *linked) insert_into_linked(data int) {
	var s = linked{value: data, next: nil}
	ty.next = &s
}

type intStack struct {
	stack []int
}

func (stack *intStack) pushBack(data int) {
	//stack.stack==(*stack).stack automatically hadnled
	(*stack).stack = append(stack.stack, data)
}

func (stack *intStack) pop() int {
	var lenStack int = len(stack.stack)
	var back int = stack.stack[lenStack-1]
	stack.stack = stack.stack[0 : lenStack-1]
	return back
}

func genpushBack[T any](slice *[]T, data T) {
	//stack.stack==(*stack).stack automatically handled
	*slice = append(*slice, data)
}

func genpop[T any](slice *[]T) T {
	var lenStack int = len(*slice)
	back := (*slice)[lenStack-1]
	*slice = (*slice)[0 : lenStack-1]
	return back
}
func main() {
	g := User{Name: "Elif", Id: 0}
	fmt.Println(g.Name)

	ttu := User{}
	ttu.createUser("ahmet")
	tu := User{}
	tu.createUser("tyu")
	abc := User{}
	abc.createUser("tyu")
	gy := User{Name: "trtElif", Id: 2}
	fmt.Println(g.Name)
	var tttyy = linked{value: 1, next: nil}
	tttyy.insert_into_linked(2)
	var y *linked = &tttyy
	for y != nil {
		fmt.Println(y.value)
		y = y.next
	}

	var st = intStack{make([]int, 0)}
	fmt.Println(st.stack)
	st.pushBack(1)
	st.pushBack(2)
	fmt.Println(st.stack)
	fmt.Println(st.pop())
	fmt.Println(st.stack)
	var u = []string{"q", "e", "r"}
	fmt.Println(u)
	genpushBack(&u, "t")
	var mmm = [][]string{}
	genpushBack(&mmm, u)
	fmt.Println(mmm)
	fmt.Println(u)
	fmt.Println(genpop(&u))
	fmt.Println(u)
	intSlice := []int{1, 2, 3}
	fmt.Println(intSlice)
	fmt.Println(genpop(&intSlice))
	fmt.Println(intSlice)
	var hh = []User{}
	genpushBack(&hh, g)
	genpushBack(&hh, gy)
	fmt.Println(hh)
	//fmt.Println(genpop(&hh))
	genpushBack(&hh, ttu)
	fmt.Println(hh)
	genpushBack(&hh, tu)
	fmt.Println((hh))
	genpushBack(&hh, abc)
	fmt.Println((hh))
	fmt.Println(hh[2].Id)
	for _, user := range hh {
		fmt.Println(user.Id)
		fmt.Println(user.Name)
	}
	var t = node{
		value: 5,
		left:  nil,
		right: nil,
	}
	fmt.Println(t)

	t.insertIntoTree(3)

	t.insertIntoTree(6)

	t.insertIntoTree(2)
	t.insertIntoTree(12)
	t.insertIntoTree(23)
	t.insertIntoTree(0)
	t.insertIntoTree(-2)
	t.insertIntoTree(8)
	t.insertIntoTree(4)
	t.PreOrder()

	var po = square{x: 4}
	var pope = rectangle{x: 4, y: 8}
	abcc(po)
	abcc(pope)

}
