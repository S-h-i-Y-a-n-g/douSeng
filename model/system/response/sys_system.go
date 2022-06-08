package response

import "project/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
