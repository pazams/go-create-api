#!/bin/bash
docker-compose up --build --force-recreate start_app_dependencies && \
docker-compose up --build --force-recreate start_test_dependencies && \
docker-compose up --build --force-recreate test
