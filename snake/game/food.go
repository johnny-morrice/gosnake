package game

import (
	"iter"
	"maps"
	"math/rand"

	"github.com/johnny-morrice/gosnake/snake/layer"
	"github.com/johnny-morrice/gosnake/snake/tiles"
)

type Food struct {
	Food     map[Point]FoodItem
	Geometry Geometry
}

type FoodItem struct {
	Nutrition int
	TileType  int
}

func NewFood(geometry Geometry) *Food {
	return &Food{
		Food:     make(map[Point]FoodItem),
		Geometry: geometry,
	}
}

func (food *Food) AddFood(occupied iter.Seq[Point]) {
	myOccupied := maps.Keys(food.Food)
	newPoint := food.Geometry.RandomPoint(myOccupied, occupied)
	// One in 10 chance of being a big food item, which gives 3 nutrition instead of 1.
	foodItem := FoodItem{
		Nutrition: 1,
		TileType:  tiles.SmallFood,
	}
	if rand.Intn(10) == 0 {
		foodItem.Nutrition = 3
		foodItem.TileType = tiles.LargeFood
	}
	food.Food[newPoint] = foodItem
}

func (food *Food) Eaten(point Point) {
	delete(food.Food, point)
}

func (food *Food) Render() layer.Layer {
	myTiles := make([]layer.Tile, 0, len(food.Food))
	for point, item := range food.Food {
		myTiles = append(myTiles, layer.Tile{
			X:     point.X,
			Y:     point.Y,
			Type:  item.TileType,
			Color: "red",
		})
	}
	return layer.Layer{
		Tiles: myTiles,
	}
}
