package review

import (
	"context"
	"github.com/QuoVadis86/go-ozon-sdk/transport"
	"os"
	"testing"
)

var ctx = context.Background()

func skipNoCreds(t *testing.T) *transport.Client {
	t.Helper()
	if os.Getenv("OZON_CLIENT_ID") == "" || os.Getenv("OZON_API_KEY") == "" {
		t.Skip("set OZON_CLIENT_ID and OZON_API_KEY to run tests")
	}
	return transport.New(os.Getenv("OZON_CLIENT_ID"), os.Getenv("OZON_API_KEY"), nil)
}

func TestCreate(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.Create(ctx, &V1QuestionAnswerCreateRequest{})
	if err != nil {
		t.Fatalf("Create() error: %v", err)
	}
	_ = resp
}

func TestCommentCreate(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.CommentCreate(ctx, &V1CommentCreateRequest{})
	if err != nil {
		t.Fatalf("CommentCreate() error: %v", err)
	}
	_ = resp
}

func TestInfo(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.Info(ctx, &V1QuestionInfoRequest{})
	if err != nil {
		t.Fatalf("Info() error: %v", err)
	}
	_ = resp
}

func TestCommentDelete(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	err := svc.CommentDelete(ctx, &V1CommentDeleteRequest{})
	_ = err
}

func TestTopSku(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.TopSku(ctx, &V1QuestionTopSkuRequest{})
	if err != nil {
		t.Fatalf("TopSku() error: %v", err)
	}
	_ = resp
}

func TestChangeStatus(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	err := svc.ChangeStatus(ctx, &V1QuestionChangeStatusRequest{})
	_ = err
}

func TestReviewInfo(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.ReviewInfo(ctx, &V1ReviewInfoRequest{})
	if err != nil {
		t.Fatalf("ReviewInfo() error: %v", err)
	}
	_ = resp
}

func TestList(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.List(ctx, &V1QuestionAnswerListRequest{})
	if err != nil {
		t.Fatalf("List() error: %v", err)
	}
	_ = resp
}

func TestReviewChangeStatus(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	err := svc.ReviewChangeStatus(ctx, &V1ReviewChangeStatusRequest{})
	_ = err
}

func TestCommentList(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.CommentList(ctx, &V1CommentListRequest{})
	if err != nil {
		t.Fatalf("CommentList() error: %v", err)
	}
	_ = resp
}

func TestCount(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.Count(ctx)
	if err != nil {
		t.Fatalf("Count() error: %v", err)
	}
	_ = resp
}

func TestReviewListV2(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.ReviewListV2(ctx, &V2ReviewListV2Request{})
	if err != nil {
		t.Fatalf("ReviewListV2() error: %v", err)
	}
	_ = resp
}

func TestListV1(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.ListV1(ctx, &V1QuestionListRequest{})
	if err != nil {
		t.Fatalf("ListV1() error: %v", err)
	}
	_ = resp
}

func TestReviewCount(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.ReviewCount(ctx)
	if err != nil {
		t.Fatalf("ReviewCount() error: %v", err)
	}
	_ = resp
}

func TestDelete(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	err := svc.Delete(ctx, &V1QuestionAnswerDeleteRequest{})
	_ = err
}

func TestReviewList(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.ReviewList(ctx, &V1ReviewListRequest{})
	if err != nil {
		t.Fatalf("ReviewList() error: %v", err)
	}
	_ = resp
}
