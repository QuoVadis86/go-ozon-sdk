package review

type V1ReviewListResponse struct {
	HasNext bool `json:"has_next"` // `true`：回复中未返回所有评价。
	LastID string `json:"last_id"` // 页面中最后一个评价的标识符。
	Reviews []interface{} `json:"reviews"` // 评价信息。
}

type V1QuestionInfoResponse struct {
	AnswersCount int64 `json:"answers_count"` // 问题的回答数量。
	ID string `json:"id"` // 问题标识符。
	ProductURL string `json:"product_url"` // 商品链接。
	QuestionLink string `json:"question_link"` // 问题链接。
	SKU int64 `json:"sku"` // Ozon 系统中的商品标识符——SKU。
	Status interface{} `json:"status"` // 问题状态： - `NEW`——新的， - `ALL`——全部问题， - `VIEWED`——已查看， - `PROCESSED`——已处理， - `UNPROCESSED`——未处理。
	AuthorName string `json:"author_name"` // 问题作者。
	PublishedAt interface{} `json:"published_at"` // 问题发布日期。
	Text string `json:"text"` // 问题文本。
}

type V1CommentCreateRequest struct {
	ReviewID string `json:"review_id"` // 评价标识符。
	Text string `json:"text"` // 评论内容。
	MarkReviewAsProcessed bool `json:"mark_review_as_processed"` // 更新评论状态： - `true` — 状态将变更为 `Processed`（已处理）； - `false` — 状态不变。
	ParentCommentID string `json:"parent_comment_id"` // 父级评论的标识符（您要回复的评论）。
}

type V1ReviewListRequest struct {
	Status string `json:"status"` // 评价状态： - `ALL` — 全部， - `UNPROCESSED` — 未处理的， - `PROCESSED` — 已处理的。
	LastID string `json:"last_id"` // 页面中最后一个评价的标识符。
	Limit int32 `json:"limit"` // 限制回复中的值数量。最少 — 20；最多 — 100。
	SortDir string `json:"sort_dir"` // 排序方向： - `ASC` — 按升序。 - `DESC` — 按降序。
}

type V1ReviewChangeStatusRequest struct {
	ReviewIds []interface{} `json:"review_ids"` // 包含评价标识符的数组（数量在1到100之间）。
	Status string `json:"status"` // 评价状态： - `PROCESSED` — 已处理。 - `UNPROCESSED` — 未处理。
}

type V1CommentCreateResponse struct {
	CommentID string `json:"comment_id"` // 评论标识符。
}

type V1CommentListResponse struct {
	Offset int32 `json:"offset"` // 搜索结果中的元素数量。
	Comments []interface{} `json:"comments"` // 评论信息。
}

type V1QuestionTopSkuResponse struct {
	SKU interface{} `json:"sku"` // Ozon 系统中的商品标识符（SKU）列表。
}

type ReviewV2ReviewListV2Response struct {
	HasNext bool `json:"has_next"` // `true`，表示响应中未返回全部评价。
	LastID string `json:"last_id"` // 页面中最后一个评价的标识符。
	Reviews []interface{} `json:"reviews"` // 评价列表。
}

type V1QuestionAnswerListRequest struct {
	LastID interface{} `json:"last_id"` // 页面上最后一个值的标识符。 如果是首次请求，请将该字段留空。 后续请求中，请传入上一次请求返回的 `last_id`。
	QuestionID string `json:"question_id"` // 问题标识符。
	SKU int64 `json:"sku"` // Ozon 系统中的商品标识符——SKU。
}

type V1QuestionListRequest struct {
	Filter interface{} `json:"filter"`
	LastID string `json:"last_id"` // 页面上最后一个值的ID。运行第一个查询时，将此字段留空。 要检索以下数值，请从上一个查询的响应中指定`last_id`。
	Limit int64 `json:"limit"` // 响应中返回的值数量。
	SortDir interface{} `json:"sort_dir"`
}

