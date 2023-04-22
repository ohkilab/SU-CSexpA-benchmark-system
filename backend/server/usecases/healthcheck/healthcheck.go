package healthcheck

type HealthcheckInteractor struct{}

func NewInteractor() *HealthcheckInteractor {
	return &HealthcheckInteractor{}
}

func (i *HealthcheckInteractor) Ping(ping string) string {
	return ping + " Pong:)"
}
