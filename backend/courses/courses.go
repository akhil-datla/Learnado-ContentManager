/*
 * File: courses.go
 * File Created: Sunday, 11th June 2023 9:57:15 pm
 * Last Modified: Friday, 23rd June 2023 1:19:57 pm
 * Author: Akhil Datla
 * Copyright © Akhil Datla 2023
 */

package courses

import (
	"bytes"
	"compress/gzip"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"main/backend/dbmanager"
	"main/backend/security"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	cp "github.com/otiai10/copy"
	uuid "github.com/satori/go.uuid"
)

// Course represents a course object.
type Course struct {
	ID       string `storm:"id"`
	Name     string `storm:"unique"`
	Filepath string
}

// CreateCourse creates a new course with the given name and filepath.
func CreateCourse(name, filepath string) (string, error) {
	// Check if the filepath exists
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		return "", fmt.Errorf("invalid filepath")
	}

	// Create a new course object
	course := &Course{
		ID:       uuid.NewV4().String(),
		Name:     name,
		Filepath: filepath,
	}

	// Save the course to the database
	err := dbmanager.Save(course)
	return course.ID, err
}

// GetCourse retrieves a course by its ID.
func GetCourse(id string) (Course, error) {
	var course Course
	err := dbmanager.Query("ID", id, &course)
	return course, err
}

// GetAllCourses retrieves all courses.
func GetAllCourses() ([]Course, error) {
	var courses []Course
	err := dbmanager.QueryAll(&courses)
	return courses, err
}

// UpdateCourse updates a course with the given ID, name, and filepath.
func UpdateCourse(id, name, filepath string) error {
	// Check if the filepath exists
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		return fmt.Errorf("invalid filepath")
	}

	// Create a new course object with updated values
	course := &Course{
		ID:       id,
		Name:     name,
		Filepath: filepath,
	}

	// Update the course in the database
	err := dbmanager.Update(course)
	return err
}

// DeleteCourse deletes a course with the given ID.
func DeleteCourse(id string) error {
	course := &Course{
		ID: id,
	}
	err := dbmanager.Delete(course)
	return err
}

// GenerateWebsite generates a website using Hugo for the specified hardware ID and course IDs.
func GenerateWebsite(hardwareID string, courseIDs []string) string {
	filePaths := make([]string, 0)
	courseNames := make([]string, 0)

	// Get filepaths and course names for the specified course IDs
	for _, id := range courseIDs {
		course, _ := GetCourse(id)
		filePaths = append(filePaths, course.Filepath)
		courseNames = append(courseNames, strings.ReplaceAll(course.Name, " ", "-"))
	}

	// Create a temporary directory for Hugo
	tempHugoDir, _ := ioutil.TempDir("", "learnado")

	// Copy the Hugo directory to the temporary directory
	cp.Copy(filepath.Join(filepath.Dir(""), "hugo"), tempHugoDir)

	// Copy course files to the temporary Hugo content directory
	for i, filePath := range filePaths {
		os.MkdirAll(filepath.Join(tempHugoDir, "content", courseNames[i]), 0755)
		cp.Copy(filePath, filepath.Join(tempHugoDir, "content", courseNames[i]))
	}

	// Copy homepage content to the temporary Hugo content directory
	homepageBytes, _ := os.ReadFile("homepage.md")
	os.WriteFile(filepath.Join(tempHugoDir, "content", "_index.md"), homepageBytes, 0666)

	// Build the Hugo website
	buildCmd := exec.Command("hugo")
	buildCmd.Dir = tempHugoDir
	buildCmd.Run()

	// Generate a unique filename for the compressed and encrypted website data
	gobFileName := uuid.NewV4().String() + ".gob"

	// Create a file map of the Hugo public directory
	m, _ := FileMapFunction(filepath.Join(tempHugoDir, "public"))

	// Compress and encrypt the file map
	compressedEncryptedMap, _ := CompressAndEncryptMap(m, hardwareID)

	// Write the compressed and encrypted data to the gob file
	os.WriteFile(gobFileName, compressedEncryptedMap, 0666)

	// Remove the temporary Hugo directory
	os.RemoveAll(tempHugoDir)

	return gobFileName
}

// FileMapFunction traverses a directory structure and creates a map with
// file/directory paths as keys and file contents as values.
func FileMapFunction(dir string) (map[string][]byte, error) {
	fileMap := make(map[string][]byte)

	err := filepath.WalkDir(dir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		relativePath, err := filepath.Rel(dir, path)
		if err != nil {
			return err
		}

		// If it's a directory, add it to the map with an empty value
		if d.IsDir() {
			fileMap[relativePath] = []byte{}
			return nil
		}

		// If it's a file, read the file and add it to the map
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		fileMap[relativePath] = data

		return nil
	})

	if err != nil {
		return nil, err
	}

	return fileMap, nil
}

// GobEncodeMap encodes a map[string][]byte to a byte slice using gob encoding.
func GobEncodeMap(m map[string][]byte) ([]byte, error) {
	buf := new(bytes.Buffer)

	enc := gob.NewEncoder(buf)
	err := enc.Encode(m)
	if err != nil {
		return nil, fmt.Errorf("failed to encode map: %w", err)
	}

	return buf.Bytes(), nil
}

// compress compresses data using gzip compression.
func compress(data []byte) ([]byte, error) {
	var buf bytes.Buffer
	w := gzip.NewWriter(&buf)
	_, err := w.Write(data)
	if err != nil {
		return nil, err
	}
	err = w.Close()
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// CompressAndEncryptMap compresses and encrypts a map using AES encryption with a given key.
func CompressAndEncryptMap(m map[string][]byte, key string) ([]byte, error) {
	// Encode the map
	encodedMap, err := GobEncodeMap(m)
	if err != nil {
		return nil, err
	}

	// Compress the encoded map
	compressedMap, err := compress(encodedMap)
	if err != nil {
		return nil, err
	}

	// Encrypt the compressed map
	encryptedMap, err := security.Encrypt(compressedMap, security.DeriveKey(key))
	if err != nil {
		return nil, err
	}

	return encryptedMap, nil
}
