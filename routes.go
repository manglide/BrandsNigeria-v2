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
		userRoutes.GET("/s/ratedProducts", ensureLoggedInJWT(), ratedProducts)
		userRoutes.GET("/s/feedback", ensureLoggedInJWT(), showFeedbackPageAuthenticated)
		userRoutes.POST("/login", ensureNotLoggedIn(), performLogin)
		userRoutes.GET("/logout", ensureLoggedInJWT(), logout)
		userRoutes.GET("/register", ensureNotLoggedIn(), showRegistrationPage)
		userRoutes.POST("/register", ensureNotLoggedIn(), register)
		userRoutes.POST("/comments", ensureLoggedInJWT(), postComments)
		userRoutes.POST("/withdrawrating", ensureLoggedInJWT(), withdrawRating)
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

	router.GET("/edit/:product_id", ensureLoggedInJWT(), editProduct)
	router.POST("/editproduct", ensureLoggedInJWT(), saveProduct)
	router.POST("/deleteproduct", ensureLoggedInJWT(), deleteProduct)

	createProductRoutes := router.Group("/new")
	{
		createProductRoutes.GET("/product", ensureLoggedInJWT(), createProductPage)
		createProductRoutes.POST("/product", ensureLoggedInJWT(), createProduct)
		createProductRoutes.GET("/productlist", ensureLoggedInJWT(), createProductListPage)
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
	}

}
