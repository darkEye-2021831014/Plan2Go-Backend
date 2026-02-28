package util

import (
	"crypto/rand"
	"fmt"
	"log"
	"net/smtp"
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
func SendOTPEmail(toEmail, otp string) error {
	auth := smtp.PlainAuth(
		"",
		"ashrafulialamraju@gmail.com",
		"uvanmhqruwesrnhc",
		"smtp.gmail.com",
	)

	msg := otp

	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"ashrafulialamraju@gmail.com",
		[]string{toEmail},
		[]byte(msg),
	)

	return err
}

// package util

// import (
// 	"context"
// 	"crypto/rand"
// 	"fmt"
// 	"log"
// 	"os"
// 	"time"

// 	brevo "github.com/getbrevo/brevo-go/lib"
// )

// // GenerateOTP returns a secure 6-digit numeric OTP
// func GenerateOTP() string {
// 	b := make([]byte, 3)
// 	_, err := rand.Read(b)
// 	if err != nil {
// 		log.Printf("Error generating random bytes for OTP: %v", err)
// 		return "000000" // fallback - don't use in production
// 	}

// 	// Combine bytes into a number and get 6 digits (0-999999)
// 	n := (int(b[0])<<16 | int(b[1])<<8 | int(b[2])) % 1_000_000
// 	return fmt.Sprintf("%06d", n)
// }

// // SendOTPEmail sends an HTML OTP email using Brevo API (free tier)
// // SendOTPEmail sends an HTML OTP email using Brevo API (free tier)
// func SendOTPEmail(toEmail, otp string) error {
// 	apiKey := os.Getenv("BREVO_API_KEY")
// 	if apiKey == "" {
// 		return fmt.Errorf("BREVO_API_KEY environment variable is not set")
// 	}

// 	senderEmail := os.Getenv("BREVO_SENDER_EMAIL")
// 	if senderEmail == "" {
// 		return fmt.Errorf("BREVO_SENDER_EMAIL environment variable is not set")
// 	}

// 	senderName := os.Getenv("BREVO_SENDER_NAME")
// 	if senderName == "" {
// 		senderName = "Plan2Go OTP" // fallback
// 	}

// 	cfg := brevo.NewConfiguration()
// 	cfg.AddDefaultHeader("api-key", apiKey)

// 	client := brevo.NewAPIClient(cfg)

// 	htmlContent := fmt.Sprintf(`
// 		<!DOCTYPE html>
// 		<html lang="en">
// 		<head>
// 			<meta charset="UTF-8">
// 			<meta name="viewport" content="width=device-width, initial-scale=1.0">
// 			<title>Your OTP Code</title>
// 		</head>
// 		<body style="font-family: Arial, Helvetica, sans-serif; margin:0; padding:0; background:#f4f4f4; color:#333;">
// 			<table width="100%%" cellpadding="0" cellspacing="0" style="background:#f4f4f4; padding:20px;">
// 				<tr>
// 					<td align="center">
// 						<table width="100%%" cellpadding="20" cellspacing="0" style="max-width:500px; background:white; border-radius:8px; box-shadow:0 2px 10px rgba(0,0,0,0.1);">
// 							<tr>
// 								<td align="center" style="padding-bottom:10px;">
// 									<h2 style="color:#2c3e50; margin:0;">Your Verification Code</h2>
// 								</td>
// 							</tr>
// 							<tr>
// 								<td align="center" style="font-size:48px; font-weight:bold; letter-spacing:10px; color:#3498db; padding:20px 0;">
// 									%s
// 								</td>
// 							</tr>
// 							<tr>
// 								<td align="center" style="color:#7f8c8d; font-size:16px;">
// 									This code expires in 10 minutes.<br>
// 									If you didn't request this, ignore this email.
// 								</td>
// 							</tr>
// 							<tr>
// 								<td align="center" style="padding-top:30px; font-size:12px; color:#95a5a6;">
// 									Sent via %s • © %d
// 								</td>
// 							</tr>
// 						</table>
// 					</td>
// 				</tr>
// 			</table>
// 		</body>
// 		</html>
// 	`, otp, senderName, time.Now().Year())

// 	email := brevo.SendSmtpEmail{
// 		Sender: &brevo.SendSmtpEmailSender{
// 			Email: senderEmail,
// 			Name:  senderName,
// 		},
// 		To: []brevo.SendSmtpEmailTo{
// 			{Email: toEmail},
// 		},
// 		Subject:     "Your OTP Code - Verify Your Account",
// 		HtmlContent: htmlContent,
// 		TextContent: fmt.Sprintf("Your OTP code is: %s\nThis code expires in 10 minutes.", otp),
// 	}

// 	resp, httpResp, err := client.TransactionalEmailsApi.SendTransacEmail(context.Background(), email)
// 	if err != nil {
// 		log.Printf("Brevo send error to %s: %v (HTTP %d)", toEmail, err, httpResp.StatusCode)
// 		return fmt.Errorf("failed to send email: %w", err)
// 	}

// 	log.Printf("OTP sent to %s | Message ID: %s (HTTP %d)", toEmail, resp.MessageId, httpResp.StatusCode)
// 	return nil
// }
