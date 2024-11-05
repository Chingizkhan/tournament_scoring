package errs

import "errors"

var (
	TournamentAlreadyExists = errors.New("tournament_already_exists")
	TournamentNotExists     = errors.New("tournament_not_exists")
	DivisionsNotFinished    = errors.New("divisions_not_finished")
	UnexpectedArgument      = errors.New("unexpected_argument")
)
