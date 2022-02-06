package handlers

import "ah/database"

type Storage interface {
	GetProviders(material database.FloorMaterial, location database.Address) ([]database.Provider, error)
}
