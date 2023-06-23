/*
 * File: main.go
 * File Created: Sunday, 11th June 2023 9:57:15 pm
 * Last Modified: Friday, 23rd June 2023 9:04:22 am
 * Author: Akhil Datla
 * Copyright © Akhil Datla 2023
 */

package main

import (
	"flag"
	"main/backend/dbmanager"
	"main/server"

	"github.com/pterm/pterm"
)

func main() {
	// Parse command-line flags
	dbnamePtr := flag.String("dbname", "backendDB.db", "name of the database")
	portPtr := flag.Int("port", 8080, "port to listen on")
	logPtr := flag.Bool("log", false, "enable logging")
	flag.Parse()

	// Open the database
	err := dbmanager.Open(*dbnamePtr)
	if err != nil {
		panic(err)
	}

	defer dbmanager.Close()

	// Display the banner
	banner()

	// Start the server
	server.Start(*portPtr, *logPtr)
}

// banner displays a banner with author information.
func banner() {
	pterm.DefaultCenter.Print(pterm.DefaultHeader.WithFullWidth().WithBackgroundStyle(pterm.NewStyle(pterm.BgLightBlue)).WithMargin(10).Sprint("Learnado: Igniting Minds, Inspiring Learning"))
	pterm.Info.Println("Content Manager Edition")
	pterm.Info.Println("(c)2023 by Akhil Datla")
}
