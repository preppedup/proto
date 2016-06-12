package routes

import (
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/preppedup/proto/otp"
)

type Route struct{
	Name string
	Method string
	Pattern string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"Otps",
		"POST",
		"/mobiles/{mobileNum}/otps",
		OtpGen,
	},
	Route{
		"Otp",
		"GET",
		"/mobiles/{mobileNum}/otps/{otp}",
		OtpVerify,
	},
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")

}

func OtpGen(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)

	//validate mobile number and other data inputs right here.
	otp.SendOTP(vars["mobileNum"])
	fmt.Fprintln(w, "Otp send attempted")

}

func OtpVerify(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	otpValue := vars["otp"]
	mobileNumber := vars["mobileNum"]
	otp.VerifyOTP(mobileNumber, otpValue)
	fmt.Fprintln(w, "You tried to verify ", otpValue)
}