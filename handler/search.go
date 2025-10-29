package handler

import (
	"errors"
	"fmt"
	"net/url"
)

func imFeelingLuck(phrase string) (user, project string, err error) {
	phrase += " site:github.com"
	// try dgg
	v := url.Values{}
	v.Set("q", "! " /*I'm feeling lucky*/ +phrase)
	if user, project, err := captureRepoLocation(("https://html.duckduckgo.com/html?" + v.Encode())); err == nil {
		return user, project, nil
	}
	// try google
	v = url.Values{}
	v.Set("btnI", "") //I'm feeling lucky
	v.Set("q", phrase)
	if user, project, err := captureRepoLocation(("https://www.google.com/search?" + v.Encode())); err == nil {
		return user, project, nil
	}
	return "", "", errors.New("not found")
}

// uses im feeling lucky and grabs the "Location"
// header from the 302, which contains the github repo
func captureRepoLocation(url string) (user, project string, err error) {

	//roundtripper doesn't follow redirects
	resp, err := httpGet(url)
	if err != nil {
		return "", "", err
	}
	defer resp.Body.Close()
	//assume redirection
	if resp.StatusCode/100 != 3 {
		return "", "", fmt.Errorf("non-redirect response: %d", resp.StatusCode)
	}
	//extract Location header URL
	loc := resp.Header.Get("Location")
	m := searchGithubRe.FindStringSubmatch(loc)
	if len(m) == 0 {
		return "", "", fmt.Errorf("github url not found in redirect: %s", loc)
	}
	user = m[1]
	project = m[2]
	return user, project, nil
}
