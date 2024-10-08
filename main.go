package main

import (
	"fmt"
	"github.com/playwright-community/playwright-go"
	"os"
	"strconv"
)

func main() {
	pw, err := playwright.Run()
	if err != nil {
		panic(err)
	}

	// Configure the browser to launch in non-headless mode
	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(false),
	})

	if err != nil {
		panic(err)
	}

	page, err := browser.NewPage()
	if err != nil {
		panic(err)
	}

	// Navigate to PakWheels website
	if _, err = page.Goto("https://www.pakwheels.com"); err != nil {
		panic(err)
	}
	fmt.Println("Navigated to:", page.URL())


	searchBox := page.Locator("input[name='home-query']")
	button := page.Locator("button[id='home-search-btn']")
	if err = searchBox.Fill("Car name"); err != nil {
		panic(err)
	}

	if err = button.Click(); err != nil {
		panic(err)
	}

	// Wait for the results page to load
	if err = page.WaitForLoadState(); err != nil {
		panic(err)
	}

	// Locate all classified listings
	classifiedListings := page.Locator("li.classified-listing")
	count, err := classifiedListings.Count()
	if err != nil {
		panic(err)
	}

	for i := 0; i < count; i++ {
		listing := classifiedListings.Nth(i)
		// Get the inner text of the classified listing
		fmt.Println("----------------------------------------------------")
		fmt.Println(listing.InnerText())
		listingText, err := listing.InnerText()
		if err != nil {
			panic(err)
		}

		// Write the inner text to a text file
		fileName := "text/listing-" + strconv.Itoa(i+1) + ".txt"
		file, err := os.Create(fileName)
		if err != nil {
			panic(err)
		}
		if _, err = file.WriteString(listingText); err != nil {
			file.Close()
			panic(err)
		}
		file.Close()
		fmt.Println("Listing data written to", fileName)

		// Taking a screenshot of the classified listing
		screenshotPath := "capture/listing-" + strconv.Itoa(i+1) + ".png"
		if _, err = listing.Screenshot(playwright.LocatorScreenshotOptions{
			Path: playwright.String(screenshotPath),
		}); err != nil {
			panic(err)
		}
		fmt.Println("Screenshot of listing taken and saved to", screenshotPath)
	}

	
	if err = browser.Close(); err != nil {
		panic(err)
	}

	if err = pw.Stop(); err != nil {
		panic(err)
	}
}
