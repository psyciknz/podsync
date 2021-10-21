package media

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	"github.com/mxpv/podsync/pkg/config"
)

const (
	DefaultDownloadTimeout = 10 * time.Minute
	UpdatePeriod           = 24 * time.Hour
)

type Plex struct {
	url       	string
	plextoken	string
	library		int
}

var (
	ErrTooManyRequests = errors.New(http.StatusText(http.StatusTooManyRequests))
)


func New(ctx context.Context, cfg config.MediaServer) (*Plex, error)  {
	
	log.Debugf("creating new plex server")

	url = cfg.Url
	plextoken = cfg.PlexToken
	library = cfg.PlexLibrary
	
	plex := &Plex{
		url:    url,
		plextoken: plextoken,
		library:	library,
	}

	return plex, nil
}


func Updatemediaserver(ctx context.Context) error {

	refresh =  fmt.Sprintf("%s/library/sections/%d/refresh?X-Plex-Token=%s", url,library,plextoken)
	emptytrash =  fmt.Sprintf("%s/library/sections/%d/emptyTrash?X-Plex-Token=%s", url,library,plextoken)

	log.Debug("Updating media server")
	//curl -v http://<server>:32400/library/sections/<library>/refresh?X-Plex-Token=<token>
	//curl -X PUT "http://<server>:32400/library/sections/<library>/emptyTrash?X-Plex-Token=<token>"
	
	resp, err := http.Get(refresh)
	if err != nil {
   		log.Fatalln(err)
	}
	log.Info(fmt.Sprintf("Updated library %d: result: %s", library,resp.Status))
	resp2, err2 := http.Get(emptytrash)
	if err2 != nil {
		log.Fatalln(err2)
 	}
	log.Info(fmt.Sprintf("Emptied Trash %d result: %s", library,resp2.Status))
	
	return nil
}


