package cmdpkg

import (
	"fmt"
	"projetContact/internal/storage"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var addName string
var addEmail string

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Ajouter un contact",
	RunE: func(cmd *cobra.Command, args []string) error {
		// assurer que la configuration est chargée
		_ = viper.ReadInConfig()
		st, err := getStore()
		if err != nil {
			return err
		}
		c := &storage.Contact{Name: addName, Email: addEmail}
		if err := st.Add(c); err != nil {
			return err
		}
		fmt.Printf("Contact ajouté avec l'ID %d\n", c.ID)
		return nil
	},
}

func init() {
	addCmd.Flags().StringVarP(&addName, "name", "n", "", "Nom du contact (requis)")
	addCmd.Flags().StringVarP(&addEmail, "email", "e", "", "Email du contact (requis)")
	addCmd.MarkFlagRequired("name")
	addCmd.MarkFlagRequired("email")
}
