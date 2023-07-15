package cmd

import (
	"io/ioutil"
	"log"
	"os"
	"shelf/db"

	"github.com/dgraph-io/badger/v3"
	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Stores strings under a specified key",
	Long:  "The 'set' command takes a string input from the standard output and stores it under the specified key in the database. The key is given as an argument. Usage: echo \"your string\" | shelf set your-key",
	Run: func(cmd *cobra.Command, args []string) {
		db := db.Instance()
		key := args[0]

		data, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			log.Fatalf("Error reading from stdin: %v", err)
		}

		err = db.Update(func(txn *badger.Txn) error {
			err := txn.Set([]byte(key), data)
			return err
		})
		if err != nil {
			log.Fatalf("Error setting key %s: %v", key, err)
		}
	},
}

func init() {
	rootCmd.AddCommand(setCmd)
}
