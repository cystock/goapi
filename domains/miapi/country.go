package miapi

import (
	"../../utils/apierrors"
	"../../utils/apiuris"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Country struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Locate string `json:"locate"`
	DecimalSeparator string `json:"decimal_separator"`
	ThousandsSeparator string `json:"thousands_separator"`
	TimeZone string `json:"time_zone"`
	GeoInformation struct{
		Location struct{
			Latitude float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"location"`
	} `json:"geo_information"`
	States []interface{} `json:"states"`
}


func (country *Country) Get() *apierrors.ApiError{
	if country.Id == "" {
		return &apierrors.ApiError{
			"Invalid Country id",
			http.StatusBadRequest,
		}
	}
	url := fmt.Sprintf("%s%s", apiuris.UrlCountries, country.Id)
	res, err := http.Get(url)
	if err != nil {
		return &apierrors.ApiError{
			err.Error(),
			http.StatusInternalServerError,
		}
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return &apierrors.ApiError{
			err.Error(),
			http.StatusInternalServerError,
		}
	}

	if err := json.Unmarshal(data, &country); err != nil{
		return &apierrors.ApiError{
			err.Error(),
			http.StatusInternalServerError,
		}
	}
	return nil
}