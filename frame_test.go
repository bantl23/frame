package frame_test

import (
	"github.com/bantl23/frame"
	"testing"
)

const (
	FRAME   = "frame"
	SYNC    = "sync"
	VERSION = "vers"
	SCID    = "scid"
	VCID    = "vcid"
	VCCNT   = "vccnt"
)

func TestGetUint64(t *testing.T) {
	f := frame.NewFrame(FRAME)
	f.Items[SYNC] = frame.NewFrameItem(0, 32)
	f.Items[VERSION] = frame.NewFrameItem(32, 2)
	f.Items[SCID] = frame.NewFrameItem(34, 8)
	f.Items[VCID] = frame.NewFrameItem(42, 6)
	f.Items[VCCNT] = frame.NewFrameItem(48, 24)

	frame := []byte{0x1a, 0xcf, 0xfc, 0x1d, 0x55, 0x57, 0x54, 0x43, 0x21}
	sync, err := f.GetUint64(SYNC, frame)
	if err == nil {
		if sync != uint64(0x1acffc1d) {
			t.Error("expected", sync, "== 0x1acffc1d")
		}
	} else {
		t.Error(err)
	}
	vers, err := f.GetUint64(VERSION, frame)
	if err == nil {
		if vers != uint64(0x01) {
			t.Error("expected", vers, "== 0x01")
		}
	} else {
		t.Error(err)
	}
	scid, err := f.GetUint64(SCID, frame)
	if err == nil {
		if scid != uint64(0x55) {
			t.Error("expected", scid, "== 0x55")
		}
	} else {
		t.Error(err)
	}
	vcid, err := f.GetUint64(VCID, frame)
	if err == nil {
		if vcid != uint64(0x17) {
			t.Error("expected", vcid, "== 0x17")
		}
	} else {
		t.Error(err)
	}
	vccn, err := f.GetUint64(VCCNT, frame)
	if err == nil {
		if vccn != uint64(0x544321) {
			t.Error("expected", vccn, "== 0x544321")
		}
	} else {
		t.Error(err)
	}

	frame = []byte{}
	sync, err = f.GetUint64(SYNC, frame)
	if err == nil {
		t.Error("expected error")
	}
}

func TestSetUint64(t *testing.T) {
	f := frame.NewFrame(FRAME)
	f.Items[SYNC] = frame.NewFrameItem(0, 32)
	f.Items[VERSION] = frame.NewFrameItem(32, 2)
	f.Items[SCID] = frame.NewFrameItem(34, 8)
	f.Items[VCID] = frame.NewFrameItem(42, 6)
	f.Items[VCCNT] = frame.NewFrameItem(48, 24)

	frame := []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	err := f.SetUint64(SYNC, frame, 0x1acffc1d)
	if err != nil {
		t.Error(err)
	}
	err = f.SetUint64(VERSION, frame, 0x01)
	if err != nil {
		t.Error(err)
	}
	err = f.SetUint64(SCID, frame, 0x55)
	if err != nil {
		t.Error(err)
	}
	err = f.SetUint64(VCID, frame, 0x17)
	if err != nil {
		t.Error(err)
	}
	err = f.SetUint64(VCCNT, frame, 0x544321)
	if err != nil {
		t.Error(err)
	}

	sync, err := f.GetUint64(SYNC, frame)
	if err == nil {
		if sync != uint64(0x1acffc1d) {
			t.Error("expected", sync, "== 0x1acffc1d")
		}
	} else {
		t.Error(err)
	}
	vers, err := f.GetUint64(VERSION, frame)
	if err == nil {
		if vers != uint64(0x01) {
			t.Error("expected", vers, "== 0x01")
		}
	} else {
		t.Error(err)
	}
	scid, err := f.GetUint64(SCID, frame)
	if err == nil {
		if scid != uint64(0x55) {
			t.Error("expected", scid, "== 0x55")
		}
	} else {
		t.Error(err)
	}
	vcid, err := f.GetUint64(VCID, frame)
	if err == nil {
		if vcid != uint64(0x17) {
			t.Error("expected", vcid, "== 0x17")
		}
	} else {
		t.Error(err)
	}
	vccn, err := f.GetUint64(VCCNT, frame)
	if err == nil {
		if vccn != uint64(0x544321) {
			t.Error("expected", vccn, "== 0x544321")
		}
	} else {
		t.Error(err)
	}

	frame = []byte{}
	err = f.SetUint64(SYNC, frame, 0x1acffc1d)
	if err == nil {
		t.Error("expected error")
	}
}
