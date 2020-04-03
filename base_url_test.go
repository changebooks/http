package http

import "testing"

func TestHasScheme(t *testing.T) {
	got := HasScheme("http://t.cn")
	if got != true {
		t.Errorf("got %v; want true", got)
	}

	got2 := HasScheme("https://t.cn")
	if got2 != true {
		t.Errorf("got %v; want true", got2)
	}

	got3 := HasScheme("t.cn")
	if got3 != false {
		t.Errorf("got %v; want false", got3)
	}
}

func TestIsDefaultPort(t *testing.T) {
	got := IsDefaultPort("http", 80)
	if got != true {
		t.Errorf("got %v; want true", got)
	}

	got2 := IsDefaultPort("https", 443)
	if got2 != true {
		t.Errorf("got %v; want true", got2)
	}

	got3 := IsDefaultPort("", 80)
	if got3 != true {
		t.Errorf("got %v; want true", got3)
	}

	got4 := IsDefaultPort("https", 80)
	if got4 != false {
		t.Errorf("got %v; want false", got4)
	}

	got5 := IsDefaultPort("http", 443)
	if got5 != false {
		t.Errorf("got %v; want false", got5)
	}
}

func TestBaseUrl(t *testing.T) {
	builder := &BaseUrlBuilder{}
	got, _ := builder.SetHost("http://t.cn:81").Build()
	if got.GetScheme() != "http" {
		t.Errorf("got %v; want %v", got.GetScheme(), "http")
	}
	if got.GetHost() != "t.cn" {
		t.Errorf("got %v; want %v", got.GetHost(), "t.cn")
	}
	if got.GetPort() != 81 {
		t.Errorf("got %v; want %v", got.GetPort(), 81)
	}
	if got.GetUrl() != "http://t.cn:81/" {
		t.Errorf("got %v; want %v", got.GetUrl(), "http://t.cn:81/")
	}

	builder2 := &BaseUrlBuilder{}
	got2, _ := builder2.SetScheme("https://").SetHost("http://t.cn:81").Build()
	if got2.GetScheme() != "https" {
		t.Errorf("got %v; want %v", got2.GetScheme(), "https")
	}
	if got2.GetHost() != "t.cn" {
		t.Errorf("got %v; want %v", got2.GetHost(), "t.cn")
	}
	if got2.GetPort() != 81 {
		t.Errorf("got %v; want %v", got2.GetPort(), 81)
	}
	if got2.GetUrl() != "https://t.cn:81/" {
		t.Errorf("got %v; want %v", got2.GetUrl(), "https://t.cn:81/")
	}

	builder3 := &BaseUrlBuilder{}
	got3, _ := builder3.SetScheme("https://").SetHost("http://t.cn:81").SetPort(82).Build()
	if got3.GetScheme() != "https" {
		t.Errorf("got %v; want %v", got3.GetScheme(), "https")
	}
	if got3.GetHost() != "t.cn" {
		t.Errorf("got %v; want %v", got3.GetHost(), "t.cn")
	}
	if got3.GetPort() != 82 {
		t.Errorf("got %v; want %v", got3.GetPort(), 82)
	}
	if got3.GetUrl() != "https://t.cn:82/" {
		t.Errorf("got %v; want %v", got3.GetUrl(), "https://t.cn:82/")
	}

	builder4 := &BaseUrlBuilder{}
	got4, _ := builder4.SetScheme("https://").SetHost("t.cn").SetPort(443).Build()
	if got4.GetUrl() != "https://t.cn/" {
		t.Errorf("got %v; want %v", got4.GetUrl(), "https://t.cn/")
	}

	builder5 := &BaseUrlBuilder{}
	got5, _ := builder5.SetScheme("http://").SetHost("t.cn").SetPort(80).Build()
	if got5.GetUrl() != "http://t.cn/" {
		t.Errorf("got %v; want %v", got4.GetUrl(), "http://t.cn/")
	}
}
