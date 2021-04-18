package collector

import (
	"encoding/json"
	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/cache"
	"github.com/geziyor/geziyor/client"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
	"regexp"
	"time"
)

const (
	subSystem  = "battlemetrics"
	metricName = "rank"
)

var (
	gez = geziyor.NewGeziyor(&geziyor.Options{
		AllowedDomains: []string{"www.battlemetrics.com"},
		CachePolicy:    cache.RFC2616,
		Timeout:        time.Second * 10,
	})
	rxState = regexp.MustCompile(`<script id="storeBootstrap" type="application/json">(.+?)</script>`)
)

type battleMetricsCollector struct {
	rank []*prometheus.Desc
}

func init() {
	Factories["rank"] = NewBattleMetricsCollector
}

// NewBattleMetricsCollector returns a new Collector exposing the current rankings.
func NewBattleMetricsCollector() (Collector, error) {
	var ranks []*prometheus.Desc
	return &battleMetricsCollector{
		rank: ranks,
	}, nil
}

func (c *battleMetricsCollector) Update(ch chan<- prometheus.Metric) error {
	res, err := fetch(gez)
	if err != nil {
		return err
	}
	for _, server := range res.State.Servers.Servers {
		rank := prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subSystem, metricName),
			"The current battlemetrics server rank.",
			nil, prometheus.Labels{
				"server": server.Name,
			})
		ch <- prometheus.MustNewConstMetric(
			rank, prometheus.GaugeValue, float64(server.Rank))
	}
	return nil
}

