package main

import (
	"reflect"
	"testing"
)

func TestYearInContent(t *testing.T) {
	cases := []struct {
		year, content string
		want          bool
	}{
		{"2024", "Statement date: 03/15/2024", true},
		{"2024", "Invoice 2024-03-15", true},
		{"2019", "Walmart 06/06/19 14:22", true},
		{"2019", "Fuel 10-16-19", true},
		{"2025", "There appears to be no visible text in this image.", false},
		{"2024", "Account 12024555", false},
		{"2016", "161310 137410", false},
		{"2019", "Total 19.00", false},
	}
	for _, c := range cases {
		if got := yearInContent(c.year, c.content); got != c.want {
			t.Errorf("yearInContent(%q, %q) = %v, want %v", c.year, c.content, got, c.want)
		}
	}
}

func TestFilterUnsupportedYearTags(t *testing.T) {
	got := filterUnsupportedYearTags([]string{"2025", "Micah", ""}, "photo of wooden planks", nil)
	want := []string{"Micah", ""}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
	got = filterUnsupportedYearTags([]string{"2026", "T.J."}, "Treatment plan 10/01/2026", nil)
	want = []string{"2026", "T.J."}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
