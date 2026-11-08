package fbs

import (
	"encoding/json"
	"testing"
)

func TestPostingFbsListRequestFilterLastChangedStatusDateJSON(t *testing.T) {
	for _, tc := range []struct {
		name     string
		date     PostingV4PostingFbsListRequestFilterLastChangedStatusDate
		wantDate bool
	}{
		{name: "unset"},
		{
			name: "set",
			date: PostingV4PostingFbsListRequestFilterLastChangedStatusDate{
				From: "2026-11-01T00:00:00Z",
				To:   "2026-11-02T00:00:00Z",
			},
			wantDate: true,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			req := PostingV4PostingFbsListRequest{
				Filter: PostingV4PostingFbsListRequestFilter{
					Since:                 "2026-11-01T00:00:00Z",
					To:                    "2026-11-02T00:00:00Z",
					LastChangedStatusDate: tc.date,
				},
				Limit: 100,
			}
			body, err := json.Marshal(req)
			if err != nil {
				t.Fatal(err)
			}
			var decoded struct {
				Filter map[string]json.RawMessage `json:"filter"`
			}
			if err := json.Unmarshal(body, &decoded); err != nil {
				t.Fatal(err)
			}
			for key, want := range map[string]string{
				"since": `"2026-11-01T00:00:00Z"`,
				"to":    `"2026-11-02T00:00:00Z"`,
			} {
				if string(decoded.Filter[key]) != want {
					t.Errorf("filter.%s = %s, want %s", key, decoded.Filter[key], want)
				}
			}
			dateJSON, present := decoded.Filter["last_changed_status_date"]
			if present != tc.wantDate {
				t.Fatalf("last_changed_status_date present = %t, want %t: %s", present, tc.wantDate, body)
			}
			if present {
				var got PostingV4PostingFbsListRequestFilterLastChangedStatusDate
				if err := json.Unmarshal(dateJSON, &got); err != nil {
					t.Fatal(err)
				}
				if got != tc.date {
					t.Errorf("last_changed_status_date = %+v, want %+v", got, tc.date)
				}
			}
		})
	}
}
