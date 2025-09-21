package graph

import (
	"log"
	"os"

	storage "github.com/supabase-community/storage-go"
)

var SupabaseClient *storage.Client

func InitSupabase() {
	url := os.Getenv("SUPABASE_URL")

	key := os.Getenv("SUPABASE_SERVICE_ROLE_KEY")
	if key == "" {

		key = os.Getenv("SUPABASE_API_KEY")
		log.Println("Warning: Using anon key instead of service role key for uploads")
	}

	log.Printf("Initializing Supabase client with URL: %s", url)
	log.Printf("Supabase key length: %d", len(key))

	SupabaseClient = storage.NewClient(url, key, nil)
	if SupabaseClient == nil {
		log.Fatal("Failed to initialize Supabase client")
	}
	log.Println("Supabase client initialized successfully")
}
