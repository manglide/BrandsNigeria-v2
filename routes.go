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
		userRoutes.GET("/s/about", ensureLoggedInJWT(), showAboutPageAuthenticated)
		userRoutes.GET("/feedback", ensureNotLoggedIn(), showFeedbackPage)
		userRoutes.GET("/s/feedback", ensureLoggedInJWT(), showFeedbackPageAuthenticated)
		userRoutes.POST("/login", ensureNotLoggedIn(), performLogin)
		userRoutes.GET("/logout", ensureLoggedInJWT(), logout)
		userRoutes.GET("/register", ensureNotLoggedIn(), showRegistrationPage)
		userRoutes.POST("/register", ensureNotLoggedIn(), register)
		userRoutes.POST("/comments", ensureLoggedInJWT(), postComments)
	}

	articleRoutes := router.Group("/article")
	{
		// route from Part 1 of the tutorial
		articleRoutes.GET("/view/:article_id", ensureLoggedInJWT(), getArticle)
		articleRoutes.GET("/views/:article_id", ensureNotLoggedIn(), getArticleUnAuthenticated)
		articleRoutes.GET("/create", ensureLoggedInJWT(), showArticleCreationPage)
		articleRoutes.POST("/create", ensureLoggedInJWT(), createArticle)
	}

	productRoutes := router.Group("/product")
	{
		productRoutes.GET("/:product_id", ensureNotLoggedIn(), getProductPage)
	}

	productRoutesAU := router.Group("/s/product")
	{
		productRoutesAU.GET("/:product_id", ensureLoggedInJWT(), getProductPageAuthenticated)
	}

	createProductRoutes := router.Group("/new")
	{
		createProductRoutes.GET("/product", ensureLoggedInJWT(), createProductPage)
		createProductRoutes.POST("/product", ensureLoggedInJWT(), createProduct)
	}

	api := router.Group("/api")
	{
		api.POST("/chartsreviewlikes", ensureNotLoggedIn(), reviewLikes)
		api.POST("/chartsreviewdislikes", ensureNotLoggedIn(), reviewDisLikes)
		api.POST("/chartsreviewrating", ensureNotLoggedIn(), reviewRatings)
		api.POST("/areasofacceptance", ensureNotLoggedIn(), getAreasOfAcceptance)
		api.POST("/areasofrejection", ensureNotLoggedIn(), getAreasOfRejection)
	}

}
