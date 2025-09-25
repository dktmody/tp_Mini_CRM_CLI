# Script PowerShell pour démonstration des 3 backends
Set-StrictMode -Version Latest
$root = Split-Path -Parent $PSScriptRoot
Set-Location $root

Write-Host "1) Build du binaire"
go build -o crm ./cmd/crm

Write-Host "`n--- Backend: memory ---"
& .\crm --config config.yaml add -n "DemoMemory" -e "memory@example.com"
& .\crm --config config.yaml list

Write-Host "`n--- Backend: json ---"
@"
storage:
  type: json
  json:
    path: demo_contacts.json
"@ | Out-File -Encoding UTF8 config.yaml
& .\crm --config config.yaml add -n "DemoJSON" -e "json@example.com"
& .\crm --config config.yaml list
Write-Host "Fichier JSON créé: demo_contacts.json"

Write-Host "`n--- Backend: gorm ---"
@"
storage:
  type: gorm
  gorm:
    path: demo_contacts.db
"@ | Out-File -Encoding UTF8 config.yaml
& .\crm --config config.yaml add -n "DemoGorm" -e "gorm@example.com"
& .\crm --config config.yaml list
if (Test-Path demo_contacts.db) { Get-Item demo_contacts.db | Format-List }

Write-Host "`nNettoyage: suppression des fichiers de démonstration"
Remove-Item -Force demo_contacts.json, demo_contacts.db, crm

Write-Host "Démonstration terminée"
