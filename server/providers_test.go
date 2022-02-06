package server

import (
	"ah/database"
	"ah/server/handlers"
	"encoding/json"
	"errors"
	jsoniter "github.com/json-iterator/go"
	. "github.com/onsi/gomega"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
	"time"
)

type MockDB struct {
	GetProvidersFunc func(material database.FloorMaterial, location database.Address) ([]database.Provider, error)
}

func (db MockDB) GetProviders(material database.FloorMaterial, location database.Address) ([]database.Provider, error) {
	return db.GetProvidersFunc(material, location)
}

var (
	db                      *MockDB
	defaultRequest          handlers.CustomerRequest
)

func TestMain(m *testing.M) {
	db = &MockDB{}
	s, err := NewServer(zap.NewNop(), db)
	if err != nil {
		panic(err)
	}
	go func() {
		_ = s.ListenAndServe()
	}()
	time.Sleep(time.Second)
	m.Run()
}

func TestGetProviders(t *testing.T) {
	dbProviders := []database.Provider{
		{
			ID: 0,
			Name: "p0",
			Address: database.Address{
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
			ID: 1,
			Name: "p1",
			Address: database.Address{
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
			ID: 2,
			Name: "p2",
			Address: database.Address{
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
			ID: 3,
			Name: "p3",
			Address: database.Address{
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
			ID: 4,
			Name: "p4",
			Address: database.Address{
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
			ID: 5,
			Name: "p5",
			Address: database.Address{
				Lat:  -26.66118,
				Long:  40.95858,
			},
			Radius: 2,
			Rating: 4.1,
			Wood:   true,
			Carpet: false,
			Tile:   true,
		},
		{
			ID: 6,
			Name: "p6",
			Address: database.Address{
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
	initTest(t, dbProviders)
	req := defaultRequest
	response, status := sendRequest(req)
	Expect(status).To(Equal(http.StatusOK))
	Expect(response).To(Equal(convertFromDBProviders(dbProviders)))
}

func TestInvalidMaterial(t *testing.T) {
	initTest(t, nil)
	req := defaultRequest
	req.Material = "invalid"
	response, status := sendRequest(req)
	Expect(status).To(Equal(http.StatusBadRequest))
	Expect(response).To(BeEmpty())
}

func TestInvalidArea(t *testing.T) {
	initTest(t, nil)
	req := defaultRequest
	req.Area = 0
	response, status := sendRequest(req)
	Expect(status).To(Equal(http.StatusBadRequest))
	Expect(response).To(BeEmpty())
}

func TestInvalidPhoneNumber(t *testing.T) {
	initTest(t, nil)
	req := defaultRequest
	req.PhoneNumber = ""
	response, status := sendRequest(req)
	Expect(status).To(Equal(http.StatusBadRequest))
	Expect(response).To(BeEmpty())
}

func TestDBError(t *testing.T) {
	initTest(t, nil)
	db.GetProvidersFunc = func(material database.FloorMaterial, location database.Address) ([]database.Provider, error) {
		return nil, errors.New("database error")
	}
	req := defaultRequest
	response, status := sendRequest(req)
	Expect(status).To(Equal(http.StatusInternalServerError))
	Expect(response).To(BeEmpty())
}

func TestGetProvidersMatching(t *testing.T) {
	dbProviders := []database.Provider{
		{
			Name:    "p0",
			Address: database.Address{
				Lat:  -26,
				Long: 40,
			},
			Radius:  10,
			Rating:  5,
			Wood:    false,
			Carpet:  false,
			Tile:    false,
		},
		{
			Name:    "p1",
			Address: database.Address{
				Lat:  -26,
				Long: 40,
			},
			Radius:  10,
			Rating:  5,
			Wood:    true,
			Carpet:  false,
			Tile:    false,
		},
		{
			Name:    "p2",
			Address: database.Address{
				Lat:  -26,
				Long: 40,
			},
			Radius:  10,
			Rating:  5,
			Wood:    false,
			Carpet:  true,
			Tile:    false,
		},
		{
			Name:    "p3",
			Address: database.Address{
				Lat:  -26,
				Long: 40,
			},
			Radius:  10,
			Rating:  5,
			Wood:    false,
			Carpet:  false,
			Tile:    true,
		},
		{
			Name:    "p4",
			Address: database.Address{
				Lat:  -26,
				Long: 40,
			},
			Radius:  10,
			Rating:  5,
			Wood:    true,
			Carpet:  true,
			Tile:    false,
		},
		{
			Name:    "p5",
			Address: database.Address{
				Lat:  -26,
				Long: 40,
			},
			Radius:  10,
			Rating:  5,
			Wood:    true,
			Carpet:  false,
			Tile:    true,
		},
		{
			Name:    "p6",
			Address: database.Address{
				Lat:  -26,
				Long: 40,
			},
			Radius:  10,
			Rating:  5,
			Wood:    false,
			Carpet:  true,
			Tile:    true,
		},
	}
	initTest(t, dbProviders)
	req := defaultRequest
	response, status := sendRequest(req)
	Expect(status).To(Equal(http.StatusOK))
	Expect(response).To(Equal(convertFromDBProviders(dbProviders)))
}

func TestGetProvidersOutOfRadius(t *testing.T) {
	dbProviders := []database.Provider{
		{
			Name: "p0",
			Address: database.Address{
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
			Address: database.Address{
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
			Address: database.Address{
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
			Address: database.Address{
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
	initTest(t, dbProviders)
	req := defaultRequest
	response, status := sendRequest(req)
	Expect(status).To(Equal(http.StatusOK))
	Expect(response).To(Equal(convertFromDBProviders(dbProviders)))
}

func TestGetProvidersOrder(t *testing.T) {
	dbProviders := []database.Provider{
		{
			Name: "p0",
			Address: database.Address{
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
			Address: database.Address{
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
			Address: database.Address{
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
			Address: database.Address{
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
	initTest(t, dbProviders)
	req := defaultRequest
	response, status := sendRequest(req)
	Expect(status).To(Equal(http.StatusOK))
	Expect(response).To(Equal(convertFromDBProviders(dbProviders)))
}

func sendRequest(request handlers.CustomerRequest) ([]handlers.Provider, int) {
	body, err := jsoniter.Marshal(request)
	Expect(err).To(BeNil())
	path := "/get_providers"
	resp := execRequest(http.MethodGet, path, string(body))
	if resp.StatusCode != http.StatusOK {
		return nil, resp.StatusCode
	}
	respBody, err := ioutil.ReadAll(resp.Body)
	Expect(err).To(BeNil())
	response := handlers.Response{
		Data: &[]handlers.Provider{},
	}

	err = json.Unmarshal(respBody, &response)
	Expect(err).To(BeNil())
	err = resp.Body.Close()
	Expect(err).To(BeNil())
	providersInResponse := *response.Data.(*[]handlers.Provider)
	return providersInResponse, resp.StatusCode
}

func execRequest(method string, path string, body string) *http.Response {
	url := generateURL(path)
	reqBody := strings.NewReader(body)
	req, err := http.NewRequest(method, url, reqBody)
	Expect(err).To(BeNil())
	req.Close = true
	req.Header.Add("Content-Type", "application/json")
	client := http.Client{Timeout: 0}
	resp, err := client.Do(req)
	Expect(err).To(BeNil())
	return resp
}

func generateURL(path string) string {
	return "http://localhost:8000" + path
}

func initTest(t *testing.T, dbProviders []database.Provider) {
	RegisterTestingT(t)

	defaultRequest = handlers.CustomerRequest{
		Material:    "wood",
		Address:     handlers.Address{
			Lat:  -26.66129,
			Long: 40.95858,
		},
		Area:        1000,
		PhoneNumber: "1-800-234673",
	}
	db.GetProvidersFunc = func(database.FloorMaterial, database.Address) ([]database.Provider, error) {
		return dbProviders, nil
	}
}

func convertFromDBProviders(dbProviders []database.Provider) []handlers.Provider {
	res := []handlers.Provider{}
	for _, dbProvider := range dbProviders {
		provider := handlers.Provider{
			Name:            dbProvider.Name,
			Experience:      nil,
			Address:         handlers.Address{Lat: dbProvider.Address.Lat, Long: dbProvider.Address.Long},
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
		res = append(res, provider)
	}
	return res
}