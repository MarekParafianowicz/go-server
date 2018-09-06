package repository

import (
	"database/sql"

	_ "github.com/lib/pq"
)

// PostgresRepository is a Postgres data base which implements the Repository interface
type PostgresRepository struct {
	connection *sql.DB
}

// New creates an instance of Repository with the given driver and dataSource names
func New(driverName, dataSourceName string) (*PostgresRepository, error) {
	connection, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}

	return &PostgresRepository{connection: connection}, nil
}

// AllSites returns all sites form DB
func (p *PostgresRepository) AllSites() ([]Site, error) {
	rows, err := p.connection.Query("SELECT * FROM sites")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	sites := make([]Site, 0)
	for rows.Next() {
		site := Site{}
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

// FindSite queries single site form DB
func (p *PostgresRepository) FindSite(id string) (Site, error) {
	row := p.connection.QueryRow("SELECT * FROM sites WHERE id = $1", id)
	site := Site{}
	err := row.Scan(&site.ID, &site.URL)
	return site, err
}

// CreateSite inserts new site with given url into database
func (p *PostgresRepository) CreateSite(url string) (Site, error) {
	st := Site{}

	row := p.connection.QueryRow("SELECT * FROM sites WHERE url = $1", url)
	err := row.Scan(&st.ID, &st.URL)
	if st.URL != "" {
		return Site{}, nil // Here implement our own error
	}

	st.URL = url
	err = p.connection.QueryRow("INSERT INTO sites (url) VALUES ($1) RETURNING id", st.URL).Scan(&st.ID)
	if err != nil {
		return st, err
	}
	return st, nil
}

// UpdateSite finds by id and edits site's url
func (p *PostgresRepository) UpdateSite(id, url string) (Site, error) {
	query := "UPDATE sites SET url = $2	WHERE id = $1 RETURNING id, url"
	st := Site{}
	err := p.connection.QueryRow(query, id, url).Scan(&st.ID, &st.URL)
	if err != nil {
		panic(err)
	}
	if st.URL != url {
		return Site{}, nil // Here implement our own error
	}
	return st, nil
}

// DeleteSite destroys in database site found by id
func (p *PostgresRepository) DeleteSite(id string) error {
	row := p.connection.QueryRow("SELECT * FROM sites WHERE id = $1", id)
	site := Site{}
	err := row.Scan(&site.ID, &site.URL)
	if err != nil {
		return err
	}

	query := "DELETE FROM sites WHERE id = $1"
	_, err = p.connection.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
