package frame_test

import (
	"fmt"
	"github.com/bantl23/frame"
)

func Example_getFrameItem() {
	// 0 == start bit
	// 32 == number of bits
	f := frame.NewFrameItem(0, 32)

	// data byte arrays must be long
	// enough to hold full frame item
	d := []byte{0x1a, 0xcf, 0xfc, 0x1d}
	value, err := f.GetUint64(d)
	if err == nil {
		fmt.Println("value", value)
	} else {
		fmt.Println("err", err)
	}
}

func Example_setFrameItem() {
	// 0 == start bit
	// 32 == number of bits
	f := frame.NewFrameItem(0, 32)

	// data byte arrays must be long
	// enough to hold full frame item
	d := []byte{0x00, 0x00, 0x00, 0x00}
	value := uint64(0x1acffc1d)
	err := f.SetUint64(d, value)
	if err == nil {
		fmt.Println("data", d)
	} else {
		fmt.Println("err", err)
	}
}

func Example_getFrame() {
	f := frame.NewFrame("FrameName")
	f.Items["FrameItemName0"] = frame.NewFrameItem(0, 16)
	f.Items["FrameItemName1"] = frame.NewFrameItem(16, 32)

	d := []byte{0x1a, 0xcf, 0xfc, 0x1d}
	value, err := f.GetUint64("FrameItemName0", d)
	if err == nil {
		fmt.Println("value", value)
	} else {
		fmt.Println("err", err)
	}
}

func Example_setFrame() {
	f := frame.NewFrame("FrameName")
	f.Items["FrameItemName0"] = frame.NewFrameItem(0, 16)
	f.Items["FrameItemName1"] = frame.NewFrameItem(16, 32)

	d := []byte{0x00, 0x00, 0x00, 0x00}
	value0 := uint64(0x1acf)
	err := f.SetUint64("FrameItemName0", d, value0)
	if err == nil {
		fmt.Println("data", d)
		value1 := uint64(0xfc1d)
		err := f.SetUint64("FrameItemName1", d, value1)
		if err == nil {
			fmt.Println("data", d)
		} else {
			fmt.Println("err", err)
		}
	} else {
		fmt.Println("err", err)
	}
}
