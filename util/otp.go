package util

import (
	"crypto/rand"
	"fmt"
	"log"
	"os"

	"github.com/resend/resend-go/v2"
)

// GenerateOTP returns a 6-digit numeric OTP
func GenerateOTP() string {
	b := make([]byte, 3)
	_, err := rand.Read(b)
	if err != nil {
		log.Println("Error generating OTP:", err)
		return "000000"
	}

	// Convert 3 bytes into a number, then modulo 1_000_000 to ensure 6 digits
	n := (int(b[0])<<16 | int(b[1])<<8 | int(b[2])) % 1000000

	return fmt.Sprintf("%06d", n) // always 6 digits with leading zeros
}

// SendOTPEmail sends an email with the OTP
// func SendOTPEmail(toEmail, otp string) error {
// 	auth := smtp.PlainAuth(
// 		"",
// 		"ashrafulialamraju@gmail.com",
// 		"uvanmhqruwesrnhc",
// 		"smtp.gmail.com",
// 	)

// 	msg := otp

// 	err := smtp.SendMail(
// 		"smtp.gmail.com:587",
// 		auth,
// 		"ashrafulialamraju@gmail.com",
// 		[]string{"ashrafulialamraju@gmail.com"},
// 		[]byte(msg),
// 	)

//		return err
//	}
func SendOTPEmail(toEmail, otp string) error {
	client := resend.NewClient(os.Getenv("RESEND_API_KEY"))

	params := &resend.SendEmailRequest{
		From:    "onboarding@resend.dev",
		To:      []string{toEmail},
		Subject: "Your OTP Code",
		Text:    "Your OTP is: " + otp,
	}

	_, err := client.Emails.Send(params)
	return err
}
