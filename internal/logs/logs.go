package logs

import (
	"github.com/muktiarafi/myriadcode-backend/internal/configs"
	"log"
	"os"
)

func SetupLogs(app *configs.AppConfig) {
	app.InfoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.WarningLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
}
