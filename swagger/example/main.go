package exmaple

import (
	"github.com/honmaple/forest"
	"github.com/honmaple/forest-contrib/swagger"
	_ "github.com/honmaple/forest-contrib/swagger/exmaple/docs"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server Petstore server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      petstore.swagger.io
// @BasePath  /v2
func main() {
	e := forest.New()

	e.GET("/swagger/*", swagger.New())
	e.Logger.Fatal(e.Start(":1323"))
}
