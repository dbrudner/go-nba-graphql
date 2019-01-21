package main

type NBAInfoJSON struct {
	_internal struct {
		EventName   string `json:"eventName"`
		PubDateTime string `json:"pubDateTime"`
		Xslt        string `json:"xslt"`
	} `json:"_internal"`
	Links struct {
		AllstarRoster               string `json:"allstarRoster"`
		AnchorDate                  string `json:"anchorDate"`
		Boxscore                    string `json:"boxscore"`
		Calendar                    string `json:"calendar"`
		CurrentDate                 string `json:"currentDate"`
		CurrentScoreboard           string `json:"currentScoreboard"`
		GameBookPdf                 string `json:"gameBookPdf"`
		LeadTracker                 string `json:"leadTracker"`
		LeagueConfStandings         string `json:"leagueConfStandings"`
		LeagueDivStandings          string `json:"leagueDivStandings"`
		LeagueLastFiveGameTeamStats string `json:"leagueLastFiveGameTeamStats"`
		LeagueMiniStandings         string `json:"leagueMiniStandings"`
		LeagueRosterCoaches         string `json:"leagueRosterCoaches"`
		LeagueRosterPlayers         string `json:"leagueRosterPlayers"`
		LeagueSchedule              string `json:"leagueSchedule"`
		LeagueTeamStatsLeaders      string `json:"leagueTeamStatsLeaders"`
		LeagueUngroupedStandings    string `json:"leagueUngroupedStandings"`
		MiniBoxscore                string `json:"miniBoxscore"`
		Pbp                         string `json:"pbp"`
		PlayerGameLog               string `json:"playerGameLog"`
		PlayerProfile               string `json:"playerProfile"`
		PlayerUberStats             string `json:"playerUberStats"`
		PlayoffSeriesLeaders        string `json:"playoffSeriesLeaders"`
		PlayoffsBracket             string `json:"playoffsBracket"`
		PreviewArticle              string `json:"previewArticle"`
		RecapArticle                string `json:"recapArticle"`
		Scoreboard                  string `json:"scoreboard"`
		TeamLeaders                 string `json:"teamLeaders"`
		TeamLeaders2                string `json:"teamLeaders2"`
		TeamRoster                  string `json:"teamRoster"`
		TeamSchedule                string `json:"teamSchedule"`
		TeamScheduleYear            string `json:"teamScheduleYear"`
		TeamScheduleYear2           string `json:"teamScheduleYear2"`
		Teams                       string `json:"teams"`
		TeamsConfig                 string `json:"teamsConfig"`
		TeamsConfigYear             string `json:"teamsConfigYear"`
		TodayScoreboard             string `json:"todayScoreboard"`
	} `json:"links"`
	SeasonScheduleYear int  `json:"seasonScheduleYear"`
	ShowPlayoffsClinch bool `json:"showPlayoffsClinch"`
}