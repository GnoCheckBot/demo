package config

import (
	"github-bot/internal/client"
	c "github-bot/internal/conditions"
	r "github-bot/internal/requirements"
)

type Teams []string

// Automatic check that will be performed by the bot.
type AutomaticCheck struct {
	Description string
	If          c.Condition   // If the condition is met, the rule is displayed and the requirement is executed.
	Then        r.Requirement // If the requirement is satisfied, the check passes.
}

// Manual check that will be performed by users.
type ManualCheck struct {
	Description string
	If          c.Condition // If the condition is met, a checkbox will be displayed on bot comment.
	Teams       Teams       // Members of these teams can check the checkbox to make the check pass.
}

// This function returns the configuration of the bot consisting of automatic and manual checks
// in which the GitHub client is injected.
func Config(gh *client.GitHub) ([]AutomaticCheck, []ManualCheck) {
	auto := []AutomaticCheck{
		{
			Description: "Maintainers must be able to edit this pull request",
			If:          c.Always(),
			Then:        r.MaintainerCanModify(),
		},
		{
			Description: "The pull request head branch must be up-to-date with its base",
			If:          c.Always(),
			Then:        r.UpToDateWith(gh, r.PR_BASE),
		},
		// {
		// 	Description: "Changes to 'docs' folder must be reviewed/authored by at least one devrel and one tech-staff",
		// 	If:          c.FileChanged(gh, "^docs/"),
		// 	Then: r.Or(
		// 		r.And(
		// 			r.AuthorInTeam(gh, "devrels"),
		// 			r.ReviewByTeamMembers(gh, "tech-staff", 1),
		// 		),
		// 		r.And(
		// 			r.AuthorInTeam(gh, "tech-staff"),
		// 			r.ReviewByTeamMembers(gh, "devrels", 1),
		// 		),
		// 	),
		// },
	}

	manual := []ManualCheck{
		// {
		// 	Description: "The pull request description provides enough details",
		// 	If:          c.Not(c.AuthorInTeam(gh, "core-contributors")),
		// 	Teams:       Teams{"core-contributors"},
		// },
		{
			Description: "Determine if infra needs to be updated before merging",
			If: c.And(
				c.BaseBranch("master"),
				c.Or(
					c.FileChanged(gh, `Dockerfile`),
					c.FileChanged(gh, `^misc/deployments`),
					c.FileChanged(gh, `^misc/docker-`),
					c.FileChanged(gh, `^.github/workflows/releaser.*\.yml$`),
					c.FileChanged(gh, `^.github/workflows/portal-loop\.yml$`),
				),
			),
			// Teams: Teams{"devops"},
		},
	}

	// Check for duplicates in manual rule descriptions (needs to be unique for the bot operations).
	unique := make(map[string]struct{})
	for _, rule := range manual {
		if _, exists := unique[rule.Description]; exists {
			gh.Logger.Fatalf("Manual rule descriptions must be unique (duplicate: %s)", rule.Description)
		}
		unique[rule.Description] = struct{}{}
	}

	return auto, manual
}
