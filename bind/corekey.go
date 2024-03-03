package bind

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	_ "github.com/xuri/excelize/v2"
	"semantiker/dto"
	"semantiker/internal/config"
	"strings"
)

type CoreKeyBind struct {
	conf map[string]string
}

func NewCoreKeyBind() *CoreKeyBind {
	return &CoreKeyBind{}
}

func (c CoreKeyBind) ReadKeyFile(name string) []*dto.Corekey {
	kp := config.GetConfig()
	p := strings.Replace(kp.CKey.KeyFilePath, "**name**", name, 1)

	fmt.Println(p)

	f, err := excelize.OpenFile(p)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	ka := []*dto.Corekey{}
	// Получить все строки в Keywords
	rows, err := f.GetRows(kp.CKey.KeySheet)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	for _, row := range rows {
		if len(row) == 0 {
			continue
		}
		dc := dto.NewCorekey(row)
		ka = append(ka, dc)
	}
	return ka
}
