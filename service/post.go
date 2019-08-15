package service

import (
	"adv/formdata"
	"adv/repo"
	"encoding/json"
	"github.com/jinzhu/gorm"
)

type PostService struct {
	repo *repo.PostRepo
}

func (ps *PostService) GetPostList(pager *formdata.Pager) map[string]interface{} {
	if pager.Page == 0 {
		pager.Page = 1
	}
	if pager.PageSize == 0 {
		pager.PageSize = 20
	}
	posts, total := ps.repo.GetPostList(pager.Page, pager.PageSize)
	return map[string]interface{}{
		"posts": posts, "total": total,
	}
}

func (ps *PostService) SetPostAdvJSON(advInfo *formdata.PostAdvInfo) bool {
	post := ps.repo.GetPost(advInfo.PostId)
	if post != nil {
		post.Ads = make([]string, 0)
		for _, adv := range advInfo.AdsInfo {
			if adv != "" {
				post.Ads = append(post.Ads, adv)
			}
		}
		return ps.repo.Save(post)
	}
	return false
}

func (ps *PostService) GetPostAds(filename string) []string {
	if post := ps.repo.GetPostByFilename(filename); post != nil {
		ads := make([]string, 0)
		json.Unmarshal([]byte(post.AdvJson), &ads)
		return ads
	} else {
		return []string{}
	}
}

func NewPostService(db *gorm.DB) *PostService {
	ps := &PostService{}
	ps.repo = repo.NewPostRepo(db)
	return ps
}
