package services

import (
	"regexp"
	"time"

	"github.com/DenisKnez/muna/data"
	"github.com/DenisKnez/muna/domains"
	errors "github.com/DenisKnez/muna/services/errors"
	"github.com/DenisKnez/muna/util"
	"github.com/google/uuid"
)

//GuessGameService guess game service
type GuessGameService struct {
	ggRepo domains.GuessGameRepository
}

//NewGuessGameService returns new guess game service
func NewGuessGameService(ggRepo domains.GuessGameRepository) domains.GuessGameService {
	return &GuessGameService{ggRepo}
}

//Check when the validity of the provided string
func (ggService *GuessGameService) Check(gameID string, guess string) (ok bool, err error) {

	err = checkStringLenght(guess)

	if err != nil {
		return false, err
	}

	isInSet := firstCharacterContainedInSet(guess)

	isPatternGuessed, err := false, nil

	if isInSet == false {
		isPatternGuessed, err = firstCharacterNotInSet(guess)
	} else {
		isPatternGuessed, err = firstCharacterInSet(guess)
	}

	if err != nil {
		return
	}

	gameUUID, err := uuid.Parse(gameID)

	if err != nil {
		return
	}

	if isPatternGuessed {
		ggService.ggRepo.ChangeInfoState(gameUUID)
	}

	ggService.logHistory(gameUUID, guess)

	return
}

func (ggService *GuessGameService) logHistory(gameID uuid.UUID, guess string) {

	historyItemUUID := util.NewUUID()
	timestamp := time.Now().Format(time.RFC3339)

	ggService.ggRepo.LogHistory(gameID, historyItemUUID, timestamp, guess)
}

//if the first character in the string is in set ['a', 'e', 'i', 'o', 'u'] returns true
func firstCharacterContainedInSet(guess string) (ok bool) {

	characters := [5]byte{'a', 'e', 'i', 'o', 'u'}
	firstCharater := guess[0]

	for _, v := range characters {
		if v == firstCharater {
			return true
		}
	}

	return false
}

func checkStringLenght(guess string) error {
	if len(guess) > 100 {
		return errors.ErrStringToLong
	}
	return nil
}

//firstCharacterInSet match the string and return ok if it matches
func firstCharacterInSet(guess string) (matched bool, err error) {

	regex, err := regexp.Compile(`^[a,e,i,o,u]([a,e,i,o,u])([a,e,i,o,u]*)[b]([#]*)[a]([#]*)[g]([#]*)[u]([#]*)[e]([#]*)[t]([#]*)[t]([#]*)[e]`)

	if err != nil {
		matched = false
		err = errors.ErrFailedRegexCompilation
		return
	}

	matched = regex.MatchString(guess)
	return
}

//FirstCharacterNotInSet match the string and return ok if the string
// does not contain any of characters in this set ["a", "e", "i", "o", "u"] and has the
// last charater as "!"
func firstCharacterNotInSet(guess string) (matched bool, err error) {

	matched, err = regexp.MatchString(`^[a,e,i,o,u].*!$`, guess)

	if err != nil {
		matched = false
		err = errors.ErrFailedRegexCompilation
		return
	}

	return
}

//NewGame create game uuid
func (ggService *GuessGameService) NewGame() uuid.UUID {

	return util.NewUUID()
}

//Stat get games stats
func (ggService *GuessGameService) Stat(gameID string) (info data.Info, err error) {
	gameUUID, err := uuid.Parse(gameID)

	if err != nil {
		return
	}

	info, err = ggService.ggRepo.Stat(gameUUID)

	if err != nil {
		return
	}

	return
}
