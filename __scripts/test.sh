#!/bin/bash
sudo docker-compose up --build start_app_dependencies && \
sudo docker-compose up --build start_test_dependencies && \
sudo docker-compose up --build --exit-code-from test test
