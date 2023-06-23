/*
 * File: licensing.go
 * File Created: Sunday, 11th June 2023 9:57:15 pm
 * Last Modified: Friday, 23rd June 2023 12:43:21 am
 * Author: Akhil Datla
 * Copyright © Akhil Datla 2023
 */

package licensing

import (
	"errors"

	uuid "github.com/satori/go.uuid"

	"main/backend/courses"
	"main/backend/dbmanager"
)

// License represents a license object.
type License struct {
	ID       string `storm:"id"`
	CourseID string `storm:"index"`
}

// Entitlement represents an entitlement object.
type Entitlement struct {
	ID         string `storm:"id"`
	CourseID   string `storm:"index"`
	HardwareID string `storm:"index"`
}

// GenerateLicense generates a license for the specified course ID.
func GenerateLicense(courseID string) (string, error) {
	var course courses.Course
	err := dbmanager.Query("ID", courseID, &course)
	if err != nil {
		return "", errors.New("invalid course id")
	}

	// Generate a new license ID
	license := &License{
		ID:       uuid.NewV4().String(),
		CourseID: courseID,
	}

	// Save the license to the database
	err = dbmanager.Save(license)
	return license.ID, err
}

// GenerateLicenses generates multiple licenses for the specified course ID and number.
func GenerateLicenses(courseID string, num int) ([]string, error) {
	licenses := make([]string, 0)
	for i := 0; i < num; i++ {
		// Generate a license for each iteration
		licenseID, err := GenerateLicense(courseID)
		if err != nil {
			return nil, err
		}
		licenses = append(licenses, licenseID)
	}
	return licenses, nil
}

// RegisterLicense registers a license with the specified license ID and hardware ID.
func RegisterLicense(licenseID, hardwareID string) error {
	var license License
	err := dbmanager.Query("ID", licenseID, &license)
	if err != nil {
		return errors.New("invalid license key")
	}

	// Create an entitlement for the registered license
	entitlement := &Entitlement{
		ID:         uuid.NewV4().String(),
		CourseID:   license.CourseID,
		HardwareID: hardwareID,
	}

	// Save the entitlement to the database
	err = dbmanager.Save(entitlement)

	// Delete the registered license
	dbmanager.Delete(&license)

	return err
}

// RevokeLicense revokes a license with the specified license ID.
func RevokeLicense(licenseID string) error {
	var license License
	err := dbmanager.Query("ID", licenseID, &license)
	if err != nil {
		return errors.New("invalid license key")
	}

	// Delete the revoked license from the database
	err = dbmanager.Delete(&license)

	return err
}

// DownloadCourses downloads courses for the specified hardware ID.
func DownloadCourses(hardwareID string) (string, error) {
	var entitlements []Entitlement
	err := dbmanager.GroupQuery("HardwareID", hardwareID, &entitlements)
	if err != nil {
		return "", err
	}

	courseIDs := make([]string, 0)
	for _, entitlement := range entitlements {
		courseIDs = append(courseIDs, entitlement.CourseID)
	}

	// Generate a website for the hardware ID and course IDs
	gobFileName := courses.GenerateWebsite(hardwareID, courseIDs)

	return gobFileName, nil
}
