package main

import "testing"

func TestCreatedDateYearSupported(t *testing.T) {
	cases := []struct {
		date, content string
		want          bool
	}{
		{"2013-12-14", "BURGER KING 12/14/2013 DRIVE THRU", true},
		{"2019-06-06", "Walmart 06/06/19", true},
		{"2021-03-04", "Handwritten note: call mom 3/4", false},
		{"2025-01-01", "There appears to be no visible text in this image.", false},
		{"", "anything", true},
		{"not a date", "anything", true},
	}
	for _, c := range cases {
		if got := createdDateYearSupported(c.date, c.content); got != c.want {
			t.Errorf("createdDateYearSupported(%q, %q) = %v, want %v", c.date, c.content, got, c.want)
		}
	}
}
