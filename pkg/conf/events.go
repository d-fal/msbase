package conf

// GetEventReceiver get the event reciever details from config file
func (confRcv *ConfigRcv) GetEventReceiver() EventRecipient {
	return ConfigList.EventRecipients
}
