package response

// swagger:model
type ReferralResponse struct {
	Success bool `json:"success"`
}

func NewReferralResponse() (ReferralResponse, HTTPMetadata) {
	respBody := ReferralResponse{
		Success: true,
	}
	return respBody, HTTPMetadata{}
}
