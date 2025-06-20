#!/usr/bin/env bash
set -euo pipefail

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Default values (customize as needed)
DB_TYPE="${DB_TYPE:-postgres}"
DB_HOST="${DB_HOST:-localhost}"
DB_PORT="${DB_PORT:-5432}"
DB_USER="${DB_USER:-quantumsynth}"
DB_PASS="${DB_PASS:-quantumsynth}"
DB_NAME="${DB_NAME:-quantumsynth_db}"
MIGRATIONS_DIR="${MIGRATIONS_DIR:-./migrations}"
MIGRATE_BIN="${MIGRATE_BIN:-migrate}"

function info {
    echo -e "${GREEN}[INFO]${NC} $1"
}

function warn {
    echo -e "${YELLOW}[WARN]${NC} $1"
}

function error {
    echo -e "${RED}[ERROR]${NC} $1"
}

function check_command {
    command -v "$1" >/dev/null 2>&1 || { error "'$1' is required but not installed."; exit 1; }
}

function show_usage {
    cat <<EOF
Usage: $0 [up|down|status|force VERSION]

Commands:
  up           Apply all up migrations
  down         Revert the last migration batch
  status       Show current migration status
  force V      Set database schema to version V (manual sync)
  help         Show this help message

Environment variables:
  DB_TYPE       [postgres|mysql|sqlite3] (default: postgres)
  DB_HOST       Database host (default: localhost)
  DB_PORT       Port (default: 5432)
  DB_USER       Username (default: quantumsynth)
  DB_PASS       Password (default: quantumsynth)
  DB_NAME       Database name (default: quantumsynth_db)
  MIGRATIONS_DIR   (default: ./migrations)
  MIGRATE_BIN      (default: migrate)
EOF
}

function build_db_url {
    case "$DB_TYPE" in
        postgres)
            echo "postgres://$DB_USER:$DB_PASS@$DB_HOST:$DB_PORT/$DB_NAME?sslmode=disable"
            ;;
        mysql)
            echo "$DB_USER:$DB_PASS@tcp($DB_HOST:$DB_PORT)/$DB_NAME"
            ;;
        sqlite3)
            echo "$DB_NAME"
            ;;
        *)
            error "Unsupported DB_TYPE: $DB_TYPE"
            exit 1
            ;;
    esac
}

# Main logic
ACTION="${1:-help}"

check_command "$MIGRATE_BIN"

DB_URL=$(build_db_url)

case "$ACTION" in
    up)
        info "Running migrations UP..."
        "$MIGRATE_BIN" -path "$MIGRATIONS_DIR" -database "$DB_URL" up
        ;;
    down)
        warn "Reverting last migration batch..."
        "$MIGRATE_BIN" -path "$MIGRATIONS_DIR" -database "$DB_URL" down 1
        ;;
    status)
        info "Current migration status:"
        "$MIGRATE_BIN" -path "$MIGRATIONS_DIR" -database "$DB_URL" version || true
        ;;
    force)
        if [ -z "${2:-}" ]; then
            error "You must specify a version for 'force'"
            exit 1
        fi
        warn "Force setting migration version to $2 (manual sync)"
        "$MIGRATE_BIN" -path "$MIGRATIONS_DIR" -database "$DB_URL" force "$2"
        ;;
    help|--help|-h)
        show_usage
        ;;
    *)
        error "Unknown command: $ACTION"
        show_usage
        exit 1
        ;;
esac

info "Migration script complete."
