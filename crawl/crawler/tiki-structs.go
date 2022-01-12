package crawler

type ProductListResponse struct {
	ProductShortList []ProductShort `json:"data"`
}
type BadgesNew struct {
	Placement  string `json:"placement"`
	Type       string `json:"type"`
	Code       string `json:"code"`
	Icon       string `json:"icon"`
	IconWidth  int    `json:"icon_width"`
	IconHeight int    `json:"icon_height"`
}
type Inventory struct {
	FulfillmentType string `json:"fulfillment_type"`
}
type StockItem struct {
	Qty          int  `json:"qty"`
	MinSaleQty   int  `json:"min_sale_qty"`
	MaxSaleQty   int  `json:"max_sale_qty"`
	PreorderDate bool `json:"preorder_date"`
}
type QuantitySold struct {
	Text  string `json:"text"`
	Value int    `json:"value"`
}
type Badges struct {
	Code string `json:"code"`
	Text string `json:"text"`
}
type OptionColor struct {
	Spid           int    `json:"spid"`
	DisplayName    string `json:"display_name"`
	IsDefault      int    `json:"is_default"`
	Price          int    `json:"price"`
	ListPrice      int    `json:"list_price"`
	OriginalPrice  int    `json:"original_price"`
	Thumbnail      string `json:"thumbnail"`
	SmallThumbnail string `json:"small_thumbnail"`
	URLPath        string `json:"url_path"`
}
type ProductShort struct {
	ID               int32         `json:"id"`
	Sku              string        `json:"sku"`
	Name             string        `json:"name"`
	URLKey           string        `json:"url_key"`
	URLPath          string        `json:"url_path"`
	Type             string        `json:"type"`
	BrandName        string        `json:"brand_name"`
	ShortDescription string        `json:"short_description"`
	Price            int           `json:"price"`
	ListPrice        int           `json:"list_price"`
	OriginalPrice    int           `json:"original_price"`
	BadgesNew        []BadgesNew   `json:"badges_new,omitempty"`
	Discount         int           `json:"discount"`
	DiscountRate     int           `json:"discount_rate"`
	RatingAverage    float64       `json:"rating_average"`
	ReviewCount      int           `json:"review_count"`
	OrderCount       int           `json:"order_count"`
	FavouriteCount   int           `json:"favourite_count"`
	ThumbnailURL     string        `json:"thumbnail_url"`
	ThumbnailWidth   int           `json:"thumbnail_width"`
	ThumbnailHeight  int           `json:"thumbnail_height"`
	HasEbook         bool          `json:"has_ebook"`
	InventoryStatus  string        `json:"inventory_status"`
	ProductsetID     int           `json:"productset_id"`
	IsFlower         bool          `json:"is_flower"`
	IsGiftCard       bool          `json:"is_gift_card"`
	Inventory        Inventory     `json:"inventory"`
	StockItem        StockItem     `json:"stock_item"`
	SalableType      string        `json:"salable_type"`
	SellerProductID  int           `json:"seller_product_id"`
	QuantitySold     QuantitySold  `json:"quantity_sold,omitempty"`
	Shippable        bool          `json:"shippable"`
	Badges           []Badges      `json:"badges,omitempty"`
	BundleDeal       bool          `json:"bundle_deal,omitempty"`
	VideoURL         string        `json:"video_url,omitempty"`
	OptionColor      []OptionColor `json:"option_color,omitempty"`
}
type Paging struct {
	Total       int `json:"total"`
	PerPage     int `json:"per_page"`
	CurrentPage int `json:"current_page"`
	LastPage    int `json:"last_page"`
	From        int `json:"from"`
	To          int `json:"to"`
}
type SortOptions struct {
	DisplayValue string `json:"display_value"`
	QueryValue   string `json:"query_value"`
	Selected     bool   `json:"selected"`
}
type Values struct {
	DisplayValue string `json:"display_value"`
	Count        int    `json:"count"`
	QueryValue   int    `json:"query_value"`
	URLKey       string `json:"url_key"`
}
type Filters struct {
	QueryName   string   `json:"query_name"`
	DisplayName string   `json:"display_name"`
	Values      []Values `json:"values"`
	Collapsed   int      `json:"collapsed"`
	MultiSelect bool     `json:"multi_select"`
	Type        string   `json:"type,omitempty"`
	Icon        string   `json:"icon,omitempty"`
	IconWidth   int      `json:"icon_width,omitempty"`
	IconHeight  int      `json:"icon_height,omitempty"`
	Sticky      bool     `json:"sticky,omitempty"`
	Min         int      `json:"min,omitempty"`
	Max         int64    `json:"max,omitempty"`
}
type ExperimentVariants struct {
}
type Amplitude struct {
	NumberOfSponsoredProductAd  string `json:"number_of_sponsored_product_ad"`
	NumberOfDisplayAd           string `json:"number_of_display_ad"`
	NumberOfBrandStoreAd        string `json:"number_of_brand_store_ad"`
	DetectedBrandName           string `json:"detected_brand_name"`
	DetectedBrandScore          string `json:"detected_brand_score"`
	DetectedPrimaryCategoryName string `json:"detected_primary_category_name"`
	DetectedL1CategoryName      string `json:"detected_l1_category_name"`
	SelectedSortOption          string `json:"selected_sort_option"`
	SellerProductSku            string `json:"seller_product_sku"`
}

