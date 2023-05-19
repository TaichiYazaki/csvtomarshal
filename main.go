package main

import (
	"fmt"
	"os"

	"github.com/gocarina/gocsv"
)

type Fruit struct {
	Id   string `csv:"Id"`
	Name string `csv:"Name"`
}

func main() {

	f, err := os.OpenFile("fruits.csv", os.O_RDWR, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	var fruit []*Fruit
	// interface{}(空のインターフェース)はどんな型でも代入可能
	gocsv.UnmarshalFile(f, &fruit)

	for _, v := range fruit {
		fmt.Println(v.Name)
	}
}
