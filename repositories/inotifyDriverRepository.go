package repositories

type INotifyDriverRepository interface {
	SendNotification() error
}
