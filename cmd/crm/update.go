package main

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
		_ = viper.ReadInConfig()
		st, err := getStore()
		if err != nil {
			return err
		}
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
	updateCmd.Flags().IntVarP(&updID, "id", "i", 0, "ID du contact à modifier (requis)")
	updateCmd.Flags().StringVarP(&updName, "name", "n", "", "Nouveau nom (optionnel)")
	updateCmd.Flags().StringVarP(&updEmail, "email", "e", "", "Nouvel email (optionnel)")
	updateCmd.MarkFlagRequired("id")
}
