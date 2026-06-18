package review

type V1QuestionTopSkuRequest struct {
	Limit int64 `json:"limit"` // 响应结果数量。
}

type V1ReviewCountResponse struct {
	Total       int32 `json:"total"`       // 评价的总数量。
	Unprocessed int32 `json:"unprocessed"` // 未处理评价的数量。
	Processed   int32 `json:"processed"`   // 已处理评价的数量。
}

type V1ReviewInfoRequest struct {
	ReviewID string `json:"review_id"` // 评价标识符。
}

type InfoResponseVideo struct {
	URL                  string `json:"url"`                     // 视频链接。
	Width                int64  `json:"width"`                   // 宽度。
	Height               int64  `json:"height"`                  // 高度。
	PreviewURL           string `json:"preview_url"`             // 预览视频链接。
	ShortVideoPreviewURL string `json:"short_video_preview_url"` // 短视频链接。
}

// Status values
type Status string

const (
	StatusNEW       Status = "NEW"       // 新的，
	StatusViewed    Status = "VIEWED"    // 已查看，
	StatusProcessed Status = "PROCESSED" // 已处理。
)

type V1QuestionChangeStatusRequest struct {
	QuestionIds []string `json:"question_ids"` // 问题标识符。
	Status      Status   `json:"status"`       // 问题状态： - `NEW`——新的， - `VIEWED`——已查看， - `PROCESSED`——已处理。
}

// 回答发布状态： - `PUBLISHED`——已发布； - `AWAITING_MODERATION`——等待审核； - `MODERATION_FAILED`——未通过审核； - `DUPLICATE`——重复。
type QuestionV1QuestionAnswerListResponseAnswerStatusPublicationEnum string

type V1QuestionAnswerListResponseAnswers struct {
	PublishedAt       any                                                             `json:"published_at"` // 回答发布时间。
	QuestionID        string                                                          `json:"question_id"`  // 问题标识符。
	SKU               int64                                                           `json:"sku"`          // Ozon 系统中的商品标识符——SKU。
	StatusPublication QuestionV1QuestionAnswerListResponseAnswerStatusPublicationEnum `json:"status_publication"`
	Text              string                                                          `json:"text"`        // 回答文本。
	AuthorName        string                                                          `json:"author_name"` // 回答作者。
	ID                string                                                          `json:"id"`          // 回答标识符。
}

// 评价状态： - `NEW`——新评价； - `VIEWED`——已查看； - `PROCESSED`——已处理。
type V2ReviewListV2ResponseReviewStatusEnum string

type V1QuestionCountResponse struct {
	Processed   int64 `json:"processed"`   // 已处理问题数量。
	Unprocessed int64 `json:"unprocessed"` // 未处理问题数量。
	Viewed      int64 `json:"viewed"`      // 已查看问题数量。
	All         int64 `json:"all"`         // 问题总数。
	New         int64 `json:"new"`         // 新问题数量。
}

// 筛选器。
type V1QuestionListRequestFilter struct {
	DateFrom string `json:"date_from"` // 时间段开始。
	DateTo   string `json:"date_to"`   // 时间段结束。
	Status   Status `json:"status"`    // 问题状态： - `NEW`——新的， - `ALL`——全部问题， - `VIEWED`——已查看， - `PROCESSED`——已处理， - `UNPROCESSED`——未处理。
}

// 排序方向： - `DESC`——降序； - `ASC`——升序。
type QuestionV1GetQuestionListRequestSortDirEnum string

type V1QuestionListRequest struct {
	Filter  V1QuestionListRequestFilter                 `json:"filter"`
	LastID  string                                      `json:"last_id"` // 页面上最后一个值的ID。运行第一个查询时，将此字段留空。 要检索以下数值，请从上一个查询的响应中指定`last_id`。
	Limit   int64                                       `json:"limit"`   // 响应中返回的值数量。
	SortDir QuestionV1GetQuestionListRequestSortDirEnum `json:"sort_dir"`
}

type V1CommentDeleteRequest struct {
	CommentID string `json:"comment_id"` // 评论标识符。
}

type V1QuestionAnswerDeleteRequest struct {
	AnswerID string `json:"answer_id"` // 回答标识符。
	SKU      int64  `json:"sku"`       // Ozon 系统中的商品标识符——SKU。
}

