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
	if _, err = page.Goto("https://www.tucarro.com.co/"); err != nil {
		panic(err)
	}
	fmt.Println("Navigated to:", page.URL())


	//Home Page Locators
	searchBox := page.Locator("input#cb1-edit")
	searchButton := page.Locator("button[type='submit']")
	cookieButton := page.Locator("button#newCookieDisclaimerButton")


	if err = cookieButton.Click(); err != nil {
		panic(err)
	}

	if err = searchBox.Fill("Volkswagen"); err != nil {
		panic(err)
	}

	if err = searchButton.Click(); err != nil {
		panic(err)
	}

	if err = page.WaitForLoadState(); err != nil {
		panic(err)
	}

	
    // Create text and capture directories 
    dirText := "text"
    dirCapture := "capture"

    err = os.Mkdir(dirText, 0755)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println("Directory created successfully")
	}
	err = os.Mkdir(dirCapture, 0755)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println("Directory created successfully")
	}

	
	// Locate all classified listings
	classifiedListings := page.Locator("li[class='ui-search-layout__item']")

	count, err := classifiedListings.Count()
	if err != nil {
		panic(err)
	}



	for i := 0; i < count; i++ {
		listing := classifiedListings.Nth(i)
		// Get the inner text of the classified listing
		fmt.Println("_______________________________________________")
		fmt.Println(listing.InnerText())
		listingText, err := listing.InnerText()
		if err != nil {
			panic(err)
		}

		// Write the inner text to a text file
		fileName := "text/CarList-" + strconv.Itoa(i+1) + ".txt"
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
		screenshotPath := "capture/CarList-" + strconv.Itoa(i+1) + ".png"
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
