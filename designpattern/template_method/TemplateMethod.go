package template_method

import "fmt"

type iOtp interface {  // One Time Password, OTP
	genRandomOTP(int) string
	saveOTPCache(string)
	getMessage(string) string
	sendNotification(string) error
	publishMetric()
}

type otp struct{
	iOtp iOtp
}
func (o *otp) genAndSendOTP(otpLength int) error {
	otp := o.iOtp.genRandomOTP(otpLength)
	o.iOtp.saveOTPCache(otp)
	message := o.iOtp.getMessage(otp)
	err := o.iOtp.sendNotification(message)
	if err!=nil {
		return err
	}
	o.iOtp.publishMetric()
	return nil
}

type sms struct {}
func (s *sms) genRandomOTP(len int) string {
	randomOTP := "1234"
	fmt.Printf("SMS: generating random otp %s\n", randomOTP)
	return randomOTP
}
func (s *sms) saveOTPCache(otp string) {
	fmt.Printf("SMS: saving otp:%s to cache\n",otp)
}
func (s *sms) getMessage(otp string) string {
	return "SMS OTP for login is "+otp
}
func(s *sms) sendNotification(message string) error {
	fmt.Printf("SMS: sending sms: %s\n", message)
	return nil
}
func (s *sms) publishMetric() {
	fmt.Printf("SMS: publishing metrics\n")
}

type email struct {}
func (e *email) genRandomOTP(len int) string {
	randomOTP := "4321"
	fmt.Printf("EMAIL: generating random otp %s\n", randomOTP)
	return randomOTP
}
func (e *email) saveOTPCache(otp string) {
	fmt.Printf("EMAIL: saving otp: %s to cache\n",otp)
}
func (e *email) getMessage(otp string) string {
	return "EMAIL OTP for login is "+otp
}
func (e *email) sendNotification(message string) error {
	fmt.Printf("EMAIL: sending email: %s\n", message)
	return nil
}
func (e *email) publishMetric() {
	fmt.Printf("EMAIL: publishing metrics\n")
}