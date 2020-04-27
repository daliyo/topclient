package responses

type TaobaoTbkDgMaterialOptionalResponse struct {
	ErrorResponse

	// TotalResults 搜索到符合条件的结果总数
	TotalResults int                                       `json:"total_results,omitempty"`
	ResultList   []TaobaoTbkDgMaterialOptionalResponseItem `json:"result_list,omitempty"`
}

type TaobaoTbkDgMaterialOptionalResponseItem struct {
	// CouponStartTime 优惠券信息-优惠券开始时间 2017-10-29
	CouponStartTime string `json:"coupon_start_time,omitempty"`
	// CouponEndTime 优惠券信息-优惠券结束时间 2017-10-29
	CouponEndTime string `json:"coupon_end_time,omitempty"`
	// InfoDxjh 商品信息-定向计划信息 	{"19013551":"2850","74510538":"2550"}
	InfoDxjh string `json:"info_dxjh,omitempty"`
	// TkTotalSales 商品信息-淘客30天推广量
	TkTotalSales string `json:"tk_total_sales,omitempty"`
	// CouponID 优惠券信息-优惠券id
	CouponID string `json:"coupon_id,omitempty"`
	// Title 商品信息-商品标题
	Title string `json:"title,omitempty"`
	// PictUrl 商品信息-商品主图
	PictUrl string `json:"pict_url,omitempty"`
	// SmallImages 商品信息-商品小图列表
	SmallImages []string `json:"small_images,omitempty"`
	// ReservePrice 商品信息-商品一口价格 102.00
	ReservePrice string `json:"reserve_price,omitempty"`
	// ZkFinalPrice 折扣价（元） 若属于预售商品，付定金时间内，折扣价=预售价 88.00
	ZkFinalPrice string `json:"zk_final_price,omitempty"`
	// UserType 店铺信息-卖家类型。0表示集市，1表示天猫
	UserType int `json:"user_type,omitempty"`
	// Provcity 商品信息-宝贝所在地
	Provcity string `json:"provcity,omitempty"`
	// ItemUrl 链接-宝贝地址
	ItemUrl string `json:"item_url,omitempty"`
	// IncludeMkt 商品信息-是否包含营销计划
	IncludeMkt string `json:"include_mkt,omitempty"`
	// IncludeDxjh	商品信息-是否包含定向计划
	IncludeDxjh string `json:"include_dxjh,omitempty"`
	// CommissionRate 商品信息-佣金比率。1550表示15.5%
	CommissionRate string `json:"commission_rate,omitempty"`
	// Volume 	商品信息-30天销量
	Volume int `json:"volume,omitempty"`
	// SellerID 店铺信息-卖家id
	SellerID int `json:"seller_id,omitempty"`
	// CouponTotalCount 优惠券信息-优惠券总量
	CouponTotalCount int `json:"coupon_total_count,omitempty"`
	// CouponRemainCount 优惠券信息-优惠券剩余量
	CouponRemainCount int `json:"coupon_remain_count,omitempty"`
	// CouponInfo 优惠券信息-优惠券满减信息
	CouponInfo string `json:"coupon_info,omitempty"`
	// CommissionType 商品信息-佣金类型。MKT表示营销计划，SP表示定向计划，COMMON表示通用计划
	CommissionType string `json:"commission_type,omitempty"`
	// ShopTitle	店铺信息-店铺名称
	ShopTitle string `json:"shop_title,omitempty"`
	// ShopDsr	店铺信息-店铺dsr评分
	ShopDsr int `json:"shop_dsr,omitempty"`
	// CouponShareUrl	链接-宝贝+券二合一页面链接
	CouponShareUrl string `json:"coupon_share_url,omitempty"`
	// Url	链接-宝贝推广链接
	Url string `json:"url,omitempty"`
	// LevelOneCategoryName 商品信息-一级类目名称
	LevelOneCategoryName string `json:"level_one_category_name,omitempty"`
	// LevelOneCategoryID 	商品信息-一级类目ID
	LevelOneCategoryID int `json:"level_one_category_id,omitempty"`
	// CategoryName 商品信息-叶子类目名称
	CategoryName string `json:"category_name,omitempty"`
	// CategoryID 	商品信息-叶子类目id
	CategoryID int `json:"category_id,omitempty"`
	// ShortTitle	商品信息-商品短标题
	ShortTitle string `json:"short_title,omitempty"`
	// WhiteImage 商品信息-商品白底图
	WhiteImage string `json:"white_image,omitempty"`
	// Oetime 拼团专用-拼团结束时间
	Oetime string `json:"oetime,omitempty"`
	// Ostime 	拼团专用-拼团开始时间
	Ostime string `json:"ostime,omitempty"`
	// JddNum 	拼团专用-拼团几人团
	JddNum int `json:"jdd_num,omitempty"`
	// JddPrice 	拼团专用-拼团拼成价，单位元
	JddPrice string `json:"jdd_price,omitempty"`
	// UvSumPreSale 	预售专用-预售数量
	UvSumPreSale int `json:"uv_sum_pre_sale,omitempty"`
	// XId 链接-物料块id(测试中请勿使用)
	XId string `json:"x_id,omitempty"`
	// CouponStartFee 优惠券信息-优惠券起用门槛，满X元可用。如：满299元减20元
	CouponStartFee string `json:"coupon_start_fee,omitempty"`
	// CouponAmount	优惠券（元） 若属于预售商品，该优惠券付尾款可用，付定金不可用
	CouponAmount string `json:"coupon_amount,omitempty"`
	// ItemDescription 商品信息-宝贝描述(推荐理由)
	ItemDescription string `json:"item_description,omitempty"`
	// Nick 店铺信息-卖家昵称
	Nick string `json:"nick,omitempty"`
	// OrigPrice 拼团专用-拼团一人价（原价)，单位元
	OrigPrice string `json:"orig_price,omitempty"`
	// TotalStock 拼团专用-拼团库存数量
	TotalStock int `json:"total_stock,omitempty"`
	// SellNum	拼团专用-拼团已售数量
	SellNum int `json:"sell_num,omitempty"`
	// Stock 拼团专用-拼团剩余库存
	Stock int `json:"stock,omitempty"`
	// TmallPlayActivityInfo 营销-天猫营销玩法
	TmallPlayActivityInfo string `json:"tmall_play_activity_info,omitempty"`
	// ItemID 商品信息-宝贝id
	ItemID int `json:"item_id,omitempty"`
	// RealPostFee 商品邮费
	RealPostFee string `json:"real_post_fee,omitempty"`
	// LockRate 锁住的佣金率
	LockRate string `json:"lock_rate,omitempty"`
	// LockRateEndTime 锁佣结束时间
	LockRateEndTime int `json:"lock_rate_end_time,omitempty"`
	// LockRateStartTime 锁佣开始时间
	LockRateStartTime int `json:"lock_rate_start_time,omitempty"`
	// PresaleDiscountFeeText 预售商品-优惠
	PresaleDiscountFeeText string `json:"presale_discount_fee_text,omitempty"`
	// PresaleTailEndTime 预售商品-付尾款结束时间（毫秒）
	PresaleTailEndTime int `json:"presale_tail_end_time,omitempty"`
	// PresaleTailStartTime 预售商品-付尾款开始时间（毫秒）
	PresaleTailStartTime int `json:"presale_tail_start_time,omitempty"`
	// PresaleEndTime 预售商品-付定金结束时间（毫秒）
	PresaleEndTime int `json:"presale_end_time,omitempty"`
	// PresaleStartTime 预售商品-付定金开始时间（毫秒）
	PresaleStartTime int `json:"presale_start_time,omitempty"`
	// PresaleDeposit 预售商品-定金（元）
	PresaleDeposit string `json:"presale_deposit,omitempty"`
	// YsylTljSendTime 预售有礼-淘礼金发放时间
	YsylTljSendTime string `json:"ysyl_tlj_send_time,omitempty"`
	// YsylCommissionRate 预售有礼-佣金比例 2030（表示20.3%）
	YsylCommissionRate string `json:"ysyl_commission_rate,omitempty"`
	// YsylTljFace 预售有礼-预估淘礼金（元）
	YsylTljFace string `json:"ysyl_tlj_face,omitempty"`
	// YsylClickUrl 预售有礼-推广链接
	YsylClickUrl string `json:"ysyl_click_url,omitempty"`
	// YsylTljUseEndTime 预售有礼-淘礼金使用结束时间
	YsylTljUseEndTime string `json:"ysyl_tlj_use_end_time,omitempty"`
	// YsylTljUseStartTime 预售有礼-淘礼金使用开始时间
	YsylTljUseStartTime string `json:"ysyl_tlj_use_start_time,omitempty"`
}
