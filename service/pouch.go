package service

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/gin-gonic/gin"
	"github.com/go-rod/rod/lib/input"
	"github.com/ysmood/rod"
)

func Pouch(ctx *gin.Context) {

	username := ctx.DefaultPostForm("username", "")
	password := ctx.DefaultPostForm("password", "")

	var buf []byte

	chromeCtx, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()
	chromeCtx, cancel = context.WithTimeout(chromeCtx, 30*time.Second)
	defer cancel()

	err := chromedp.Run(chromeCtx,
		chromedp.Navigate(`https://bsignin.104.com.tw/login`),
		chromedp.WaitVisible(`.BaseInput__view`),
		chromedp.SendKeys(`.BaseInput__view[type="text"]`, username, chromedp.ByQuery),
		chromedp.SendKeys(`.BaseInput__view[type="password"]`, password, chromedp.ByQuery),
		chromedp.Click(`.BaseButton`, chromedp.NodeVisible),
		chromedp.WaitVisible(`.Product__product`),
		chromedp.Navigate(`https://pro.104.com.tw/psc2`),
		chromedp.WaitVisible(`.btn.btn-white.btn-lg.btn-block`),
		chromedp.Click(`.fa.fa-times`, chromedp.NodeVisible),
		chromedp.Click(`.btn.btn-white.btn-lg.btn-block`, chromedp.NodeVisible),
		chromedp.CaptureScreenshot(&buf),
	)
	if err != nil {
		log.Fatal(err)
		ctx.JSON(http.StatusOK, gin.H{"code": -1, "result": err})
	}
	if err := ioutil.WriteFile("fullScreenshot.png", buf, 0644); err != nil {
		log.Fatal(err)
	}

	ctx.JSON(http.StatusOK, gin.H{"code": 0, "result": "success"})
}

func PouchRod(ctx *gin.Context) {

	username := ctx.DefaultPostForm("username", "")
	password := ctx.DefaultPostForm("password", "")
	go func() {
		page := rod.New().MustConnect().MustPage("https://bsignin.104.com.tw/login")
		page.MustElement(`.BaseInput__view[type="text"]`).MustInput(username)
		page.MustElement(`.BaseInput__view[type="password"]`).MustInput(password).MustPress(input.Enter)
		page.MustElement(`.Product__product`)
		wait := page.MustWaitNavigation()
		page.MustNavigate(`https://pro.104.com.tw/psc2`)
		wait()
		page.MustElement(`.fa.fa-times`).MustClick()
		page.MustElement(`.btn.btn-white.btn-lg.btn-block`).MustClick()
	}()

	ctx.JSON(http.StatusOK, gin.H{"code": 0, "result": "success"})
}