type V1QuestionInfoRequest struct {
	QuestionID string `json:"question_id"` // 问题标识符。
}

type V1QuestionInfoResponse struct {
	AnswersCount int64  `json:"answers_count"` // 问题的回答数量。
	ID           string `json:"id"`            // 问题标识符。
	ProductURL   string `json:"product_url"`   // 商品链接。
	QuestionLink string `json:"question_link"` // 问题链接。
	SKU          int64  `json:"sku"`           // Ozon 系统中的商品标识符——SKU。
	AuthorName   string `json:"author_name"`   // 问题作者。
	PublishedAt  any    `json:"published_at"`  // 问题发布日期。
	Status       any    `json:"status"`        // 问题状态： - `NEW`——新的， - `ALL`——全部问题， - `VIEWED`——已查看， - `PROCESSED`——已处理， - `UNPROCESSED`——未处理。
	Text         string `json:"text"`          // 问题文本。
}

type V1CommentCreateRequest struct {
	ParentCommentID       string `json:"parent_comment_id"`        // 父级评论的标识符（您要回复的评论）。
	ReviewID              string `json:"review_id"`                // 评价标识符。
	Text                  string `json:"text"`                     // 评论内容。
	MarkReviewAsProcessed bool   `json:"mark_review_as_processed"` // 更新评论状态： - `true` — 状态将变更为 `Processed`（已处理）； - `false` — 状态不变。
}

type InfoResponsePhoto struct {
	Height int32  `json:"height"` // 高度。
	URL    string `json:"url"`    // 图片链接。
	Width  int32  `json:"width"`  // 宽度。
}

// OrderStatus values
type OrderStatus string

const (
	OrderStatusDelivered OrderStatus = "DELIVERED" // 已送达，
	OrderStatusCancelled OrderStatus = "CANCELLED" // 已取消。
)

type V1ReviewInfoResponse struct {
	CommentsAmount      int32               `json:"comments_amount"`       // 评价的回复数量。
	PhotosAmount        int32               `json:"photos_amount"`         // 评价中的图片数量。
	LikesAmount         int32               `json:"likes_amount"`          // 评价的点赞数量。
	Photos              []InfoResponsePhoto `json:"photos"`                // 图片信息。
	PublishedAt         string              `json:"published_at"`          // 评价的发布日期。
	Status              Status              `json:"status"`                // 评价状态： - `UNPROCESSED` — 未处理， - `PROCESSED` — 已处理。
	Text                string              `json:"text"`                  // 评价文字。
	DislikesAmount      int32               `json:"dislikes_amount"`       // 评价的踩数量。
	ID                  string              `json:"id"`                    // 评价标识符。
	OrderStatus         OrderStatus         `json:"order_status"`          // 买家留下评价的订单状态： - `DELIVERED`— 已送达， - `CANCELLED` — 已取消。
	SKU                 int64               `json:"sku"`                   // Ozon系统中的商品识别符——SKU。
	IsRatingParticipant bool                `json:"is_rating_participant"` // `true`：评论是由官方人员留下的；`false`：评论是由买家留下的。
	Rating              int32               `json:"rating"`                // 评价评分。
	Videos              []InfoResponseVideo `json:"videos"`                // 视频信息。
	VideosAmount        int32               `json:"videos_amount"`         // 评价中的视频数量。
}

// 评价状态： - `ALL`——全部； - `NEW`——新的； - `VIEWED`——已查看； - `PROCESSED`——已处理。
type V2ReviewListV2RequestFiltersStatusEnum string

// 买家留下评价的订单状态： - `ALL`——全部； - `DELIVERED`——已送达； - `CANCELLED`——已取消。
type V2ReviewListV2RequestFiltersOrderStatusEnum string

// 用于搜索评价的筛选条件。
type V2ReviewListV2RequestFilters struct {
	Skus          []string                                    `json:"skus"` // 商品在Ozon系统中的标识符——SKU。
	Status        V2ReviewListV2RequestFiltersStatusEnum      `json:"status"`
	OrderStatus   V2ReviewListV2RequestFiltersOrderStatusEnum `json:"order_status"`
	PublishedFrom string                                      `json:"published_from"` // 时间段开始日期。将返回在此日期之后创建的评价。
	PublishedTo   string                                      `json:"published_to"`   // 时间段结束日期。将返回在此日期之前创建的评价。
}

