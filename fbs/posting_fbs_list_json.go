package fbs

import "encoding/json"

// MarshalJSON omits an unset status change date without changing the public
// value type of LastChangedStatusDate in the generated request filter.
func (f PostingV4PostingFbsListRequestFilter) MarshalJSON() ([]byte, error) {
	type filterAlias PostingV4PostingFbsListRequestFilter
	var date *PostingV4PostingFbsListRequestFilterLastChangedStatusDate
	if f.LastChangedStatusDate != (PostingV4PostingFbsListRequestFilterLastChangedStatusDate{}) {
		date = &f.LastChangedStatusDate
	}
	return json.Marshal(struct {
		*filterAlias
		LastChangedStatusDate *PostingV4PostingFbsListRequestFilterLastChangedStatusDate `json:"last_changed_status_date,omitempty"`
	}{
		filterAlias:           (*filterAlias)(&f),
		LastChangedStatusDate: date,
	})
}
