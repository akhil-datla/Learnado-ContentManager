/*
 * File: routes.go
 * File Created: Sunday, 11th June 2023 9:57:15 pm
 * Last Modified: Friday, 23rd June 2023 3:29:31 pm
 * Author: Akhil Datla
 * Copyright © Akhil Datla 2023
 */

package server

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/labstack/echo/v4"
	middleware "github.com/labstack/echo/v4/middleware"
)

var e *echo.Echo

// Start starts the server on the specified port, with optional logging.
func Start(port int, log bool) {
	e = echo.New()
	e.HideBanner = true

	// Render GUI
	e.Use(middleware.Gzip())
	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:   filepath.Join(filepath.Dir(""), "frontend", "build"),
		Index:  "index.html",
		Browse: false,
		HTML5:  true,
	}))

	if log {
		e.Use(middleware.Logger())
	}
	e.Use(middleware.Recover())

	// Configure CORS
	defaultCORSConfig := middleware.CORSConfig{
		Skipper:      middleware.DefaultSkipper,
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete, http.MethodOptions},
	}
	e.Use(middleware.CORSWithConfig(defaultCORSConfig))

	initializeRoutes()

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", port)))
}

// initializeRoutes initializes all the routes and associates them with their respective handlers.
func initializeRoutes() {
	e.POST("/courses/create", createCourse)
	e.GET("/courses/info", getCourse)
	e.GET("/courses/all", getAllCourses)
	e.POST("/courses/update", updateCourse)
	e.DELETE("/courses/delete", deleteCourse)
	e.POST("/licenses/create", generateLicenses)
	e.POST("/licenses/register", registerLicense)
	e.DELETE("/licenses/revoke", revokeLicense)
	e.POST("/download", downloadCourses)
}
