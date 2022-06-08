package response

import "project/model/system"

type SysAuthorityResponse struct {
	Authority system.SysAuthority `json:"authority"`
}
