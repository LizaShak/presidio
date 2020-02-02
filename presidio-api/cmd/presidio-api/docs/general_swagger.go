package docs

import types "github.com/Microsoft/presidio-genproto/golang"

// swagger:route GET /fieldTypes fields getFieldTypes
//
// Get all available field types
//
// responses:
//   200: fieldTypesGetResponse

// swagger:response fieldTypesGetResponse
type fieldTypesResponseWrapper struct {
	// The response including a list of field types
	// in:body
	Types []types.FieldTypes
}
