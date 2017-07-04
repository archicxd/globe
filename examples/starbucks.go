// +build starbucks

package main

import (
	"encoding/json"
	"image/color"
	"io/ioutil"

	"github.com/mmcloughlin/globe"
)

type CoffeeShop struct {
	Lat float64 `json:"latitude"`
	Lng float64 `json:"longitude"`
}

func LoadCoffeeShops(filename string) ([]CoffeeShop, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	shops := []CoffeeShop{}
	err = json.Unmarshal(data, &shops)
	if err != nil {
		return nil, err
	}

	return shops, nil
}

func Build() (*globe.Globe, error) {
	// BEGIN
	shops, err := LoadCoffeeShops("./starbucks.json")
	if err != nil {
		return nil, err
	}

	green := color.NRGBA{0x00, 0x64, 0x3c, 192}
	g := globe.New()
	g.DrawGraticule(10.0)
	for _, s := range shops {
		g.DrawDot(s.Lat, s.Lng, 0.05, globe.Color(green))
	}
	g.CenterOn(40.645423, -73.903879)
	// END

	return g, nil
}
