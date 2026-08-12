package main

import (
	"fmt"
	"net/url"
	"os"
	"sort"
	"strings"

	"github.com/olivedapp/olived/readme/gen"
)

// star

func main() {
	sites := gen.Sites
	sitesTable := buildSitesTable(sites)

	const outPath = "sites.md"
	if err := os.WriteFile(outPath, []byte(sitesTable), 0644); err != nil {
		fmt.Println("write sites.md failed:", err)
		return
	}
	fmt.Println("generated", outPath, "with", len(sites), "sites")

	const readmePath = "../README.md"
	if err := updateREADME(readmePath, sitesTable); err != nil {
		fmt.Println("update README.md failed:", err)
		return
	}
	fmt.Println("updated", readmePath, "Supported Sites section")

	const sitesTxtPath = "sites.txt"
	homeURLs := collectHomeURLs(sites)
	if err := os.WriteFile(sitesTxtPath, []byte(strings.Join(homeURLs, "\n")+"\n"), 0644); err != nil {
		fmt.Println("write sites.txt failed:", err)
		return
	}
	fmt.Println("generated", sitesTxtPath, "with", len(homeURLs), "domains")
}

func buildSitesTable(sites []gen.SiteInfo) string {
	var b strings.Builder
	b.WriteString("| Site ID | Site Name | Live URL Format |\n")
	b.WriteString("| --- | --- | --- |\n")

	for _, site := range sites {
		domains := make([]string, len(site.Domain))
		for i, d := range site.Domain {
			domains[i] = escapeMDCell(d)
		}
		domainCell := strings.Join(domains, "<br>")
		if site.IsNSFW {
			domainCell += " (NSFW)"
		}

		b.WriteString(fmt.Sprintf(
			"| %s | %s | %s |\n",
			escapeMDCell(site.ID),
			escapeMDCell(site.Name),
			domainCell,
		))
	}
	return b.String()
}

func updateREADME(readmePath, sitesTable string) error {
	const startMarker = "<!-- supported-sites:start -->"
	const endMarker = "<!-- supported-sites:end -->"

	content, err := os.ReadFile(readmePath)
	if err != nil {
		return err
	}

	text := string(content)
	start := strings.Index(text, startMarker)
	end := strings.Index(text, endMarker)
	if start == -1 || end == -1 || end < start {
		return fmt.Errorf("README markers not found")
	}

	end += len(endMarker)
	var b strings.Builder
	b.WriteString(text[:start+len(startMarker)])
	b.WriteByte('\n')
	b.WriteString(sitesTable)
	b.WriteString(text[end-len(endMarker):])

	return os.WriteFile(readmePath, []byte(b.String()), 0644)
}

func collectHomeURLs(sites []gen.SiteInfo) []string {
	seen := make(map[string]struct{})
	for _, site := range sites {
		for _, d := range site.Domain {
			home, err := extractHomeURL(d)
			if err != nil {
				fmt.Println("skip invalid domain:", d, err)
				continue
			}
			seen[home] = struct{}{}
		}
	}
	urls := make([]string, 0, len(seen))
	for u := range seen {
		urls = append(urls, u)
	}
	sort.Strings(urls)
	return urls
}

func extractHomeURL(raw string) (string, error) {
	u, err := url.Parse(raw)
	if err != nil {
		return "", err
	}
	if u.Scheme == "" || u.Host == "" {
		return "", fmt.Errorf("missing scheme or host")
	}
	return u.Scheme + "://" + u.Host, nil
}

func escapeMDCell(s string) string {
	return strings.ReplaceAll(s, "|", "\\|")
}
