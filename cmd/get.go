package cmd

import (
	"log"
	"os"
	"shelf/db"

	"github.com/dgraph-io/badger/v3"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieve strings stored under a specified key",
	Long:  "The 'get' command retrieves the string stored under the specified key in the database. The key is given as an argument. Usage: shelf get your-key",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		db := db.Instance()
		key := args[0]

		err := db.View(func(txn *badger.Txn) error {
			item, err := txn.Get([]byte(key))
			if err != nil {
				log.Fatalf("Error getting key %s: %v", key, err)
			}

			val, err := item.ValueCopy(nil)
			if err != nil {
				log.Fatal(err)
			}

			os.Stdout.Write(val)
			return nil
		})

		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
