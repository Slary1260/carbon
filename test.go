package carbon

// SetTestNow sets a test Carbon instance (real or mock) to be returned when a "now" instance is created.
func SetTestNow(carbon Carbon) Carbon {
	return NewCarbon().SetTestNow(carbon)
}

// SetTestNow sets a test Carbon instance (real or mock) to be returned when a "now" instance is created.
func (c Carbon) SetTestNow(carbon Carbon) Carbon {
	c.time = carbon.time
	c.loc = carbon.loc
	c.weekStartsAt = carbon.weekStartsAt
	c.lang = carbon.lang
	c.isTestNow = true
	return c
}

// ClearTestNow clears a test Carbon instance (real or mock) to be returned when a "now" instance is created.
func ClearTestNow() Carbon {
	return NewCarbon().ClearTestNow()
}

// ClearTestNow clears a test Carbon instance (real or mock) to be returned when a "now" instance is created.
func (c Carbon) ClearTestNow() Carbon {
	c.isTestNow = false
	return c
}

// IsTestNow reports whether the test now time is set.
func IsTestNow() bool {
	return NewCarbon().IsTestNow()
}

// IsTestNow reports whether the test now time is set.
func (c Carbon) IsTestNow() bool {
	return c.isTestNow
}
