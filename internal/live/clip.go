package live

import (
	"context"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/zibbp/ganymede/ent/channel"
	"github.com/zibbp/ganymede/ent/live"
	"github.com/zibbp/ganymede/ent/vod"
	"github.com/zibbp/ganymede/internal/archive"
	"github.com/zibbp/ganymede/internal/platform"
	"github.com/zibbp/ganymede/internal/utils"
)

// CheckWatchedChannelClips checks watched channels for clips to arhive
func (s *Service) CheckWatchedChannelClips(ctx context.Context, logger zerolog.Logger) error {
	// Get channels from DB
	channels, err := s.Store.Client.Live.Query().Where(live.WatchClips(true)).WithChannel().All(context.Background())
	if err != nil {
		return err
	}

	if len(channels) == 0 {
		logger.Info().Msg("no channels to check")
		return nil
	}

	logger.Info().Msgf("checking %d channels for new clips", len(channels))

	for _, watchedChannel := range channels {

		if watchedChannel.ClipsLimit == 0 || watchedChannel.ClipsIntervalDays == 0 {
			logger.Error().Err(err).Str("channel", watchedChannel.Edges.Channel.Name).Msg("clip limit and clips interval must be greater than 0")
			continue
		}

		// Check if clips need to be checked for channel
		now := time.Now()
		lastChecked := watchedChannel.ClipsLastChecked

		// Round times down to the nearest hour to avoid partial hour comparisons
		roundedNow := now.Truncate(time.Hour)
		roundedLastChecked := lastChecked.Truncate(time.Hour)

		// Calculate full hours between checks
		diffHours := roundedNow.Sub(roundedLastChecked).Hours()
		channelIntervalHours := watchedChannel.ClipsIntervalDays * 24

		if diffHours < float64(channelIntervalHours) {
			logger.Info().
				Str("channel", watchedChannel.Edges.Channel.DisplayName).
				Float64("hours_passed", diffHours).
				Int("hours_required", channelIntervalHours).
				Msg("skipping clip check, not enough time has passed")
			continue
		}

		startedAt := now.AddDate(0, 0, -watchedChannel.ClipsIntervalDays)

		// Get clips
		clips, err := s.PlatformTwitch.GetChannelClips(ctx, watchedChannel.Edges.Channel.ExtID, platform.ClipsFilter{
			Limit:     watchedChannel.ClipsLimit,
			StartedAt: startedAt,
			EndedAt:   now,
		})
		if err != nil {
			logger.Error().Err(err).Str("channel", watchedChannel.Edges.Channel.Name).Msg("error getting clips for channel")
			continue
		}

		// Fetch all videos from DB
		dbVideos, err := s.Store.Client.Vod.Query().Where(vod.HasChannelWith(channel.ID(watchedChannel.Edges.Channel.ID))).All(context.Background())
		if err != nil {
			logger.Error().Str("channel", watchedChannel.Edges.Channel.Name).Err(err).Msg("error getting videos from database")
			continue
		}

		// Check if video is already in DB
		for _, clip := range clips {
			// Video is not in DB
			if !contains(dbVideos, clip.ID) {
				// Archive clip
				input := archive.ArchiveClipInput{
					ID:          clip.ID,
					Quality:     utils.VodQuality(watchedChannel.Resolution),
					ArchiveChat: watchedChannel.ArchiveChat,
					RenderChat:  watchedChannel.RenderChat,
				}
				err = s.ArchiveService.ArchiveClip(ctx, input)
				if err != nil {
					log.Error().Err(err).Str("clip_id", clip.ID).Msgf("error archiving clip")
					continue
				}
				logger.Info().Str("clip_id", clip.ID).Msgf("archiving clip")
			}
		}

		// Set clip last checked time on channel
		err = s.Store.Client.Live.Update().Where(live.ID(watchedChannel.ID)).SetClipsLastChecked(now).Exec(ctx)
		if err != nil {
			logger.Error().Err(err).Str("channe", watchedChannel.Edges.Channel.Name).Msg("error setting clip last check time on channel")
		}
	}
	return nil
}