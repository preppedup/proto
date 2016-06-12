package sms

import (
	"net/http"
	"net/url"
	"fmt"
)

type Smsgateway interface{
	SendSMS(text string, mobileNumber string) string
}

// implements SMSgateway interface
type smsVendor struct {
	Name string
}

var primaryVendor *smsVendor;

// initializes the vendor sms gateway(s) if its not already initialized and returns a reference to it.
func Gateway() Smsgateway{
	if primaryVendor == nil{
		primaryVendor = &smsVendor{"msg91"}
	}
	return primaryVendor
}

func(vendor *smsVendor) SendSMS(text string, mobileNumber string) string{
	if vendor.Name == "msg91"{
		var Url *url.URL
		Url, err := url.Parse("https://control.msg91.com")
		if (err != nil){
			panic("Vendor Url has failed parse attempt")
		}

		Url.Path += "/api/sendhttp.php"
		parameters := url.Values{}
		parameters.Add("authkey","114997ACigZfy8H5753b8c8")
		parameters.Add("mobiles", mobileNumber)
		parameters.Add("message", text)
		parameters.Add("sender", "PREPUP")
		parameters.Add("route","4")
		parameters.Add("country","91")

		Url.RawQuery = parameters.Encode()
		fmt.Printf(Url.String())

		resp, err := http.Get(Url.String())
		fmt.Print(resp)
		if err != nil{
			fmt.Printf(err.Error())
		}
	}
	return "whatever";
}
