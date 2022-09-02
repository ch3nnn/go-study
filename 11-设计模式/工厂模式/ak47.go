package 工厂模式

type Ak47 struct {
	Gun
}

func NewAk47() *Ak47 {
	return &Ak47{Gun: Gun{
		name:  "AK47 gun",
		power: 4,
	}}
}
