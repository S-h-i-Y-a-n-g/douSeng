package douseng

const UserFavoriteTableName = "ds_user_video_action"
const user_rowname = "user_id"
const video_rowname = "video_id"

type Favorite struct{}

func (v *Favorite) UerRowName() string {
	return user_rowname
}
func (v *Favorite) VideoRowName() string {
	return video_rowname
}

func (v *Favorite) GetUserFavoriteTableName() string {
	return UserFavoriteTableName
}
