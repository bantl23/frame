package frame

import (
	"errors"
)

type FrameItem struct {
	BitLoc   uint64
	BitLen   uint64
	byteLoc  uint64
	byteLen  uint64
	shiftAmt uint64
	mask     uint64
}

type Frame struct {
	Name  string
	Items map[string]*FrameItem
}

func NewFrameItem(bitLoc uint64, bitLen uint64) *FrameItem {
	f := new(FrameItem)
	f.BitLoc = bitLoc
	f.BitLen = bitLen

	f.byteLoc = f.BitLoc / 8
	f.shiftAmt = 8 - ((f.BitLoc + f.BitLen) % 8)
	if f.shiftAmt == 8 {
		f.shiftAmt = 0
	}
	f.byteLen = (((f.BitLoc % 8) + f.BitLen - 1) / 8) + 1
	for i := uint64(0); i < f.BitLen; i++ {
		f.mask = f.mask | (uint64(0x01) << i)
	}
	return f
}

func NewFrame(name string) *Frame {
	f := new(Frame)
	f.Name = name
	f.Items = make(map[string]*FrameItem)
	return f
}

func (f FrameItem) GetUint64(data []byte) (uint64, error) {
	if (f.byteLoc + f.byteLen) <= uint64(len(data)) {
		val := uint64(0)
		for i := uint64(0); i < f.byteLen; i++ {
			mv := uint64((f.byteLen - i - 1) * 8)
			val = val | (uint64(data[i+f.byteLoc]) << mv)
		}
		return (val >> f.shiftAmt) & f.mask, nil
	} else {
		return 0, errors.New("byte array too small")
	}
}

func (f FrameItem) SetUint64(data []byte, val uint64) error {
	if (f.byteLoc + f.byteLen) <= uint64(len(data)) {
		val = (val & f.mask) << f.shiftAmt
		for i := uint64(0); i < f.byteLen; i++ {
			mv := uint64((f.byteLen - i - 1) * 8)
			data[i+f.byteLoc] = data[i+f.byteLoc] | byte(val>>mv)
		}
		return nil
	} else {
		return errors.New("byte array too small")
	}
}

func (f Frame) GetUint64(name string, data []byte) (uint64, error) {
	return f.Items[name].GetUint64(data)
}

func (f Frame) SetUint64(name string, data []byte, val uint64) error {
	return f.Items[name].SetUint64(data, val)
}
