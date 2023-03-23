package envs

import "os"

var (
	SMTP_Servername = os.Getenv("SMTP_Servername")
	SMTP_Hostname   = os.Getenv("SMTP_Hostname")
)
