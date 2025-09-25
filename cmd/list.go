package cmdpkg

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lister tous les contacts",
	RunE: func(cmd *cobra.Command, args []string) error {
		_ = viper.ReadInConfig()
		st, err := getStore()
		if err != nil {
			return err
		}
		contacts, err := st.GetAll()
		if err != nil {
			return err
		}
		if len(contacts) == 0 {
			fmt.Println("Aucun contact")
			return nil
		}
		for _, c := range contacts {
			fmt.Printf("ID:%d Name:%s Email:%s\n", c.ID, c.Name, c.Email)
		}
		return nil
	},
}
