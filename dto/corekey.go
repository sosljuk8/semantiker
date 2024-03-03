package dto

type Corekey struct {
	Key   string
	Name  string
	Addon string
}

func NewCorekey(s []string) *Corekey {
	return &Corekey{
		Key:   s[0],
		Name:  s[1],
		Addon: s[2],
	}
}
