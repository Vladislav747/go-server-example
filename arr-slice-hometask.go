package main

import ("fmt")

const pi = 3.14

func main(){
	var phrase = "Go Массивы и Слайсы Домашнее задание"
	fmt.Println(phrase);

	// Возвращает количество скопированных элементов, 
	// которое будет минимумом len(dst) и len(src). 
	// Результат не зависит от того, перекрываются ли аргументы.

	//Встроенная функция copy копирует элементы в целевой срез dst из исходного среза src.


	//func copy(dst, src []Type) int

	var ex = make([]int, 3)
	n := copy(ex, []int{0, 1, 2, 3}) 
	fmt.Printf("%d\n", n)
	// n == 3, ex == []int{0, 1, 2}



	// make используется для создания слайсов
	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)

	/*
		Как бы вы написали функцию, которая переворачивает все элементы массива или
		среза местами?
	*/
	s := []int{5, 2, 6, 3, 1, 4}

    

    fmt.Println(s)
	reverseArr(s);
	fmt.Println(s)

	/* COPY SLICE*/
	//https://github.com/golang/go/wiki/SliceTricks
	// b := make([]int, len(a))
	
	// copy(b, a)

	// These two are often a little slower than the above one,
	// but they would be more efficient if there are more
	// elements to be appended to b after copying.
	// b = append([]T(nil), a...)
	// b = append(a[:0:0], a...)

	// This one-line implementation is equivalent to the above
	// two-line make+copy implementation logically. But it is
	// actually a bit slower (as of Go toolchain v1.16).
	
	//b = append(make([]T, 0, len(a)), a...)
}

/* Перевернуть массив */ 
func reverseArr(arr []int) []int  {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
        arr[i], arr[j] = arr[j], arr[i]
    }
	return arr
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
