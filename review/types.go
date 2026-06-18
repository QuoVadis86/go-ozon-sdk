package review

type ReviewV2ReviewListV2Response struct {
	LastID string `json:"last_id"`
	Reviews []interface{} `json:"reviews"`
	HasNext bool `json:"has_next"`
}

type V1ReviewInfoRequest struct {
	ReviewID string `json:"review_id"`
}

type V1QuestionCountResponse struct {
	All int64 `json:"all"`
	New int64 `json:"new"`
	Processed int64 `json:"processed"`
	Unprocessed int64 `json:"unprocessed"`
	Viewed int64 `json:"viewed"`
}

type V1QuestionInfoResponse struct {
	ProductURL string `json:"product_url"`
	QuestionLink string `json:"question_link"`
	SKU int64 `json:"sku"`
	Status interface{} `json:"status"`
	Text string `json:"text"`
	AnswersCount int64 `json:"answers_count"`
	AuthorName string `json:"author_name"`
	ID string `json:"id"`
	PublishedAt interface{} `json:"published_at"`
}

type V1QuestionAnswerListRequest struct {
	LastID interface{} `json:"last_id"`
	QuestionID string `json:"question_id"`
	SKU int64 `json:"sku"`
}

type V1CommentCreateResponse struct {
	CommentID string `json:"comment_id"`
}

type V1ReviewChangeStatusRequest struct {
	Status string `json:"status"`
	ReviewIds []interface{} `json:"review_ids"`
}

type ReviewV2ReviewListV2Request struct {
	Filters interface{} `json:"filters"`
	LastID string `json:"last_id"`
	Limit int32 `json:"limit"`
	SortDir interface{} `json:"sort_dir"`
}

type V1QuestionAnswerCreateResponse struct {
	AnswerID string `json:"answer_id"`
}

type V1QuestionListRequest struct {
	SortDir interface{} `json:"sort_dir"`
	Filter interface{} `json:"filter"`
	LastID string `json:"last_id"`
	Limit int64 `json:"limit"`
}

type V1CommentDeleteRequest struct {
	CommentID string `json:"comment_id"`
}

type V1QuestionAnswerCreateRequest struct {
	QuestionID string `json:"question_id"`
	SKU int64 `json:"sku"`
	Text string `json:"text"`
}

type V1CommentCreateRequest struct {
	MarkReviewAsProcessed bool `json:"mark_review_as_processed"`
	ParentCommentID string `json:"parent_comment_id"`
	ReviewID string `json:"review_id"`
	Text string `json:"text"`
}

type V1QuestionAnswerDeleteRequest struct {
	AnswerID string `json:"answer_id"`
	SKU int64 `json:"sku"`
}

type V1CommentListResponse struct {
	Comments []interface{} `json:"comments"`
	Offset int32 `json:"offset"`
}

type V1QuestionAnswerListResponse struct {
	Answers interface{} `json:"answers"`
	LastID string `json:"last_id"`
}

type V1QuestionInfoRequest struct {
	QuestionID string `json:"question_id"`
}

type V1ReviewListResponse struct {
	HasNext bool `json:"has_next"`
	LastID string `json:"last_id"`
	Reviews []interface{} `json:"reviews"`
}

type V1QuestionListResponse struct {
	Questions interface{} `json:"questions"`
	LastID string `json:"last_id"`
	HasNext bool `json:"has_next"`
}

type V1CommentListRequest struct {
	Limit int32 `json:"limit"`
	Offset int32 `json:"offset"`
	ReviewID string `json:"review_id"`
	SortDir interface{} `json:"sort_dir"`
}

type V1QuestionChangeStatusRequest struct {
	QuestionIds interface{} `json:"question_ids"`
	Status string `json:"status"`
}

type V1ReviewCountResponse struct {
	Processed int32 `json:"processed"`
	Total int32 `json:"total"`
	Unprocessed int32 `json:"unprocessed"`
}

type V1ReviewInfoResponse struct {
	Rating int32 `json:"rating"`
	SKU int64 `json:"sku"`
	Videos []interface{} `json:"videos"`
	CommentsAmount int32 `json:"comments_amount"`
	Photos []interface{} `json:"photos"`
	PublishedAt string `json:"published_at"`
	Text string `json:"text"`
	VideosAmount int32 `json:"videos_amount"`
	ID string `json:"id"`
	IsRatingParticipant bool `json:"is_rating_participant"`
	OrderStatus string `json:"order_status"`
	LikesAmount int32 `json:"likes_amount"`
	PhotosAmount int32 `json:"photos_amount"`
	Status string `json:"status"`
	DislikesAmount int32 `json:"dislikes_amount"`
}

type V1QuestionTopSkuRequest struct {
	Limit int64 `json:"limit"`
}

type V1QuestionTopSkuResponse struct {
	SKU interface{} `json:"sku"`
}

type V1ReviewListRequest struct {
	SortDir string `json:"sort_dir"`
	Status string `json:"status"`
	LastID string `json:"last_id"`
	Limit int32 `json:"limit"`
}
