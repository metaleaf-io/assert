package assert

import "testing"

func TestWith(t *testing.T) {
	assert := With(t)

	if assert == nil {
		t.Error("With returned nil.")
	}
}

func TestMatcher_That(t *testing.T) {
	assert := With(t).That(nil)

	if assert == nil {
		t.Error("That returned nil.")
	}
}

func TestMatcher_That_Panics(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("That did not panic")
		}
	}()

	assert := new(Matcher)
	assert.That(nil)
	t.Error("That did not panic.")
}

func TestMatcher_IsNil_WithNil(t *testing.T) {
	assert := With(t).That(nil).IsNil()

	if assert == nil {
		t.Error("IsNil returned nil")
		return
	}

	if assert.match == false {
		t.Error("IsNil matcher failed.")
	}
}

func TestMatcher_IsNil_WithInt(t *testing.T) {
	assert := With(t).That(0).IsNil()

	if assert == nil {
		t.Error("IsNil returned nil")
		return
	}

	if assert.match == true {
		t.Error("IsNil matcher failed.")
	}
}

func TestMatcher_IsNil_WithString(t *testing.T) {
	assert := With(t).That("String").IsNil()

	if assert == nil {
		t.Error("IsNil returned nil")
		return
	}

	if assert.match == true {
		t.Error("IsNil matcher failed.")
	}
}

func TestMatcher_IsNil_WithSlice(t *testing.T) {
	assert := With(t).That(make([]byte, 0)).IsNil()

	if assert == nil {
		t.Error("IsNil returned nil")
		return
	}

	if assert.match == true {
		t.Error("IsNil matcher failed.")
	}
}

func TestMatcher_IsNil_WithObject(t *testing.T) {
	assert := With(t).That(new(Matcher)).IsNil()

	if assert == nil {
		t.Error("IsNil returned nil")
		return
	}

	if assert.match == true {
		t.Error("IsNil matcher failed.")
	}
}

func TestMatcher_IsNotNil_WithNil(t *testing.T) {
	assert := With(t).That(nil).IsNotNil()

	if assert == nil {
		t.Error("IsNil returned nil")
		return
	}

	if assert.match == true {
		t.Error("IsNotNil matcher failed.")
	}
}

func TestMatcher_IsNotNil_WithInt(t *testing.T) {
	assert := With(t).That(0).IsNotNil()

	if assert == nil {
		t.Error("IsNil returned nil")
		return
	}

	if assert.match == false {
		t.Error("IsNotNil matcher failed.")
	}
}

func TestMatcher_IsNotNil_WithString(t *testing.T) {
	assert := With(t).That("String").IsNotNil()

	if assert == nil {
		t.Error("IsNil returned nil")
		return
	}

	if assert.match == false {
		t.Error("IsNotNil matcher failed.")
	}
}

func TestMatcher_IsNotNil_WithSlice(t *testing.T) {
	assert := With(t).That(make([]byte, 0)).IsNotNil()

	if assert == nil {
		t.Error("IsNil returned nil")
		return
	}

	if assert.match == false {
		t.Error("IsNotNil matcher failed.")
	}
}

func TestMatcher_IsNotNil_WithObject(t *testing.T) {
	assert := With(t).That(new(Matcher)).IsNotNil()

	if assert == nil {
		t.Error("IsNil returned nil")
		return
	}

	if assert.match == false {
		t.Error("IsNotNil matcher failed.")
	}
}

func TestMatcher_Equals_WithNil(t *testing.T) {
	assert := With(t).That(nil).IsEqualTo(nil)

	if assert == nil {
		t.Error("IsEqualTo returned nil")
		return
	}

	if assert.match == false {
		t.Error("IsEqualTo matcher failed")
	}
}


func TestMatcher_Equals_WithBool(t *testing.T) {
	assert := With(t).That(true).IsEqualTo(true)

	if assert == nil {
		t.Error("IsEqualTo returned nil")
		return
	}

	if assert.match == false {
		t.Error("IsEqualTo matcher failed")
	}
}

func TestMatcher_Equals_WithComplex(t *testing.T) {
	assert := With(t).That(complex(1.0, 1.0)).IsEqualTo(complex(1.0, 1.0))

	if assert == nil {
		t.Error("IsEqualTo returned nil")
		return
	}

	if assert.match == false {
		t.Error("IsEqualTo matcher failed")
	}
}

func TestMatcher_Equals_WithFloat(t *testing.T) {
	assert := With(t).That(3.14159).IsEqualTo(3.14159)

	if assert == nil {
		t.Error("IsEqualTo returned nil")
		return
	}

	if assert.match == false {
		t.Error("IsEqualTo matcher failed")
	}
}

func TestMatcher_Equals_WithInt(t *testing.T) {
	assert := With(t).That(-1073741824).IsEqualTo(-1073741824)

	if assert == nil {
		t.Error("IsEqualTo returned nil")
		return
	}

	if assert.match == false {
		t.Error("IsEqualTo matcher failed")
	}
}

func TestMatcher_Equals_WithString(t *testing.T) {
	assert := With(t).That("The quick brown fox jumps over the lazy dog").IsEqualTo("The quick brown fox jumps over the lazy dog")

	if assert == nil {
		t.Error("IsEqualTo returned nil")
		return
	}

	if assert.match == false {
		t.Error("IsEqualTo matcher failed")
	}
}

func TestMatcher_Equals_WithUint(t *testing.T) {
	assert := With(t).That(uint(1073741824)).IsEqualTo(uint(1073741824))

	if assert == nil {
		t.Error("IsEqualTo returned nil")
		return
	}

	if assert.match == false {
		t.Error("IsEqualTo matcher failed")
	}
}

func TestMatcher_Equals_WithDifferentTypes(t *testing.T) {
	tt := new(testing.T)
	assert := With(tt).That(true).IsEqualTo(1.0)

	if assert == nil {
		t.Error("IsEqualTo returned nil")
		return
	}

	if assert.match == true {
		t.Error("IsEqualTo matcher failed")
	}
}
