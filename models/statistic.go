// Copyright 2021 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package models

import (
	"code.gitea.io/gitea/models/db"
	"code.gitea.io/gitea/models/login"
)

// Statistic contains the database statistics
type Statistic struct {
	Counter struct {
		User, Org, PublicKey,
		Repo, Watch, Star, Action, Access,
		Issue, IssueClosed, IssueOpen,
		Comment, Oauth, Follow,
		Mirror, Release, LoginSource, Webhook,
		Milestone, Label, HookTask,
		Team, UpdateTask, Project,
		ProjectBoard, Attachment int64
	}
}

// GetStatistic returns the database statistics
func GetStatistic() (stats Statistic) {
	e := db.GetEngine(db.DefaultContext)
	stats.Counter.User = CountUsers()
	stats.Counter.Org = CountOrganizations()
	stats.Counter.PublicKey, _ = e.Count(new(PublicKey))
	stats.Counter.Repo = CountRepositories(true)
	stats.Counter.Watch, _ = e.Count(new(Watch))
	stats.Counter.Star, _ = e.Count(new(Star))
	stats.Counter.Action, _ = e.Count(new(Action))
	stats.Counter.Access, _ = e.Count(new(Access))

	type IssueCount struct {
		Count    int64
		IsClosed bool
	}
	issueCounts := []IssueCount{}

	_ = e.Select("COUNT(*) AS count, is_closed").Table("issue").GroupBy("is_closed").Find(&issueCounts)
	for _, c := range issueCounts {
		if c.IsClosed {
			stats.Counter.IssueClosed = c.Count
		} else {
			stats.Counter.IssueOpen = c.Count
		}
	}

	stats.Counter.Issue = stats.Counter.IssueClosed + stats.Counter.IssueOpen

	stats.Counter.Comment, _ = e.Count(new(Comment))
	stats.Counter.Oauth = 0
	stats.Counter.Follow, _ = e.Count(new(Follow))
	stats.Counter.Mirror, _ = e.Count(new(Mirror))
	stats.Counter.Release, _ = e.Count(new(Release))
	stats.Counter.LoginSource = login.CountSources()
	stats.Counter.Webhook, _ = e.Count(new(Webhook))
	stats.Counter.Milestone, _ = e.Count(new(Milestone))
	stats.Counter.Label, _ = e.Count(new(Label))
	stats.Counter.HookTask, _ = e.Count(new(HookTask))
	stats.Counter.Team, _ = e.Count(new(Team))
	stats.Counter.Attachment, _ = e.Count(new(Attachment))
	stats.Counter.Project, _ = e.Count(new(Project))
	stats.Counter.ProjectBoard, _ = e.Count(new(ProjectBoard))
	return
}
