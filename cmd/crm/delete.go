package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var delID int

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Supprimer un contact",
	RunE: func(cmd *cobra.Command, args []string) error {
		_ = viper.ReadInConfig()
		st, err := getStore()
		if err != nil {
			return err
		}
		if err := st.Delete(delID); err != nil {
			return err
		}
		fmt.Println("Contact supprimé")
		return nil
	},
}

func init() {
	deleteCmd.Flags().IntVarP(&delID, "id", "i", 0, "ID du contact à supprimer (requis)")
	deleteCmd.MarkFlagRequired("id")
}
