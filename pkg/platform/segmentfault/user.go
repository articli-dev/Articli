package segmentfault

import "github.com/juju/errors"

type User struct {
	ID         int64  `json:"id"`
	Slug       string `json:"slug"`
	Name       string `json:"name"`
	URL        string `json:"url"`
	Rank       int    `json:"rank"`
	ProfileURL string `json:"profile_url"`
	AvatarURL  string `json:"avatar_url"`
	Created    int64  `json:"created"`
}

func (u *User) GetURL() string {
	return BuildURL(u.URL)
}

type Roles struct {
	IsAdmin  bool `json:"is_admin"`
	IsEditor bool `json:"is_editor"`
}

type GetUserResponse struct {
	User      *User  `json:"user"`
	Roles     *Roles `json:"roles"`
	IsChange  bool   `json:"is_change"`
	CanChange bool   `json:"can_change"`
}

func (c *Client) GetMe() (resp *GetUserResponse, err error) {
	endpoint := "/user/@me"
	err = c.Get(endpoint, nil, &resp)
	err = errors.Trace(err)
	return
}
