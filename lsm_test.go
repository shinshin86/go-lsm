package lsm

import "testing"

func TestNewLsm(t *testing.T) {
	lsmk := New()

	got := lsmk.Length()
	if got != 0 {
		t.Errorf("lsm.Key() = %d; want 0", got)
	}
}

func TestLengthLsm(t *testing.T) {
	lsmk := New()

	got := lsmk.Length()
	if got != 0 {
		t.Errorf("lsm.Length() = %d; want 0", got)
	}

	s, _ := lsmk.Set("testkey", "testvalue")

	if s["testkey"] != "testvalue" {
		t.Errorf("lsm.Set(\"testkey\") = %s; want testvalue", s["testkey"])
	}

	got2 := lsmk.Length()
	if got2 != 1 {
		t.Errorf("lsm.Length() = %d; want 1", got2)
	}
}

func TestSetLsm(t *testing.T) {
	lsmk := New()

	got := lsmk.Length()
	if got != 0 {
		t.Errorf("lsm.Length() = %d; want 0", got)
	}

	s, _ := lsmk.Set("testkey", "testvalue")

	if s["testkey"] != "testvalue" {
		t.Errorf("lsm.Set(\"testkey\") = %s; want testvalue", s["testkey"])
	}

	got2 := lsmk.Length()
	if got2 != 1 {
		t.Errorf("lsm.Length() = %d; want 1", got2)
	}

	s2, _ := lsmk.Set("testkey2", "testvalue2")

	if s2["testkey2"] != "testvalue2" {
		t.Errorf("lsm.Set(\"testkey\") = %s; want testvalue", s["testkey"])
	}

	got3 := lsmk.Length()
	if got3 != 2 {
		t.Errorf("lsm.Length() = %d; want 1", got3)
	}
}

func TestGetLsm(t *testing.T) {
	lsmk := New()

	s, _ := lsmk.Set("testkey", "testvalue")

	if s["testkey"] != "testvalue" {
		t.Errorf("lsm.Set(\"testkey\") = %s; want testvalue", s["testkey"])
	}

	v, _ := lsmk.Get("testkey")
	if v != "testvalue" {
		t.Errorf("lsm.Get(\"testkey\") = %s; want testvalue", s["testkey"])
	}

	s2, _ := lsmk.Set("testkey2", "testvalue2")

	if s2["testkey2"] != "testvalue2" {
		t.Errorf("lsm.Set(\"testkey\") = %s; want testvalue", s["testkey"])
	}

	got3 := lsmk.Length()
	if got3 != 2 {
		t.Errorf("lsm.Length() = %d; want 1", got3)
	}

	v2, _ := lsmk.Get("testkey2")
	if v2 != "testvalue2" {
		t.Errorf("lsm.Get(\"testkey\") = %s; want testvalue", s["testkey"])
	}

	_, err := lsmk.Get("testkey3")
	if err == nil {
		t.Errorf("want error message")
	}
}

func TestRemoveLsm(t *testing.T) {
	lsmk := New()

	s, _ := lsmk.Set("testkey", "testvalue")

	if s["testkey"] != "testvalue" {
		t.Errorf("lsm.Set(\"testkey\") = %s; want testvalue", s["testkey"])
	}

	v, _ := lsmk.Get("testkey")
	if v != "testvalue" {
		t.Errorf("lsm.Get(\"testkey\") = %s; want testvalue", s["testkey"])
	}

	lsmk.Remove("testkey")

	got := lsmk.Length()
	if got != 0 {
		t.Errorf("lsm.Length() = %d; want 0", got)
	}

	v2, err := lsmk.Get("testkey")
	if err == nil {
		t.Errorf("lsmk.Get(\"testkey\") = %s; want nil", v2)
	}
}

func TestClearLsm(t *testing.T) {
	lsmk := New()

	s, _ := lsmk.Set("testkey", "testvalue")

	if s["testkey"] != "testvalue" {
		t.Errorf("lsm.Set(\"testkey\") = %s; want testvalue", s["testkey"])
	}

	v, _ := lsmk.Get("testkey")
	if v != "testvalue" {
		t.Errorf("lsm.Get(\"testkey\") = %s; want testvalue", s["testkey"])
	}

	s2, _ := lsmk.Set("testkey2", "testvalue2")

	if s2["testkey2"] != "testvalue2" {
		t.Errorf("lsm.Set(\"testkey\") = %s; want testvalue", s["testkey"])
	}

	got3 := lsmk.Length()
	if got3 != 2 {
		t.Errorf("lsm.Length() = %d; want 1", got3)
	}

	v2, _ := lsmk.Get("testkey2")
	if v2 != "testvalue2" {
		t.Errorf("lsm.Get(\"testkey\") = %s; want testvalue", s["testkey"])
	}

	lsmk.Clear()

	got := lsmk.Length()
	if got != 0 {
		t.Errorf("lsm.Length() = %d; want 0", got)
	}

	v3, err := lsmk.Get("testkey")
	if err == nil {
		t.Errorf("lsmk.Get(\"testkey\") = %s; want nil", v3)
	}

	v4, err := lsmk.Get("testkey2")
	if err == nil {
		t.Errorf("lsmk.Get(\"testkey\") = %s; want nil", v4)
	}
}
