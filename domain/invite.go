package domain

import "time"

type Invite struct {
	ID            int64      `db:"id"`
	Code          string     `db:"code"`
	InvitedUserID string     `db:"invited_user_uuid"`
	CreatedAt     *time.Time `db:"created_at"`
}
