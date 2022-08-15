package configs

import (
	"github.com/TendonT52/tendon-api/services"
	"github.com/TendonT52/tendon-api/drivers"
)

type App struct{
	MongoDB  *drivers.DB
	JwtSecret *services.JwtServices
}
