#!/bin/bash
docker compose down --remove-orphans --volumes --rmi all --timeout 0
docker compose up --build --force-recreate --remove-orphans