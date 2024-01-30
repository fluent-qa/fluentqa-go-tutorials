package jiracli

// docker run -it --rm ghcr.io/ankitpokhrel/jira-cli:latest
//https://github.com/ankitpokhrel/jira-cli/
import (
	"github.com/ankitpokhrel/jira-cli/api"
)

var DEBUG = true

func JiraClient() {
	api.DefaultClient(DEBUG).ServerInfo()
}