// 排序方向： - `ASC` — 按升序。 - `DESC` — 按降序。
type V1CommentSort string

type CommentListResponseComment struct {
	IsOwner         bool   `json:"is_owner"`          // `true`：评论是由卖家留下的；`false`：评论是由买家留下的。
	ParentCommentID string `json:"parent_comment_id"` // 父级评论的标识符（需要对此评论进行回复）。
	PublishedAt     string `json:"published_at"`      // 评论发布日期。
	Text            string `json:"text"`              // 评论内容。
	ID              string `json:"id"`                // 评论标识符。
	IsOfficial      bool   `json:"is_official"`       // `true`：评论是由官方人员留下的；`false`：评论是由买家留下的。
}

type V1ReviewChangeStatusRequest struct {
	ReviewIds []string `json:"review_ids"` // 包含评价标识符的数组（数量在1到100之间）。
	Status    Status   `json:"status"`     // 评价状态： - `PROCESSED` — 已处理。 - `UNPROCESSED` — 未处理。
}

type V1CommentListResponse struct {
	Offset   int32                        `json:"offset"`   // 搜索结果中的元素数量。
	Comments []CommentListResponseComment `json:"comments"` // 评论信息。
}

type V1QuestionAnswerListRequest struct {
	QuestionID string `json:"question_id"` // 问题标识符。
	SKU        int64  `json:"sku"`         // Ozon 系统中的商品标识符——SKU。
	LastID     any    `json:"last_id"`     // 页面上最后一个值的标识符。 如果是首次请求，请将该字段留空。 后续请求中，请传入上一次请求返回的 `last_id`。
}

// SortDir values
type SortDir string

const (
	SortDirASC  SortDir = "ASC"  // 按升序。
	SortDirDesc SortDir = "DESC" // 按降序。
)

type V1ReviewListRequest struct {
	LastID  string  `json:"last_id"`  // 页面中最后一个评价的标识符。
	Limit   int32   `json:"limit"`    // 限制回复中的值数量。最少 — 20；最多 — 100。
	SortDir SortDir `json:"sort_dir"` // 排序方向： - `ASC` — 按升序。 - `DESC` — 按降序。
	Status  Status  `json:"status"`   // 评价状态： - `ALL` — 全部， - `UNPROCESSED` — 未处理的， - `PROCESSED` — 已处理的。
}

type ListResponseReview struct {
	OrderStatus         OrderStatus `json:"order_status"`          // 买家留下评价的订单状态： - `DELIVERED`— 已送达， - `CANCELLED` — 已取消。
	Text                string      `json:"text"`                  // 评价文字。
	VideosAmount        int32       `json:"videos_amount"`         // 评价的视频数量。
	CommentsAmount      int32       `json:"comments_amount"`       // 评价的评论数量。
	ID                  string      `json:"id"`                    // 评价标识符。
	IsRatingParticipant bool        `json:"is_rating_participant"` // `true`：如果评价被计入评级计算。
	PhotosAmount        int32       `json:"photos_amount"`         // 评价中的图片数量。
	PublishedAt         string      `json:"published_at"`          // 评价的发布时间。
	Rating              int32       `json:"rating"`                // 评价的评分。
	SKU                 int64       `json:"sku"`                   // Ozon系统中的商品识别符——SKU。
	Status              Status      `json:"status"`                // 评价状态： - `UNPROCESSED` — 未处理， - `PROCESSED` — 已处理。
}

type V1QuestionTopSkuResponse struct {
	SKU []string `json:"sku"` // Ozon 系统中的商品标识符（SKU）列表。
}

// 排序方向： - `ASC`——升序； - `DESC`——降序。
type V2ReviewListV2RequestSortDirEnum string

type V2ReviewListV2Request struct {
	LastID  string                           `json:"last_id"` // 响应中最后一条评价的标识符。
	Limit   int32                            `json:"limit"`   // 响应中的评价数量。
	SortDir V2ReviewListV2RequestSortDirEnum `json:"sort_dir"`
	Filters V2ReviewListV2RequestFilters     `json:"filters"`
}

