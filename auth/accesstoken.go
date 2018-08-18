// Copyright 2018 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

// Package authsnippets contains Google Cloud authentication snippets.
// This sample demonstrates AccessTokenCredentials:
// https://godoc.org/golang.org/x/oauth2/google#JWTAccessTokenSourceFromJSON

// To use, create a service accountJSON file and allow atleast Pub/Sub Viewer IAM
// permissions on allow listing Topics on a project.
package authsnippets

import (
	"io/ioutil"
	"log"

	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"

	"cloud.google.com/go/pubsub"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

// audience values for other services can be found in the repo here similar to
// PubSub
// https://github.com/googleapis/googleapis/blob/master/google/pubsub/pubsub.yaml
const aud string = "https://pubsub.googleapis.com/google.pubsub.v1.Publisher"

func authsnippets() {

	projectID := "YOUR_PROJECT"
	keyfile := "service_account.json"

	ctx := context.Background()
	keyBytes, err := ioutil.ReadFile(keyfile)
	if err != nil {
		log.Fatalf("Unable to read service account key file  %v", err)
	}
	tokenSource, err := google.JWTAccessTokenSourceFromJSON(keyBytes, aud)
	if err != nil {
		log.Fatalf("Error building JWT access token source: %v", err)
	}
	jwt, err := tokenSource.Token()
	if err != nil {
		log.Fatalf("Unable to generate JWT token: %v", err)
	}
	log.Println(jwt.AccessToken)

	client, err := pubsub.NewClient(ctx, projectID, option.WithTokenSource(tokenSource))
	if err != nil {
		log.Fatalf("Could not create pubsub Client: %v", err)
	}
	topics := client.Topics(ctx)
	for {
		topic, err := topics.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Error listing topics %v", err)
		}
		log.Println(topic)
	}
}
