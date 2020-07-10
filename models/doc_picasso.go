package models

// 点击分类选项卡 ------------------------------------------------------------------
// curl -X GET
// -H "Host:service.picasso.adesk.com"
// -H "Connection:Keep-Alive"
// -H "User-Agent:picasso,268,tencent"
// -H "Accept-Encoding:gzip"
// -H "X-Transmission-Session-Id:"
// "http://service.picasso.adesk.com/v1/vertical/category?adult=false&first=1"

// 点击某一分类, 相当于点击 "最新" 选项卡
// curl -X GET
// -H "Host:service.picasso.adesk.com"
// -H "Connection:Keep-Alive"
// -H "User-Agent:picasso,268,tencent"
// -H "Accept-Encoding:gzip"
// -H "X-Transmission-Session-Id:"
// "http://service.picasso.adesk.com/v1/vertical/category/4e4d610cdf714d2966000000/vertical?limit=30&adult=false&first=1&order=new"

// 手动点击 "热门" 选项卡
// curl -X GET
// -H "Host:service.picasso.adesk.com"
// -H "Connection:Keep-Alive"
// -H "User-Agent:picasso,268,tencent"
// -H "Accept-Encoding:gzip"
// -H "X-Transmission-Session-Id:"
// "http://service.picasso.adesk.com/v1/vertical/category/4e4d610cdf714d2966000000/vertical?limit=30&adult=false&first=1&order=hot"

// 翻页
// curl -X GET
// -H "Host:service.picasso.adesk.com"
// -H "Connection:Keep-Alive"
// -H "User-Agent:picasso,268,tencent"
// -H "Accept-Encoding:gzip"
// -H "X-Transmission-Session-Id:"
// "http://service.picasso.adesk.com/v1/vertical/category/4e4d610cdf714d2966000000/vertical?limit=30&skip=30&adult=false&first=0&order=hot"
// first 貌似一值为0

// 点击图片, 之后进入的是 "查看较大分辨率的图片" 页面
// 之后再点击 "较大", 进入 "查看最大分辨率的图片" 页面 (全屏)
// http://img5.adesk.com/5eff4198e7bce75e7ef4c25e (不详细解释这步, 因为这里的实际 url 比较奇怪)
// 较大 url: GET /5f03f14ce7bce75d8f8d0ee1?imageMogr2/thumbnail/!720x1280r/gravity/Center/crop/720x1280 HTTP/1.1
// 最大 url: GET /5f03f14ce7bce75d8f8d0ee1?imageMogr2/thumbnail/!1080x1920r/gravity/Center/crop/1080x1920&adult=false HTTP/1.1

// 点击专辑选项卡 --------------------------------------------------------------------------
// curl -X GET
// -H "Host:service.picasso.adesk.com"
// -H "Connection:Keep-Alive"
// -H "User-Agent:picasso,268,tencent"
// -H "Accept-Encoding:gzip"
// -H "X-Transmission-Session-Id:"
// "http://service.picasso.adesk.com/v1/wallpaper/album?limit=10&adult=false&first=1&order=new"

// 翻页
// curl -X GET
// -H "Host:service.picasso.adesk.com"
// -H "Connection:Keep-Alive"
// -H "User-Agent:picasso,268,tencent"
// -H "Accept-Encoding:gzip"
// -H "X-Transmission-Session-Id:"
//"http://service.picasso.adesk.com/v1/wallpaper/album?limit=10&skip=10&adult=false&first=0&order=new"
