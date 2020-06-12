package myiblt

type testcommon interface {
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Fail()
	FailNow()
	Failed() bool
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
	Helper()
	Log(args ...interface{})
	Logf(format string, args ...interface{})
	Name() string
	Skip(args ...interface{})
	SkipNow()
	Skipf(format string, args ...interface{})
	Skipped() bool
}

func assertEqual(t testcommon, a, b interface{}) {
	t.Helper()
	if a != b {
		t.Errorf("Not Equal. %d %d", a, b)
	}
}

func assertTrue(t testcommon, a bool) {
	t.Helper()
	if !a {
		t.Errorf("Not True %t", a)
	}
}

func assertFalse(t testcommon, a bool) {
	t.Helper()
	if a {
		t.Errorf("Not True %t", a)
	}
}

func assertNil(t testcommon, a interface{}) {
	t.Helper()
	if a != nil {
		t.Error("Not Nil")
	}
}

func assertNotNil(t testcommon, a interface{}) {
	t.Helper()
	if a == nil {
		t.Error("Is Nil")
	}
}