type gtRules struct {
	TfSpawnGlowsDuration                 string `json:"tf_spawn_glows_duration"`
	MpMatchEndAtTimelimit                string `json:"mp_match_end_at_timelimit"`
	MpScrambleteamsAuto                  string `json:"mp_scrambleteams_auto"`
	SvNoclipspeed                        string `json:"sv_noclipspeed"`
	TfPasstimeThrowspeedEngineer         string `json:"tf_passtime_throwspeed_engineer"`
	Coop                                 string `json:"coop"`
	TfGamemodeArena                      string `json:"tf_gamemode_arena"`
	TfPasstimeOvertimeIdleSec            string `json:"tf_passtime_overtime_idle_sec"`
	TfWeaponCriticalsMelee               string `json:"tf_weapon_criticals_melee"`
	Nextlevel                            string `json:"nextlevel"`
	TfPasstimeBallMass                   string `json:"tf_passtime_ball_mass"`
	SvMaxusrcmdprocessticks              string `json:"sv_maxusrcmdprocessticks"`
	TfOvertimeNag                        string `json:"tf_overtime_nag"`
	MpTournamentReadymodeTeamSize        string `json:"mp_tournament_readymode_team_size"`
	SmClassrestrictVersion               string `json:"sm_classrestrict_version"`
	TfPasstimeBallSphereRadius           string `json:"tf_passtime_ball_sphere_radius"`
	MpTournamentReadymode                string `json:"mp_tournament_readymode"`
	MetamodVersion                       string `json:"metamod_version"`
	SvGravity                            string `json:"sv_gravity"`
	SvFootsteps                          string `json:"sv_footsteps"`
	RJeepViewDampenDamp                  string `json:"r_JeepViewDampenDamp"`
	TfSpecXray                           string `json:"tf_spec_xray"`
	SvWaterfriction                      string `json:"sv_waterfriction"`
	TfGamemodePd                         string `json:"tf_gamemode_pd"`
	MpFlashlight                         string `json:"mp_flashlight"`
	MpTournamentReadymodeMin             string `json:"mp_tournament_readymode_min"`
	TfPasstimeThrowspeedSoldier          string `json:"tf_passtime_throwspeed_soldier"`
	SvStepsize                           string `json:"sv_stepsize"`
	SvVoteQuorumRatio                    string `json:"sv_vote_quorum_ratio"`
	TfMmStrict                           string `json:"tf_mm_strict"`
	TfArenaForceClass                    string `json:"tf_arena_force_class"`
	TfPasstimeSpeedboostOnGetBallTime    string `json:"tf_passtime_speedboost_on_get_ball_time"`
	TfPasstimeBallSeekRange              string `json:"tf_passtime_ball_seek_range"`
	SvBounce                             string `json:"sv_bounce"`
	TfPasstimeThrowarcDemoman            string `json:"tf_passtime_throwarc_demoman"`
	TfPasstimeThrowarcSniper             string `json:"tf_passtime_throwarc_sniper"`
	TfServerIdentityDisableQuickplay     string `json:"tf_server_identity_disable_quickplay"`
	SvSpecaccelerate                     string `json:"sv_specaccelerate"`
	TfPasstimeThrowspeedSpy              string `json:"tf_passtime_throwspeed_spy"`
	TvRelaypassword                      string `json:"tv_relaypassword"`
	SvAlltalk                            string `json:"sv_alltalk"`
	CronjobsVersion                      string `json:"cronjobs_version"`
	TfPasstimeThrowspeedMedic            string `json:"tf_passtime_throwspeed_medic"`
	TfPasstimeBallInertiaScale           string `json:"tf_passtime_ball_inertia_scale"`
	SvAccelerate                         string `json:"sv_accelerate"`
	TfMmServermode                       string `json:"tf_mm_servermode"`
	TfPasstimeThrowarcEngineer           string `json:"tf_passtime_throwarc_engineer"`
	SvVoiceenable                        string `json:"sv_voiceenable"`
	MpTeamplay                           string `json:"mp_teamplay"`
	RVehicleViewDampen                   string `json:"r_VehicleViewDampen"`
	TfPasstimeBallSeekSpeedFactor        string `json:"tf_passtime_ball_seek_speed_factor"`
	TfGamemodeTc                         string `json:"tf_gamemode_tc"`
	TfArenaOverrideCapEnableTime         string `json:"tf_arena_override_cap_enable_time"`
	MpFadetoblack                        string `json:"mp_fadetoblack"`
	TfBotCount                           string `json:"tf_bot_count"`
	MpWindifference                      string `json:"mp_windifference"`
	MpTournamentStopwatch                string `json:"mp_tournament_stopwatch"`
	SmNextmap                            string `json:"sm_nextmap"`
	SvRegistrationMessage                string `json:"sv_registration_message"`
	MpWindifferenceMin                   string `json:"mp_windifference_min"`
	TfClasslimit                         string `json:"tf_classlimit"`
	TfPasstimeBallModel                  string `json:"tf_passtime_ball_model"`
	TfPasstimeThrowspeedScout            string `json:"tf_passtime_throwspeed_scout"`
	MpDisableRespawnTimes                string `json:"mp_disable_respawn_times"`
	MpHighlander                         string `json:"mp_highlander"`
	MpScrambleteamsAutoWindifference     string `json:"mp_scrambleteams_auto_windifference"`
	RAirboatViewDampenFreq               string `json:"r_AirboatViewDampenFreq"`
	SvRegistrationSuccessful             string `json:"sv_registration_successful"`
	TfArenaRoundTime                     string `json:"tf_arena_round_time"`
	SmNtVerison                          string `json:"sm_nt_verison"`
	TfPasstimePackRange                  string `json:"tf_passtime_pack_range"`
	TfPasstimeBallDampingScale           string `json:"tf_passtime_ball_damping_scale"`
	TfGamemodePayload                    string `json:"tf_gamemode_payload"`
	TfMvmDeathPenalty                    string `json:"tf_mvm_death_penalty"`
	TfGamemodeSd                         string `json:"tf_gamemode_sd"`
	TfPasstimePowerballDecaysecNeutral   string `json:"tf_passtime_powerball_decaysec_neutral"`
	TfArenaFirstBlood                    string `json:"tf_arena_first_blood"`
	TfPasstimeScoresPerRound             string `json:"tf_passtime_scores_per_round"`
	SvRollangle                          string `json:"sv_rollangle"`
	MpWinlimit                           string `json:"mp_winlimit"`
	MpAllowNPCs                          string `json:"mp_allowNPCs"`
	SvFriction                           string `json:"sv_friction"`
	MpAutoteambalance                    string `json:"mp_autoteambalance"`
	TfPasstimeExperimentTelepass         string `json:"tf_passtime_experiment_telepass"`
	TvPassword                           string `json:"tv_password"`
	TfPasstimeExperimentAutopass         string `json:"tf_passtime_experiment_autopass"`
	RAirboatViewDampenDamp               string `json:"r_AirboatViewDampenDamp"`
	MpFraglimit                          string `json:"mp_fraglimit"`
	TfPasstimeModeHomingSpeed            string `json:"tf_passtime_mode_homing_speed"`
	MpForcerespawn                       string `json:"mp_forcerespawn"`
	TfMmTrusted                          string `json:"tf_mm_trusted"`
	TfPasstimeSaveStats                  string `json:"tf_passtime_save_stats"`
	TfPasstimeThrowspeedPyro             string `json:"tf_passtime_throwspeed_pyro"`
	TfGamemodeRd                         string `json:"tf_gamemode_rd"`
	MpFootsteps                          string `json:"mp_footsteps"`
	SvAiraccelerate                      string `json:"sv_airaccelerate"`
	MpFriendlyfire                       string `json:"mp_friendlyfire"`
	TfPasstimePackSpeed                  string `json:"tf_passtime_pack_speed"`
	TfPasstimePowerballAirtimebonus      string `json:"tf_passtime_powerball_airtimebonus"`
	TfPasstimeStealOnMelee               string `json:"tf_passtime_steal_on_melee"`
	TfPasstimeExperimentInstapassCharge  string `json:"tf_passtime_experiment_instapass_charge"`
	RAirboatViewZHeight                  string `json:"r_AirboatViewZHeight"`
	SvPassword                           string `json:"sv_password"`
	TfArenaChangeLimit                   string `json:"tf_arena_change_limit"`
	TfPasstimePowerballDecayamount       string `json:"tf_passtime_powerball_decayamount"`
	TfPasstimePowerballDecayDelay        string `json:"tf_passtime_powerball_decay_delay"`
	MpTeamlist                           string `json:"mp_teamlist"`
	MpHolidayNogifts                     string `json:"mp_holiday_nogifts"`
	TfPasstimeThrowarcSpy                string `json:"tf_passtime_throwarc_spy"`
	MpFalldamage                         string `json:"mp_falldamage"`
	MpTournamentReadymodeCountdown       string `json:"mp_tournament_readymode_countdown"`
	TfPasstimeThrowarcSoldier            string `json:"tf_passtime_throwarc_soldier"`
	TfMedievalAutorp                     string `json:"tf_medieval_autorp"`
	TfPasstimeScoreCritSec               string `json:"tf_passtime_score_crit_sec"`
	TfHalloweenAllowTruceDuringBossEvent string `json:"tf_halloween_allow_truce_during_boss_event"`
	MpMaxrounds                          string `json:"mp_maxrounds"`
	TfAllowPlayerUse                     string `json:"tf_allow_player_use"`
	TfPasstimeModeHomingLockSec          string `json:"tf_passtime_mode_homing_lock_sec"`
	MpStalemateEnable                    string `json:"mp_stalemate_enable"`
	TfPasstimeBallSphereCollision        string `json:"tf_passtime_ball_sphere_collision"`
	TfDamageDisablespread                string `json:"tf_damage_disablespread"`
	TfPasstimeFlinchBoost                string `json:"tf_passtime_flinch_boost"`
	TfMvmMinPlayersToStart               string `json:"tf_mvm_min_players_to_start"`
	TfPasstimeThrowarcHeavy              string `json:"tf_passtime_throwarc_heavy"`
	SvMaxspeed                           string `json:"sv_maxspeed"`
	TfPasstimeBallRotdampingScale        string `json:"tf_passtime_ball_rotdamping_scale"`
	TfPasstimePowerballDecaysec          string `json:"tf_passtime_powerball_decaysec"`
	TfPasstimePackHpPerSec               string `json:"tf_passtime_pack_hp_per_sec"`
	MpWeaponstay                         string `json:"mp_weaponstay"`
	TfUseFixedWeaponspreads              string `json:"tf_use_fixed_weaponspreads"`
	TfArenaUseQueue                      string `json:"tf_arena_use_queue"`
	TfBetaContent                        string `json:"tf_beta_content"`
	ExtendedmapconfigVersion             string `json:"extendedmapconfig_version"`
	Deathmatch                           string `json:"deathmatch"`
	TfPasstimeBallResetTime              string `json:"tf_passtime_ball_reset_time"`
	MpTournament                         string `json:"mp_tournament"`
	Decalfrequency                       string `json:"decalfrequency"`
	TfPasstimeBallTakedamage             string `json:"tf_passtime_ball_takedamage"`
	RJeepViewDampenFreq                  string `json:"r_JeepViewDampenFreq"`
	SvContact                            string `json:"sv_contact"`
	TfGravetalk                          string `json:"tf_gravetalk"`
	TfPasstimePowerballThreshold         string `json:"tf_passtime_powerball_threshold"`
	SmTidychatVersion                    string `json:"sm_tidychat_version"`
	TfCtfBonusTime                       string `json:"tf_ctf_bonus_time"`
	TfPasstimeTeammateStealTime          string `json:"tf_passtime_teammate_steal_time"`
	TfForceHolidaysOff                   string `json:"tf_force_holidays_off"`
	TfPowerupMode                        string `json:"tf_powerup_mode"`
	TfArenaPreroundTime                  string `json:"tf_arena_preround_time"`
	TfPasstimeThrowarcMedic              string `json:"tf_passtime_throwarc_medic"`
	TfMaxChargeSpeed                     string `json:"tf_max_charge_speed"`
	TvEnable                             string `json:"tv_enable"`
	TfPasstimeThrowarcPyro               string `json:"tf_passtime_throwarc_pyro"`
	SvRollspeed                          string `json:"sv_rollspeed"`
	TfPlayergib                          string `json:"tf_playergib"`
	TfBirthday                           string `json:"tf_birthday"`
	TfPasstimePlayerReticlesEnemies      string `json:"tf_passtime_player_reticles_enemies"`
	TfGamemodePasstime                   string `json:"tf_gamemode_passtime"`
	TfPasstimeThrowarcScout              string `json:"tf_passtime_throwarc_scout"`
	CrontabVersion                       string `json:"crontab_version"`
	SourcemodVersion                     string `json:"sourcemod_version"`
	TfArenaMaxStreak                     string `json:"tf_arena_max_streak"`
	TfPasstimePlayerReticlesFriends      string `json:"tf_passtime_player_reticles_friends"`
	TfPlayerNameChangeTime               string `json:"tf_player_name_change_time"`
	MpStalemateMeleeonly                 string `json:"mp_stalemate_meleeonly"`
	MpTimelimit                          string `json:"mp_timelimit"`
	TfGamemodeMvm                        string `json:"tf_gamemode_mvm"`
	SvStopspeed                          string `json:"sv_stopspeed"`
	SvCheats                             string `json:"sv_cheats"`
	TfGamemodeCtf                        string `json:"tf_gamemode_ctf"`
	NativevotesVersion                   string `json:"nativevotes_version"`
	TfPasstimeExperimentInstapass        string `json:"tf_passtime_experiment_instapass"`
	TfPasstimeThrowspeedHeavy            string `json:"tf_passtime_throwspeed_heavy"`
	SvNoclipaccelerate                   string `json:"sv_noclipaccelerate"`
	SvSteamgroup                         string `json:"sv_steamgroup"`
	MpForceautoteam                      string `json:"mp_forceautoteam"`
	TfPasstimeBallTakedamageForce        string `json:"tf_passtime_ball_takedamage_force"`
	TfPasstimeThrowspeedDemoman          string `json:"tf_passtime_throwspeed_demoman"`
	TfPasstimeThrowspeedVelocityScale    string `json:"tf_passtime_throwspeed_velocity_scale"`
	TfPasstimeThrowspeedSniper           string `json:"tf_passtime_throwspeed_sniper"`
	RJeepViewZHeight                     string `json:"r_JeepViewZHeight"`
	SvTags                               string `json:"sv_tags"`
	TfPasstimeBallDragCoefficient        string `json:"tf_passtime_ball_drag_coefficient"`
	SvSpecnoclip                         string `json:"sv_specnoclip"`
	MpAutocrosshair                      string `json:"mp_autocrosshair"`
	TfPasstimePowerballMaxairtimebonus   string `json:"tf_passtime_powerball_maxairtimebonus"`
	TfGamemodeMisc                       string `json:"tf_gamemode_misc"`
	TfWeaponCriticals                    string `json:"tf_weapon_criticals"`
	SvWateraccelerate                    string `json:"sv_wateraccelerate"`
	TfPasstimePowerballPasspoints        string `json:"tf_passtime_powerball_passpoints"`
	TfSpellsEnabled                      string `json:"tf_spells_enabled"`
	SvPausable                           string `json:"sv_pausable"`
	TfGamemodeCp                         string `json:"tf_gamemode_cp"`
	SvSpecspeed                          string `json:"sv_specspeed"`
	NativevotesMapchooserVersion         string `json:"nativevotes_mapchooser_version"`
	MpRespawnwavetime                    string `json:"mp_respawnwavetime"`
	TfMedieval                           string `json:"tf_medieval"`
}

