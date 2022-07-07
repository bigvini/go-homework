package main

import (
	"fmt"
)

func main() {

	AnimalsOnFerm := []Animals{
		Dog{weight: 25, eatPerKilo: 2},
		Elephant{weight: 2000, eatPerKilo: 50},
		Cow{weight: 400, eatPerKilo: 25},
		Cat{7, 7},
	}

	animalEatPerMounts(AnimalsOnFerm)

}

func animalEatPerMounts(animal []Animals) {
	foodTotal := 0
	for _, value := range animal {
		foodTotal += value.getEatForMonth()
		fmt.Printf("\nThe %v eats of food per months %v kg and and weighs %v kg", value.getName(), value.getEatForMonth(), value.getWeight())
	}
	fmt.Printf("\nThe total amount of feed for the farm is %v kg", foodTotal)

}

// --------------------------------------

type Animals interface {
	getterFoodPerMounts
	getterName
	getterWeight
}

type getterWeight interface {
	getWeight() int
}
type getterFoodPerMounts interface {
	getEatForMonth() int
}
type getterName interface {
	getName() string
}

type Dog struct {
	weight     int
	eatPerKilo int
}

func (d Dog) getEatForMonth() int {
	eatFor := d.weight * d.eatPerKilo
	return eatFor
}
func (d Dog) getName() string {
	return "Dog"
}

func (d Dog) getWeight() int {
	return d.weight
}

type Cat struct {
	weight     int
	eatPerKilo int
}

func (d Cat) getWeight() int {
	return d.weight
}
func (d Cat) getName() string {

	return "Cat"
}

func (d Cat) getEatForMonth() int {
	eatFor := d.weight * d.eatPerKilo
	return eatFor
}

type Cow struct {
	weight     int
	eatPerKilo int
}

func (d Cow) getWeight() int {
	return d.weight
}
func (d Cow) getName() string {

	return "Cow"
}

func (d Cow) getEatForMonth() int {
	eatFor := d.weight * d.eatPerKilo
	return eatFor
}

type Elephant struct {
	weight     int
	eatPerKilo int
}

func (d Elephant) getWeight() int {
	return d.weight
}

func (d Elephant) getEatForMonth() int {
	eatFor := d.weight * d.eatPerKilo
	return eatFor
}
func (d Elephant) getName() string {

	return "Elephant"
}
