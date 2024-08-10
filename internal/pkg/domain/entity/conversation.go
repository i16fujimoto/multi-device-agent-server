package entity

const (
	DeviceIOS      = "ios"
	DeviceAndroid  = "android"
	DevicePC       = "pc"
	DeviceWatch    = "watch"
	DeviceHologram = "hologram"
)

type Conversation struct {
	Time   int64   `json:"time" validate:"required"`
	User   string  `json:"user" validate:"required"`
	Agent  string  `json:"agent" validate:"required"`
	Device *string `json:"device"`
}

func AllDevices() []string {
	return []string{
		DeviceIOS,
		DeviceAndroid,
		DevicePC,
		DeviceWatch,
		DeviceHologram,
	}
}
