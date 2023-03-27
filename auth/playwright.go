package auth

import (
	"fmt"
	"github.com/playwright-community/playwright-go"
	"net/url"
	"time"
)

func Token() (string, error) {
	pw, err := playwright.Run()
	if err != nil {
		if err = playwright.Install(); err != nil {
			return "", fmt.Errorf("could not install playwright: %v", err)
		}
		pw, err = playwright.Run()
		if err != nil {
			return "", fmt.Errorf("could not start playwright: %v", err)
		}
	}
	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{Headless: playwright.Bool(false)})
	if err != nil {
		return "", fmt.Errorf("could not launch browser: %v", err)
	}
	defer func() {
		if err = browser.Close(); err != nil {
			panic(fmt.Errorf("could not close browser: %v", err))
		}
		if err = pw.Stop(); err != nil {
			panic(fmt.Errorf("could not stop Playwright: %v", err))
		}
	}()

	page, err := browser.NewPage()
	if err != nil {
		return "", fmt.Errorf("could not create page: %v", err)
	}
	if _, err = page.Goto("https://www.floatplane.com/login"); err != nil {
		return "", fmt.Errorf("could not goto: %v", err)
	}

	// Wait for user to log in
	for {
		currentUrl := page.URL()
		if currentUrl == "https://www.floatplane.com/" {
			break
		}
		time.Sleep(time.Millisecond * 100)
	}
	cookies, err := page.Context().Cookies("https://www.floatplane.com")
	if err != nil {
		return "", fmt.Errorf("could not goto: %v", err)
	}
	err = fmt.Errorf("could not find `sails.id` cookie")
	var sailsSid string
	for _, c := range cookies {
		if c.Name == "sails.sid" {
			if sailsSid, err = url.QueryUnescape(c.Value); err != nil {
				return "", err
			}
			err = nil
			break
		}
	}

	return sailsSid, err
}