type V1QuestionAnswerListResponse struct {
	Answers []V1QuestionAnswerListResponseAnswers `json:"answers"` // 回答。
	LastID  string                                `json:"last_id"` // 页面上最后一个值的标识符。 要获取下一个批次的数据，请在下一个请求的 `last_id` 参数中传递上次获取的值。
}

// 买家留下评价的订单状态： - `DELIVERED`——已送达； - `CANCELLED`——已取消。
type V2ReviewListV2ResponseReviewOrderStatusEnum string

type V2ReviewListV2ResponseReview struct {
	OrderStatus         V2ReviewListV2ResponseReviewOrderStatusEnum `json:"order_status"`
	PhotosAmount        int32                                       `json:"photos_amount"` // 评价中的图片数量。
	Rating              int32                                       `json:"rating"`        // 评价评分。
	SKU                 int64                                       `json:"sku"`           // 商品在Ozon系统中的标识符——SKU。
	PublishedAt         string                                      `json:"published_at"`  // 评价发布日期。
	Status              V2ReviewListV2ResponseReviewStatusEnum      `json:"status"`
	Text                string                                      `json:"text"`                  // 评价文字。
	VideosAmount        int32                                       `json:"videos_amount"`         // 评价中的视频数量。
	CommentsAmount      int32                                       `json:"comments_amount"`       // 评价的评论数量。
	ID                  string                                      `json:"id"`                    // 评价标识符。
	IsRatingParticipant bool                                        `json:"is_rating_participant"` // `true`表示评价参与评分计算。
}

type V1ReviewListResponse struct {
	HasNext bool                 `json:"has_next"` // `true`：回复中未返回所有评价。
	LastID  string               `json:"last_id"`  // 页面中最后一个评价的标识符。
	Reviews []ListResponseReview `json:"reviews"`  // 评价信息。
}

type V1CommentCreateResponse struct {
	CommentID string `json:"comment_id"` // 评论标识符。
}

type V1QuestionListResponseQuestions struct {
	PublishedAt  any    `json:"published_at"`  // 问题发布日期。
	QuestionLink string `json:"question_link"` // 问题链接。
	Status       any    `json:"status"`        // 问题状态： - `NEW`——新的， - `ALL`——全部问题， - `VIEWED`——已查看， - `PROCESSED`——已处理， - `UNPROCESSED`——未处理。
	AnswersCount int64  `json:"answers_count"` // 问题的回答数量。
	AuthorName   string `json:"author_name"`   // 问题作者姓名。
	ID           string `json:"id"`            // 问题标识符。
	ProductURL   string `json:"product_url"`   // 商品链接。
	SKU          int64  `json:"sku"`           // Ozon 系统中的商品标识符——SKU。
	Text         string `json:"text"`          // 问题文本。
}

type V1QuestionListResponse struct {
	LastID    string                            `json:"last_id"`   // 页面上最后一个值的标识符。 要获取下一个批次的数据，请在下一个请求的 `last_id` 参数中传递上次获取的值。
	HasNext   bool                              `json:"has_next"`  // 如果响应中未返回所有问题，则为`true`。
	Questions []V1QuestionListResponseQuestions `json:"questions"` // 问题。
}

type V2ReviewListV2Response struct {
	HasNext bool                           `json:"has_next"` // `true`，表示响应中未返回全部评价。
	LastID  string                         `json:"last_id"`  // 页面中最后一个评价的标识符。
	Reviews []V2ReviewListV2ResponseReview `json:"reviews"`  // 评价列表。
}

type V1CommentListRequest struct {
	Offset   int32         `json:"offset"`    // 从列表开头跳过的元素数量：例如，如果 `offset = 10`，那么回复将从找到的第11个元素开始。
	ReviewID string        `json:"review_id"` // 评价标识符。
	SortDir  V1CommentSort `json:"sort_dir"`
	Limit    int32         `json:"limit"` // 限制回复中的值数量。 最少 — 20；最多 — 100。
}

type V1QuestionAnswerCreateRequest struct {
	QuestionID string `json:"question_id"` // 问题标识符。
	SKU        int64  `json:"sku"`         // Ozon 系统中的商品标识符——SKU。
	Text       string `json:"text"`        // 回答文本，长度为 2 至 3000 个字符。
}

type V1QuestionAnswerCreateResponse struct {
	AnswerID string `json:"answer_id"` // 问题回答标识符。
}
