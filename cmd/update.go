package cmdpkg

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var updID int
var updName string
var updEmail string

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Modifier un contact",
	RunE: func(cmd *cobra.Command, args []string) error {
		// assurer que la configuration est chargée
		_ = viper.ReadInConfig()
		st, err := getStore()
		if err != nil {
			return err
		}
		// s'assurer que le contact existe
		_, err = st.GetByID(updID)
		if err != nil {
			return err
		}
		if err := st.Update(updID, updName, updEmail); err != nil {
			return err
		}
		fmt.Println("Contact mis à jour")
		return nil
	},
}

func init() {
	updateCmd.Flags().IntVarP(&updID, "id", "i", 0, "ID of contact to update (required)")
	updateCmd.Flags().StringVarP(&updName, "name", "n", "", "New name (optional)")
	updateCmd.Flags().StringVarP(&updEmail, "email", "e", "", "New email (optional)")
	updateCmd.MarkFlagRequired("id")
}
