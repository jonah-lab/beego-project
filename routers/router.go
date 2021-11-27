package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/jonah-lab/beego-project/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
