package database

type FloorMaterial string

const (
	FloorWood   FloorMaterial = "wood"
	FloorCarpet FloorMaterial = "carpet"
	FloorTile   FloorMaterial = "tile"
)

type Address struct {
	Lat  float64
	Long float64
}

type ID int64

type Provider struct {
	ID      ID
	Name    string
	Address Address
	Radius  float64
	Rating  float64
	Wood    bool
	Carpet  bool
	Tile    bool
}
