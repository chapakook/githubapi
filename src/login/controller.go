package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"time"

	"github.com/gofiber/fiber/v2"
)

func Index(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"BASE_URL":     BASE_URL,
		"CLIENT_ID":    CLIENT_ID,
		"REDIRECT_URI": REDIRECT_URI,
	})
}

func Oauth(c *fiber.Ctx) error {
	code := c.Query("code")
	params := url.Values{
		"client_id":     []string{CLIENT_ID},
		"client_secret": []string{CLIENT_SECRESTS},
		"code":          []string{code},
		"redirect_uri":  []string{REDIRECT_URI},
	}
	req, err := http.NewRequest("POST", BASE_URL+"/login/oauth/access_token", bytes.NewBufferString(params.Encode()))
	CheckErr(err)
	req.Header.Set("Accept", "application/json")

	client := http.Client{}
	resp, err := client.Do(req)
	CheckErr(err)
	CheckStatus(resp)

	var auth AuthResult
	err = json.NewDecoder(resp.Body).Decode(&auth)
	CheckErr(err)

	cookie := new(fiber.Cookie)
	cookie.Name = "accesstoken"
	cookie.Value = auth.AccessToken
	cookie.Expires = time.Now().Add(1 * time.Hour)
	c.Cookie(cookie)

	cookie.Name = "tokentype"
	cookie.Value = auth.TokenType
	cookie.Expires = time.Now().Add(1 * time.Hour)
	c.Cookie(cookie)

	return c.Redirect("http://localhost:3000/end")
}

func End(c *fiber.Ctx) error {
	return c.SendString("end")
}
