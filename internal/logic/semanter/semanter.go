package semanter

import "semantiker/internal/api"

type (
	CoreKeyBind interface {
		ReadKeyFile(name string) []*api.CoreKeyBind
	}
)

type AdsManage interface {
	//CreateAds([]*dto.CoreKey) ([]*dto.ads, error)
}

type PromoManage interface {
	//CreatePromo([]*dto.ads) ([]*dto.promo, error)
}

type PromoBind interface {
	//SavePromo([]*bind.promo) error
}

///// func exec (name string, args ...string)

// open file by name

//create []dto.keys splitted by count 200

// for each dto.keys cluster create []dto.ads

// create template and fill it with dto.ads

// save template to file

type Semanter struct {
	keys CoreKeyBind
}

func NewSemanter(keys CoreKeyBind) Semanter {
	return Semanter{
		keys: keys,
	}
}
