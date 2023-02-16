package prototype

import (
	"errors"
	"fmt"
)

const (
	White = iota + 1
	Black
	Blue
)

type ItemInfoGetter interface {
	GetInfo() string
}

type ShirtCloner interface {
	GetClone(s int) (ItemInfoGetter, error)
}

func GetShirtsCloner() ShirtCloner {
	return &ShirtsCache{}
}

type ShirtsCache struct {
}

func (s *ShirtsCache) GetClone(c int) (ItemInfoGetter, error) {
	switch c {
	case White:
		newItem := *whitePrototype
		return &newItem, nil
	case Black:
		newItem := *blackPrototype
		return &newItem, nil
	case Blue:
		newItem := *bluePrototype
		return &newItem, nil
	default:
		return nil, errors.New("shirt model not recognized")
	}
}

type ShirtColor byte

type Shirt struct {
	Price float32
	SKU   string
	Color ShirtColor
}

func (s *Shirt) GetInfo() string {
	return fmt.Sprintf("Shirt with SKU '%s' and Color id %d that costs %f\n", s.SKU, s.Color, s.Price)
}

func (s *Shirt) GetPrice() float32 {
	return 0.0
}

func GetShirtCloner() ShirtCloner {
	return &ShirtsCache{}
}

var whitePrototype = &Shirt{
	Price: 15.00,
	SKU:   "empty",
	Color: White,
}

var blackPrototype = &Shirt{
	Price: 16.00,
	SKU:   "empty",
	Color: Black,
}

var bluePrototype = &Shirt{
	Price: 17.00,
	SKU:   "empty",
	Color: Blue,
}
