#!/bin/bash

# Application Settings
export LOG_LEVEL="info"                        # Log level for the application (e.g., info, debug, warn, error)
export SERVER_PORT=8080                        # Port the application server listens on
export PRODUCT_NAME="SwiftSchool"              # Product name
export PRODUCT_VERSION="0.0.1"                 # Product version
export TIMEOUTS_IN_SECONDS=10                  # Timeout value for requests in seconds

# PostgreSQL Database Configuration
export POSTGRES_HOST="localhost"               # Hostname or IP address of the PostgreSQL database
export POSTGRES_PORT=5432                      # Port of the PostgreSQL database
export POSTGRES_USER="postgres"                # PostgreSQL username
export POSTGRES_PASSWORD="postgres"            # PostgreSQL password
export POSTGRES_DB="postgres"                  # Database name
export POSTGRES_SSL_MODE="disable"             # Disable SSL (use 'require' for production with SSL)
export POSTGRES_CONN_MAX_LIFETIME=3600         # Maximum lifetime of a connection in seconds (default 1 hour)
export POSTGRES_MAX_OPEN=10                    # Maximum number of open database connections in the pool
export POSTGRES_MAX_IDLE=5                     # Maximum number of idle connections in the pool

# Authentication Settings
export AUTH_SOURCE="local"                     # Authentication source (e.g., local, ldap, etc.)

# Run the application
go run main.go
