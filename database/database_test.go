package database

import (
	. "github.com/onsi/gomega"
	"testing"
)

var (
	db *DataBase
)

func TestMain(m *testing.M) {
	var err error
	db, err = Connect()
	if err != nil {
		panic(err)
	}
	m.Run()
}

func PopulateDB(providers []Provider) {
	err := db.Clear()
	Expect(err).To(BeNil())
	for i := range providers {
		id, err := db.AddProvider(providers[i])
		Expect(err).To(BeNil())
		providers[i].ID = id
	}
}

func TestMatchingMaterial(t *testing.T) {
	RegisterTestingT(t)
	providers := []Provider{
		{
			Name: "p0",
			Address: Address{
				Lat:  -26,
				Long: 40,
			},
			Radius: 10,
			Rating: 5,
			Wood:   false,
			Carpet: false,
			Tile:   false,
		},
		{
			Name: "p1",
			Address: Address{
				Lat:  -26,
				Long: 40,
			},
			Radius: 10,
			Rating: 5,
			Wood:   true,
			Carpet: false,
			Tile:   false,
		},
		{
			Name: "p2",
			Address: Address{
				Lat:  -26,
				Long: 40,
			},
			Radius: 10,
			Rating: 5,
			Wood:   false,
			Carpet: true,
			Tile:   false,
		},
		{
			Name: "p3",
			Address: Address{
				Lat:  -26,
				Long: 40,
			},
			Radius: 10,
			Rating: 5,
			Wood:   false,
			Carpet: false,
			Tile:   true,
		},
		{
			Name: "p4",
			Address: Address{
				Lat:  -26,
				Long: 40,
			},
			Radius: 10,
			Rating: 5,
			Wood:   true,
			Carpet: true,
			Tile:   false,
		},
		{
			Name: "p5",
			Address: Address{
				Lat:  -26,
				Long: 40,
			},
			Radius: 10,
			Rating: 5,
			Wood:   true,
			Carpet: false,
			Tile:   true,
		},
		{
			Name: "p6",
			Address: Address{
				Lat:  -26,
				Long: 40,
			},
			Radius: 10,
			Rating: 5,
			Wood:   false,
			Carpet: true,
			Tile:   true,
		},
	}
	PopulateDB(providers)
	res, err := db.GetProviders(FloorWood, Address{Lat: -26, Long: 40})
	Expect(err).To(BeNil())
	Expect(res).To(ConsistOf([]Provider{providers[1], providers[4], providers[5]}))
	res, err = db.GetProviders(FloorCarpet, Address{Lat: -26, Long: 40})
	Expect(err).To(BeNil())
	Expect(res).To(ConsistOf([]Provider{providers[2], providers[4], providers[6]}))
	res, err = db.GetProviders(FloorTile, Address{Lat: -26, Long: 40})
	Expect(err).To(BeNil())
	Expect(res).To(ConsistOf([]Provider{providers[3], providers[5], providers[6]}))
}

func TestExcludeOutOfRadius(t *testing.T) {
	RegisterTestingT(t)
	providers := []Provider{
		{
			Name: "p0",
			Address: Address{
				Lat:  -26.66119,
				Long: 40.95858,
			},
			Radius: 10,
			Rating: 5,
			Wood:   true,
			Carpet: true,
			Tile:   true,
		},
		{
			Name: "p1",
			Address: Address{
				Lat:  -26.66129,
				Long: 40.95858,
			},
			Radius: 10,
			Rating: 5,
			Wood:   true,
			Carpet: true,
			Tile:   true,
		},
		{
			Name: "p2",
			Address: Address{
				Lat:  -26.66119,
				Long: 40.95868,
			},
			Radius: 10,
			Rating: 5,
			Wood:   true,
			Carpet: true,
			Tile:   true,
		},
		{
			Name: "p3",
			Address: Address{
				Lat:  -26.66129,
				Long: 40.95868,
			},
			Radius: 10,
			Rating: 5,
			Wood:   true,
			Carpet: true,
			Tile:   true,
		},
	}
	PopulateDB(providers)
	res, err := db.GetProviders(FloorWood, Address{Lat: -26.66119, Long: 40.95858})
	Expect(err).To(BeNil())
	Expect(res).To(ConsistOf([]Provider{providers[0], providers[1]}))
}

