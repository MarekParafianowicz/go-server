package sites

import (
	"fmt"

	"github.com/marekparafianowicz/go-server/config"
)

type site struct {
	ID  int
	URL string
}

func allSites() ([]site, error) {
	rows, err := config.DB.Query("SELECT * FROM sites")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	sites := make([]site, 0)
	for rows.Next() {
		site := site{}
		err := rows.Scan(&site.ID, &site.URL)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		sites = append(sites, site)
	}
	if err = rows.Err(); err != nil {
		fmt.Println(err)
		return nil, err
	}
	return sites, nil
}
