package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	workloadservice "github.com/yametech/fuxi/pkg/service/workload"
	"io"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	watch "k8s.io/apimachinery/pkg/watch"
	"log"
	"net/url"
	"strings"
	"sync"
)

type Event struct {
	Type   watch.EventType `json:"type"`
	Object runtime.Object  `json:"object"`
	Url    string          `json:"url"`
	Status int             `json:"status"`
}

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

func parseForApiUrl(apiUrl string) (gvr *schema.GroupVersionResource, namespace string, resourceVersion string, err error) {
	u, err := url.Parse(apiUrl)
	if err != nil {
		return nil, "", "", err
	}
	gvr = &schema.GroupVersionResource{}
	paths := trimPrefixSuffixSpace(strings.Split(u.Path, "/"))
	if len(paths) == 0 {
		return nil, "", "", fmt.Errorf("parse url error")
	}
	switch paths[0] {
	case "api":
		gvr.Group = ""
		gvr.Version = paths[1]
		gvr.Resource = paths[2]
		// /api/v1/watch/namespaces/{namespaces}/pods
		if len(paths) >= 6 && paths[3] == "namespaces" {
			namespace = paths[4]
		}
	case "apis":
		if paths[1] == "crd" {
			// /apis/crd/nuwa.nip.io/v1/waters?watch=1&resourceVersion=5738162
			remove(paths, 1)
		}
		gvr.Group = paths[1]
		gvr.Version = paths[2]
		gvr.Resource = paths[3]
		// /apis/apps/v1/watch/namespaces/{namespaces}/deployments
		if len(paths) >= 7 && paths[4] == "namespaces" {
			namespace = paths[5]
		}

	default:
		return nil, "", "", fmt.Errorf("parser url (%s) error", apiUrl)
	}

	m, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		return nil, "", "", err
	}
	if len(m["resourceVersion"]) > 0 {
		resourceVersion = m["resourceVersion"][0]
	}

	return
}

func trimPrefixSuffixSpace(slice []string) []string {
	for i, s := range slice {
		if s == "" {
			slice = slice[i+1:]
		}
	}
	return slice
}

func listenByApis(event *workloadservice.Generic, g *gin.Context, eventChan chan Event, closed chan struct{}) {
	defer close(eventChan)
	apis := g.QueryArray("api")
	wg := sync.WaitGroup{}
	wg.Add(len(apis))
	for _, api := range apis {
		gvr, ns, rv, err := parseForApiUrl(api)
		if err != nil {
			log.Printf("parser api url error %s", api)
			return
		}
		event.SetGroupVersionResource(*gvr)
		k8sWatchChan, err := event.Watch(ns, rv, 0, nil)
		if err != nil {
			log.Printf("watch for gvr: %s stream error: %s for api request %s \r\n", gvr, err, api)
			continue
		}
		go func() {
			defer wg.Done()
			for {
				select {
				case _, ok := <-closed:
					if !ok {
						return
					}
				case item, ok := <-k8sWatchChan:
					if !ok {
						return
					}
					eventChan <- Event{
						Type:   item.Type,
						Object: item.Object,
					}
				}
			}
		}()
	}
	wg.Wait()
}

// watchStream watch api request resource group and the version
// after server timeout then close send closed event to client side server watcher close
func (w *WorkloadsAPI) WatchStream(g *gin.Context) {
	eventChan := make(chan Event, 32)
	closed := make(chan struct{})
	go listenByApis(w.generic, g, eventChan, closed)

	streamEndEvent := Event{
		Type:   watch.EventType("STREAM_END"),
		Url:    g.Request.URL.String(),
		Status: 410,
	}
	g.Stream(func(w io.Writer) bool {
		select {
		case <-g.Writer.CloseNotify():
			close(closed)
			g.SSEvent("", streamEndEvent)
			return false
		case event, ok := <-eventChan:
			if !ok {
				g.SSEvent("", streamEndEvent)
				return false
			}
			g.SSEvent("", event)
		}
		return true
	})
}
