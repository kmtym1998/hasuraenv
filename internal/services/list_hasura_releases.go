package services

import (
	"net/http"

	"github.com/kmtym1998/hasuraenv/internal/httpc"
)

func ListHasuraReleases() {
	httpc.SendRequest(
		http.MethodGet,
		"https://api.github.com/repos/hasura/graphql-engine/releases",
		nil,
		map[string]string{
			"Content-Type": "application/json",
		},
	)
}
