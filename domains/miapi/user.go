package miapi

import (
	"../../utils/apierrors"
	"../../utils/apiuris"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)


type User struct {
	Id               int64    `json:"id"`
	Nickname         string `json:"nickname"`
	RegistrationDate string `json:"registration_date"`
	CountryID        string `json:"country_id"`
	SiteID           string      `json:"site_id"`
	Permalink        string      `json:"permalink"`
	SellerReputation struct {
		LevelID           interface{} `json:"level_id"`
		PowerSellerStatus interface{} `json:"power_seller_status"`
		Transactions      struct {
			Canceled  int    `json:"canceled"`
			Completed int    `json:"completed"`
			Period    string `json:"period"`
			Ratings   struct {
				Negative int `json:"negative"`
				Neutral  int `json:"neutral"`
				Positive int `json:"positive"`
			} `json:"ratings"`
			Total int `json:"total"`
		} `json:"transactions"`
	} `json:"seller_reputation"`
	BuyerReputation struct {
		Tags []interface{} `json:"tags"`
	} `json:"buyer_reputation"`
	Status struct {
		SiteStatus string `json:"site_status"`
	} `json:"status"`
}


func (user *User) Get() *apierrors.ApiError {
	if user.Id == 0 {
		return &apierrors.ApiError{
			"Invalid user id",
			http.StatusBadRequest,
		}
	}
	url := fmt.Sprintf("%s%d", apiuris.UrlUsers, user.Id)
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

	if err := json.Unmarshal(data, &user); err != nil{
		return &apierrors.ApiError{
			err.Error(),
			http.StatusInternalServerError,
		}
	}
	return nil
}