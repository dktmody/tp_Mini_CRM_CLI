# Projet Efrei Contact

# Mini-CRM CLI (Go)

Ce dépôt contient une mini-application CLI pour gérer des contacts (CRUD).
L'application supporte trois backends de stockage interchangeables : mémoire, fichier JSON et GORM/SQLite.  
Pour avoir plus d'infos concernant les consignes, voir le fichier à la racine nommé consignes.md.

## Prérequis

- Go 1.25+ installé
- (optionnel) git

## Installer les dépendances

Depuis la racine du projet, lancez :

```bash
go get gorm.io/gorm
go get gorm.io/driver/sqlite
go get github.com/spf13/cobra
go get github.com/spf13/viper
```

Cela mettra à jour votre module `go.mod` et téléchargera les packages nécessaires.

## Configuration

Le fichier `config.yaml` à la racine permet de choisir le backend de stockage. Exemple fourni :

```yaml
storage:
	type: memory # options: memory, json, gorm
	json:
		path: contacts.json
	gorm:
		path: contacts.db
```

Pour utiliser SQLite/GORM, définissez `storage.type: gorm`. Le fichier SQLite sera créé (par défaut `contacts.db`).

Pour utiliser le stockage JSON, mettez `storage.type: json`. Le fichier JSON sera créé (par défaut `contacts.json`).

Si aucun fichier de configuration n'est trouvé, l'application utilisera `memory` par défaut.

Vous pouvez aussi indiquer un fichier de config personnalisé avec `--config ./monconfig.yaml`.

## Commandes

Le binaire s'appelle `crm`. Les commandes principales sont :

- `crm add -n NAME -e EMAIL` : ajoute un contact
- `crm list` : liste tous les contacts
- `crm update -i ID [-n NAME] [-e EMAIL]` : met à jour un contact
- `crm delete -i ID` : supprime un contact

Exemples :

```bash
go run ./cmd/crm --config config.yaml add -n "Alice" -e "alice@example.com"
go run ./cmd/crm --config config.yaml list
go run ./cmd/crm --config config.yaml update -i 1 -n "Alice Dupont"
go run ./cmd/crm --config config.yaml delete -i 1
```

Remarque : quand `storage.type` est `gorm`, utilisez le même `config.yaml` et vérifiez que `contacts.db` est créé et qu'il contient des données après `add`.

## Quick start — commandes copy‑paste

Copiez le fichier d'exemple et suivez UNE des deux méthodes ci‑dessous : Bash (Git Bash / WSL) ou PowerShell.

1. Préparer la config (optionnel) :

```bash
cp config.yaml.example config.yaml
```

2. Construire le binaire (optionnel) :

```bash
go build -o crm ./cmd/crm
```

3A) Tester (Bash / Git Bash / WSL)

```bash
# Backend mémoire (par défaut)
./crm --config config.yaml list
./crm --config config.yaml add -n "MemoryUser" -e "m@example.com"
./crm --config config.yaml list

# Passer au backend JSON (modifie config.yaml)
cat > config.yaml <<'YAML'
storage:
  type: json
  json:
    path: demo_contacts.json
YAML
./crm --config config.yaml add -n "JsonUser" -e "j@example.com"
./crm --config config.yaml list
cat demo_contacts.json

# Passer au backend GORM/SQLite (modifie config.yaml)
cat > config.yaml <<'YAML'
storage:
  type: gorm
  gorm:
    path: demo_contacts.db
YAML
./crm --config config.yaml add -n "GormUser" -e "g@example.com"
./crm --config config.yaml list
ls -l demo_contacts.db

# Nettoyage (facultatif)
rm -f demo_contacts.json demo_contacts.db crm
```

3B) Tester (PowerShell)

```powershell
# Copier l'exemple
Copy-Item config.yaml.example config.yaml

# Backend mémoire (par défaut)
.
\crm --config config.yaml list
.
\crm --config config.yaml add -n "MemoryUser" -e "m@example.com"
.
\crm --config config.yaml list

# Backend JSON
@"
storage:
  type: json
  json:
    path: demo_contacts.json
"@ | Out-File -Encoding UTF8 config.yaml
.
\crm --config config.yaml add -n "JsonUser" -e "j@example.com"
.
\crm --config config.yaml list
Get-Content demo_contacts.json

# Backend GORM
@"
storage:
  type: gorm
  gorm:
    path: demo_contacts.db
"@ | Out-File -Encoding UTF8 config.yaml
.
\crm --config config.yaml add -n "GormUser" -e "g@example.com"
.
\crm --config config.yaml list
Get-Item demo_contacts.db

# Nettoyage
Remove-Item -Force demo_contacts.json, demo_contacts.db, crm
```

