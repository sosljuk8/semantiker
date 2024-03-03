package bind

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"os"
	"semantiker/dto"
	"semantiker/internal/config"
	"strings"
)

type PromoBind struct {
	Conf *config.SemanterConfig
}

func NewPromobind() *PromoBind {
	return &PromoBind{}

}

func (p PromoBind) CreateFile(n string, ind string, ads map[int]*dto.Ads) {
	kp := config.GetConfig()
	f, errf := excelize.OpenFile(kp.Promo.TemplatePath)
	if errf != nil {
		fmt.Println(errf)
		return
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	var (
		err   error
		addr  string
		sheet = kp.Promo.Sheet
	)

	//ads := ak.GetAds()
	// Set value of a cell.
	for kr, a := range ads {
		for k, ad := range a.Cells {
			if addr, err = excelize.JoinCellName(k, kr); err != nil {
				fmt.Println(err)
				return
			}

			f.SetCellValue(sheet, addr, ad)
		}
	}

	dir := strings.Replace(kp.Promo.DirPath, "**name**", n, 1)

	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		fmt.Println(err)
	}

	if err := f.SaveAs(dir + ind + ".xlsx"); err != nil {
		fmt.Println(err)
	}
}
