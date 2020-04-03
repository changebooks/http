package http

import "testing"

func TestUrlPath(t *testing.T) {
	params := make(map[string]string)
	params["a"] = "1"
	params["b"] = "2"

	got := UrlPath("", params)
	want1 := "a=1&b=2"
	want2 := "b=2&a=1"
	if got != want1 && got != want2 {
		t.Errorf("got %q; want %q or %q", got, want1, want2)
	}

	got2 := UrlPath("t.cn", params)
	want21 := "t.cn?a=1&b=2"
	want22 := "t.cn?b=2&a=1"
	if got2 != want21 && got2 != want22 {
		t.Errorf("got %q; want %q or %q", got2, want21, want22)
	}

	got3 := UrlPath("t.cn?123", params)
	want31 := "t.cn?123&a=1&b=2"
	want32 := "t.cn?123&b=2&a=1"
	if got3 != want31 && got3 != want32 {
		t.Errorf("got %q; want %q or %q", got3, want31, want32)
	}
}

func TestUrlQuery(t *testing.T) {
	params := make(map[string]string)
	params["a"] = "1"

	got := UrlQuery(params)
	want := "a=1"
	if got != want {
		t.Errorf("got %q; want %q", got, want)
	}

	params["b"] = "2"
	got2 := UrlQuery(params)
	want21 := "a=1&b=2"
	want22 := "b=2&a=1"
	if got2 != want21 && got2 != want22 {
		t.Errorf("got %q; want %q or %q", got2, want21, want22)
	}
}