func TestOrder(t *testing.T) {
	RegisterTestingT(t)
	providers := []Provider{
		{
			Name: "p0",
			Address: Address{
				Lat:  -26.66129,
				Long: 40.95858,
			},
			Radius: 100,
			Rating: 3.5,
			Wood:   true,
			Carpet: true,
			Tile:   true,
		},
		{
			Name: "p1",
			Address: Address{
				Lat:  -26.66139,
				Long: 40.95858,
			},
			Radius: 100,
			Rating: 3.5,
			Wood:   true,
			Carpet: true,
			Tile:   true,
		},
		{
			Name: "p2",
			Address: Address{
				Lat:  -26.66139,
				Long: 40.95878,
			},
			Radius: 100,
			Rating: 5,
			Wood:   true,
			Carpet: true,
			Tile:   true,
		},
		{
			Name: "p3",
			Address: Address{
				Lat:  -26.66119,
				Long: 40.95858,
			},
			Radius: 10,
			Rating: 3.0,
			Wood:   true,
			Carpet: true,
			Tile:   true,
		},
	}
	PopulateDB(providers)
	res, err := db.GetProviders(FloorWood, Address{Lat: -26.66119, Long: 40.95858})
	Expect(err).To(BeNil())
	Expect(res).To(Equal([]Provider{providers[2], providers[0], providers[1], providers[3]}))
}

func TestMultipleChecks(t *testing.T) {
	RegisterTestingT(t)
	providers := []Provider{
		{
			Name: "p0",
			Address: Address{
				Lat:  -26.66119,
				Long: 40.95858,
			},
			Radius: 2,
			Rating: 3.5,
			Wood:   true,
			Carpet: true,
			Tile:   true,
		},
		{
			Name: "p1",
			Address: Address{
				Lat:  -26.66120,
				Long: 40.95858,
			},
			Radius: 2,
			Rating: 4.5,
			Wood:   false,
			Carpet: true,
			Tile:   true,
		},
		{
			Name: "p2",
			Address: Address{
				Lat:  -26.66116,
				Long: 40.95858,
			},
			Radius: 2,
			Rating: 4.5,
			Wood:   true,
			Carpet: false,
			Tile:   false,
		},
		{
			Name: "p3",
			Address: Address{
				Lat:  -26.66117,
				Long: 40.95858,
			},
			Radius: 2,
			Rating: 4.7,
			Wood:   true,
			Carpet: true,
			Tile:   false,
		},
		{
			Name: "p4",
			Address: Address{
				Lat:  -26.66115,
				Long: 40.95858,
			},
			Radius: 2,
			Rating: 4.5,
			Wood:   true,
			Carpet: false,
			Tile:   false,
		},
		{
			Name: "p5",
			Address: Address{
				Lat:  -26.66118,
				Long: 40.95858,
			},
			Radius: 2,
			Rating: 4.1,
			Wood:   true,
			Carpet: false,
			Tile:   true,
		},
		{
			Name: "p6",
			Address: Address{
				Lat:  -26.66116,
				Long: 40.95858,
			},
			Radius: 2,
			Rating: 4.8,
			Wood:   true,
			Carpet: false,
			Tile:   false,
		},
	}
	PopulateDB(providers)
	res, err := db.GetProviders(FloorWood, Address{Lat: -26.66119, Long: 40.95858})
	Expect(err).To(BeNil())
	Expect(res).To(Equal([]Provider{providers[3], providers[5], providers[0]}))
}
