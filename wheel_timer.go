package wheel_timer

type Slot map[int][]string
type Indicator map[string]int

type TickCallBack = func(dataList []string)
type RoundCallBack = func()

type Timer struct {
	Length      int // 刻度长度
	Index       int // 当前刻度
	Round       int // 转了第n圈
	Map         Indicator
	Slot        Slot
	onTickFunc  TickCallBack
	onRoundFunc RoundCallBack
}

func New(length int) (t *Timer) {
	t = &Timer{
		Length: length,
		Index:  0,
		Round:  0,
		Map:    make(Indicator, 10),
		Slot:   make(Slot, 10),
	}

	for i := 0; i < length; i++ {
		slot := t.Slot
		slot[i] = make([]string, 0)
	}

	return
}

func (t *Timer) Tick() *Timer {
	arr := t.Slot[t.Index]

	values := make([]string, len(arr))

	for i, v := range arr {
		values[i] = v
	}

	if t.onTickFunc != nil {
		t.onTickFunc(values)
	}

	delete(t.Slot, t.Index)

	// cursor
	if t.Index >= t.Length-1 {
		t.Round++
		t.Index = 0
		if t.onRoundFunc != nil {
			t.onRoundFunc()
		}
	} else {
		t.Index++
	}

	return t
}

func (t *Timer) Add(data string) *Timer {

	i := t.Index - 1

	if i < 0 {
		i = t.Length - 1
	}

	t.Slot[i] = append(t.Slot[i], data)
	t.Map[data] = i

	return t
}

func (t *Timer) OnTick(cb func(dataList []string)) {
	t.onTickFunc = cb
}

func (t *Timer) OnRound(cb func()) {
	t.onRoundFunc = cb
}
