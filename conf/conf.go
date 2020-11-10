/*Package conf is the configuration infrastructure for the applciation.
It provides utilites to read environment variables and Initializes the global configuration object.
*/
package conf

import (
	"fmt"
	"sync"

	// "github.com/caarlos0/env"
	"github.com/caarlos0/env/v6"
)

// Mode Values
const (
	MODEDEV       string = "dev"
	MODESTAG      string = "stag"
	MODEPROD      string = "prod"
	MODEDEBUG     string = "debug"
	LOGFORMATTXT  string = "txt"
	LOGFORMATJSON string = "json"
)

// Vars represents the environment variables used by this App during Startup.
type vars struct {
	App          string `env:""`
	Mode         string `env:"CRUD_MODE"`  //required:"true" default:"dev"`
	Port         string `env:"CRUD_PORT" ` //default:"8080"`
	UserName     string `env:"CRUD_USERNAME"`
	Password     string `env:"CRUD_PASSWORD"`
	Host         string `env:"CRUD_HOST"`
	DBPort       string `env:"CRUD_DBPORT"`
	Database     string `env:"CRUD_DB"`
	GraceTimeout int    `env:"CRUD_GRACE_TIMEOUT"` // default:"30"`
	WriteTimeout int    `env:"CRUD_WRITE_TIMEOUT"` // default:"300"`
	TokenTimeout int    `env:"CRUD_TOKEN_TIMEOUT"` // default:"1800"`
	Combined     string `env:"CRUD_LOG_COMBINED"`  // required:"true"`
	Format       string `env:"CRUD_LOG_FORMAT"`    // default:"txt"`
	MaxActive    int    `env:"CRUD_MYSQL_MAX_ACTIVE"`
	MaxIdle      int    `env:"CRUD_MYSQL_MAX_IDLE"`
}

//Current stores the variables
var Current vars

//Load envs
func Load() {

	err := env.Parse(&Current)
	if err != nil {
		fmt.Println("Failed", err)
	}
}

const appname = "CRUD"

func init() {
	once := new(sync.Once)
	once.Do(func() {
		Load()
	})
}
