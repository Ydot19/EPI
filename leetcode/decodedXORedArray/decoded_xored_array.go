package decodedXORedArray

type DecodedXORedArray interface {
	Optimized(encode []int, first int) []int
}

type Impl struct{}

func NewDecodedXORedArrayImpl() DecodedXORedArray {
	return &Impl{}
}

func (d *Impl) Optimized(encoded []int, first int) []int {
	targetLen := len(encoded) + 1
	decoded := make([]int, targetLen)
	for i, encodeVal := range encoded {
		decoded[i] = first
		first = encodeVal ^ first
	}
	decoded[targetLen-1] = first
	return decoded
}
