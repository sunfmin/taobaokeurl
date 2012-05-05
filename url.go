package taobaokeurl

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"regexp"
)

func ResolveTaobaokeUrl(taobaokeurl string) (detailurl string, numiid string) {
	client := &http.Client{
		CheckRedirect: catchNumiid,
	}
	req, _ := http.NewRequest("GET", taobaokeurl, nil)
	req.Header.Set("Content-Type", "text/html")
	req.Header.Set("User-Agent", "Googlebot/2.1 (+http://www.googlebot.com/bot.html)")
	req.Header.Set("Referer", "http://s.click.taobao.com/t_js?tu="+url.QueryEscape(taobaokeurl))

	_, err := client.Do(req)

	if err != nil {
		urlerr, ok := err.(*url.Error)
		if ok {
			foundiid, ok1 := urlerr.Err.(*FoundNumiid)
			if ok1 {
				detailurl = foundiid.URL
				numiid = foundiid.Numiid
				return
			}

		}
	}

	if err != nil {
		log.Println(err)
	}

	return
}

var regexpTaobaoItem = regexp.MustCompile(`item\.htm[\?\&]*id=(\d+)`)

type FoundNumiid struct {
	Numiid string
	URL    string
}

func (f *FoundNumiid) Error() string {
	return fmt.Sprintf("found %s", f.Numiid)
}

func catchNumiid(req *http.Request, via []*http.Request) error {
	found := regexpTaobaoItem.FindStringSubmatch(req.URL.String())
	if len(found) > 0 {
		return &FoundNumiid{
			Numiid: found[1],
			URL:    req.URL.String(),
		}
	}

	if len(via) >= 10 {
		return errors.New("stopped after 10 redirects")
	}

	return nil
}
