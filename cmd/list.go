/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"shelf/db"

	"github.com/dgraph-io/badger/v3"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all the stored keys",
	Long:  `This command will retrieve and print out all the keys currently stored in the database.`,
	Run: func(cmd *cobra.Command, args []string) {
		db := db.Instance()

		err := db.View(func(txn *badger.Txn) error {
			it := txn.NewIterator(badger.DefaultIteratorOptions)
			defer it.Close()

			for it.Seek([]byte("")); it.Valid(); it.Next() {
				item := it.Item()
				k := item.Key()
				fmt.Println(string(k))
			}
			return nil
		})

		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
