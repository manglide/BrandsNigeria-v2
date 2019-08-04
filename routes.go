package main

func initializeRoutes() {
	// Auth Middleware
	router.Use(setUserStatus())

	// Handle the index route
	router.GET("/", showIndexPage)

	userRoutes := router.Group("/u")
	{
		userRoutes.GET("/login", ensureNotLoggedIn(), showLoginPage)
		userRoutes.GET("/about", ensureNotLoggedIn(), showAboutPage)
		userRoutes.GET("/s/about", ensureLoggedIn(), showAboutPageAuthenticated)
		userRoutes.GET("/feedback", ensureNotLoggedIn(), showFeedbackPage)
		userRoutes.GET("/s/ratedProducts", ensureLoggedIn(), ratedProducts)
		userRoutes.GET("/s/feedback", ensureLoggedIn(), showFeedbackPageAuthenticated)
		userRoutes.POST("/login", ensureNotLoggedIn(), performLogin)
		userRoutes.GET("/logout", ensureLoggedIn(), logout)
		userRoutes.GET("/register", ensureNotLoggedIn(), showRegistrationPage)
		userRoutes.POST("/register", ensureNotLoggedIn(), register)
		userRoutes.POST("/comments", ensureLoggedIn(), postComments)
		userRoutes.POST("/withdrawrating", ensureLoggedIn(), withdrawRating)
	}

	articleRoutes := router.Group("/article")
	{
		// route from Part 1 of the tutorial
		articleRoutes.GET("/view/:article_id", ensureLoggedIn(), getArticle)
		articleRoutes.GET("/views/:article_id", ensureNotLoggedIn(), getArticleUnAuthenticated)
		articleRoutes.GET("/create", ensureLoggedIn(), showArticleCreationPage)
		articleRoutes.POST("/create", ensureLoggedIn(), createArticle)
	}

	productRoutes := router.Group("/product")
	{
		productRoutes.GET("/:product_id", ensureNotLoggedIn(), getProductPage)
	}

	productRoutesAU := router.Group("/s/product")
	{
		productRoutesAU.GET("/:product_id", ensureLoggedIn(), getProductPageAuthenticated)
	}

	router.GET("/edit/:product_id", ensureLoggedIn(), editProduct)
	router.POST("/editproduct", ensureLoggedIn(), saveProduct)
	router.POST("/deleteproduct", ensureLoggedIn(), deleteProduct)
	router.POST("/restoreproduct", ensureLoggedIn(), restoreProduct)
	router.GET("/sitemap", neutral(), genSitemap)

	createProductRoutes := router.Group("/new")
	{
		createProductRoutes.GET("/product", ensureLoggedIn(), createProductPage)
		createProductRoutes.POST("/product", ensureLoggedIn(), createProduct)
		createProductRoutes.GET("/productlist", ensureLoggedIn(), createProductListPage)
		createProductRoutes.GET("/deletedProductlist", ensureLoggedIn(), createDeletedProductListPage)
	}

	api := router.Group("/api")
	{
		api.POST("/chartsreviewlikes", neutral(), reviewLikes)
		api.POST("/chartsreviewdislikes", neutral(), reviewDisLikes)
		api.POST("/chartsreviewrating", neutral(), reviewRatings)
		api.POST("/areasofacceptance", neutral(), getAreasOfAcceptance)
		api.POST("/areasofrejection", neutral(), getAreasOfRejection)
		api.POST("/productrecommendation", neutral(), getProductRecommendation)
		api.POST("/productsAPICompetitor", neutral(), pCompetitor)
		api.POST("/approveRating", ensureLoggedIn(), approveRating)
		api.POST("/disapproveRating", ensureLoggedIn(), disapproveRating)
	}

}
