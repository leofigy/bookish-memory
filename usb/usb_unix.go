//go:build !windows

package usb

// TODO just placeholder to avoid breaking the functionality on nix
type NixInspector struct{}

func NewInspektor() USBInspektor {
	return &NixInspector{}
}

func (s *NixInspector) GetDevicesHistory() ([]Info, error) {
	return make([]Info, 0), nil
}
