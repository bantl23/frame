package frame

import (
	"errors"
)

// FrameItem struct contains the information
// to get and set frame sub elements
type FrameItem struct {
	BitLoc   uint64
	BitLen   uint64
	byteLoc  uint64
	byteLen  uint64
	shiftAmt uint64
	mask     uint64
}

// Frame struct contains a list of frame
// items to mimic a complete data frame
type Frame struct {
	Name  string
	Items map[string]*FrameItem
}

// NewFrameItem creates and initializes
// a new frame item struct by receiving
// bitLoc (bit start location) and bitLen
// (number of bits).
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

// NewFrame creates and initializes
// a new frame which contains an
// empty list of frame items
func NewFrame(name string) *Frame {
	f := new(Frame)
	f.Name = name
	f.Items = make(map[string]*FrameItem)
	return f
}

// GetUint64 obtains a value from a data byte array
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

// SetUint64 assigns a value to a data byte array
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

// GetUint64 obtains a value from a data byte array using frame item name
func (f Frame) GetUint64(name string, data []byte) (uint64, error) {
	return f.Items[name].GetUint64(data)
}

// SetUint64 assigns a value to a data byte array using a frame item name
func (f Frame) SetUint64(name string, data []byte, val uint64) error {
	return f.Items[name].SetUint64(data, val)
}
