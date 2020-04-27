package responses

type TaobaoTbkUatmFavoritesItemGetResponse struct {
	ErrorResponse

	Results      []UatmTbkItem `json:"results,omitempty"`
	TotalResults int           `json:"total_results,omitempty"`
}

// FavoriteItem	招商宝贝信息
type UatmTbkItem struct {
	// NumIid 商品ID
	NumIid int `json:"num_iid,omitempty"`
	//Title 商品标题
	Title string `json:"title,omitempty"`
	//PictUrl 商品主图
	PictUrl string `json:"pict_url,omitempty"`
	//SmallImages 商品小图列表
	SmallImages []string `json:"small_images,omitempty"`
	//ReservePrice 	商品一口价格
	ReservePrice string `json:"reserve_price,omitempty"`
	//ZkFinalPrice 商品折扣价格
	ZkFinalPrice string `json:"zk_final_price,omitempty"`
	//UserType 卖家类型，0表示集市，1表示商城
	UserType int `json:"user_type,omitempty"`
	//Provcity 	宝贝所在地
	Provcity string `json:"provcity,omitempty"`
	//ItemUrl 商品地址
	ItemUrl string `json:"item_url,omitempty"`
	// ClickUrl x	淘客地址
	ClickUrl string `json:"click_url,omitempty"`
	// Nick 	卖家昵称
	Nick string `json:"nick,omitempty"`
	// SellerId 	卖家id
	SellerId int `json:"seller_id,omitempty"`
	// Volume 30天销量
	Volume int `json:"volume,omitempty"`
	// TkRate 	收入比例，举例，取值为20.00，表示比例20.00%
	TkRate string `json:"tk_rate,omitempty"`
	// ZkFinalPriceWap 无线折扣价，即宝贝在无线上的实际售卖价格。
	ZkFinalPriceWap string `json:"zk_final_price_wap,omitempty"`
	// ShopTitle 店铺名称
	ShopTitle string `json:"shop_title,omitempty"`
	// EventStartTime 招商活动开始时间；如果该宝贝取自普通选品组，则取值为1970-01-01 00:00:00；
	EventStartTime string `json:"event_start_time,omitempty"`
	// EventEndTime 招行活动的结束时间；如果该宝贝取自普通的选品组，则取值为1970-01-01 00:00:00
	EventEndTime string `json:"event_end_time,omitempty"`
	// Type 宝贝类型：1 普通商品； 2 鹊桥高佣金商品；3 定向招商商品；4 营销计划商品;
	Type int `json:"type,omitempty"`
	// Status 宝贝状态，0失效，1有效；注：失效可能是宝贝已经下线或者是被处罚不能在进行推广
	Status int `json:"status,omitempty"`
	// Category 后台一级类目
	Category int `json:"category,omitempty"`
	// CouponClickUrl 商品优惠券推广链接
	CouponClickUrl string `json:"coupon_click_url,omitempty"`
	// CouponEndTime 优惠券结束时间
	CouponEndTime string `json:"coupon_end_time,omitempty"`
	// CouponInfo 优惠券面额
	CouponInfo string `json:"coupon_info,omitempty"`
	// CouponStartTime 	优惠券开始时间
	CouponStartTime string `json:"coupon_start_time,omitempty"`
	// CouponTotalCount 优惠券总量
	CouponTotalCount int `json:"coupon_total_count,omitempty"`
	// CouponRemainCount 优惠券剩余量
	CouponRemainCount int `json:"coupon_remain_count,omitempty"`
}
