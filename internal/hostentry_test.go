package internal

import (
	"reflect"
	"regexp"
	"testing"
)

func TestApplyRegexp(t *testing.T) {
	re := regexp.MustCompile(`^(?P<thing>.+)$`)
	res := ApplyRegexp(re, "loempia")

	if len(res) != 1 {
		t.Fatalf("Contains more then 1 key: %v", res)
	}

	if res["thing"] != "loempia" {
		t.Fatalf("Result was something else then 'loempia': %s", res["thing"])
	}
}

func TestCanResolve(t *testing.T) {
	tests := map[string]struct {
		input string
		he    HostEntry
		want  bool
	}{
		"empty hostentry":     {input: "example.com", he: HostEntry{}, want: false},
		"empty input":         {input: "", he: HostEntry{}, want: false},
		"empty input with he": {input: "", he: HostEntry{Name: "example.com"}, want: false},
		"same name":           {input: "example.com", he: HostEntry{Name: "example.com"}, want: true},
		"he is subdomain":     {input: "example.com", he: HostEntry{Name: "sub.example.com"}, want: false},
		"input is subdomain":  {input: "sub.example.com", he: HostEntry{Name: "example.com"}, want: true},
	}

	for name, tc := range tests {
		got := tc.he.CanResolve(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("%s: expected: %v, got: %v", name, tc.want, got)
		}
	}
}

func TestWouldResolve(t *testing.T) {
	tests := map[string]struct {
		input string
		he    HostEntry
		want  bool
	}{
		"empty hostentry":     {input: "example.com", he: HostEntry{}, want: false},
		"empty input":         {input: "", he: HostEntry{}, want: false},
		"empty input with he": {input: "", he: HostEntry{Name: "example.com"}, want: false},
		"same name":           {input: "example.com", he: HostEntry{Name: "example.com"}, want: true},
		"he is subdomain":     {input: "example.com", he: HostEntry{Name: "sub.example.com"}, want: true},
		"input is subdomain":  {input: "sub.example.com", he: HostEntry{Name: "example.com"}, want: false},
	}

	for name, tc := range tests {
		got := tc.he.WouldResolve(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("%s: expected: %v, got: %v", name, tc.want, got)
		}
	}
}

func TestSwapIn(t *testing.T) {
	he := HostEntry{Name: "old"}
	he.SwapInName("new")

	if len(he.Aliasses) != 1 {
		t.Fatalf("Expected aliasses to contain an entry: %v", he)
	}
	if he.Aliasses[0] != "old" {
		t.Fatalf("Expected aliasses to contain the 'old' entry: %v", he)
	}
	if he.Name != "new" {
		t.Fatalf("Expected name to be the 'new' entry: %v", he)
	}
}
