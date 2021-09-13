package handlers

import (
	"WeatherByCoordinates/api/mapbox"
	"WeatherByCoordinates/api/weatherstack"
	"WeatherByCoordinates/repository"
	"fmt"
	"strings"
	"time"

	"github.com/pkg/errors"
)

func (repo *UseCase) FullResult(city string) (*repository.UserReqRes, error) {
	now := time.Now()
	res1, err := mapbox.Geocode(city)
	if err != nil {
		return nil, errors.Wrap(err, "Error request")
	}
	res2, err := weatherstack.Forecast(fmt.Sprintf("%v", res1.Latitude), fmt.Sprintf("%v", res1.Longitude))
	counter, err := repo.Repo.NumberOfRecords()
	if err != nil {
		return nil, errors.Wrap(err, "записей нет")
	}
	fullData := repository.UserReqRes{
		Data_id:             fmt.Sprint(counter + 1),
		Request:             strings.ToLower(city),
		City:                res2.Region,
		Latitude:            fmt.Sprint(res1.Latitude),
		Longitude:           fmt.Sprint(res1.Longitude),
		Temperature:         fmt.Sprint(res2.Temperature),
		Weatherdescriptions: fmt.Sprint(res2.Weather_Descriptions),
		Humidity:            fmt.Sprint(res2.Humidity),
		Data:                fmt.Sprintf("%d-%d-%d", now.Year(), now.Month(), now.Day()),
	}

	if err != nil {
		return nil, errors.Wrap(err, "Error request")
	}

	return &fullData, err
}

func (repo *UseCase) AddData(structData *repository.UserReqRes) (string, error) {
	err := repo.Repo.CreateUsersReqRes(
		structData.Request,
		structData.City,
		structData.Latitude,
		structData.Longitude,
		structData.Temperature,
		structData.Weatherdescriptions,
		structData.Humidity,
		structData.Data)
	if err != nil {
		return "", err
	}
	return "ok", err

}
