package database

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/ilyakaznacheev/cleanenv"
)

type DataBase struct {
	db *sql.DB
}

var (
	ErrDuplicateEntry = errors.New("duplicate entry")
	ErrNotFound       = errors.New("not found")
	ErrInvalid        = errors.New("invalid operation")
)

func parseError(err error) error {
	var mysqlErr *mysql.MySQLError
	if errors.Is(err, sql.ErrNoRows) {
		return ErrNotFound
	}
	if errors.As(err, &mysqlErr) {
		switch mysqlErr.Number {
		case 1062:
			return ErrDuplicateEntry
		case 1452:
			return ErrInvalid
		default:
			return err
		}
	}
	return err
}

func Connect() (*DataBase, error) {
	var config Config
	err := cleanenv.ReadEnv(&config)
	if err != nil {
		return nil, err
	}

	db, err := sql.Open("mysql", config.ConnectionString)
	if err != nil {
		return nil, parseError(err)
	}
	return &DataBase{db}, nil
}

func (db *DataBase) Clear() error {
	_, err := db.db.Exec("delete from Provider")
	return parseError(err)
}

func (db *DataBase) GetProviders(material FloorMaterial, location Address) ([]Provider, error) {
	query := "select p.Id, p.Name, ST_X(p.Address) AS Latitude, ST_Y(p.Address) AS Longitude, p.Radius, p.Rating, p.Wood, p.Carpet, p.Tile, st_distance_sphere(point(?, ?), p.Address) as dist from Provider p"
	filter := " where "
	switch material {
	case FloorWood: filter += "p.Wood = 1"
	case FloorCarpet: filter += "p.Carpet = 1"
	case FloorTile: filter += "p.Tile = 1"
	default:
		filter = ""
	}
	limitAndOrder := " having dist < Radius order by Rating desc"
	rows, err := db.db.Query(query + filter + limitAndOrder, location.Lat, location.Long)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	defer func() { _ = rows.Close() }()
	res := []Provider{}
	for rows.Next() {
		var (
			item Provider
			distance float64
		)
		err := rows.Scan(&item.ID, &item.Name, &item.Address.Lat, &item.Address.Long, &item.Radius, &item.Rating, &item.Wood, &item.Carpet, &item.Tile, &distance)
		if err != nil {
			return nil, err
		}
		res = append(res, item)
	}
	return res, nil
}

func (db *DataBase) AddProvider(p Provider) (ID, error) {
	pointStr := fmt.Sprintf("POINT(%f %f)", p.Address.Lat, p.Address.Long)
	query := `insert into Provider values(NULL, ?, ST_GeomFromText(?), ?, ?, ?, ?, ?)`
	result, err := db.db.Exec(query, p.Name, pointStr, p.Radius, p.Rating, p.Wood, p.Carpet, p.Tile)

	if err != nil {
		return 0, parseError(err)
	}
	id, err := result.LastInsertId()
	return ID(id), parseError(err)
}

