package services

import (
	"encoding/json"
	"fmt"
	"regexp"
	"time"

	"github.com/DenisKnez/muna/data"
	"github.com/DenisKnez/muna/domains"
	errors "github.com/DenisKnez/muna/services/errors"
	"github.com/DenisKnez/muna/util"
	"github.com/go-redis/redis"
	"github.com/google/uuid"
)

const (
	statRedisID = "stats"
)

//GuessGameService guess game service
type GuessGameService struct {
	ggRepo domains.GuessGameRepository
	redis  *redis.Client
}

//NewGuessGameService returns new guess game service
func NewGuessGameService(ggRepo domains.GuessGameRepository, redis *redis.Client) domains.GuessGameService {
	return &GuessGameService{ggRepo, redis}
}

//Check when the validity of the provided string
func (ggService *GuessGameService) Check(gameID uuid.UUID, guess string) (isPatternGuessed bool, err error) {

	err = checkStringLenght(guess)

	if err != nil {
		return false, err
	}

	isInSet := firstCharacterContainedInSet(guess)

	if isInSet == false {
		isPatternGuessed, err = firstCharacterNotInSet(guess)
	} else {
		isPatternGuessed, err = firstCharacterInSet(guess)
	}

	if err != nil {
		fmt.Println(err)
		return
	}

	if isPatternGuessed {
		err = ggService.ggRepo.ChangeInfoState(gameID)

		if err != nil {
			fmt.Println(err)
			return
		}
	}

	err = ggService.logHistory(gameID, guess)

	if err != nil {
		fmt.Println(err)

	}

	return
}

//logs the rounds played
func (ggService *GuessGameService) logHistory(gameID uuid.UUID, guess string) (err error) {

	historyItemUUID := util.NewUUID().String()
	timestamp := time.Now().Format(time.RFC3339)

	err = ggService.ggRepo.LogHistory(gameID.String(), historyItemUUID, timestamp, guess)

	return

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

//checks if there a game with that id in the database, if there is returns
func (ggService *GuessGameService) gameExists(gameID uuid.UUID) (ok bool) {
	return ggService.ggRepo.GameExists(gameID)
}

//NewGame create game uuid
func (ggService *GuessGameService) NewGame() (newID uuid.UUID) {
	newID = util.NewUUID()
	ggService.ggRepo.NewGame(newID, data.StateUnsolved)
	return
}

//Stat get games stats
func (ggService *GuessGameService) Stat(gameID uuid.UUID) (info data.Info, err error) {

	jsonString, err := ggService.redis.Get(statRedisID).Result()

	if err == redis.Nil {
		err = ggService.getStats(gameID, &info)

		if err != nil {
			fmt.Println(err)
		}

		return

	} else if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(jsonString)

	err = json.Unmarshal([]byte(jsonString), &info)

	if err != nil {
		fmt.Println(err)
		return
	}

	return
}

func (ggService *GuessGameService) getStats(gameID uuid.UUID, info *data.Info) (err error) {

	expirationTime := 10 * time.Second

	*info, err = ggService.ggRepo.Stat(gameID)

	if err != nil {
		fmt.Println(err)
		return
	}

	jsonInfo, err := json.Marshal(*info)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = ggService.redis.Set(statRedisID, string(jsonInfo), expirationTime).Err()

	if err != nil {
		fmt.Println(err)
	}

	return

}
