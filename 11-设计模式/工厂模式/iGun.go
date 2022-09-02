package 工厂模式

type IGun interface {
	setName(name string)
	setPower(power int)
	getName() string
	getPower() int
}
