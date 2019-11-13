#!/bin/bash
docker-compose up --build start_app_dependencies && \
docker-compose up --build start_test_dependencies && \
docker-compose up --build --exit-code-from test test
