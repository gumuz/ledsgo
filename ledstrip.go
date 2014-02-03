package ledsgo

type LEDStrip interface {
	Reset()
	Set(i int, color Color)
	Fill(color Color)
	Update()
}
