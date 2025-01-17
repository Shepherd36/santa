// SPDX-License-Identifier: ice License 1.0

package badges

import (
	"context"
	"math"
	"sort"

	"github.com/pkg/errors"

	storage "github.com/ice-blockchain/wintr/connectors/storage/v2"
)

func (r *repository) GetBadges(ctx context.Context, groupType GroupType, userID string) ([]*Badge, error) {
	if ctx.Err() != nil {
		return nil, errors.Wrap(ctx.Err(), "unexpected deadline")
	}
	stats, err := r.getStatistics(ctx, groupType)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to getStatistics for %v", groupType)
	}
	userProgress, err := r.getProgress(ctx, userID, true)
	if err != nil && !errors.Is(err, ErrRelationNotFound) {
		return nil, errors.Wrapf(err, "failed to getProgress for userID:%v", userID)
	}
	if userProgress != nil && (userProgress.HideBadges && requestingUserID(ctx) != userID) {
		return nil, ErrHidden
	}

	return userProgress.buildBadges(groupType, stats), nil
}

func (r *repository) GetSummary(ctx context.Context, userID string) ([]*BadgeSummary, error) {
	if ctx.Err() != nil {
		return nil, errors.Wrap(ctx.Err(), "unexpected deadline")
	}
	userProgress, err := r.getProgress(ctx, userID, true)
	if err != nil && !errors.Is(err, ErrRelationNotFound) {
		return nil, errors.Wrapf(err, "failed to getProgress for userID:%v", userID)
	}
	if userProgress != nil && (userProgress.HideBadges && requestingUserID(ctx) != userID) {
		return nil, ErrHidden
	}

	return userProgress.buildBadgeSummaries(), nil
}

//nolint:revive // .
func (r *repository) getProgress(ctx context.Context, userID string, tolerateOldData bool) (res *progress, err error) {
	if ctx.Err() != nil {
		return nil, errors.Wrap(ctx.Err(), "unexpected deadline")
	}
	sql := `SELECT * FROM badge_progress WHERE user_id = $1`
	if tolerateOldData {
		res, err = storage.Get[progress](ctx, r.db, sql, userID)
	} else {
		res, err = storage.ExecOne[progress](ctx, r.db, sql, userID)
	}

	if res == nil {
		return nil, ErrRelationNotFound
	}

	return res, errors.Wrapf(err, "can't get badge progress for userID:%v", userID)
}

func (r *repository) getStatistics(ctx context.Context, groupType GroupType) (map[Type]float64, error) {
	if ctx.Err() != nil {
		return nil, errors.Wrap(ctx.Err(), "unexpected deadline")
	}
	sql := `SELECT *
				FROM badge_statistics
				WHERE badge_group_type = $1`
	res, err := storage.Select[statistics](ctx, r.db, sql, string(groupType))
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get BADGE_STATISTICS for groupType:%v", groupType)
	}

	return r.calculateUnachievedPercentages(groupType, res), nil
}

//nolint:funlen,revive,gocognit // calculation logic, it is better to keep in one place
func (*repository) calculateUnachievedPercentages(groupType GroupType, res []*statistics) map[Type]float64 {
	allTypes := AllGroups[groupType]
	var totalUsers, totalAchievedBy uint64
	achievedByForEachType, resp := make(map[Type]uint64, cap(res)-1), make(map[Type]float64, cap(res)-1)
	sort.SliceStable(res, func(i, j int) bool {
		return AllTypeOrder[res[i].Type] < AllTypeOrder[res[j].Type]
	})
	for idx, row := range res {
		if row.Type == Type(row.GroupType) {
			totalUsers = row.AchievedBy
		} else {
			// Previous cannot be less than current, trying to fix the data.
			maxNextItem := row.AchievedBy
			for nextItemsIdx := idx; nextItemsIdx < len(res); nextItemsIdx++ {
				if res[nextItemsIdx].AchievedBy > maxNextItem {
					maxNextItem = res[nextItemsIdx].AchievedBy
				}
			}
			achievedByForEachType[row.Type] = uint64(math.Max(float64(row.AchievedBy), float64(maxNextItem)))
			totalAchievedBy += achievedByForEachType[row.Type]
		}
	}
	if totalUsers == 0 {
		return resp
	}
	if totalAchievedBy > totalUsers {
		totalAchievedBy = totalUsers
	}
	for ind, currentBadgeType := range allTypes {
		currentBadgeAchievedBy := math.Min(float64(achievedByForEachType[allTypes[ind]]), float64(totalUsers))
		if currentBadgeType == allTypes[0] {
			resp[currentBadgeType] = percent100 * ((float64(totalAchievedBy) - currentBadgeAchievedBy) / float64(totalUsers))
			if totalAchievedBy < totalUsers {
				resp[currentBadgeType] = percent100 - (percent100 * currentBadgeAchievedBy / float64(totalUsers))
			}

			continue
		}
		usersWhoOwnsPreviousBadge := math.Min(float64(achievedByForEachType[allTypes[ind-1]]), float64(totalUsers))
		usersInProgressWithBadge := usersWhoOwnsPreviousBadge
		usersInProgressWithBadge -= currentBadgeAchievedBy
		resp[currentBadgeType] = percent100 * (usersInProgressWithBadge / float64(totalUsers))
	}

	return resp
}

func (p *progress) buildBadges(groupType GroupType, stats map[Type]float64) []*Badge {
	resp := make([]*Badge, 0, len(AllGroups[groupType]))
	for _, badgeType := range AllGroups[groupType] {
		resp = append(resp, &Badge{
			AchievingRange:              Milestones[badgeType],
			Name:                        AllNames[groupType][badgeType],
			Type:                        badgeType,
			GroupType:                   groupType,
			PercentageOfUsersInProgress: stats[badgeType],
		})
	}
	if p == nil || p.AchievedBadges == nil || len(*p.AchievedBadges) == 0 {
		return resp
	}
	achievedBadges := make(map[Type]bool, len(resp))
	for _, achievedBadge := range *p.AchievedBadges {
		achievedBadges[achievedBadge] = true
	}
	for _, badge := range resp {
		badge.Achieved = achievedBadges[badge.Type]
	}

	return resp
}

func (p *progress) buildBadgeSummaries() []*BadgeSummary { //nolint:gocognit,revive // .
	resp := make([]*BadgeSummary, 0, len(AllGroups))
	for _, groupType := range &GroupsOrderSummaries {
		types := AllGroups[groupType]
		lastAchievedIndex := 0
		if p != nil && p.AchievedBadges != nil {
			for ix, badgeType := range types {
				for _, achievedBadge := range *p.AchievedBadges {
					if badgeType == achievedBadge {
						lastAchievedIndex = ix
					}
				}
			}
		}
		resp = append(resp, &BadgeSummary{
			Name:      AllNames[groupType][types[lastAchievedIndex]],
			GroupType: groupType,
			Index:     uint64(lastAchievedIndex),
			LastIndex: uint64(len(types) - 1),
		})
	}

	return resp
}
