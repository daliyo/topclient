package responses

type TaobaoTbkUatmFavoritesGetResponse struct {
	ErrorResponse

	Results      []TbkFavorite `json:"results,omitempty"`
	TotalResults int           `json:"total_results,omitempty"`
}

type TbkFavorite struct {
	// Type 选品库类型，1：普通类型，2高佣金类型
	Type int `json:"type,omitempty"`
	// FavoritesID 选品组id
	FavoritesID int `json:"favorites_id,omitempty"`
	// FavoritesTitle 选品组名称
	FavoritesTitle string `json:"favorites_title,omitempty"`
}
