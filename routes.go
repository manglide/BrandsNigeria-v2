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
		userRoutes.GET("/s/feedback", ensureLoggedIn(), showFeedbackPageAuthenticated)
		userRoutes.POST("/login", ensureNotLoggedIn(), performLogin)
		userRoutes.GET("/logout", ensureLoggedIn(), logout)
		userRoutes.GET("/register", ensureNotLoggedIn(), showRegistrationPage)
		userRoutes.POST("/register", ensureNotLoggedIn(), register)
	}

	articleRoutes := router.Group("/article")
	{
		// route from Part 1 of the tutorial
		articleRoutes.GET("/view/:article_id", ensureLoggedIn(), getArticle)
		articleRoutes.GET("/views/:article_id", ensureNotLoggedIn(), getArticleUnAuthenticated)
		articleRoutes.GET("/create", ensureLoggedIn(), showArticleCreationPage)
		articleRoutes.POST("/create", ensureLoggedIn(), createArticle)
	}

}
