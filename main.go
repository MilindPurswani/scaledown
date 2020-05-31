package main

import (
	"bufio"
	"fmt"
	"log"
	"net/url"
	"os"
)

type urlList []string
type urlElement string

func main() {
	// url List from the user input
	urls := urlList{}
	// Unique url list
	uu := urlList{}
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		urls = append(urls, sc.Text())
	}
	if err := sc.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "failed to read input from file %s\n", err)
	}
	for _, u := range urls {
		u, err := url.Parse(u)
		if err != nil {
			// Invalid character, we can't do anything so just ignore.
			//fmt.Fprintf(os.Stderr, "Invalid Character encountered at %s\n", err)
			continue
		}
		// escaped urls
		eu := u.Scheme + "://" + u.Hostname() + u.RequestURI()
		_, found := find(uu, eu)
		if !found {
			uu = append(uu, eu)
			fmt.Println(eu)
		} else {
			// TODO: call checkParams()

		}

	}

}

//TODO: Implement checking for already existing parameters and append them to one that is non existent
func (eu urlElement) checkParams(u string) bool {
	fmt.Println(eu)
	return false
}

// Looks for a url eu in urlList ul.
// if escaped url/current url matches any element in ul,
// then return position and true
func find(ul urlList, eu string) (int, bool) {
	cu, err := url.Parse((string(eu)))
	if err != nil {
		//Might have parsing issues, we can't do anything so ignore.
		return -1, true
	}
	// Scan for common urls only for last 10 occurances in newurls.
	// A small trade-off to gain performance.
	// Should be enough. Just a small Jugaad :P
	count := 10
	if len(ul) < count {
		count = len(ul)
	}
	for i := len(ul) - count; i < len(ul); i++ {
		//for i, u := range ul {
		u, err := url.Parse(ul[i])
		if err != nil {
			log.Fatal(err)
		}
		if u.EscapedPath() == cu.EscapedPath() {
			return i, true
		}
	}
	return -1, false
}
