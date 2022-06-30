#!/bin/bash
docker build -t gin-example .
docker run -p 8000:8000 gin-example