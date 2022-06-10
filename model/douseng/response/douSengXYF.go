package request

type FavoriteResponse struct {
	StatusCode int    `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

type GetFavoriteResponse struct {
	FavoriteResponse
}

const UserTableName = "ds_user"
const VideoTableName = "ds_video"
