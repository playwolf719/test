package routers

import (
	"github.com/astaxie/beego"
	"github.com/playwolf719/test/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/test", &controllers.TestController{})
	beego.Router("/search_match", &controllers.SMController{})
}
