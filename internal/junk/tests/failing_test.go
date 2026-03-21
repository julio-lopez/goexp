package junk_test

import "testing"

func TestGotestsumOutputWithFailingTest(t *testing.T) {
	t.Skip("this test always fails, comment out the 'Skip' to have it run and fail")

	t.Log("test output")
	t.Error("fail test")
}