type V1ReviewInfoResponse struct {
	Text string `json:"text"` // 评价文字。
	Videos []interface{} `json:"videos"` // 视频信息。
	Rating int32 `json:"rating"` // 评价评分。
	Status string `json:"status"` // 评价状态： - `UNPROCESSED` — 未处理， - `PROCESSED` — 已处理。
	VideosAmount int32 `json:"videos_amount"` // 评价中的视频数量。
	LikesAmount int32 `json:"likes_amount"` // 评价的点赞数量。
	CommentsAmount int32 `json:"comments_amount"` // 评价的回复数量。
	DislikesAmount int32 `json:"dislikes_amount"` // 评价的踩数量。
	ID string `json:"id"` // 评价标识符。
	IsRatingParticipant bool `json:"is_rating_participant"` // `true`：评论是由官方人员留下的；`false`：评论是由买家留下的。
	Photos []interface{} `json:"photos"` // 图片信息。
	PhotosAmount int32 `json:"photos_amount"` // 评价中的图片数量。
	PublishedAt string `json:"published_at"` // 评价的发布日期。
	SKU int64 `json:"sku"` // Ozon系统中的商品识别符——SKU。
	OrderStatus string `json:"order_status"` // 买家留下评价的订单状态： - `DELIVERED`— 已送达， - `CANCELLED` — 已取消。
}

type V1QuestionAnswerDeleteRequest struct {
	AnswerID string `json:"answer_id"` // 回答标识符。
	SKU int64 `json:"sku"` // Ozon 系统中的商品标识符——SKU。
}

type V1QuestionAnswerCreateResponse struct {
	AnswerID string `json:"answer_id"` // 问题回答标识符。
}

type V1QuestionTopSkuRequest struct {
	Limit int64 `json:"limit"` // 响应结果数量。
}

type V1ReviewCountResponse struct {
	Unprocessed int32 `json:"unprocessed"` // 未处理评价的数量。
	Processed int32 `json:"processed"` // 已处理评价的数量。
	Total int32 `json:"total"` // 评价的总数量。
}

type V1QuestionInfoRequest struct {
	QuestionID string `json:"question_id"` // 问题标识符。
}

type V1ReviewInfoRequest struct {
	ReviewID string `json:"review_id"` // 评价标识符。
}

type V1CommentListRequest struct {
	Limit int32 `json:"limit"` // 限制回复中的值数量。 最少 — 20；最多 — 100。
	Offset int32 `json:"offset"` // 从列表开头跳过的元素数量：例如，如果 `offset = 10`，那么回复将从找到的第11个元素开始。
	ReviewID string `json:"review_id"` // 评价标识符。
	SortDir interface{} `json:"sort_dir"`
}

type V1QuestionChangeStatusRequest struct {
	QuestionIds interface{} `json:"question_ids"` // 问题标识符。
	Status string `json:"status"` // 问题状态： - `NEW`——新的， - `VIEWED`——已查看， - `PROCESSED`——已处理。
}

type V1QuestionListResponse struct {
	Questions interface{} `json:"questions"` // 问题。
	LastID string `json:"last_id"` // 页面上最后一个值的标识符。 要获取下一个批次的数据，请在下一个请求的 `last_id` 参数中传递上次获取的值。
	HasNext bool `json:"has_next"` // 如果响应中未返回所有问题，则为`true`。
}

type ReviewV2ReviewListV2Request struct {
	SortDir interface{} `json:"sort_dir"`
	Filters interface{} `json:"filters"`
	LastID string `json:"last_id"` // 响应中最后一条评价的标识符。
	Limit int32 `json:"limit"` // 响应中的评价数量。
}

type V1QuestionAnswerCreateRequest struct {
	QuestionID string `json:"question_id"` // 问题标识符。
	SKU int64 `json:"sku"` // Ozon 系统中的商品标识符——SKU。
	Text string `json:"text"` // 回答文本，长度为 2 至 3000 个字符。
}

type V1QuestionCountResponse struct {
	All int64 `json:"all"` // 问题总数。
	New int64 `json:"new"` // 新问题数量。
	Processed int64 `json:"processed"` // 已处理问题数量。
	Unprocessed int64 `json:"unprocessed"` // 未处理问题数量。
	Viewed int64 `json:"viewed"` // 已查看问题数量。
}

type V1QuestionAnswerListResponse struct {
	Answers interface{} `json:"answers"` // 回答。
	LastID string `json:"last_id"` // 页面上最后一个值的标识符。 要获取下一个批次的数据，请在下一个请求的 `last_id` 参数中传递上次获取的值。
}

type V1CommentDeleteRequest struct {
	CommentID string `json:"comment_id"` // 评论标识符。
}
