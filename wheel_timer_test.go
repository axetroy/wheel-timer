package wheel_timer

import (
	"testing"
	"fmt"
)

func Test_WheelTimer(t *testing.T) {
	timer := New(10)

	if timer.Length != 10 {
		t.Errorf("The Length should be, not %v", timer.Length)
		return
	}

	if timer.Index != 0 {
		t.Errorf("The Index should be 0, not %v", timer.Index)
		return
	}

	if timer.Round != 0 {
		t.Errorf("The Round should be 0, not %v", timer.Round)
		return
	}

	for i := 0; i < 10; i++ {
		if len(timer.Slot[i]) != 0 {
			t.Errorf("The slot should be empty")
			return
		}
	}

	timer.OnTick(func(dataList []string) {

	})

	timer.Add("hello")

	if i, ok := timer.Map["hello"]; !ok {
		t.Errorf("The key of hello should exist in the map")
		return
	} else if i != 9 {
		fmt.Println(i)
		t.Errorf("The first key should be 9, the last length")
		return
	}

	for i := 0; i < 10-1; i++ {
		timer.Tick()
	}

	timer.OnRound(func() {

	})

	if timer.Index != 9 {
		t.Errorf("The Index should be 9, not %v", timer.Index)
		return
	}

	// tick and delete
	timer.Tick()

	if timer.Index != 0 {
		t.Errorf("The Index should be 0, not %v", timer.Index)
		return
	}

	timer.Tick()

	if timer.Index != 1 {
		t.Errorf("The Index should be 1, not %v", timer.Index)
		return
	}

	fmt.Print(timer)

}
