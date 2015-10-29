package main

import (
	"bytes"
	"encoding/json"
	"google.golang.org/appengine"
	"google.golang.org/appengine/memcache"
	"google.golang.org/appengine/datastore"
	"io"
	"net/http"
)

func serveTemplate(res http.ResponseWriter, req *http.Request, templateName string) {
	memItem := getSession(req)
	ctx := appengine.NewContext(req)
	if len(memItem.Value) > 0 {
		var sd SessionData
		json.Unmarshal(memItem.Value, &sd)
		sd.LoggedIn = true
		q := datastore.NewQuery("Tweet").Filter("UserName =", sd.UserName)
		t := q.Run(ctx)
		count := 0

		var ts []Tweet

		for {
        _, err := t.Next(&ts[count])
        if err == datastore.Done {
                break // No further entities match the query.
        }
        if err != nil {
                break
        }
				count++
			}

		mt := mainTemp {}
		mt.SessionData = sd
		mt.Tweets = ts

		tpl.ExecuteTemplate(res, templateName, mt)
	} else {
		ctx := appengine.NewContext(req)
		i, err := memcache.Get(ctx, templateName)
		if err != nil {
			buf := new(bytes.Buffer)
			writ := io.MultiWriter(res, buf)
			tpl.ExecuteTemplate(writ, templateName, SessionData{})
			memcache.Set(ctx, &memcache.Item{
				Key:   templateName,
				Value: buf.Bytes(),
			})
			return
		}
		io.WriteString(res, string(i.Value)) // we're serving the page from memcache
	}
}
