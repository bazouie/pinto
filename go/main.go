// Note, for this to work at least 1 commit is needed, so you if you use this
// after creating a repository you might want to make sure you set `AutoInit` to
// `true`. It could be more!
package main

import (
	"flag"
)

var (
	sourceOwner   = flag.String("source-owner", "", "Name of the owner (user or org) of the repo to create the commit in.")
	sourceRepo    = flag.String("source-repo", "", "Name of repo to create the commit in.")
	commitMessage = flag.String("commit-message", "", "Content of the commit message.")
	commitBranch  = flag.String("commit-branch", "", "Name of branch to create the commit in. If it does not already exists, it will be created using the `base-branch` parameter")
	baseBranch    = flag.String("base-branch", "master", "Name of branch to create the `commit-branch` from.")
	prRepoOwner   = flag.String("merge-repo-owner", "", "Name of the owner (user or org) of the repo to create the PR against. If not specified, the value of the `-source-owner` flag will be used.")
	prRepo        = flag.String("merge-repo", "", "Name of repo to create the PR against. If not specified, the value of the `-source-repo` flag will be used.")
	prBranch      = flag.String("merge-branch", "master", "Name of branch to create the PR against (the one you want to merge your branch in via the PR).")
	prSubject     = flag.String("pr-title", "", "Title of the pull request. If not specified, no pull request will be created.")
	prDescription = flag.String("pr-text", "", "Text to put in the description of the pull request.")
	sourceFiles   = flag.String("files", "", `Comma-separated list of files to commit and their location.
The local file is separated by its target location by a semi-colon.
If the file should be in the same location with the same name, you can just put the file name and omit the repetition.
Example: README.md,main.go:github/examples/commitpr/main.go`)
	authorName  = flag.String("author-name", "", "Name of the author of the commit.")
	authorEmail = flag.String("author-email", "", "Email of the author of the commit.")
)
