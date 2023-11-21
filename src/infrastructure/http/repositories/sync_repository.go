package repositories

type SyncRepository interface {
	GetIDRanges() (string, error)
}
