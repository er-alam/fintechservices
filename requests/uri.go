package uri

type Uri struct {
	Id   int64 `uri:"lead_id" binding:"required"`
	Name int64 `uri:"name" binding:"required"`
}
