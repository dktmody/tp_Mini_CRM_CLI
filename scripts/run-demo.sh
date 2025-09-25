#!/usr/bin/env bash
# Script de démonstration pour tester les 3 backends (bash)
set -euo pipefail
ROOT=$(cd "$(dirname "$0")/.." && pwd)
cd "$ROOT"

echo "1) Build du binaire"
go build -o crm ./cmd/crm

# Memory (par défaut)
echo "\n--- Backend: memory ---"
./crm --config config.yaml add -n "DemoMemory" -e "memory@example.com"
./crm --config config.yaml list

# JSON
echo "\n--- Backend: json ---"
cat > config.yaml <<'YAML'
storage:
  type: json
  json:
    path: demo_contacts.json
YAML
./crm --config config.yaml add -n "DemoJSON" -e "json@example.com"
./crm --config config.yaml list
echo "Fichier JSON créé: demo_contacts.json"

# GORM/SQLite
echo "\n--- Backend: gorm ---"
cat > config.yaml <<'YAML'
storage:
  type: gorm
  gorm:
    path: demo_contacts.db
YAML
./crm --config config.yaml add -n "DemoGorm" -e "gorm@example.com"
./crm --config config.yaml list
ls -l demo_contacts.db || true

echo "\nNettoyage: suppression des fichiers de démonstration"
rm -f demo_contacts.json demo_contacts.db crm

echo "Démonstration terminée"
