package handlers

import "ah/database"

// Storage database contract required for handlers
type Storage interface {
	GetProviders(material database.FloorMaterial, location database.Address) ([]database.Provider, error)
}
