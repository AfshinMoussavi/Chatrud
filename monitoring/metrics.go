package monitoring

import "github.com/prometheus/client_golang/prometheus"

var (
	LoginCounter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "app_login_total",
			Help: "Total number of successful user logins",
		})
	CreateCounter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "app_create_total",
			Help: "Total number of successful user creates",
		})
	GetUsersCounter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "app_get_users_total",
			Help: "Total number of requests to GET /users",
		})
)

func InitMetrics() {
	prometheus.MustRegister(LoginCounter, CreateCounter, GetUsersCounter)
}
