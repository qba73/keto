package main

import (
	"fmt"
	"time"
)


type Sugar struct {
	Date time.Time
	Record float64
	Unit string
}

type SpO2 struct {
	Date time.Time
	Record float64
	Unit string
}

type Preassure struct {
	Date time.Time
	Record float64
	Unit string
}

type HR struct {
	Date time.Time
	Record float64
	Unit string
}

func main() {
	fmt.Println("Add data...")
}