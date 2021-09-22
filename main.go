package main

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type myError struct {
	err string
	num int
}

func (e *myError) Error() string {
	return fmt.Sprintf("Rand number: %d Error: %s", e.num, e.err)
}

func main() {
	var wg sync.WaitGroup
	pls := time.Now()
	var b = 20 // capacity
	var a = make([]int, 0, b)
	var d = make([]int, b, b)
	for i := 0; i < b; i++ {
		num := rand.Intn(b) + 1
		a = append(a, chrand(a, num))
	}
	copy(d, a)
	fmt.Println("Random result: ", a)
	wg.Add(2)
	go func() {
		defer wg.Done()
		err := qsort(a)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Quick sorting result: ", a)
		}
		fmt.Println("Quick time: ", time.Since(pls))
	}()
	go func() {
		defer wg.Done()
		err := bubblesort(d)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Bubble sorting result: ", d)
		}
		fmt.Println("Bubble time: ", time.Since(pls))
	}()
	wg.Wait()
}

func chrand(a []int, num int) int {
	for i := 0; i < len(a); i++ {
		if a[i] == num {
			/*e := myError{err: "random repeated on element: " + strconv.Itoa(i), num: num}
			fmt.Println(e.Error())*/
			rand.Seed(time.Now().UnixNano()) //new random pool
			num = rand.Intn(cap(a)) + 1
			i = 0
		}
	}
	return num
}

func qsort(a []int) error {
	if len(a) < 2 {
		return errors.New("Slice have less than 2 elements")
	}
	mid := a[len(a)/2]
	left := 0
	right := len(a) - 1
	for left <= right {
		for a[left] < mid {
			left++
		}
		for a[right] > mid {
			right--
		}
		if left <= right {
			a[left], a[right] = a[right], a[left]
			left++
			right--
		}
	}
	//fmt.Println("After sort: ", a)
	if left <= len(a)-1 {
		qsort(a[:left])
		qsort(a[left:])
	}
	return nil
}

func bubblesort(d []int) error {
	if len(d) < 2 {
		return errors.New("Slice have less than 2 elements")
	}
	for i := 0; i < len(d)-1; i++ {
		for j := 0; j < len(d)-1; j++ {
			if d[j] > d[j+1] {
				d[j], d[j+1] = d[j+1], d[j]
				j--
			}
		}
	}
	return nil
}
