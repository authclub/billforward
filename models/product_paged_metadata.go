package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/strfmt"
	"github.com/go-swagger/go-swagger/swag"
)

/*ProductPagedMetadata ProductPagedMetadata product paged metadata

swagger:model ProductPagedMetadata
*/
type ProductPagedMetadata struct {

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
	NextPage *string `json:"nextPage,omitempty"`

	/* {"default":10,"description":"Paging parameter. Describes how many records you requested.","verbs":["GET","PUT","POST"]}
	 */
	RecordsRequested int32 `json:"recordsRequested,omitempty"`

	/* {"description":"Describes how many records were returned by your query.","verbs":["GET","PUT","POST"]}
	 */
	RecordsReturned int32 `json:"recordsReturned,omitempty"`

	/* {"description":"The results returned by your query.","verbs":["GET","PUT","POST"]}
	 */
	Results []*Product `json:"results,omitempty"`
}

// Validate validates this product paged metadata
func (m *ProductPagedMetadata) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResults(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProductPagedMetadata) validateResults(formats strfmt.Registry) error {

	if swag.IsZero(m.Results) { // not required
		return nil
	}

	for i := 0; i < len(m.Results); i++ {

		if m.Results[i] != nil {

			if err := m.Results[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}
