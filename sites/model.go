package sites

import (
	"github.com/marekparafianowicz/go-server/config"
)

type site struct {
	ID  int    `json:"id"`
	URL string `json:"url"`
}

func allSites() ([]site, error) {
	rows, err := config.DB.Query("SELECT * FROM sites")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	sites := make([]site, 0)
	for rows.Next() {
		site := site{}
		err := rows.Scan(&site.ID, &site.URL)
		if err != nil {
			return nil, err
		}
		sites = append(sites, site)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return sites, nil
}

func findSite(id string) (site, error) {
	row := config.DB.QueryRow("SELECT * FROM sites WHERE id = $1", id)
	site := site{}
	err := row.Scan(&site.ID, &site.URL)
	return site, err
}
