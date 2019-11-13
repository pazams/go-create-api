#!/bin/bash

#Currently it expects: API_TOKEN=<the-api-token> POSTGRES_GCP_CONNECTION_NAME=<project:zone:instance> POSTGRES_PASSWORD=<the-password>
envsubst < "app-subst.yaml" > "app.yaml"

GO111MODULE=on gcloud app deploy --project=$GCP_PROJECT