type gtServer struct {
	ID         string    `json:"id"`
	GameID     string    `json:"game_id"`
	Name       string    `json:"name"`
	Address    string    `json:"address"`
	IP         string    `json:"ip"`
	Port       int       `json:"port"`
	Players    int       `json:"players"`
	MaxPlayers int       `json:"maxPlayers"`
	Rank       int       `json:"rank"`
	Location   []float64 `json:"location"`
	Country    string    `json:"country"`
	Status     string    `json:"status"`
	Details    struct {
		GameMode      string  `json:"gameMode"`
		Rules         gtRules `json:"rules"`
		ServerSteamID string  `json:"serverSteamId"`
		Password      bool    `json:"password"`
		Numbots       int     `json:"numbots"`
		Map           string  `json:"map"`
		Tags          string  `json:"tags"`
	} `json:"details"`
}
type gameTrackerState struct {
	Tracker []interface{} `json:"tracker"`
	State   struct {
		LoadManager struct {
			Account struct {
				Loaded   bool  `json:"loaded"`
				LoadedAt int64 `json:"loadedAt"`
				Loading  bool  `json:"loading"`
			} `json:"account"`
			Games struct {
				Tf2 struct {
					Loading  bool  `json:"loading"`
					Loaded   bool  `json:"loaded"`
					LoadedAt int64 `json:"loadedAt"`
				} `json:"tf2"`
			} `json:"games"`
			GameFeatures struct {
				Tf2 struct {
					Loading  bool  `json:"loading"`
					Loaded   bool  `json:"loaded"`
					LoadedAt int64 `json:"loadedAt"`
				} `json:"tf2"`
			} `json:"gameFeatures"`
			Servers struct {
				K2689418472 struct {
					Loading  bool  `json:"loading"`
					Loaded   bool  `json:"loaded"`
					LoadedAt int64 `json:"loadedAt"`
				} `json:"k:2689418472"`
			} `json:"servers"`
		} `json:"loadManager"`
		Account struct {
			Country  string        `json:"country"`
			Location []float64     `json:"location"`
			TzZones  []string      `json:"tzZones"`
			TzLinks  []interface{} `json:"tzLinks"`
		} `json:"account"`
		Servers struct {
			Servers map[string]gtServer `json:"servers"`
			Lists   struct {
				K2689418472 struct {
					NoAuth bool `json:"noAuth"`
					Params struct {
						Page struct {
						} `json:"page"`
						Fields struct {
							Server string `json:"server"`
						} `json:"fields"`
						Relations struct {
							Server string `json:"server"`
						} `json:"relations"`
						Filter struct {
							Game      string        `json:"game"`
							Search    string        `json:"search"`
							Countries []interface{} `json:"countries"`
							Players   struct {
							} `json:"players"`
							Ids struct {
							} `json:"ids"`
						} `json:"filter"`
					} `json:"params"`
					Links struct {
						Next string `json:"next"`
					} `json:"links"`
					ItemIds []string `json:"itemIds"`
				} `json:"k:2689418472"`
			} `json:"lists"`
			Groups struct {
			} `json:"groups"`
		} `json:"servers"`
		Games struct {
			Games struct {
				Tf2 struct {
					ID      string `json:"id"`
					Pretty  string `json:"pretty"`
					Options struct {
						Appid   int    `json:"appid"`
						Gamedir string `json:"gamedir"`
					} `json:"options"`
					Players          int            `json:"players"`
					Servers          int            `json:"servers"`
					ServersByCountry map[string]int `json:"serversByCountry"`
					PlayersByCountry map[string]int `json:"playersByCountry"`
					MaxPlayers30D    int            `json:"maxPlayers30D"`
					MaxPlayers7D     int            `json:"maxPlayers7D"`
					MaxPlayers24H    int            `json:"maxPlayers24H"`
					MinPlayers30D    int            `json:"minPlayers30D"`
					MinPlayers7D     int            `json:"minPlayers7D"`
					MinPlayers24H    int            `json:"minPlayers24H"`
				} `json:"tf2"`
			} `json:"games"`
			Features       interface{} `json:"-"`
			FeaturesByGame interface{} `json:"-"`
		} `json:"games"`
		Env interface{} `json:"-"`
	} `json:"state"`
}

func fetch(gz *geziyor.Geziyor) (gameTrackerState, error) {
	var result gameTrackerState
	doneChan := make(chan bool)
	parse := func(g *geziyor.Geziyor, r *client.Response) {
		m := rxState.FindSubmatch(r.Body)
		if len(m) == 0 {
			logrus.Errorf("Failed to find json data")
			doneChan <- true
			return
		}
		var state gameTrackerState
		if err := json.Unmarshal(m[1], &state); err != nil {
			logrus.Errorf("Failed to decode: %v", err)
			doneChan <- true
			return
		}
		result = state
		doneChan <- true
	}
	gz.GetRendered("https://www.battlemetrics.com/servers/tf2?q=uncletopia&sort=score", parse)
	<-doneChan
	return result, nil
}
