package main

import (
	"encoding/json"
	"fmt"
	"time"
)

func ArrayDT() {
	// fixed sized and homogeneous type

	var ar [5]int // array of 5 integers, ar will have 0,0,0,0,0
	for _, item := range ar {
		fmt.Printf("%d ", item)
	}
	fmt.Println()

	// array literal
	fmt.Println("string literal, fixed length:")
	var ap [5]int = [5]int{4, 5, 6, 7}
	for _, item := range ap {
		fmt.Printf("%d ", item)
	}
	fmt.Println()

	// array literal, 2nd way
	var al = [5]int{42, 55, 68, 79}
	for _, item := range al {
		fmt.Printf("%d ", item)
	}
	fmt.Println()

	// array literal with ellipsis. the array length is determined by the number of initializers
	var ab = [...]int{45, 78, 99}
	fmt.Println("string literal, dynamic length:")
	for _, item := range ab {
		fmt.Printf("%d ", item)
	}
	fmt.Println()

	// Arrays in go are passed by value and NOT reference
	/* modifyArray(ab)
	func modifyArray(u [3]int){
		u[0] = 55
	}
	*/

	// Arrays : pass as Reference
	/* modifyArray(&ab)
	func modifyArray(u *[3]int){
		u[0] = 55
	}
	*/
}

func SliceDT() {
	// variable sized and homogeneous type
	// slice looks just like arrays without the size

	// initialise an int slice. 3 methods:

	var _ []int = []int{8, 5, 423, 2}
	var _ = []int{8, 5, 423, 2}
	ap2 := []int{8, 5, 423, 2}

	// slice are passed by reference
	// reverse will reverse the actual slice
	reverse(ap2)

	for _, item := range ap2 {
		fmt.Printf("%d ", item)
	}
	fmt.Println()

	// Length vs Capacity of a Slice
	fmt.Println(len(ap2))
	fmt.Println(cap(ap2))

	// Append elements to Slice

	ap2 = append(ap2, 56)
	fmt.Println("append 56 to my slice")
	for _, item := range ap2 {
		fmt.Printf("%d ", item)
	}
	fmt.Println()

	ap2 = append([]int{67}, ap2...)
	fmt.Println("append 67 to the start of a slice")
	for _, item := range ap2 {
		fmt.Printf("%d ", item)
	}
	fmt.Println()

	// In-Place Slice Techniques

}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func MapDT() {
	// hash tables

	ages := make(map[string]int)
	var _ map[string]int = make(map[string]int)

	// nil map
	var _ map[string]int
	/*
		Attempting to use the nil map, such as trying to assign a value to a key or access a key, will result in a runtime
		panic because a nil map does not have an underlying data structure to store key-value pairs
	*/

	// map literal
	ages = map[string]int{
		"alice": 56,
		"bob":   58,
	}

	// add to map

	ages["sakshi"] = 21
	ages["rachit"] = 24

	for key, value := range ages {
		fmt.Printf("key: %s, value: %d\n", key, value)
	}

	// delete

	delete(ages, "bob")
	fmt.Println("delete bob")

	for key, value := range ages {
		fmt.Printf("key: %s, value: %d\n", key, value)
	}

	// size of a map
	fmt.Printf("size of the map ages: %d/n", len(ages))

	// check value in Map
	_, ok := ages["mayank"]
	if !ok {
		fmt.Println("mayank does not exist in ages")
	}
}

func SetDT() {
	/*
		Go does not have a built-in "set" data structure.
		However, you can implement a set data structure in Go using a map with keys as the
		set elements.

	*/

	var set = make(map[string]bool)

	// add element to set
	var name = "gusain"
	set[name] = true
	fmt.Println(set)

	// remove element from the set
	delete(set, name)
	fmt.Println(set)

	// check if the set contains an element or not
	_, exists := set[name]
	if !exists {
		fmt.Printf("set does not contains: %v /n", name)
	}
}

func StructDT() {

	// type definition:
	type employee struct {
		ID       int
		Name     string
		Address  string
		DoB      time.Time
		Position string
	}

	var emma employee
	emma.ID = 18923479

	fmt.Println(emma)

	// A named struct type Tree cant declare a field of the same type tree
	// have to use pointers
	type tree struct {
		data                  int
		leftChild, rightChild *tree
	}

	// struct literals
	verender := employee{
		ID:       12312,
		Name:     "Verender",
		Address:  "asdfasd",
		DoB:      time.Date(1992, time.August, 12, 0, 0, 0, 0, time.UTC),
		Position: "Director",
	}

	yuvi := employee{
		12312,
		"Yuvraj",
		"asdfasd",
		time.Date(1992, time.August, 12, 0, 0, 0, 0, time.UTC),
		"Sr. Director",
	}

	fmt.Printf("struct literals : %v \n", verender)
	fmt.Printf("struct literals : %v \n", yuvi)
}

func JsonDT() {
	type Movie struct {
		Title  string
		Year   int  `jons:"released"`
		Color  bool `json:"color,omitempty"`
		Actors []string
	}

	var jungleBook = Movie{
		"The Jungle Book",
		1995, true,
		[]string{"asa", "bbb"},
	}

	// converting a Go data structure to JSON is called Marshalling

	//data, err := json.Marshal(jungleBook)
	data, err := json.MarshalIndent(jungleBook, "", "\t")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))

	// some common field tags:
	// 1. json:"fieldName": This tag specifies the name of the JSON field when marshaling and unmarshaling
	// 2. omitempty option tells the JSON encoder to omit the field from the JSON output if the field's value is the
	//zero value for its type.
	// 3. json:"-": A hyphen (-) as the field name tells the JSON encoder to skip this field during both marshaling
	// and unmarshalling
	// 4. json:",string": This tag is used to marshal a numeric field as a JSON string
	var jsonM Movie
	if err := json.Unmarshal([]byte(data), &jsonM); err != nil {
		fmt.Print(err)
	}

	fmt.Println(jsonM)

}
