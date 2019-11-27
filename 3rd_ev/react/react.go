package react

type MyReactor struct {
	MyCell
}

type MyCell struct {
	val       int
	lastVal   int
	listeners []func(int)
	callbacks map[int]func(int)
}

type MyInputCell struct {
	MyCell
}

type MyComputeCell struct {
	MyCell
}

type MyCanceler struct {
	index int
	c     *MyCell
}

func New() *MyReactor {
	reactor := new(MyReactor)
	return reactor
}

func (r MyReactor) CreateInput(v int) InputCell {
	d := MyCell{val: v, callbacks: map[int]func(int){}}
	return &d
}

func (r MyReactor) CreateCompute1(c Cell, callback func(int) int) ComputeCell {
	computeCell := MyCell{}
	computeCell.val = callback(c.Value())
	c.(*MyCell).AddListener(func(v int) { computeCell.SetValue(callback(v)) })

	return &computeCell
}

func (c *MyCell) AddListener(listener func(int)) {
	c.listeners = append(c.listeners, listener)
}

func (r MyReactor) CreateCompute2(c1 Cell, c2 Cell, callback func(int, int) int) ComputeCell {
	computeCell := MyCell{}
	computeCell.val = callback(c1.Value(), c2.Value())
	c1.(*MyCell).AddListener(func(v int) {
		computeCell.lastVal = computeCell.val
		computeCell.val = callback(v, c2.Value())
	})
	c2.(*MyCell).AddListener(func(v int) { computeCell.SetComputeValue(callback(c1.Value(), v)) })

	return &computeCell
}

func (ic *MyCell) SetValue(v int) {
	for _, listener := range ic.listeners {
		listener(v)
	}

	if ic.val != v {
		for _, callback := range ic.callbacks {
			callback(v)
		}
	}
	ic.val = v
}

func (ic *MyCell) SetComputeValue(v int) {
	for _, listener := range ic.listeners {
		listener(v)
	}

	if ic.lastVal != v {
		for _, callback := range ic.callbacks {
			callback(v)
		}
	}
	ic.val = v
}

func (c *MyCell) Value() int {
	return c.val
}

func (c *MyCell) AddCallback(callback func(int)) Canceler {
	callbackSize := len(c.callbacks)
	if c.callbacks == nil {
		c.callbacks = make(map[int]func(int))
	}
	c.callbacks[callbackSize] = callback
	return &MyCanceler{index: callbackSize, c: c}
}

func (can *MyCanceler) Cancel() {
	delete(can.c.callbacks, can.index)
}