// ProductDetail is the struct for product detail
type ProductDetail struct {
	ID                      int                    `json:"id"`
	MasterID                int                    `json:"master_id"`
	Sku                     string                 `json:"sku"`
	Name                    string                 `json:"name"`
	URLKey                  string                 `json:"url_key"`
	URLPath                 string                 `json:"url_path"`
	Type                    string                 `json:"type"`
	BookCover               interface{}            `json:"book_cover"`
	ShortDescription        string                 `json:"short_description"`
	Price                   int                    `json:"price"`
	ListPrice               int                    `json:"list_price"`
	OriginalPrice           int                    `json:"original_price"`
	Discount                int                    `json:"discount"`
	DiscountRate            int                    `json:"discount_rate"`
	RatingAverage           float64                `json:"rating_average"`
	ReviewCount             int                    `json:"review_count"`
	ReviewText              string                 `json:"review_text"`
	FavouriteCount          int                    `json:"favourite_count"`
	ThumbnailURL            string                 `json:"thumbnail_url"`
	HasEbook                bool                   `json:"has_ebook"`
	InventoryStatus         string                 `json:"inventory_status"`
	InventoryType           string                 `json:"inventory_type"`
	ProductsetGroupName     string                 `json:"productset_group_name"`
	IsFresh                 bool                   `json:"is_fresh"`
	Seller                  interface{}            `json:"seller"`
	IsFlower                bool                   `json:"is_flower"`
	HasBuynow               bool                   `json:"has_buynow"`
	IsGiftCard              bool                   `json:"is_gift_card"`
	SalableType             interface{}            `json:"salable_type"`
	DataVersion             int                    `json:"data_version"`
	DayAgoCreated           int                    `json:"day_ago_created"`
	AllTimeQuantitySold     int                    `json:"all_time_quantity_sold"`
	MetaTitle               string                 `json:"meta_title"`
	MetaDescription         string                 `json:"meta_description"`
	MetaKeywords            string                 `json:"meta_keywords"`
	IsBabyMilk              bool                   `json:"is_baby_milk"`
	IsAcoholicDrink         bool                   `json:"is_acoholic_drink"`
	Description             string                 `json:"description"`
	Images                  []Images               `json:"images"`
	WarrantyPolicy          interface{}            `json:"warranty_policy"`
	Brand                   Brand                  `json:"brand"`
	CurrentSeller           CurrentSeller          `json:"current_seller"`
	OtherSellers            []interface{}          `json:"other_sellers"`
	SellerSpecifications    []SellerSpecifications `json:"seller_specifications"`
	Specifications          []Specifications       `json:"specifications"`
	ProductLinks            []interface{}          `json:"product_links"`
	ServicesAndPromotions   []interface{}          `json:"services_and_promotions"`
	Promitions              []interface{}          `json:"promitions"`
	StockItem               StockItem              `json:"stock_item"`
	QuantitySold            QuantitySold           `json:"quantity_sold"`
	Categories              Categories             `json:"categories"`
	Breadcrumbs             []Breadcrumbs          `json:"breadcrumbs"`
	InstallmentInfo         interface{}            `json:"installment_info"`
	IsSellerInChatWhitelist bool                   `json:"is_seller_in_chat_whitelist"`
	Inventory               Inventory              `json:"inventory"`
	WarrantyInfo            []interface{}          `json:"warranty_info"`
	ReturnAndExchangePolicy string                 `json:"return_and_exchange_policy"`
	DealSpecs               DealSpecs              `json:"deal_specs"`
	BestPriceGuaranteed     bool                   `json:"best_price_guaranteed"`
	Benefits                []Benefits             `json:"benefits"`
}
type Images struct {
	BaseURL      string      `json:"base_url"`
	IsGallery    bool        `json:"is_gallery"`
	Label        interface{} `json:"label"`
	LargeURL     string      `json:"large_url"`
	MediumURL    string      `json:"medium_url"`
	Position     interface{} `json:"position"`
	SmallURL     string      `json:"small_url"`
	ThumbnailURL string      `json:"thumbnail_url"`
}
type Brand struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}
type CurrentSeller struct {
	ID                            int         `json:"id"`
	Sku                           string      `json:"sku"`
	Name                          string      `json:"name"`
	Link                          string      `json:"link"`
	Logo                          string      `json:"logo"`
	Price                         int         `json:"price"`
	ProductID                     string      `json:"product_id"`
	StoreID                       int         `json:"store_id"`
	IsBestStore                   bool        `json:"is_best_store"`
	IsOfflineInstallmentSupported interface{} `json:"is_offline_installment_supported"`
}
type SellerSpecifications struct {
	Name  string      `json:"name"`
	URL   interface{} `json:"url"`
	Value interface{} `json:"value"`
}
type Attributes struct {
	Code  string `json:"code"`
	Name  string `json:"name"`
	Value string `json:"value"`
}
type Specifications struct {
	Name       string       `json:"name"`
	Attributes []Attributes `json:"attributes"`
}
type Categories struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	IsLeaf bool   `json:"is_leaf"`
}
type Breadcrumbs struct {
	URL        string `json:"url"`
	Name       string `json:"name"`
	CategoryID int    `json:"category_id"`
}
type DealSpecs struct {
	DealID        int    `json:"deal_id"`
	IsHotDeal     bool   `json:"is_hot_deal"`
	Qty           int    `json:"qty"`
	MaxSaleQty    int    `json:"max_sale_qty"`
	QtyOrdered    int    `json:"qty_ordered"`
	SpecialToDate int    `json:"special_to_date"`
	Price         int    `json:"price"`
	ListPrice     int    `json:"list_price"`
	OriginalPrice int    `json:"original_price"`
	DiscountRate  int    `json:"discount_rate"`
	Progress      int    `json:"progress"`
	ProgressText  string `json:"progress_text"`
	V5            int    `json:"v5"`
}
type Extra struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
type Benefits struct {
	Icon        string  `json:"icon"`
	Text        string  `json:"text"`
	ExtraText   string  `json:"extra_text,omitempty"`
	ExtraHeader string  `json:"extra_header,omitempty"`
	Extra       []Extra `json:"extra,omitempty"`
}
