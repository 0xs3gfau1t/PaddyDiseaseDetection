package config

import (
	"os"

	"github.com/supabase-community/storage-go"
)

func NewSupabaseClient() *storage_go.Client {
	return storage_go.NewClient(os.Getenv("SUPABASE_CONN_STRING"), os.Getenv("SUPABASE_KEY"), nil)
}
