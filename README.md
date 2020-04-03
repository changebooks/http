# http
http request
==

<pre>
urlBuilder := http.BaseUrlBuilder{}
baseUrl, err := urlBuilder.SetHost("https://www.baidu.com").Build()
if err != nil {
    fmt.Println(err)
    return
}

talkBuilder := http.TalkBuilder{}
talk, err := talkBuilder.SetBaseUrl(baseUrl).SetUseTrace(true).Build()
if err != nil {
    fmt.Println(err)
    return
}

params := make(map[string]string)
params["wd"] = "abc"

req, url, err := talk.NewRequest(http.MethodGet, "s", params)
if err != nil {
    fmt.Println(url)
    fmt.Println(err)
    return
}

s := talk.Request(req, 0)
fmt.Println(s.ToString())
if s.Elapsed != nil {
    fmt.Println(s.Elapsed.ToString())
}
</pre>