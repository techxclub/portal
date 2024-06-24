package domain

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"

	"github.com/techx/portal/utils"
)

var (
	_ sql.Scanner   = (*MentorConfig)(nil)
	_ driver.Valuer = MentorConfig{}
)

type MentorConfig struct {
	IsMentor      bool     `json:"is_mentor"`
	IsPreApproved bool     `json:"is_pre_approved"`
	Status        string   `json:"status"`
	CalendalyLink string   `json:"calendaly_link,omitempty"`
	Description   string   `json:"description,omitempty"`
	IsAvailable   bool     `json:"is_available,omitempty"`
	Tags          []string `json:"tags,omitempty"`
	Domain        string   `json:"domain,omitempty"`
}

func (m MentorConfig) Value() (driver.Value, error) {
	return json.Marshal(m)
}

func (m *MentorConfig) Scan(data interface{}) error {
	return utils.ScanJSON(data, m)
}
