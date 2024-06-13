#!/bin/bash

# Function to generate a timestamp
generate_timestamp() {
  date +"%Y%m%d%H%M%S"
}

# Function to create a migration file
create_migration_file() {
  local timestamp=$1
  local name=$2
  local filename="migrations/${timestamp}_${name}.go"

  # Create migrations directory if it doesn't exist
  mkdir -p migrations

  cat <<EOF > "$filename"
package migrations

import (
  "github.com/go-gormigrate/gormigrate/v2"
  "gorm.io/gorm"
)

func init() {
  m := &gormigrate.Migration{
    ID: "$timestamp",
    Migrate: func(tx *gorm.DB) error {
      // Add migration logic here
      return nil
    },
    Rollback: func(tx *gorm.DB) error {
      // Add rollback logic here
      return nil
    },
  }

  Migrations = append(Migrations, m)
}
EOF

  echo "Migration file created: $filename"
}

# Main script
echo "Enter the migration name (use underscores for spaces):"
read -r migration_name

if [[ -z "$migration_name" ]]; then
  echo "Migration name cannot be empty."
  exit 1
fi

timestamp=$(generate_timestamp)
create_migration_file "$timestamp" "$migration_name"