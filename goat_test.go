package rod_test

import (
	"log"
	"testing"
	"time"

	"github.com/go-rod/rod"
)

func TestMidnightHuntFoil(t *testing.T) {
	// Launch a new browser with default options, and connect to it.
	browser := rod.New().MustConnect()
	defer browser.MustClose()

	// Create a new page
	page := browser.MustPage("https://www.goatbots.com/redeemable-sets")

	// Click on the appropriate row in the sets table
	page.MustElement("li:nth-child(5) .name").MustClick()
	page.MustWaitRequestIdle()

	// Verify we're on the right page
	heading := page.MustElement("h1").MustText()
	log.Println(heading)

	// Wait until the JS finishes, then steal the stock notification
	time.Sleep(time.Second)
	text := page.MustElement("#card-stock").MustClick().MustText()

	log.Println(text)
	page.MustScreenshotFullPage("./screenshot.png")

	want := "Innistrad: Midnight Hunt Redeemable SetOut of stock"
	if heading+text != want {
		t.Errorf("Got %q, want %q", heading+text, want)
	}
}
