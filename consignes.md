# Mini-CRM CLI 

Un gestionnaire de contacts simple et efficace en ligne de commande, écrit en Go. Ce projet a été conçu comme un cas pratique pour illustrer les bonnes pratiques de développement Go, incluant :
* Une architecture en packages découplés.
* L'injection de dépendances via les interfaces.
* La création d'une CLI professionnelle avec Cobra.
* La gestion de configuration externe avec Viper.
* Plusieurs couches de persistance, notamment avec GORM et SQLite.

## Contexte du devoir 

Nous avons fait évoluer notre application Mini-CRM depuis un simple script jusqu'à un programme modulaire utilisant des packages et la persistance de données via un fichier JSON. L'objectif de ce devoir est de finaliser cette transformation pour en faire un véritable outil en ligne de commande, robuste et configurable.

Ce devoir est divisé en deux grandes étapes indépendantes mais complémentaires.

## Partie 1 : Intégration d'une Base de Données avec GORM/SQLite (45%)

**Objectifld** : Remplacer notre système de stockage JSON par une base de données SQL via l'ORM GORM. Grâce à notre architecture basée sur les interfaces, cette modification devrait se faire sans impacter la logique métier.

### Étapes 

1. **Ajouter les dépendances nécessaires**, à la racine du projet ajoutez GORM et son driver SQLite : 
```bash
    go get gorm.io/gorm
    go get gorm.io/driver/sqlite
```

2. Mettre à jour la struct `Contact`
3. Créer le `GORMStore`
4. Implémenter l'interface `Storer`
5. Intégrer dans `cmd/root.go`


## Partie 2 : Création d'une CLI Professionnelle avec Cobra & Viper (55%)

### Étapes

1. Ajouter les dépendances nécessaire :
```bash
    go get github.com/spf13/cobra
    go get github.com/spf13/viper
```
2. Réorganiser les projets : Adoptez une structure de projet orientée Cobra
3. Créer le fichier de configuration (`.yaml`) qui permettra de choisir le stockage
4. Implémenter la commande Racine (`cmd/root.go`)
5. Implémenter les sous-commandes 
   * Pour chaque fonctionnalité (ajouter, lister, mettre à jour, supprimer), créez un fichier .go dédié dans le package cmd.

## Critères de Réussite

* Le programme **compile et s'exécute sans erreur**.
* Toutes les **sous-commandes** (add, list, update, delete) sont fonctionnelles.
* L'application utilise bien la base de données **SQLite** (un fichier .db est créé et mis à jour) lorsque type: "gorm" est configuré dans config.yaml.
* Il est possible de **basculer sur le stockage json ou memory en modifiant simplement le fichier config.yaml**, sans recompiler.
* Le code est propre, formaté avec gofmt, et raisonnablement commenté.
* Un documentation (readme) claire et complète

## Fonctionnalités finales attendues

* **Gestion complète des contacts (CRUD)** : Ajouter, Lister, Mettre à jour et Supprimer des contacts.
* **Interface en ligne de commande** : Commandes et sous-commandes claires et standardisées.
* **Configuration externe** : Le comportement de l'application (notamment le type de stockage) peut être modifié sans recompiler.
* **Persistance des données** : Support de multiples backends de stockage :
  * GORM/SQLite : Une base de données SQL robuste contenue dans un simple fichier.
  * Fichier JSON : Une sauvegarde simple et lisible.
  * En mémoire : Un stockage éphémère pour les tests.



