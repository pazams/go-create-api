#!/bin/bash

cat "app-subst-sec-conf.yaml" > "app.yaml"

GO111MODULE=on gcloud app deploy --project=$GCP_PROJECT
