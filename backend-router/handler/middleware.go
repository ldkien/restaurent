package handler

import (
	"bytes"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/time/rate"
	"io/ioutil"
	"net"
	"net/http"
	"regexp"
	"restaurant/backend-base/app"
	"restaurant/backend-base/entity"
	backendError "restaurant/backend-base/error"
	log "restaurant/backend-base/logger"
	"restaurant/backend-entity/entities"
	"sync"
	"time"
)

type visitor struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

var visitors = make(map[string]*visitor)
var mu sync.RWMutex

func init() {
	go cleanupVisitors()
}

func cleanupVisitors() {
	for {
		time.Sleep(time.Minute)

		mu.Lock()
		for ip, v := range visitors {
			if time.Since(v.lastSeen) > 3*time.Minute {
				delete(visitors, ip)
			}
		}
		mu.Unlock()
	}
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//rate limit
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			log.Logger.Error(err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		limiter := getVisitor(ip)
		if limiter.Allow() == false {
			http.Error(w, http.StatusText(429), http.StatusTooManyRequests)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		url := r.RequestURI
		match, _ := regexp.MatchString(app.API_PUBLIC+"*", url)
		if match {
			next.ServeHTTP(w, r)
			return
		}
		claims := &entity.Claims{}
		var baseRequest entities.BaseRequest
		err = json.NewDecoder(r.Body).Decode(&baseRequest)
		if err != nil || baseRequest.Common == nil {
			response := entities.BaseResponse{
				Error: backendError.GetError(backendError.UNAUTHENTICATED),
			}
			data := []byte(app.ConvertToJson(&response))
			w.Write(data)
			return
		}

		common := *baseRequest.Common
		tokenHeader := common.Token
		if len(tokenHeader) == 0 {
			response := entities.BaseResponse{
				Error: backendError.GetError(backendError.UNAUTHENTICATED),
			}
			data := []byte(app.ConvertToJson(&response))
			w.Write(data)
			return
		}

		tkn, err := jwt.ParseWithClaims(tokenHeader, claims, func(token *jwt.Token) (interface{}, error) {
			return app.JWT_KEY, nil
		})

		if err != nil {
			log.Logger.Error(err)
			response := entities.BaseResponse{
				Error: backendError.GetError(backendError.UNAUTHENTICATED),
			}
			data := []byte(app.ConvertToJson(&response))
			w.Write(data)
			return
		}
		if !tkn.Valid {
			response := entities.BaseResponse{
				Error: backendError.GetError(backendError.UNAUTHENTICATED),
			}
			data := []byte(app.ConvertToJson(&response))
			w.Write(data)
			return
		}

		if common.User == nil || common.User.Username != claims.Username {
			response := entities.BaseResponse{
				Error: backendError.GetError(backendError.NOT_USER),
			}
			data := []byte(app.ConvertToJson(&response))
			w.Write(data)
			return
		}

		common.User.Username = claims.Username
		common.User.Group = claims.Group
		body := []byte(app.ConvertToJson(&baseRequest))
		r.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		w.WriteHeader(http.StatusOK)
		next.ServeHTTP(w, r)
	})
}

func getVisitor(ip string) *rate.Limiter {
	mu.Lock()
	defer mu.Unlock()

	v, exists := visitors[ip]
	if !exists {
		limiter := rate.NewLimiter(1, 3)
		visitors[ip] = &visitor{limiter, time.Now()}
		return limiter
	}

	v.lastSeen = time.Now()
	return v.limiter
}
