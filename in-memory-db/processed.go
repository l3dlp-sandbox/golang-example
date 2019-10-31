package inMemoryDb
type ClientPro struct {
	ClientID string
	ConnectedAt string
	Keepalive   int
	Node        string
}
type SessionsPro struct {
	CreatedAt        string
	ClientID         string
	Subscriptions    int
}
type SubscriptionsPro struct {
	ClientID string
	Qos      int
	Topic    string
}

type Processed struct {
	ClientPro []ClientPro
	SessionPro []SessionsPro
	SubscriptionPro []SubscriptionsPro
}

