package cmdpkg

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
		// Load config
		if cfgFile != "" {
			viper.SetConfigFile(cfgFile)
		} else {
			viper.SetConfigName("config")
			viper.SetConfigType("yaml")
			viper.AddConfigPath(".")
		}
		if err := viper.ReadInConfig(); err != nil {
			// If no config found, use defaults
			fmt.Println("No config file found, using defaults (memory store)")
			viper.SetDefault("storage.type", "memory")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		// select store based on config
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
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to initialize store: %v\n", err)
			os.Exit(1)
		}
		app.Run(st)
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ./config.yaml)")
	// Add subcommands
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(updateCmd)
	rootCmd.AddCommand(deleteCmd)
}

// getStore choisit et initialise le Storer en fonction de viper
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
