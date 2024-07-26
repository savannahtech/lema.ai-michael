package server

import (
	"github.com/dilly3/houdini/internal/server/response"
	"net"
	"net/http"
	"sync"
	"time"
)

var (
	defaultLimiter *Limiter
)

func GetLimiter() *Limiter {
	return defaultLimiter
}

type Limiter struct {
	visitors map[string]*Visitor
	mu       sync.Mutex
	duration time.Duration
}

type Visitor struct {
	lastSeen time.Time
	requests int
}

func NewRateLimiter(d time.Duration) *Limiter {
	defaultLimiter = &Limiter{visitors: make(map[string]*Visitor), duration: d, mu: sync.Mutex{}}
	return defaultLimiter
}
func (l *Limiter) CleanUp() {
	for {
		time.Sleep(l.duration)
		l.mu.Lock()
		for ip, v := range l.visitors {
			if time.Since(v.lastSeen) > l.duration {
				delete(l.visitors, ip)
			}
		}
		l.mu.Unlock()
	}
}

func (l *Limiter) IPRateLimit(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			response.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		l.mu.Lock()

		visitor, found := l.visitors[ip]
		if !found {
			visitor = &Visitor{lastSeen: time.Now(), requests: 0}
			l.visitors[ip] = visitor
		}
		visitor.requests++
		visitor.lastSeen = time.Now()

		l.mu.Unlock()
		if visitor.requests > 5 {
			response.RespondWithError(w, http.StatusTooManyRequests, "rate limit exceeded")
			return
		} else {
			next.ServeHTTP(w, r)
		}

	})
}
