#!/bin/bash

#See readme for each env vars to have set before running this script
envsubst < "app-subst.yaml" > "app.yaml"

GO111MODULE=on gcloud app deploy --project=$GCP_PROJECT
