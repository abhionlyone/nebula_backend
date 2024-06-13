#!/bin/bash

# Load environment variables from .env file
if [ -f .env ]; then
  export $(cat .env | grep -v '#' | awk '/=/ {print $1}')
else
  echo ".env file not found!"
  exit 1
fi

# Function to create the database
create_db() {
  PGPASSWORD=$DB_PASSWORD psql -h $DB_HOST -U $DB_USER -c "CREATE DATABASE $DB_NAME;"
  if [ $? -eq 0 ]; then
    echo "Database $DB_NAME created successfully."
  else
    echo "Failed to create database $DB_NAME."
  fi
}

# Function to drop the database
drop_db() {
  read -p "Are you sure you want to drop the database $DB_NAME? This action cannot be undone. (y/n): " confirm
  if [ "$confirm" = "y" ]; then
    PGPASSWORD=$DB_PASSWORD psql -h $DB_HOST -U $DB_USER -c "DROP DATABASE $DB_NAME;"
    if [ $? -eq 0 ]; then
      echo "Database $DB_NAME dropped successfully."
    else
      echo "Failed to drop database $DB_NAME."
    fi
  else
    echo "Database drop operation cancelled."
  fi
}

# Check the input parameter
case "$1" in
  --create)
    create_db
    ;;
  --drop)
    drop_db
    ;;
  *)
    echo "Usage: $0 --create | --drop"
    exit 1
    ;;
esac
