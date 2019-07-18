package miapi

import (
	"../../domains/miapi"
	"../../utils/apierrors"
)

func GetSiteFromApi(siteId string) (*miapi.Site, *apierrors.ApiError) {
	site := &miapi.Site{
		Id: siteId,
	}

	if err := site.Get(); err != nil{
		return nil, err
	}
	return site, nil
}

/*func GetSitesFromApi() (*[]miapi.Site, *apierrors.ApiError) {
	site := &miapi.Site{}
	sites := &miapi.Sites

	if err := site.GetSites(); err != nil{
		return nil, err
	}
	fmt.Println(sites)
	return sites, nil
}*/
