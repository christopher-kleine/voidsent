package main

import (
	"log"
	"net/http"
	"time"

	"github.com/christopher-kleine/sse"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/joho/godotenv/autoload"

	"github.com/gridanias-helden/voidsent/pkg/config"
	"github.com/gridanias-helden/voidsent/pkg/services/session"
	"github.com/gridanias-helden/voidsent/pkg/storage/memory"
)

func main() {
	appConfig, err := config.New()
	if err != nil {
		log.Fatalln(err)
	}

	hub := sse.New()
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	sessionService := memory.NewSessions(24 * time.Hour)

	// discordHandler := &session.Discord{
	// 	OAuth: &oauth2.Config{
	// 		RedirectURL:  appConfig.RedirectURL,
	// 		Scopes:       []string{discord.ScopeIdentify},
	// 		Endpoint:     discord.Endpoint,
	// 		ClientID:     appConfig.DiscordClientID,
	// 		ClientSecret: appConfig.DiscordClientSecret,
	// 	},
	// 	KV:       make(map[string]time.Time),
	// 	Sessions: sessionService,
	// }
	guestHandler := &session.GuestLogin{
		Sessions: sessionService,
	}

	// wrappedRouter := middleware.Chain(
	// 	router,
	// 	middleware.WithLogging,
	// 	middleware.WithSession(sessionService),
	// )

	r.Handle("/*", http.FileServer(http.Dir(appConfig.Static)))
	r.Route("/auth", func(r chi.Router) {
		// router.HandleFunc("/discord", discordHandler.Auth) // Disable for now
		// router.HandleFunc("/discord/callback", discordHandler.Callback)
		r.Get("/guest", guestHandler.Register)
		// router.Get("/logout", discordHandler.Logout)
	})
	r.Route("/api", func(r chi.Router) {
		r.Handle("/sse", hub)
	})

	log.Printf("Listening on %s", appConfig.Bind)
	log.Fatal(http.ListenAndServe(appConfig.Bind, r))
}
