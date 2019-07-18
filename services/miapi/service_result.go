package miapi

import (
	"../../domains/miapi"
	"../../utils/apierrors"
	"sync"
)

func GetResultFromApi(userId int64) (*miapi.Result, *apierrors.ApiError) {
	user := &miapi.User{
		Id: userId,
	}
	if err := user.Get(); err != nil{
		return nil, err
	}

	site := &miapi.Site{
		Id: user.SiteID,
	}
	country := &miapi.Country{
		Id: user.CountryID,
	}

	if err := site.Get(); err != nil{
		return nil, err
	}

	if err := country.Get(); err != nil{
		return nil, err
	}

	result := &miapi.Result{
		User: user,
		Site: site,
		Country: country,
	}
	return result, nil
}

func GetResultGoroutineFromApi(userId int64) (*miapi.Result, *apierrors.ApiError)  {
	result := &miapi.Result{}

	var wg sync.WaitGroup
	wg.Add(2)

	result.User, _ = createUser(userId) //TODO return err


	go createSite(result, &wg)
	go createCountry(result, &wg)

	wg.Wait()

	return result, nil
}

func createUser(userId int64) (*miapi.User, *apierrors.ApiError) {
	user := &miapi.User{
		Id: userId,
	}
	if err := user.Get(); err != nil{
		return nil, err
	}
	return user, nil
}

func createSite(result *miapi.Result, wg *sync.WaitGroup) *apierrors.ApiError {
	result.Site = &miapi.Site {
		Id: result.User.SiteID,
	}
	if err := result.Site.Get(); err != nil{
		return err
	}
	wg.Done()
	return nil
}

func createCountry( result *miapi.Result, wg *sync.WaitGroup) *apierrors.ApiError  {
	result.Country = &miapi.Country {
		Id: result.User.CountryID,
	}
	if err := result.Country.Get(); err != nil{
		return err
	}
	wg.Done()
	return nil
}

/// ------------------------------- ///
func GetResultGChannelFromApi(userId int64) *miapi.Result {
	result := miapi.Result{}

	result.User, _ = createUser(userId)

	c := make(chan *miapi.Result)
	//defer c

	var wg sync.WaitGroup
	wg.Add(2)

	go createSiteChannel(result.User.SiteID, c)
	go createCountryChannel(result.User.CountryID, c)

	go func() {
		for i:=0; i<2; i++{
			if result.User != nil && result.Site != nil && result.Country != nil {
				break
			}
			elem := <- c
			wg.Done()
			if elem.Country != nil{
				result.Country = elem.Country
			} else if elem.Site != nil{
				result.Site = elem.Site
			}

		}
	}()

	wg.Wait()
	return &result
}

func createSiteChannel(siteId string, c chan *miapi.Result) {
	site := miapi.Site{
		Id: siteId,
	}

	if err := site.Get(); err != nil{
		return
	}

	result := &miapi.Result {
		Site: &site,
	}
	c <- result
}

func createCountryChannel(countryId string, c chan *miapi.Result) {
	country := miapi.Country{
		Id: countryId,
	}

	if err := country.Get(); err != nil{
		return
	}

	result := &miapi.Result {
		Country: &country,
	}

	c <- result
}


