package ledsgo

import (
	"bufio"
	"os"
)

var (
	GAMMA = [256]int{
		128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128,
		128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 129, 129, 129, 129,
		129, 129, 129, 129, 129, 129, 129, 129, 129, 129, 129, 129, 130, 130, 130, 130,
		130, 130, 130, 130, 130, 131, 131, 131, 131, 131, 131, 131, 131, 132, 132, 132,
		132, 132, 132, 132, 133, 133, 133, 133, 133, 134, 134, 134, 134, 134, 135, 135,
		135, 135, 135, 136, 136, 136, 136, 137, 137, 137, 137, 138, 138, 138, 138, 139,
		139, 139, 140, 140, 140, 141, 141, 141, 141, 142, 142, 142, 143, 143, 144, 144,
		144, 145, 145, 145, 146, 146, 146, 147, 147, 148, 148, 149, 149, 149, 150, 150,
		151, 151, 152, 152, 152, 153, 153, 154, 154, 155, 155, 156, 156, 157, 157, 158,
		158, 159, 160, 160, 161, 161, 162, 162, 163, 163, 164, 165, 165, 166, 166, 167,
		168, 168, 169, 169, 170, 171, 171, 172, 173, 173, 174, 175, 175, 176, 177, 178,
		178, 179, 180, 180, 181, 182, 183, 183, 184, 185, 186, 186, 187, 188, 189, 190,
		190, 191, 192, 193, 194, 195, 195, 196, 197, 198, 199, 200, 201, 202, 202, 203,
		204, 205, 206, 207, 208, 209, 210, 211, 212, 213, 214, 215, 216, 217, 218, 219,
		220, 221, 222, 223, 224, 225, 226, 227, 228, 229, 230, 232, 233, 234, 235, 236,
		237, 238, 239, 241, 242, 243, 244, 245, 246, 248, 249, 250, 251, 253, 254, 255}
)

type LPD8806Color struct {
	R int
	G int
	B int
}

func (c *LPD8806Color) GetR() int {
	return GAMMA[c.R]
}

func (c *LPD8806Color) GetG() int {
	return GAMMA[c.G]
}

func (c *LPD8806Color) GetB() int {
	return GAMMA[c.B]
}

var (
	SPI_DEVICE = "/dev/spidev0.0"
)

type LPD8806Strip struct {
	length int
	buffer []byte
	device *bufio.Writer
}

func (s *LPD8806Strip) Reset() {
	s.device.Write([]byte{0})
	s.Fill(&LPD8806Color{0, 0, 0})
	s.Update()
}

func (s *LPD8806Strip) Set(i int, color Color) {
	s.buffer[i*3] = byte(color.GetG())
	s.buffer[i*3+1] = byte(color.GetR())
	s.buffer[i*3+2] = byte(color.GetB())
}

func (s *LPD8806Strip) Fill(color Color) {
	for i := 0; i < len(s.buffer); i += 3 {
		s.buffer[i] = byte(color.GetG())
		s.buffer[i+1] = byte(color.GetR())
		s.buffer[i+2] = byte(color.GetB())
	}
}

func (s *LPD8806Strip) Update() {
	s.device.Write(s.buffer)
	s.device.Write([]byte{0})
	s.device.Flush()
}

func NewLPD8806Strip(length int) *LPD8806Strip {
	spidev, err := os.Create(SPI_DEVICE)
	if err != nil {
		panic(err)
	}

	return &LPD8806Strip{
		length: length,
		buffer: make([]byte, length*3),
		device: bufio.NewWriterSize(spidev, length*3),
	}
}
