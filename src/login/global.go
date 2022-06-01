package main

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	PORT             string
	GITHUB_API_TOKEN string
	CLIENT_ID        string
	CLIENT_SECRESTS  string
	BASE_URL         string
	REDIRECT_URI     string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	GITHUB_API_TOKEN = os.Getenv("GITHUB_API_TOKEN")
	CLIENT_ID = os.Getenv("CLIENT_ID")
	CLIENT_SECRESTS = os.Getenv("CLIENT_SECRESTS")
	BASE_URL = os.Getenv("BASE_URL")
	REDIRECT_URI = os.Getenv("REDIRECT_URI")

	PORT = ":3000"
}
