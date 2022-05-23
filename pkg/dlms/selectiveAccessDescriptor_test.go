package dlms

import (
	"bytes"
	"testing"
	"time"
)

func TestSelectiveAccessDescriptor_Encode(t *testing.T) {
	a := *CreateSelectiveAccessDescriptor(AccessSelectorEntry, []uint32{0, 5})
	t1, e := a.Encode()
	if e != nil {
		t.Errorf("t1 Encode Failed. err: %v", e)
	}
	result := []byte{2, 2, 4, 6, 0, 0, 0, 0, 6, 0, 0, 0, 5, 18, 0, 0, 18, 0, 0}

	res := bytes.Compare(t1, result)
	if res != 0 {
		t.Errorf("Test AccessSelectorEntry failed. get: %d, should:%v", t1, result)
	}

	timeStart := time.Date(2020, time.January, 1, 10, 0, 0, 0, time.UTC)
	timeEnd := time.Date(2020, time.January, 1, 11, 0, 0, 0, time.UTC)
	b := *CreateSelectiveAccessDescriptor(AccessSelectorRange, []time.Time{timeStart, timeEnd})
	t2, e := b.Encode()
	if e != nil {
		t.Errorf("t2 Encode Failed. err: %v", e)
	}
	result = []byte{1, 2, 4, 2, 4, 18, 0, 8, 9, 6, 0, 0, 1, 0, 0, 255, 15, 2, 18, 0, 0, 9, 12, 7, 228, 1, 1, 3, 10, 0, 0, 0, 0, 0, 0, 9, 12, 7, 228, 1, 1, 3, 11, 0, 0, 0, 0, 0, 0, 1, 0}

	res = bytes.Compare(t2, result)
	if res != 0 {
		t.Errorf("Test AccessSelectorRange failed. get: %d, should:%v", t2, result)
	}
}

func TestSelectiveAccessDescriptor_Decode(t *testing.T) {
	// ------------------------ AccessSelectorEntry
	src := []byte{2, 2, 4, 6, 0, 0, 0, 0, 6, 0, 0, 0, 5, 18, 0, 0, 18, 0, 0}
	b := *CreateSelectiveAccessDescriptor(AccessSelectorEntry, []uint32{0, 5})

	a, e := DecodeSelectiveAccessDescriptor(&src)
	if e != nil {
		t.Errorf("t1 Failed to Decode. err:%v", e)
	}
	if a.AccessSelector != b.AccessSelector {
		t.Errorf("t1 AccessSelector Failed. get: %d, should:%v", a.AccessSelector, b.AccessSelector)
	}

	aByte, _ := a.AccessParameter.Encode()
	bByte, _ := b.AccessParameter.Encode()

	res := bytes.Compare(aByte, bByte)
	if res != 0 {
		t.Errorf("t1 AccessParameter Failed. get: %d, should:%v", a.AccessParameter.Value, b.AccessParameter.Value)
	}

	// ------------------------ AccessSelectorRange

	src = []byte{1, 2, 4, 2, 4, 18, 0, 8, 9, 6, 0, 0, 1, 0, 0, 255, 15, 2, 18, 0, 0, 9, 12, 7, 228, 1, 1, 3, 10, 0, 0, 0, 0, 0, 0, 9, 12, 7, 228, 1, 1, 3, 11, 0, 0, 0, 0, 0, 0, 1, 0}
	timeStart := time.Date(2020, time.January, 1, 10, 0, 0, 0, time.UTC)
	timeEnd := time.Date(2020, time.January, 1, 11, 0, 0, 0, time.UTC)
	b = *CreateSelectiveAccessDescriptor(AccessSelectorRange, []time.Time{timeStart, timeEnd})

	a, e = DecodeSelectiveAccessDescriptor(&src)
	if e != nil {
		t.Errorf("t2 Failed to Decode. err:%v", e)
	}
	if a.AccessSelector != b.AccessSelector {
		t.Errorf("t2 AccessSelector Failed. get: %d, should:%v", a.AccessSelector, b.AccessSelector)
	}

	aByte, _ = a.AccessParameter.Encode()
	bByte, _ = b.AccessParameter.Encode()

	res = bytes.Compare(aByte, bByte)
	if res != 0 {
		t.Errorf("t2 AccessParameter Failed. get: %d, should:%v", a.AccessParameter.Value, b.AccessParameter.Value)
	}

	// --- making sure src wont change if decode fail
	src = []byte{2, 2, 4, 6, 0, 0, 0, 0, 6, 0, 0, 0, 5, 255, 0, 0, 18, 0, 0}
	oriLength := len(src)
	a, e = DecodeSelectiveAccessDescriptor(&src)
	if e == nil {
		t.Errorf("t3 should fail")
	}
	if len(src) != oriLength {
		t.Errorf("t3. src should not change on fail (%v)", src)
	}
}
