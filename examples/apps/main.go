package main

import (
	"context"
	"log"
	"os"

	"golang.org/x/oauth2"

	"github.com/micnncim/appstoreconnect-go/appstoreconnect"
	"github.com/micnncim/appstoreconnect-go/auth"
)

var (
	issuerID   = os.Getenv("ISSUER_ID")
	keyID      = os.Getenv("KEY_ID")
	apiKeyPath = os.Getenv("API_KEY_PATH")
)

func main() {
	ctx := context.Background()

	ts, err := auth.NewTokenSource(issuerID, keyID, auth.WithAPIKeyPath(apiKeyPath))
	if err != nil {
		log.Fatal(err)
	}

	cfg := appstoreconnect.NewConfiguration()
	cfg.HTTPClient = oauth2.NewClient(ctx, ts)

	c := appstoreconnect.NewAPIClient(cfg)

	resp, _, err := c.AppsApi.AppsGetCollectionExecute(c.AppsApi.AppsGetCollection(ctx))
	if err != nil {
		log.Fatal(err)
	}

	for _, d := range resp.GetData() {
		a := d.Attributes
		log.Printf("Name: %s, Bundle ID: %s", a.GetName(), a.GetBundleId())
	}
}
