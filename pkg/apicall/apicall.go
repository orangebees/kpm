package apicall

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/valyala/fasthttp"
)

const (
	None = iota
	InQuery
	InHeader
)

type CallToken struct {
	name string
	val  string
}
type ApiCall struct {
	addr        string
	tokenName   string
	TokenVal    any
	tokenStats  int
	getToken    func() string
	compression bool
	//token位置，名字，结果
}
type ApiCtx struct {
	Req  *fasthttp.Request
	Resp *fasthttp.Response
}

func New(addr string) *ApiCall {
	return &ApiCall{addr: addr,
		tokenName:   "accesstoken",
		compression: true,
		getToken: func() string {
			return ""
		},
	}
}
func (a *ApiCall) SetGetTokenFunction(f func() string) {
	a.getToken = f
}
func (a *ApiCall) SetTokenInQuery() {
	a.tokenStats = InQuery
}
func (a *ApiCall) SetTokenInHeader() {
	a.tokenStats = InHeader
}
func (a *ApiCall) SetTokenNone() {
	a.tokenStats = None
}
func (a *ApiCall) SetCompression() {
	a.compression = true
}
func (a *ApiCtx) SetRequestJsonArgs(val any) error {
	bytes, err := json.Marshal(val)
	if err != nil {
		return err
	}
	a.Req.SetBody(bytes)
	return nil
}
func AcquireApiCtx() *ApiCtx {
	ac := &ApiCtx{}
	ac.Req = fasthttp.AcquireRequest()
	ac.Resp = fasthttp.AcquireResponse()
	return ac
}
func ReleaseApiCtx(ac *ApiCtx) {
	fasthttp.ReleaseRequest(ac.Req)
	fasthttp.ReleaseResponse(ac.Resp)
}

func Test() {
	apic := New("https://www.baidu.com")
	apic.SetTokenInQuery()
	ac := AcquireApiCtx()
	ac.SetRequestJsonArgs(ac)
	defer ReleaseApiCtx(ac)
	err := apic.Call("/5", ac, ac, UseGetMethod, PrintRequest)
	if err != nil {
		fmt.Println(err)
		return
	}

}

type ApiOption func(*ApiCtx)

func UseGetMethod(ac *ApiCtx) {
	ac.Req.Header.SetMethod("GET")
}
func PrintRequest(ac *ApiCtx) {
	fmt.Println(ac.Req.Header.String())
	fmt.Println(ac.Req.Body())
}

// CallLite 轻量版无依赖apicall
func (a *ApiCall) CallLite(apiName string, ac *ApiCtx, option ...ApiOption) error {
	ac.Req.Header.SetMethod("POST")
	ac.Req.SetRequestURI(a.addr + apiName)
	//ac.Req.Header.SetContentType("text/html")
	optlen := len(option)
	for i := 0; i < optlen; i++ {
		option[i](ac)
	}
	//fmt.Println(ac.Req.URI().String())
	if err := fasthttp.Do(ac.Req, ac.Resp); err != nil {
		//连接错误
		return err
	}
	if ac.Resp.StatusCode() != 200 {
		//业务错误

		return errors.New("意外响应")
	}
	return nil
}
func (a *ApiCall) Call(apiName string, ac *ApiCtx, Resp any, option ...ApiOption) error {
	ac.Req.Header.SetMethod("POST")
	ac.Req.SetRequestURI(a.addr + apiName)
	ac.Req.Header.SetContentType("application/json")
	if a.compression {
		ac.Req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	}
	switch a.tokenStats {
	case None:
	case InHeader:
		ac.Req.Header.Set(a.tokenName, a.getToken())
	case InQuery:
		ac.Req.URI().QueryArgs().Set(a.tokenName, a.getToken())
	default:

	}
	optlen := len(option)
	for i := 0; i < optlen; i++ {
		option[i](ac)
	}
	fmt.Println(ac.Req.URI().String())
	if err := fasthttp.Do(ac.Req, ac.Resp); err != nil {
		//连接错误
		return err
	}
	//fmt.Println(ac.Resp.Header.String())
	//fmt.Println(ac.Resp.Body())
	if ac.Resp.StatusCode() != 200 {
		//业务错误
		return errors.New("意外响应")
	}
	body, err := ac.Resp.BodyUncompressed()
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, Resp)
	if err != nil {
		return err
	}
	return nil
}

//func Call1(addr string, router string, arg *fasthttp.Args, cookie []byte) ([]byte, int) {
//	Req := fasthttp.AcquireRequest()
//	defer fasthttp.ReleaseRequest(Req)
//	//fmt.Println("kk")
//	Req.Header.SetMethod("POST")
//	if cookie != nil {
//		Req.Header.SetBytesV("Cookie", cookie)
//	}
//	Req.SetRequestURI(addr + router)
//	Req.Header.SetContentType("application/x-www-form-urlencoded")
//	if arg != nil {
//		if _, err := arg.WriteTo(Req.BodyWriter()); err != nil {
//		}
//	}
//	Resp := fasthttp.AcquireResponse()
//	defer fasthttp.ReleaseResponse(Resp) // 用完需要释放资源
//	if err := fasthttp.Do(Req, Resp); err != nil {
//		//请求失败
//		//fmt.Println(Req.Header.String())
//		//fmt.Println(Resp.Header.String())
//		//fmt.Println(err)
//		//fmt.Println("jj")
//		return nil, CALLERR
//	}
//	//fmt.Println(Req.Header.String())
//	//fmt.Println(Resp.Header.String())
//	switch Resp.StatusCode() {
//	case 200:
//		return Resp.Header.Peek("Set-Cookie"), OK
//		//业务正常
//	case 401:
//		return nil, ERR
//		//授权失败，登录失效
//	case 500:
//		return nil, AUTHERR
//		//内部处理异常
//	default:
//		return nil, ERR
//		//意料之外的异常
//	}
//
//}
