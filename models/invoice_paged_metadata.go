package models

import "github.com/go-swagger/go-swagger/strfmt"

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

/*
InvoicePagedMetadata invoice paged metadata

swagger:model InvoicePagedMetadata
*/
type InvoicePagedMetadata struct {

	/* {"description":"Paging parameter. 0-indexed. Describes your current location within a pageable list of query results.","verbs":["GET","PUT","POST"]}
	 */
	CurrentOffset int32 `json:"currentOffset,omitempty"`

	/* {"description":"Paging parameter. 0-indexed. Describes which page (given a page size of `recordsRequested`) of the result set you are viewing.","verbs":["GET","PUT","POST"]}
	 */
	CurrentPage int32 `json:"currentPage,omitempty"`

	/* {"description":"Number of milliseconds taken by API to calculate response.","verbs":["GET","PUT","POST"]}
	 */
	ExecutionTime int64 `json:"executionTime,omitempty"`

	/* {"description":"Paging parameter. URL fragment that can be used to fetch next page of results.","verbs":["GET","PUT","POST"]}
	 */
	NextPage string `json:"nextPage,omitempty"`

	/* {"default":10,"description":"Paging parameter. Describes how many records you requested.","verbs":["GET","PUT","POST"]}
	 */
	RecordsRequested int32 `json:"recordsRequested,omitempty"`

	/* {"description":"Describes how many records were returned by your query.","verbs":["GET","PUT","POST"]}
	 */
	RecordsReturned int32 `json:"recordsReturned,omitempty"`

	/* {"description":"The results returned by your query.","verbs":["GET","PUT","POST"]}
	 */
	Results []*Invoice `json:"results,omitempty"`
}

// Validate validates this invoice paged metadata
func (m *InvoicePagedMetadata) Validate(formats strfmt.Registry) error {
	return nil
}
