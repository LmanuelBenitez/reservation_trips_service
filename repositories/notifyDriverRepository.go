package repositories

type NotifyDriverRepository struct {
}

func NewNotifyDriverRepository() *NotifyDriverRepository {
	return &NotifyDriverRepository{}
}

func (repository *NotifyDriverRepository) SendNotification() error {
	return nil
}
