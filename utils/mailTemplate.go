// utils/templates.go
package utils

import "fmt"

func OtpVerificationTemplate(name, otp string) string {
	return fmt.Sprintf(`
	<!DOCTYPE html>
	<html lang="en">
	<head>
	  <meta charset="UTF-8">
	  <meta name="viewport" content="width=device-width, initial-scale=1.0">
	  <title>OTP Verification</title>
	  <style>
		body {
		  font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
		  background-color: #f9f9f9;
		  color: #333;
		  margin: 0;
		  padding: 0;
		}

		.container {
		  max-width: 600px;
		  margin: 20px auto;
		  padding: 20px 30px;
		  background-color: #ffffff;
		  border-radius: 8px;
		  box-shadow: 0 5px 20px rgba(0, 0, 0, 0.1);
		  border-left: 5px solid #007bff;
		}

		h1 {
		  color: #007bff;
		  font-size: 28px;
		  margin-bottom: 20px;
		}

		p {
		  font-size: 16px;
		  line-height: 1.6;
		  margin-bottom: 15px;
		}

		.otp-code {
		  font-size: 32px;
		  color: #28a745;
		  font-weight: bold;
		  background-color: #e6f7e1;
		  padding: 10px 20px;
		  border-radius: 6px;
		  display: inline-block;
		  margin: 20px 0;
		  letter-spacing: 2px;
		}

		.footer {
		  margin-top: 30px;
		  font-size: 14px;
		  color: #777;
		  text-align: center;
		  border-top: 1px solid #ddd;
		  padding-top: 15px;
		}

		a {
		  color: #007bff;
		  text-decoration: none;
		}

		a:hover {
		  text-decoration: underline;
		}

	  </style>
	</head>
	<body>
	  <div class="container">
		<h1>OTP Verification</h1>
		<p>Dear %s,</p>
		<p>We received a request to verify your email for Track_Us. Please use the OTP below to complete the verification process:</p>
		<p class="otp-code">%s</p>
		<p>This OTP is valid for 5 minutes. Please do not share this code with anyone.</p>
		<div class="footer">
		  <p>If you did not request this OTP, you can safely ignore this email.</p>
		  <p>Thank you,<br>Track_Us Team</p>
		  <p><a href="https://trackus.example.com">Visit Track_Us</a></p>
		</div>
	  </div>
	</body>
	</html>
	`, name, otp)
}



func AccountRegistrationTemplate(name string) string {
	return fmt.Sprintf(`
	<!DOCTYPE html>
	<html lang="en">
	<head>
	  <meta charset="UTF-8">
	  <meta name="viewport" content="width=device-width, initial-scale=1.0">
	  <title>Welcome to Portfolio Builder</title>
	  <style>
		body {
		  font-family: Arial, sans-serif;
		  background-color: #f9f9f9;
		  color: #333;
		  margin: 0;
		  padding: 0;
		}

		.container {
		  max-width: 600px;
		  margin: 20px auto;
		  padding: 20px 30px;
		  background-color: #ffffff;
		  border-radius: 8px;
		  box-shadow: 0 5px 20px rgba(0, 0, 0, 0.1);
		  border-left: 5px solid #007bff;
		}

		h1 {
		  color: #007bff;
		  font-size: 26px;
		  margin-bottom: 20px;
		}

		p {
		  font-size: 16px;
		  line-height: 1.6;
		  margin-bottom: 15px;
		}

		.footer {
		  margin-top: 30px;
		  font-size: 14px;
		  color: #777;
		  text-align: center;
		  border-top: 1px solid #ddd;
		  padding-top: 15px;
		}

		a {
		  color: #007bff;
		  text-decoration: none;
		}

		a:hover {
		  text-decoration: underline;
		}

	  </style>
	</head>
	<body>
	  <div class="container">
		<h1>Welcome to Portfolio Builder, %s!</h1>
		<p>Thank you for verifying your email. Your account is now fully activated, and you can start building and showcasing your professional portfolio.</p>
		<p>We are excited to have you on board!</p>
		<p>Get started by creating your first project, adding your skills, and sharing your work with the world.</p>
		<div class="footer">
		  <p>Thank you,<br>Portfolio Builder Team</p>
		  <p><a href="https://portfolio-builder.example.com">Visit Portfolio Builder</a></p>
		</div>
	  </div>
	</body>
	</html>
	`, name)
}
