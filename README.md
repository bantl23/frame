FRAME
=====

[![Build Status](https://travis-ci.org/bantl23/frame.svg?branch=master)](https://travis-ci.org/bantl23/frame)
[![Build Status](https://drone.io/github.com/bantl23/frame/status.png)](https://drone.io/github.com/bantl23/frame/latest)
[![Coverage Status](https://coveralls.io/repos/github/bantl23/frame/badge.svg?branch=master)](https://coveralls.io/github/bantl23/frame?branch=master)

This library implements a generic data frame

## Usage

```
import (
  "github.com/bantl23/frame"
)

func main() {
  f := frame.NewFrame("FrameName")

  // 0 == bit start location
  // 32 == bit length
  f.Items["FrameItemName"] = frame.NewFrameItem(0, 32)
  f.Items["FrameItemName"].GetUint64()
  x := f.GetUint64("FrameItemName")
  f.Items["FrameItemName"].SetUint64(x)
  f.SetUint64("FrameItemName", x)
}
```
