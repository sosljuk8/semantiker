package main

import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"semantiker/bind"
	"semantiker/cmd"
	"semantiker/internal/logic/ads"
	"strconv"
)

func init() {
	cmd.MigrateSql()
}

func main() {
	name := "hydac"

	b := bind.NewCoreKeyBind()
	keys := b.ReadKeyFile(name)

	aa := ads.NewAds()
	aa.MakeAds(keys)

	for i, v := range aa.Ads {
		p := bind.NewPromobind()
		ind := strconv.Itoa(i)
		p.CreateFile(name, ind, v)
	}
}
