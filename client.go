package vip_gosdk

import (
	"crypto/hmac"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/fatih/structs"
	jsoniter "github.com/json-iterator/go"
	"github.com/spf13/cast"
	"net/http"
	"net/url"
	"reflect"
	"sort"
	"strings"
	"time"
)

// The client saves some private required fields.
// Before executing the request, a new Client must be created.
type Client struct {
	appKey     string
	secret     string
	userNumber string
}

// NewClient returns a new Client.
// appKey, secret, userNumber can not be empty.
func NewClient(appKey, secret, userNumber string) *Client {
	if appKey == "" || secret == "" || userNumber == "" {
		panic("required parameters cannot be empty")
	}
	return &Client{
		appKey:     appKey,
		secret:     secret,
		userNumber: userNumber,
	}
}

// The Request describes a request's method and request request type
type Request interface {
	ApiMethod() string
	RequestMethod() string
}

// The Request describes a request's return information
type Response interface{}

// General request.
func (client *Client) Do(req Request, params map[string]interface{}, res Response) (err error) {
	r := map[string]interface{}{"request": params}
	// Body
	body, err := jsoniter.MarshalToString(r)
	if err != nil {
		return
	}
	// Initialization timestamp
	t := cast.ToString(time.Now().Unix())
	// Get signature
	sign := client.getVipSign(body, req.ApiMethod(), t)
	// Request URL: k-v key-value pair splicing
	v := url.Values{}
	v.Set("appKey", client.appKey)
	v.Set("format", Format)
	v.Set("method", req.ApiMethod())
	v.Set("service", Service)
	v.Set("sign", sign)
	v.Set("timestamp", t)
	v.Set("version", Version)
	domain := Domain
	_url, err := url.Parse(domain)
	if err != nil {
		return
	}
	_url.RawQuery = v.Encode()
	urlPath := _url.String()
	// NewRequest
	request, err := http.NewRequest(req.RequestMethod(), urlPath, strings.NewReader(body))
	if err != nil {
		return
	}
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(res)
	if err != nil {
		return
	}
	return
}

// Signature encryption string.
type VipApiEncrypt struct {
	AppKey    string `json:"appKey"`
	Format    string `json:"format"`
	Method    string `json:"method"`
	Service   string `json:"service"`
	Timestamp string `json:"timestamp"`
	Version   string `json:"version"`
}

// Obtain VIP signature, internal implementation.
func (client *Client) getVipSign(reqString, apiMethod, t string) (signString string) {
	p := VipApiEncrypt{
		AppKey:    client.appKey,
		Format:    Format,
		Method:    apiMethod,
		Service:   Service,
		Timestamp: t,
		Version:   Version,
	}
	resMap := structs.Map(p)
	sortedKeys := make([]string, 0, len(resMap))
	for k := range resMap {
		sortedKeys = append(sortedKeys, k)
	}
	sort.Strings(sortedKeys)

	for _, k := range sortedKeys {
		value := fmt.Sprintf("%v", cast.ToString(resMap[k]))
		if value != "" {
			signString += getKeyName(k, &p) + value
		}
	}
	return strings.ToUpper(Hmac(client.secret, signString+reqString))
}

// Generate hmac-md 5 encryption.
func Hmac(key, data string) string {
	h := hmac.New(md5.New, []byte(key))
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum([]byte("")))
}

// Get the corresponding Key Name according to the tag name.
func getKeyName(key string, req *VipApiEncrypt) string {
	s := reflect.TypeOf(req).Elem()
	f, _ := s.FieldByName(key)
	return f.Tag.Get("json")
}
