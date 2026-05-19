package middleware

import (
	"net/http"
	"sync"

	"golang.org/x/time/rate"
)

type client struct {
	limiter *rate.Limiter
}

var (
	clients = make(map[string]*client)
	mu      sync.Mutex
)

func getClient(ip string) *rate.Limiter {
	mu.Lock()
	defer mu.Unlock()

	if c, ok := clients[ip]; ok {
		return c.limiter
	}

	//5 requests po sekundi po 10 puta
	limiter := rate.NewLimiter(5, 10)
	clients[ip] = &client{limiter: limiter}
	return limiter
}

func RateLimiterMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		limiter := getClient(r.RemoteAddr)
		if !limiter.Allow() {
			http.Error(w, "too many requests", http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// CleanupClients se poziva periodično da bi se očistili klijenti koji su možda prestali da šalju zahteve
func CleanupClients() {
	mu.Lock()
	defer mu.Unlock()
	clients = make(map[string]*client)
}
