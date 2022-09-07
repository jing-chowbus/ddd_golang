package value

type Percentage struct {
	Percent int64
}

func (percentage Percentage) ToFloat() float64 {
	return float64(percentage.Percent / 100)
}
