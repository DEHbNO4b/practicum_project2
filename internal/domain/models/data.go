package models

type Data struct {
	lpd []LogPassData
	cd  []Card
	td  []TextData
	bd  []BinaryData
}

func NewData() *Data {

	data := Data{
		lpd: make([]LogPassData, 0),
		cd:  make([]Card, 0),
		td:  make([]TextData, 0),
		bd:  make([]BinaryData, 0),
	}

	return &data
}

func (d *Data) AddLpd(lp LogPassData) {
	d.lpd = append(d.lpd, lp)
}
func (d *Data) AddCd(cd Card) {
	d.cd = append(d.cd, cd)
}
func (d *Data) AddTd(td TextData) {
	d.td = append(d.td, td)
}
func (d *Data) AddBd(bd BinaryData) {
	d.bd = append(d.bd, bd)
}

func (d *Data) Lpd() []LogPassData {
	ans := make([]LogPassData, 0)

	ans = append(ans, d.lpd...)

	return ans
}
func (d *Data) Cd() []Card {

	ans := make([]Card, 0)

	ans = append(ans, d.cd...)

	return ans
}
func (d *Data) Td() []TextData {
	ans := make([]TextData, 0)

	ans = append(ans, d.td...)

	return ans
}
func (d *Data) Bd() []BinaryData {

	ans := make([]BinaryData, 0)

	ans = append(ans, d.bd...)

	return ans
}

func (d *Data) String() string {

	str := "Data: \n"

	str += "LogPass data: \n"

	for _, el := range d.lpd {
		str += el.String()
	}

	str += "Cards data: \n"

	for _, el := range d.cd {
		str += el.String()
	}

	str += "Text data: \n"

	for _, el := range d.td {
		str += el.String()
	}

	str += "Binary data: \n"

	for _, el := range d.bd {
		str += el.String()
	}

	return str
}
