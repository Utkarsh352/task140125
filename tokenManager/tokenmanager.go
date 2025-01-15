package tokenManager

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	model "github.com/utkarsh352/task140125/model"
)

type TokenManager struct {
	tokens    []model.Token
	lock      sync.Mutex
	ticker    *time.Ticker
	stopReset chan bool
}

// To initialize a new TokenManager with 1000 tokens
func NewTokenManager() *TokenManager {
	tokens := make([]model.Token, 1000)
	for i := 0; i < 1000; i++ {
		tokens[i] = model.Token{
			ID:    fmt.Sprintf("Token %d", i+1),
			Usage: 0,
		}
	}

	tm := &TokenManager{
		tokens:    tokens,
		stopReset: make(chan bool),
	}

	tm.startResetRoutine()
	return tm
}

// To select a token with the least usage count, randomly among ties
func (tm *TokenManager) SelectToken() *model.Token {
	tm.lock.Lock()
	defer tm.lock.Unlock()

	minUsage := tm.tokens[0].Usage
	candidates := []*model.Token{}

	// Identify the tokens with the least usage
	for i := range tm.tokens {
		if tm.tokens[i].Usage < minUsage {
			minUsage = tm.tokens[i].Usage
			candidates = []*model.Token{&tm.tokens[i]}
		} else if tm.tokens[i].Usage == minUsage {
			candidates = append(candidates, &tm.tokens[i])
		}
	}

	// Select randomly among candidates
	selected := candidates[rand.Intn(len(candidates))]
	selected.Usage++
	return selected
}

// To reset all tokens' usage counts to zero
func (tm *TokenManager) ResetTokens() {
	tm.lock.Lock()
	defer tm.lock.Unlock()
	for i := range tm.tokens {
		tm.tokens[i].Usage = 0
	}
}

// To Start the reset routine to reset tokens every 24 hours
func (tm *TokenManager) startResetRoutine() {
	tm.ticker = time.NewTicker(24 * time.Hour)
	go func() {
		for {
			select {
			case <-tm.ticker.C:
				tm.ResetTokens()
			case <-tm.stopReset:
				tm.ticker.Stop()
				return
			}
		}
	}()
}

// To Stop the reset routine
func (tm *TokenManager) Stop() {
	tm.stopReset <- true
}

// To Simulate token usage for a specified number of operations
func (tm *TokenManager) Simulate(operations int) {
	for i := 0; i < operations; i++ {
		tm.SelectToken()
	}
}

// returns all tokens in the TokenManager.
func (tm *TokenManager) GetAllTokens() []model.Token {
	tm.lock.Lock()
	defer tm.lock.Unlock()
	return append([]model.Token{}, tm.tokens...) // Return a copy to prevent modification
}

// returns the current list of tokens
func (tm *TokenManager) GetTokens() []model.Token {
	tm.lock.Lock()
	defer tm.lock.Unlock()
	return tm.tokens
}

// To gather token usage statistics for presentation
func (tm *TokenManager) GetStatistics() ([]model.Token, int) {
	tm.lock.Lock()
	defer tm.lock.Unlock()

	leastUsage := tm.tokens[0].Usage
	leastUsedTokens := []model.Token{}

	for _, token := range tm.tokens {
		if token.Usage < leastUsage {
			leastUsage = token.Usage
			leastUsedTokens = []model.Token{token}
		} else if token.Usage == leastUsage {
			leastUsedTokens = append(leastUsedTokens, token)
		}
	}

	return leastUsedTokens, leastUsage
}
