package middleware

import "net/http"

func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")

		// 2. Обработка Preflight-запроса (OPTIONS)
		if r.Method == http.MethodOptions {
			// Разрешаем все нужные вам методы
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			w.WriteHeader(http.StatusOK)
			return
		}

		// Проверяем JSON ТОЛЬКО если это методы изменения данных (POST, PUT, DELETE)
		if r.Method == http.MethodPost || r.Method == http.MethodPut || r.Method == http.MethodDelete {
			if r.Header.Get("Content-Type") != "application/json" {
				http.Error(w, `{"error": "Ожидался JSON (Content-Type: application/json)"}`, http.StatusUnsupportedMediaType)
				return
			}
		}

		next.ServeHTTP(w, r)
	})
}
