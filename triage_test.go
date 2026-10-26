package main

import "testing"

func TestTriageComplete(t *testing.T) {
	req := GenerateSuggestionsRequest{GenerateTitles: true, GenerateTags: true, GenerateDocumentTypes: true, IsAutoProcessing: true}
	tags := []string{"2026", "Atari", "needstriage"}
	types := []string{"Receipts", "Bills"}
	cases := []struct {
		name    string
		title   string
		llmTags []string
		docType string
		want    bool
	}{
		{"complete", "Canyon View - Receipt", []string{"2026"}, "Receipts", true},
		{"old year only, model gave none", "T", []string{""}, "Receipts", false},
		{"no type", "T", []string{"2026"}, "", false},
		{"type not in list", "T", []string{"2026"}, "Invoices", false},
		{"no title", " ", []string{"2026"}, "Receipts", false},
		{"year not an existing tag", "T", []string{"1999"}, "Receipts", false},
		{"non-year tag only", "T", []string{"Atari"}, "Receipts", false},
	}
	for _, c := range cases {
		if got := triageComplete(req, c.title, c.llmTags, tags, c.docType, types); got != c.want {
			t.Errorf("%s: got %v want %v", c.name, got, c.want)
		}
	}
	off := req
	off.GenerateDocumentTypes = false
	if triageComplete(off, "T", []string{"2026"}, tags, "Receipts", types) {
		t.Error("doc type generation disabled should not clear triage")
	}
}
