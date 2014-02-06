package main

import (
	"fmt"
	"reflect"
)

type AbstractFactory interface {
	CreateProductA() AbstractProductA
	CreateProductB() AbstractProductB
}

type ConcreteFactory1 struct{}

func (f *ConcreteFactory1) CreateProductA() AbstractProductA {
	return new(ProductA1)
}

func (f *ConcreteFactory1) CreateProductB() AbstractProductB {
	return new(ProductB1)
}

type ConcreteFactory2 struct{}

func (f *ConcreteFactory2) CreateProductA() AbstractProductA {
	return new(ProductA2)
}

func (f *ConcreteFactory2) CreateProductB() AbstractProductB {
	return new(ProductB2)
}

type AbstractProductA interface{}

type AbstractProductB interface{}

type ProductA1 struct{}

type ProductA2 struct{}

type ProductB1 struct{}

type ProductB2 struct{}

func main() {
	factory1 := new(ConcreteFactory1)
	productA1 := factory1.CreateProductA()
	productB1 := factory1.CreateProductB()
	fmt.Println(reflect.TypeOf(productA1))
	fmt.Println(reflect.TypeOf(productB1))
	factory2 := new(ConcreteFactory2)
	productA2 := factory2.CreateProductA()
	productB2 := factory2.CreateProductB()
	fmt.Println(reflect.TypeOf(productA2))
	fmt.Println(reflect.TypeOf(productB2))
}
