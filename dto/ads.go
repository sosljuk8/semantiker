package dto

import (
	"strings"
)

type Cells struct {
	// A-Z-AT strings
	A  string
	B  string
	C  string
	D  string
	E  string
	F  string
	G  string
	H  string
	I  string
	J  string
	K  string
	L  string
	M  string
	N  string
	O  string
	P  string
	Q  string
	R  string
	S  string
	T  string
	U  string
	V  string
	W  string
	X  string
	Y  string
	Z  string
	AA string
	AB string
	AC string
	AD string
	AE string
	AF string
	AG string
	AH string
	AI string
	AJ string
	AK string
	AL string
	AM string
	AN string
	AO string
	AP string
	AQ string
	AR string
	AS string
	AT string
	AU string
	AV string
}

type Ads struct {
	Cells map[string]string
}

func NewAds() *Ads {
	return &Ads{
		Cells: NewCells(),
	}
}

func NewCells() map[string]string {
	c := map[string]string{
		"A":  "-",
		"B":  "Текстово-графическое",
		"C":  "-",
		"D":  "5409766352",
		"E":  "DOLD%%%%",
		"F":  "1",
		"G":  "-",
		"H":  "5409766352",
		"I":  "bh5922%%%%%%%",
		"J":  "15887716369",
		"K":  "Оригинальные реле Dold!%%%%%%",
		"L":  "Прямые поставки из Европы!",
		"M":  "Замена оборудования в короткие сроки! Работаем с заводами по всей России! Звоните!",
		"N":  "0",
		"O":  "0",
		"P":  "0",
		"Q":  "https://snabsystem.ru/offers/dold?utm_source=yandex&utm_medium=cpc&utm_campaign={campaign_id}&utm_content={ad_id}&utm_term={keyword}",
		"R":  "Поставки-Dold%%%%%%",
		"S":  "Россия",
		"T":  "54866432559",
		"U":  "100",
		"V":  "",
		"W":  "-",
		"X":  "",
		"Y":  "Работает везде",
		"Z":  "Датчики и реле  ||Тормозные реле||О компании||Заказать онлайн",
		"AA": "||||||Оформите заявку в один клик",
		"AB": "http://snabsystem.ru/catalog/rele-i-datchiki/dold?utm_source=yandex&utm_medium=cpc&utm_campaign={campaign_id}&utm_content={ad_id}&utm_term={keyword}||http://snabsystem.ru/catalog/rele-i-datchiki/tormoznye-rele-dold?utm_source=yandex&utm_medium=cpc&utm_campaign={campaign_id}&utm_content={ad_id}&utm_term={keyword}||http://snabsystem.ru/catalog/rele-i-datchiki/e-dold-sohne-kg?utm_source=yandex&utm_medium=cpc&utm_campaign={campaign_id}&utm_content={ad_id}&utm_term={keyword}||http://snabsystem.ru/ajaxy/call-me/nojs?utm_source=yandex&utm_medium=cpc&utm_campaign={campaign_id}&utm_content={ad_id}&utm_term={keyword}",
		"AC": "",
		"AD": "",
		"AE": "",
		"AF": "",
		"AG": "",
		"AH": "",
		"AI": "",
		"AJ": "",
		"AK": "Страховка поставки||Гарантия 6 месяцев||Express доставка по РФ||Перс. менеджер||Фиксация цен ||Страховка на поставки",
		"AL": "",
		"AM": "",
		"AN": "",
		"AO": "",
		"AP": "",
		"AQ": "",
		"AR": "",
		"AS": "",
		"AT": "",
		"AU": "",
		"AV": "",
	}
	return c
}

func (a *Ads) SetKeys(c *Corekey) {

	caps := a.GetUppers(c.Name)

	a.Cells["E"] = caps["upper"]
	a.Cells["I"] = c.Key
	a.Cells["K"] = "Оригинальные " + c.Addon + " " + caps["up"] + "!"
	a.Cells["R"] = "Поставки-" + caps["up"]
	a.Cells["S"] = "Россия"
	a.Cells["U"] = "100"
	a.Cells["Z"] = "Датчики и реле  ||Тормозные реле||О компании||Заказать онлайн"

}

func (a *Ads) GetUppers(s string) map[string]string {
	u := map[string]string{}

	u["lower"] = strings.ToLower(s)
	u["upper"] = strings.ToUpper(s)
	u["up"] = strings.ToUpper(s[:1]) + s[1:]

	return u

}

type AdsKit struct {
	Ads  []*Ads
	Kits []map[int]*Ads
	TT   [][]*Ads
}

func NewAdsKit() *AdsKit {
	return &AdsKit{
		Ads:  []*Ads{},
		Kits: []map[int]*Ads{},
		TT:   [][]*Ads{},
	}
}

func (a *AdsKit) AddAds(ad *Ads) {
	a.Ads = append(a.Ads, ad)
}

func (a *AdsKit) AddKit() LastKit {
	n := make(map[int]*Ads)
	a.Kits = append(a.Kits, n)
	return a.GeLastKit()
}

type LastKit struct {
	Last   map[int]*Ads
	Status bool
}

func (a *AdsKit) GeLastKit() LastKit {
	if len(a.Kits) == 0 {
		return a.AddKit()
	}
	l := LastKit{
		Last:   a.Kits[len(a.Kits)-1],
		Status: false,
	}
	if len(a.Kits) < 200 {
		l.Status = true
	}
	return l

}

func (a *AdsKit) AddKits() {
	//a.Kits = append(a.Kits, k)

	for i, ad := range a.Ads {
		mi := i + 12
		if mi > 200 {
			mi = mi - 200
		}
		kit := a.GeLastKit()
		if !kit.Status {
			kit = a.AddKit()
		}

		kit.Last[mi] = ad
	}

}

func (a *AdsKit) GetAds() map[int]*Ads {

	m := map[int]*Ads{}

	for i, ad := range a.Ads {
		m[i+12] = ad
	}

	return m
}

func (a *AdsKit) Chank() {
	//// Подготовка исходного слайса
	//sourceSize := len(a.Ads)
	//source := make([]int, sourceSize)
	//for i := 0; i < sourceSize; i++ {
	//	source[i] = i + 1
	//}

	// Разбиение слайса
	chunkSize := 200
	result := make([][]*Ads, 0)
	var first, last int
	for i := 0; i < len(a.Ads)/chunkSize+1; i++ {
		first = i * chunkSize
		last = i*chunkSize + chunkSize
		if last > len(a.Ads) {
			last = len(a.Ads)
		}
		if first == last {
			break
		}

		result = append(result, a.Ads[first:last])
	}

	a.TT = result
	// Вывод результата
	//for _, res := range result {
	//	fmt.Println(res)
	//}
}
