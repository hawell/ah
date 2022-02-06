package database

// FloorMaterial material for the floor
type FloorMaterial string

const (
	// FloorWood wood material
	FloorWood FloorMaterial = "wood"
	// FloorCarpet carpet material
	FloorCarpet FloorMaterial = "carpet"
	// FloorTile tile material
	FloorTile FloorMaterial = "tile"
)

// Address is a location on map
type Address struct {
	Lat  float64
	Long float64
}

// ID database ID type
type ID int64

// Provider holds information aboud a provider in db
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
