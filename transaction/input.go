package transaction

import "golang-crowdfunding/user"

type GetCampaignTransationsInput struct {
	ID   int `uri:"id" binding:"required"`
	User user.User
}
