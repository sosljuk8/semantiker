package ads

import (
	"semantiker/dto"
)

type Ads struct {
	Ads []map[int]*dto.Ads
}

func NewAds() *Ads {
	return &Ads{}
}

func (a *Ads) MakeAds(c []*dto.Corekey) {

	ak := []*dto.Ads{}

	for _, v := range c {
		ad := dto.NewAds()
		ad.SetKeys(v)
		ak = append(ak, ad)
	}

	a.Chank(ak)
}

func (d *Ads) Chank(a []*dto.Ads) {

	// Разбиение слайса
	chunkSize := 200
	result := make([][]*dto.Ads, 0)
	var first, last int
	for i := 0; i < len(a)/chunkSize+1; i++ {
		first = i * chunkSize
		last = i*chunkSize + chunkSize
		if last > len(a) {
			last = len(a)
		}
		if first == last {
			break
		}

		result = append(result, a[first:last])
	}

	d.Ads = Mapperise(result)
}

func Mapperise(a [][]*dto.Ads) []map[int]*dto.Ads {

	arr := []map[int]*dto.Ads{}

	for _, item := range a {
		ads := map[int]*dto.Ads{}

		for i, ad := range item {
			ads[i+12] = ad
		}
		arr = append(arr, ads)
	}
	return arr
}
