package routers

import (
     "net/http"
    "github.com/gin-gonic/gin"
    "assignment_project/controllers"
    "assignment_project/config"
    "assignment_project/models"
)

func SetupRoutes(r *gin.Engine) {
    r.POST("/login", controllers.Login)

    auth := r.Group("/")
    auth.Use(config.AuthMiddleware())
    {
        auth.GET("/", func(c *gin.Context) {
            c.JSON(http.StatusOK, gin.H{"message": "Welcome!"})
        })

        // Customer Management
        auth.GET("/customers", config.AuthorizeRoles(string(models.RoleSales), string(models.RoleAdministrator)), controllers.GetCustomers)
        auth.POST("/customers", config.AuthorizeRoles(string(models.RoleSales), string(models.RoleAdministrator)), controllers.CreateCustomer)

        // Billing Management
        auth.GET("/billing", config.AuthorizeRoles(string(models.RoleSales), string(models.RoleAdministrator)), controllers.GetBillingInfo)
        auth.POST("/billing", config.AuthorizeRoles(string(models.RoleSales), string(models.RoleAdministrator)), controllers.CreateBillingEntry)

        // Payroll Management
        auth.GET("/payroll", config.AuthorizeRoles(string(models.RoleHR), string(models.RoleAdministrator)), controllers.GetPayrollInfo)
        auth.POST("/payroll", config.AuthorizeRoles(string(models.RoleHR), string(models.RoleAdministrator)), controllers.CreatePayrollEntry)

        // User Management
        auth.GET("/users", config.AuthorizeRoles(string(models.RoleAdministrator)), controllers.GetUsers)
        auth.POST("/users", config.AuthorizeRoles(string(models.RoleAdministrator)), controllers.CreateUser)
    }
}
