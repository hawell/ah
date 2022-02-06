package handlers

import (
	"ah/database"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Address struct {
	Lat float64 `json:"lat" binding:"required"`
	Long float64 `json:"long" binding:"required"`
}

type Provider struct {
	Name string `json:"name"`
	Experience []string `json:"experience"`
	Address Address `json:"address"`
	OperatingRadius float64 `json:"operating_radius"`
	Rating float64 `json:"rating"`
}

type CustomerRequest struct {
	Material string `json:"material" binding:"required,oneof=wood carpet tile"`
	Address Address `json:"address" binding:"required"`
	Area float64 `json:"area" binding:"required,gt=0"`
	PhoneNumber string `json:"phone_number" binding:"required"`
}

func GetProviders(ctx *gin.Context) {
	var req CustomerRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, "binding request failed", err)
		return
	}
	db, exists := ctx.Get("db")
	if !exists {
		ErrorResponse(ctx, http.StatusInternalServerError, "storage instance is not present", nil)
		return
	}

	var (
		material database.FloorMaterial
		location database.Address
	)
	switch req.Material {
	case "wood":material = database.FloorWood
	case "carpet":material = database.FloorCarpet
	case "tile":material = database.FloorTile
	default:
		ErrorResponse(ctx, http.StatusBadRequest, "floor material is not supported", nil)
		return
	}
	location.Lat = req.Address.Lat
	location.Long = req.Address.Long
	storage := db.(Storage)
	dbProviders, err := storage.GetProviders(material, location)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, "db error", err)
		return
	}
	resp := []Provider{}
	for _, dbProvider := range dbProviders {
		provider := Provider{
			Name:            dbProvider.Name,
			Address:         Address{
				Lat: dbProvider.Address.Lat,
				Long: dbProvider.Address.Long,
			},
			OperatingRadius: dbProvider.Radius,
			Rating:          dbProvider.Rating,
		}
		if dbProvider.Wood {
			provider.Experience = append(provider.Experience, "wood")
		}
		if dbProvider.Carpet {
			provider.Experience = append(provider.Experience, "carpet")
		}
		if dbProvider.Tile {
			provider.Experience = append(provider.Experience, "tile")
		}
		resp = append(resp, provider)
	}

	SuccessResponse(ctx, http.StatusOK, "list of providers", resp)
}
