package repository

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/techx/portal/client/db"
	"github.com/techx/portal/constants"
	"github.com/techx/portal/domain"
)

const (
	insertInviteQuery = `INSERT INTO invites (created_time, code, invited_user_uuid) VALUES (:created_time, :code, :invited_user_uuid) RETURNING id`
)

type InvitesRepository interface {
	Insert(ctx context.Context, invite domain.Invite) (*domain.Invite, error)
}

type invitesRepository struct {
	dbClient db.Client
}

func NewInvitesRepository(dbClient db.Client) InvitesRepository {
	return &invitesRepository{
		dbClient: dbClient,
	}
}

type InvitesReturning struct {
	InviteID int64 `db:"id"`
}

func (r invitesRepository) Insert(ctx context.Context, invite domain.Invite) (*domain.Invite, error) {
	var returning InvitesReturning
	now := time.Now()
	invite.CreatedAt = &now
	err := r.dbClient.DBRunInTxContext(ctx, func(tx *sqlx.Tx) error {
		return r.dbClient.DBNamedExecReturningInTx(ctx, tx, &returning, insertInviteQuery, map[string]interface{}{
			constants.ParamCreateTime:    invite.CreatedAt,
			constants.ParamCode:          invite.Code,
			constants.ParamInvitedUserID: invite.InvitedUserID,
		})
	})
	if err != nil {
		return nil, err
	}

	invite.ID = returning.InviteID
	return &invite, nil
}
