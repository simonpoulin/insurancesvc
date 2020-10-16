package admin

import "github.com/labstack/echo"

// RouteAdmin ...
func RouteAdmin(e *echo.Echo) {
	group := e.Group("/admin")

	adminUser(group)
	adminEmployee(group)
	adminCar(group)
}
