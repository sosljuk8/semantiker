package core

type BrandsKeys struct {
	cores []string
}

func NewBrandsKeys() *BrandsKeys {
	return &BrandsKeys{
		cores: []string{
			"dold",
			"hydac",
			"omron",
		},
	}
}

func (b *BrandsKeys) GetCoresAds() []string {
	return b.cores
}
