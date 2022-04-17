package model

type account struct {
	numeber  string
	balance  float64
	password string
}

func (a *account) SetNumber(num int) {
	if len(num) >= 6 && len(num) <= 10 {
		a.num = number
	}

}