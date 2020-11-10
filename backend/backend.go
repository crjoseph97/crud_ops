/*Package backend provides various Respoistories to handle database operations related to Domain
  It has various implementations to satisfy backend/repo interafaces.
*/
package backend

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

//Infra defines the infrastructure for all the database connection and logger objects.
type Infra struct {
	// S3            *s3.S3
	DBr *sqlx.DB
	DBw *sqlx.DB
}

// Repository to handle domain specific database operations
var (
	Users UsersRepo
)

//Init ...
func Init(infra *Infra) {
	Users = &users{Infra: infra}

	fmt.Println("Infra initialiazed")
}
