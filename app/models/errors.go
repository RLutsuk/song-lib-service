package models

import "github.com/pkg/errors"

var (
	ErrSongNotFound   = errors.New("Song not found")
	ErrBadData        = errors.New("Bad request")
	ErrInternalServer = errors.New(("Internal server error"))
	ErrInvalidInput   = errors.New("Invalid input")
	// ErrTenderNotFound      = errors.New("Tender not found")
	// ErrBidNotFound         = errors.New("Bid not found")
	// ErrBidWasRejected      = errors.New("Bid was rejected or closed")
	// ErrUserHasmadeDecision = errors.New("the user has already made a decision")
)
