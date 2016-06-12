package otp

import(
	//"github.com/preppedup/proto/sms"
	"github.com/preppedup/proto/database"
	"math/rand"
	"log"
	"github.com/preppedup/proto/sms"
	"strconv"
)

type otp struct{
	id int64
	value int
	mobileNumber string
	status string
}

// TODO check if a live otp is already there and resend it if it is
//Generates an otp if required and sends it to the user's mobile
func SendOTP(mobileNumber string){
	//connect to db to see if there is an active otp against this use case for this number

	//if not create a new otp and add it to the data base
	otp := rand.Intn(8999) + 1000
	db := database.Pool()

	result, err2 := db.Exec("INSERT INTO otps (id, otp_value, mobile_number, status) VALUES (DEFAULT, $1, $2, 'active')", otp, mobileNumber)
	if err2 != nil {
		log.Fatal("Could not insert generated OTP in to the db: " + err2.Error())
	}

	rowsAffected, err3 := result.RowsAffected()
	if err3 != nil {
		log.Fatal("Could not retrieve inserted OTP row from db: " + err3.Error())
	}
	log.Print(rowsAffected)

	sms.Gateway().SendSMS(strconv.Itoa(otp) + " is your PreppedUP verification code.",mobileNumber)

	//send the otp sms to the user.
}


func VerifyOTP(mobileNumber string, otpValue string){
	// fetch the last active otp against this user
	db := database.Pool()
	otprows, err := db.Query("SELECT * FROM otps WHERE mobile_number = $1 AND status = 'active'", mobileNumber)
	if err != nil{
		log.Print("Could not retrieve otp record: " + err.Error())
	}
	defer otprows.Close()

	// if no results then return resource not valid http error
	if otprows == nil{
		log.Print("No valid otps found for this mobile number")
		return
	}

	// if results found compare with otp provided
	otpObject := new(otp)
	numRows := 0
	for otprows.Next(){
		err := otprows.Scan(&otpObject.id, &otpObject.value, &otpObject.mobileNumber, &otpObject.status)
		if err != nil{
			log.Print("Error while scanning an otp row: " + err.Error())
			return
		}
		numRows = numRows + 1
	}

	if err = otprows.Err(); err != nil {
		log.Print("Error while scanning otp rows: " + err.Error())
		return
	}

	if numRows == 0{
		log.Print("No active otps for this mobile number")
		return
	}

	// if there is a match then return 200 OK
	if otpObject.value != 0 {
		actualotp, converr := strconv.Atoi(otpValue)
		if converr != nil{
			log.Print()
		} else {
			if otpObject.value == actualotp{
				log.Print("OTP successfully verified")
				return
			} else {
				// if there is no match then return resource not found error
				log.Print("OTP did not match")
				return
			}
		}
	}
}