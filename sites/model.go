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

func createSite(url string) (site, error) {
	st := site{}

	row := config.DB.QueryRow("SELECT * FROM sites WHERE url = $1", url)
	err := row.Scan(&st.ID, &st.URL)
	if st.URL != "" {
		return site{}, nil // Here implement our own error
	}

	st.URL = url
	err = config.DB.QueryRow("INSERT INTO sites (url) VALUES ($1) RETURNING id", st.URL).Scan(&st.ID)
	if err != nil {
		return st, err
	}
	return st, nil
}

func updateSite(id, url string) (site, error) {
	query := "UPDATE sites SET url = $2	WHERE id = $1 RETURNING id, url"
	st := site{}
	err := config.DB.QueryRow(query, id, url).Scan(&st.ID, &st.URL)
	if err != nil {
		panic(err)
	}
	if st.URL != url {
		return site{}, nil // Here implement our own error
	}
	return st, nil
}

func deleteSite(id string) error {
	row := config.DB.QueryRow("SELECT * FROM sites WHERE id = $1", id)
	site := site{}
	err := row.Scan(&site.ID, &site.URL)
	if err != nil {
		return err
	}

	query := "DELETE FROM sites WHERE id = $1"
	_, err = config.DB.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
