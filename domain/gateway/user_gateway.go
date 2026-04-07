package gateway

type UserGateway interface {
	ExistsActiveSuperAdmin(username string) bool
}
