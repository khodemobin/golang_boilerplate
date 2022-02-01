package response

import "github.com/khodemobin/pio/provider/internal/domain"

type DeviceResource struct {
	ID     *int    `json:"nrc,omitempty"`
	Params *string `json:"p,omitempty"`
}

func NewDeviceResource(u *domain.Sample) *DeviceResource {
	return &DeviceResource{
		ID:     u.ID,
		Params: u.Params,
	}
}
