package main

import (
	"fmt"
	f "github.com/izqui/functional"
	"github.com/izqui/helpers"
	"time"
)

type Race struct {
	Time int
	Cars []Car
}

type Car struct {
	Position int
}

func draw(r Race) {

	drawing := f.DoMap(func(car Car) string {

		s := ""
		for i := 0; i < car.Position; i++ {

			s += "-"
		}

		return s
	}, r.Cars).([]string)

	fmt.Println("......................")
	f.DoMap(func(a string) interface{} {

		fmt.Println(a)
		return a
	}, drawing)
}

func raceLoop(r Race) Race {

	r.Time -= 1
	r.Cars = f.DoMap(func(car Car) Car {

		if helpers.RandomInt(0, 3) == 0 {

			car.Position += 1
		}

		return car
	}, r.Cars).([]Car)

	return r
}

func race(r Race) {

	if r.Time > 0 {

		draw(r)
		time.Sleep(100 * time.Millisecond)
		race(raceLoop(r))

	} else {

		fmt.Println("Race finished")
	}
}

func main() {

	r := Race{Time: 100, Cars: []Car{Car{}, Car{}, Car{}}}
	race(r)
}
