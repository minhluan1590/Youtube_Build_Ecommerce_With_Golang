package routes

import (
	"github.com/minhluan1590/Youtube_Build_Ecommerce_With_Golang/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
    userRoutes := router.Group("/users")
    {
        userRoutes.POST("/signup", controllers.SignUp)
        userRoutes.POST("/login", controllers.Login)
        userRoutes.GET("/product_view", controllers.ProductView)
        userRoutes.GET("/search", controllers.Search)
    }

    adminRoutes := router.Group("/admin")
    {
        adminRoutes.POST("/add_product", controllers.AddProduct)
    }
}


func CartRoutes(router *gin.Engine) {
    cartRoutes := router.Group("/cart")
    {
        cartRoutes.POST("/add_to_cart", controllers.AddToCart)
        cartRoutes.POST("/remove_item", controllers.RemoveItem)
        cartRoutes.POST("/cart_checkout", controllers.CartCheckout)
        cartRoutes.POST("/instant_buy", controllers.InstantBuy)
    }
}


