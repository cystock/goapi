package miapi

import (
	"../../utils/apierrors"
	"../../utils/apiuris"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Site struct {
	Id		   		string `json:"id"`
	Name	   		string `json:"name"`
	CountryId		string `json:"country_id"`
	SaleFeesMode 	string `json:"sale_fees_mode"`
	MercadopagoVersion int8 `json:"mercadopago_version"`
	DefaultCurrencyId string `json:"default_currency_id"`
	ImmediatePayment string `json:"immediate_payment"`
	PaymentMethodIds []interface{} `json:"payment_method_ids"`
	Settings struct{
		IdentificationType []interface{} `json:"identification_type"`
		TaxpayerTypes []interface{} `json:"taxpayer_types"`
		IdentificationTypesRules []interface{} `json:"identification_types_rules"`
	} `json:"settings"`
	Currencies []interface{}  `json:"currencies"`
	Categories []interface{} `json:"categories"`
}


//var Sites []Site


func (site *Site) Get() *apierrors.ApiError{
	if site.Id == "" {
		return &apierrors.ApiError{
			"Invalid site id",
			http.StatusBadRequest,
		}
	}
	url := fmt.Sprintf("%s%s", apiuris.UrlSites, site.Id)
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

	if err := json.Unmarshal(data, &site); err != nil{
		return &apierrors.ApiError{
			err.Error(),
			http.StatusInternalServerError,
		}
	}
	return nil
}

/*func (site *Site) GetSites() *apierrors.ApiError {
	res, err := http.Get(urlSites)
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

	if err := json.Unmarshal(data, &Sites); err != nil{
		return &apierrors.ApiError{
			err.Error(),
			http.StatusInternalServerError,
		}
	}
	fmt.Println(&Sites)
	return nil
}*/