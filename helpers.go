package sportmonks

type (
	aggregatedAssistScorerData struct {
		Data []AssistScorer `json:"data"`
	}

	aggregatedCardScorerData struct {
		Data []CardScorer `json:"data"`
	}

	aggregatedGoalScorerData struct {
		Data []GoalScorer `json:"data"`
	}

	aggregatedMatchOddsData struct {
		Data []MatchOdds `json:"data"`
	}

	assistScorerData struct {
		Data []AssistScorer `json:"data"`
	}

	bookmakerOddsData struct {
		Data []BookmakerOdds `json:"data"`
	}

	cardEventsData struct {
		Data []CardEvent `json:"data"`
	}

	cardScorerData struct {
		Data []CardScorer `json:"data"`
	}

	coachData struct {
		Data *Coach `json:"data"`
	}

	cornerEventsData struct {
		Data []CornerEvent `json:"data"`
	}

	countryData struct {
		Data *Country `json:"data"`
	}

	countriesData struct {
		Data []Country `json:"data"`
	}

	continentData struct {
		Data *Continent `json:"data"`
	}

	fixtureData struct {
		Data *Fixture `json:"data"`
	}

	fixturesData struct {
		Data []Fixture `json:"data"`
	}

	goalEventsData struct {
		Data []GoalEvent `json:"data"`
	}

	goalScorerData struct {
		Data []GoalScorer `json:"data"`
	}

	groupsData struct {
		Data []Group `json:"data"`
	}

	leagueData struct {
		Data *League `json:"data"`
	}

	leaguesData struct {
		Data []League `json:"data"`
	}

	leagueStandingData struct {
		Data []LeagueStanding `json:"data"`
	}

	matchCommentaryData struct {
		Data []Commentary `json:"data"`
	}

	matchHighlightsData struct {
		Data []VideoHighlights `json:"data"`
	}

	matchEventsData struct {
		Data []MatchEvent `json:"data"`
	}

	matchOfficialData struct {
		Data *MatchOfficial `json:"data"`
	}

	matchTVStationsData struct {
		Data []TVStation `json:"data"`
	}

	oddsData struct {
		Data []Odds `json:"data"`
	}

	playerData struct {
		Data *Player `json:"data"`
	}

	playerLineupData struct {
		Data []PlayerStats `json:"data"`
	}

	playerSeasonStatsData struct {
		Data []PlayerSeasonStats `json:"data"`
	}

	positionData struct {
		Data *Position `json:"data"`
	}

	rankingData struct {
		Data *Ranking `json:"data"`
	}

	roundData struct {
		Data *Round `json:"data"`
	}

	roundsData struct {
		Data []Round `json:"data"`
	}

	seasonData struct {
		Data *Season `json:"data"`
	}

	seasonsData struct {
		Data []Season `json:"data"`
	}

	sidelinedData struct {
		Data []Sidelined `json:"data"`
	}

	squadPlayerData struct {
		Data []SquadPlayer `json:"data"`
	}

	stageData struct {
		Data *Stage `json:"data"`
	}

	stagesData struct {
		Data []Stage `json:"data"`
	}

	substitutionEventsData struct {
		Data []SubstitutionEvent `json:"data"`
	}

	teamData struct {
		Data *Team `json:"data"`
	}

	teamSeasonStatsData struct {
		Data *TeamSeasonStats `json:"data"`
	}

	teamsStatsData struct {
		Data []TeamStats `json:"data"`
	}

	transfersData struct {
		Data []Transfer `json:"data"`
	}

	trophyData struct {
		Data []Trophy `json:"data"`
	}

	venueData struct {
		Data *Venue `json:"data"`
	}
)
