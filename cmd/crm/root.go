package main

import (
	"fmt"
	"os"
	"projetContact/internal/app"
	"projetContact/internal/storage"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "crm",
	Short: "Mini CRM CLI",
	Long:  "Un petit gestionnaire de contacts en ligne de commande.",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		// Si un chemin de configuration est fourni via le flag, l'utiliser (absolu ou relatif)
		if cfgFile != "" {
			viper.SetConfigFile(cfgFile)
			if err := viper.ReadInConfig(); err != nil {
				// tenter d'accepter une extension manquante en ajoutant .yaml
				if _, ok := err.(viper.ConfigFileNotFoundError); ok {
					// non trouvé au chemin donné : on tente d'autres emplacements
				} else {
					// autre erreur de lecture, renvoyer
					return err
				}
			}
		}

		// Si aucune config n'a été lue, tenter des emplacements par défaut
		if viper.ConfigFileUsed() == "" {
			viper.SetConfigName("config")
			viper.SetConfigType("yaml")
			viper.AddConfigPath(".")
			if err := viper.ReadInConfig(); err != nil {
				// pas de config trouvée : définir des valeurs par défaut raisonnables
				viper.SetDefault("storage.type", "memory")
				// Rappel utile à l'utilisateur : si vous voulez de la persistance,
				// copiez le fichier d'exemple et modifiez-le. Le dépôt ignore `config.yaml`.
				fmt.Fprintln(os.Stderr, "Aucun fichier de configuration trouvé. Le backend 'memory' (éphémère) sera utilisé.")
				fmt.Fprintln(os.Stderr, "Pour utiliser un backend persistant, copiez config.yaml.example en config.yaml puis réessayez:")
				fmt.Fprintln(os.Stderr, "  cp config.yaml.example config.yaml")
			}
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		st, err := getStore()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Échec de l'initialisation du stockage : %v\n", err)
			os.Exit(1)
		}
		app.Run(st)
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// ajouter un raccourci -c pour faciliter l'utilisation
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "fichier de configuration (par défaut ./config.yaml)")
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(updateCmd)
	rootCmd.AddCommand(deleteCmd)
}

func getStore() (storage.Storer, error) {
	stype := viper.GetString("storage.type")
	var st storage.Storer
	var err error
	switch stype {
	case "gorm":
		dbPath := viper.GetString("storage.gorm.path")
		if dbPath == "" {
			dbPath = "contacts.db"
		}
		st, err = storage.NewGORMStore(dbPath)
	case "json":
		file := viper.GetString("storage.json.path")
		if file == "" {
			file = "contacts.json"
		}
		st, err = storage.NewJSONStore(file)
	default:
		st = storage.NewMemoryStore()
	}
	return st, err
}