Notes :

- Vous pouvez aussi exécuter la démonstration automatisée fournie :
  - Bash : `./scripts/run-demo.sh`
  - PowerShell : `.\scripts\run-demo.bat` (ou `PowerShell -ExecutionPolicy Bypass -File .\scripts\run-demo.ps1`)

## Notes de développement

- Le code est organisé en packages : `internal/app` contient la logique interactive; `internal/storage` contient l'interface `Storer` et les implémentations (`memory`, `json`, `gorm`).
- Les commandes Cobra sont sous `cmd/crm`.

## Choix du driver SQLite

Par défaut ce projet utilise un pilote SQLite écrit en Go pur (`modernc.org/sqlite`) pour éviter la dépendance à CGO. Cela présente deux avantages :

- Portabilité : pas besoin d'un compilateur C (gcc) sur la machine de développement.
- Simplicité pour les environnements CI/Windows où CGO peut poser des problèmes.

Si vous préférez utiliser le pilote natif `github.com/mattn/go-sqlite3` (implémentation C), notez que vous devrez disposer d'un compilateur C et compiler avec CGO activé :

1. Installer un toolchain C (par exemple MinGW sur Windows) et s'assurer que `gcc` est disponible dans le PATH.
2. Dans votre environnement, activer CGO lors de la compilation :

```bash
export CGO_ENABLED=1
go build ./cmd/crm
```

3. Remplacez l'import du driver dans le code (si vous voulez explicitement utiliser `go-sqlite3`) et ajustez `go.mod`.

Nous avons choisi `modernc.org/sqlite` pour rendre le projet plus simple à exécuter sur différentes plateformes sans configuration supplémentaire.

## Démonstration pas‑à‑pas : tester chaque backend

Ces instructions supposent que vous êtes dans la racine du projet (`c:\Users\diakh\Desktop\tp1`) et que Go est installé.

1. Construire la CLI (optionnel, `go run` fonctionne aussi) :

```bash
go build -o crm ./cmd/crm
```

Si vous préférez utiliser `go run` directement, remplacez `./crm` par `go run ./cmd/crm --config config.yaml` dans les commandes ci‑dessous.

2. Backend en mémoire (temporaire)

```bash
# utiliser le mode mémoire (par défaut si config manquante)
./crm --config config.yaml list
./crm --config config.yaml add -n "MemoryUser" -e "m@example.com"
./crm --config config.yaml list
# le stockage est éphémère et disparaît après fermeture
```

3. Backend JSON (persistant dans un fichier)

Préparez `config.json.yaml` ou modifiez `config.yaml` :

```yaml
storage:
  type: json
  json:
    path: test_contacts.json
```

Puis exécutez :

```bash
# ajouter un contact
./crm --config config.yaml add -n "JsonUser" -e "j@example.com"
# lister les contacts
./crm --config config.yaml list
# vérifier le fichier JSON créé
cat test_contacts.json
```

4. Backend GORM/SQLite (persistant dans un fichier .db)

Préparez `config.yaml` avec :

```yaml
storage:
  type: gorm
  gorm:
    path: test_contacts.db
```

Puis exécutez :

```bash
# ajouter un contact
./crm --config config.yaml add -n "GormUser" -e "g@example.com"
# lister
./crm --config config.yaml list
# vérifier que le fichier .db a été créé
ls -l test_contacts.db
```

## Commandes copy‑paste prêtes (Bash et PowerShell)

Si vous voulez exécuter la démonstration automatiquement :

Bash (Linux / macOS / Git Bash sur Windows) :

```bash
./scripts/run-demo.sh
```

PowerShell (Windows) :

```powershell
.\scripts\run-demo.ps1
```

Le fichier d'exemple `config.yaml.example` est fourni ; copiez‑le en `config.yaml` si vous voulez partir d'une configuration propre :

```bash
cp config.yaml.example config.yaml
```

Remarque importante : le dépôt ignore `config.yaml` (fichier de config local). Ne committez pas votre `config.yaml` — gardez `config.yaml.example` dans le repo comme modèle.

Notes :

- Les chemins (test_contacts.json / test_contacts.db) peuvent être relatifs ; ils seront créés dans le répertoire courant.
- Si vous souhaitez utiliser `go run` au lieu du binaire construit :

```bash
go run ./cmd/crm --config config.yaml add -n "Tmp" -e "t@e.com"
```

