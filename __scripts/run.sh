#!/bin/bash

# use --force-recreate if needed
docker-compose up --build start_app_dependencies && \ 
docker-compose up --build app
