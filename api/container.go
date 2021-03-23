package api

import (
	"encoding/json"
	"errors"

	"github.com/imroc/req"
	"github.com/patrickmn/go-cache"
)

// GetContainerId of uid
func (api *WeiboAPI) GetContainerId(uid string) (string, error) {
	key := "weibo:containerid:" + uid
	if value, found := api.cache.Get(key); found {
		return value.(string), nil
	}
	res, err := req.Get(
		"https://m.weibo.cn/api/container/getIndex",
		req.QueryParam{
			"type":  "uid",
			"value": uid,
		},
		req.Header{
			"Referer": "https://m.weibo.cn/",
		},
	)
	if err != nil {
		return "", err
	}
	body := &WeiboUserIndex{}
	if err := res.ToJSON(body); err != nil {
		return "", err
	}

	for _, tab := range body.Data.TabsInfo.Tabs {
		if tab.TabKey == "weibo" {
			value := tab.Containerid
			api.cache.Set(key, value, cache.DefaultExpiration)
			return value, nil
		}
	}

	return "", errors.New("not found correct container for 'weibo'")

}

func UnmarshalWeiboUserIndex(data []byte) (WeiboUserIndex, error) {
	var r WeiboUserIndex
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *WeiboUserIndex) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type WeiboUserIndex struct {
	Ok   int64 `json:"ok"`
	Data Data  `json:"data"`
}

type Data struct {
	AvatarGuide  []interface{} `json:"avatar_guide"`
	IsStarStyle  int64         `json:"isStarStyle"`
	UserInfo     DataUserInfo  `json:"userInfo"`
	FansScheme   string        `json:"fans_scheme"`
	FollowScheme string        `json:"follow_scheme"`
	TabsInfo     TabsInfo      `json:"tabsInfo"`
	Scheme       string        `json:"scheme"`
	ShowAppTips  int64         `json:"showAppTips"`
}

type TabsInfo struct {
	SelectedTab int64 `json:"selectedTab"`
	Tabs        []Tab `json:"tabs"`
}

type Tab struct {
	ID              int64            `json:"id"`
	TabKey          string           `json:"tabKey"`
	MustShow        int64            `json:"must_show"`
	Hidden          int64            `json:"hidden"`
	Title           string           `json:"title"`
	TabType         string           `json:"tab_type"`
	Containerid     string           `json:"containerid"`
	Apipath         *string          `json:"apipath,omitempty"`
	URL             *string          `json:"url,omitempty"`
	FilterGroup     []FilterGroup    `json:"filter_group,omitempty"`
	FilterGroupInfo *FilterGroupInfo `json:"filter_group_info,omitempty"`
}

type FilterGroup struct {
	Name        string `json:"name"`
	Containerid string `json:"containerid"`
	Title       string `json:"title"`
	Scheme      string `json:"scheme"`
}

type FilterGroupInfo struct {
	Title      string `json:"title"`
	Icon       string `json:"icon"`
	IconName   string `json:"icon_name"`
	IconScheme string `json:"icon_scheme"`
}

type DataUserInfo struct {
	ID              int64         `json:"id"`
	ScreenName      string        `json:"screen_name"`
	ProfileImageURL string        `json:"profile_image_url"`
	ProfileURL      string        `json:"profile_url"`
	StatusesCount   int64         `json:"statuses_count"`
	Verified        bool          `json:"verified"`
	VerifiedType    int64         `json:"verified_type"`
	VerifiedTypeEXT int64         `json:"verified_type_ext"`
	VerifiedReason  string        `json:"verified_reason"`
	CloseBlueV      bool          `json:"close_blue_v"`
	Description     string        `json:"description"`
	Gender          string        `json:"gender"`
	Mbtype          int64         `json:"mbtype"`
	Urank           int64         `json:"urank"`
	Mbrank          int64         `json:"mbrank"`
	FollowMe        bool          `json:"follow_me"`
	Following       bool          `json:"following"`
	FollowersCount  int64         `json:"followers_count"`
	FollowCount     int64         `json:"follow_count"`
	CoverImagePhone string        `json:"cover_image_phone"`
	AvatarHD        string        `json:"avatar_hd"`
	Like            bool          `json:"like"`
	LikeMe          bool          `json:"like_me"`
	ToolbarMenus    []ToolbarMenu `json:"toolbar_menus"`
}

type ToolbarMenu struct {
	Type      string               `json:"type"`
	Name      string               `json:"name"`
	Params    ToolbarMenuParams    `json:"params"`
	Actionlog Actionlog            `json:"actionlog"`
	UserInfo  *ToolbarMenuUserInfo `json:"userInfo,omitempty"`
	SubType   *int64               `json:"sub_type,omitempty"`
	Scheme    *string              `json:"scheme,omitempty"`
}

type Actionlog struct {
	ActCode string `json:"act_code"`
	Fid     string `json:"fid"`
	OID     string `json:"oid"`
	Cardid  string `json:"cardid"`
	EXT     string `json:"ext"`
}

type ToolbarMenuParams struct {
	Uid       *int64     `json:"uid,omitempty"`
	Extparams *Extparams `json:"extparams,omitempty"`
	Scheme    *string    `json:"scheme,omitempty"`
	MenuList  []MenuList `json:"menu_list,omitempty"`
}

type Extparams struct {
	Followcardid interface{} `json:"followcardid"`
}

type MenuList struct {
	Type      string         `json:"type"`
	Name      string         `json:"name"`
	Params    MenuListParams `json:"params"`
	Actionlog Actionlog      `json:"actionlog"`
	Scheme    string         `json:"scheme"`
}

type MenuListParams struct {
	Scheme string `json:"scheme"`
}

type ToolbarMenuUserInfo struct {
	ID              int64  `json:"id"`
	Idstr           string `json:"idstr"`
	ScreenName      string `json:"screen_name"`
	ProfileImageURL string `json:"profile_image_url"`
	Following       bool   `json:"following"`
	Verified        bool   `json:"verified"`
	VerifiedType    int64  `json:"verified_type"`
	Remark          string `json:"remark"`
	AvatarLarge     string `json:"avatar_large"`
	AvatarHD        string `json:"avatar_hd"`
	VerifiedTypeEXT int64  `json:"verified_type_ext"`
	FollowMe        bool   `json:"follow_me"`
	Mbtype          int64  `json:"mbtype"`
	Mbrank          int64  `json:"mbrank"`
	Level           int64  `json:"level"`
	Type            int64  `json:"type"`
	StoryReadState  int64  `json:"story_read_state"`
	AllowMsg        int64  `json:"allow_msg"`
	SpecialFollow   bool   `json:"special_follow"`
}
