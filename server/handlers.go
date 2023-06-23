/*
 * File: handlers.go
 * File Created: Sunday, 11th June 2023 9:57:15 pm
 * Last Modified: Friday, 23rd June 2023 12:45:59 am
 * Author: Akhil Datla
 * Copyright © Akhil Datla 2023
 */

package server

import (
	"encoding/json"
	"main/backend/courses"
	"main/backend/licensing"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
)

// createCourse handles the creation of a new course.
func createCourse(c echo.Context) error {
	// Parse the request body to JSON
	jsonMap := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&jsonMap)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error parsing request body")
	}

	// Extract the name and filepath from the JSON map
	name := jsonMap["name"].(string)
	filepath := jsonMap["filepath"].(string)

	// Create a new course
	id, err := courses.CreateCourse(name, filepath)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error creating course")
	}

	// Return the created course ID
	return c.JSON(http.StatusOK, map[string]string{"courseID": id})
}

// getCourse retrieves a specific course by its ID.
func getCourse(c echo.Context) error {
	// Parse the request body to JSON
	jsonMap := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&jsonMap)
	if err != nil {
		return err
	}

	// Extract the course ID from the JSON map
	id := jsonMap["id"].(string)

	// Retrieve the course
	course, err := courses.GetCourse(id)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error getting course")
	}

	// Return the retrieved course
	return c.JSON(http.StatusOK, course)
}

// getAllCourses retrieves all courses.
func getAllCourses(c echo.Context) error {
	// Retrieve all courses
	courses, err := courses.GetAllCourses()
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error getting courses")
	}

	// Return the retrieved courses
	return c.JSON(http.StatusOK, courses)
}

// updateCourse updates the details of a course.
func updateCourse(c echo.Context) error {
	// Parse the request body to JSON
	jsonMap := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&jsonMap)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error parsing request body")
	}

	// Extract the course details from the JSON map
	id := jsonMap["id"].(string)
	name := jsonMap["name"].(string)
	filepath := jsonMap["filepath"].(string)

	// Update the course
	err = courses.UpdateCourse(id, name, filepath)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error updating course")
	}

	return c.String(http.StatusOK, "Course updated")
}

// deleteCourse deletes a course by its ID.
func deleteCourse(c echo.Context) error {
	// Parse the request body to JSON
	jsonMap := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&jsonMap)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error parsing request body")
	}

	// Extract the course ID from the JSON map
	id := jsonMap["id"].(string)

	// Delete the course
	err = courses.DeleteCourse(id)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error deleting course")
	}

	return c.String(http.StatusOK, "Course deleted")
}

// generateLicenses generates licenses for a specific course.
func generateLicenses(c echo.Context) error {
	// Parse the request body to JSON
	jsonMap := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&jsonMap)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error parsing request body")
	}

	// Extract the course ID and number of licenses from the JSON map
	courseID := jsonMap["courseID"].(string)
	num, err := strconv.Atoi(jsonMap["num"].(string))
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error parsing number of licenses")
	}

	// Generate the licenses
	licenses, err := licensing.GenerateLicenses(courseID, num)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error generating licenses")
	}

	// Return the generated license keys
	return c.JSON(http.StatusOK, map[string][]string{"licenseKeys": licenses})
}

// registerLicense registers a license with a specific license key and hardware ID.
func registerLicense(c echo.Context) error {
	// Parse the request body to JSON
	jsonMap := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&jsonMap)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error parsing request body")
	}

	// Extract the license key and hardware ID from the JSON map
	licenseKey := jsonMap["licenseKey"].(string)
	hardwareID := jsonMap["hardwareID"].(string)

	// Register the license
	err = licensing.RegisterLicense(licenseKey, hardwareID)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error registering license")
	}

	return c.String(http.StatusOK, "License registered")
}

// revokeLicense revokes a license with a specific license key.
func revokeLicense(c echo.Context) error {
	// Parse the request body to JSON
	jsonMap := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&jsonMap)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error parsing request body")
	}

	// Extract the license key from the JSON map
	licenseKey := jsonMap["licenseKey"].(string)

	// Revoke the license
	err = licensing.RevokeLicense(licenseKey)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error revoking license")
	}

	return c.String(http.StatusOK, "License revoked")
}

// downloadCourses downloads courses for a specific hardware ID.
func downloadCourses(c echo.Context) error {
	// Parse the request body to JSON
	jsonMap := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&jsonMap)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error parsing request body")
	}

	// Extract the hardware ID from the JSON map
	hardwareID := jsonMap["hardwareID"].(string)

	// Download the courses for the hardware ID
	file, err := licensing.DownloadCourses(hardwareID)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error downloading courses")
	}

	defer os.Remove(file)

	return c.File(file)
}
