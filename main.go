package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var db *sql.DB

var cfgFile string

var cmd = &cobra.Command{
	Use:   "simple-http-server",
	Short: "Simple HTTP Server",
	Long:  "Simple HTTP Server",
	Run: func(cmd *cobra.Command, args []string) {
		runServer()
	},
}

func init() {
	cobra.OnInitialize(initConfig)

	cmd.PersistentFlags().StringVar(&cfgFile, "config", "", "Configuration file (default is ./simple-http-server.yaml)")
	cmd.PersistentFlags().IntP("port", "p", 8080, "Port on localhost on which to run server")
	cmd.PersistentFlags().StringP("db", "d", "", "The postgres connection string (e.g. postgresql://<username>:<password>@<database_ip>)")
	viper.BindPFlag("port", cmd.PersistentFlags().Lookup("port"))
	viper.BindPFlag("db", cmd.PersistentFlags().Lookup("db"))
}

// initConfig Read in config file and ENV variables if set.
func initConfig() {
	if cfgFile == "" {
		// Look for config file in working directory
		viper.SetConfigName("simple-http-server")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(".")
	} else {
		// Use configuration file specified on the command line.
		viper.SetConfigFile(cfgFile)
	}

	// Read in environment variables that match the flag names (but upper case)
	viper.AutomaticEnv()

	// If a config file is found, read it in.
	err := viper.ReadInConfig()
	if err == nil {
		fmt.Printf("Using config file: %s\n", viper.ConfigFileUsed())
	}
}

func runServer() {
	var err error

	// Connect to DB
	connStr := viper.GetString("db")
	fmt.Printf("Connecting to Database at %s...\n", connStr)
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// Setup HTTP Router
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/board/{boardID}", getBoard)
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	// Start HTTP Server
	port := viper.GetInt("port")
	address := fmt.Sprintf(":%d", port)
	fmt.Printf("Starting HTTP server on port %d...\n", port)
	err = http.ListenAndServe(address, r)
	if err != nil {
		fmt.Printf("Failed to run server: %s\n", err)
	}
}

func main() {
	cmd.Execute()
}
