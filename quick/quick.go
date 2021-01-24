package quick

func Add(a, b uint16) uint16 {
	var i uint16
	for i = 0; i < b; i++ {
		a++
	}
	return a
}
