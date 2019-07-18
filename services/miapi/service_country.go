package miapi

import (
	"../../domains/miapi"
	"../../utils/apierrors"
)

func GetCountryFromApi(countryId string) (*miapi.Country, *apierrors.ApiError){
	country := &miapi.Country{
		Id: countryId,
	}

	if err := country.Get(); err != nil{
		return nil, err
	}

	return country, nil
}