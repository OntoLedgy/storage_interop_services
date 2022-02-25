package databases

import (
	"fmt"
	discogs "github.com/OntoLedgy/domain_ontologies/code/data_models/discogs"
	musicbrainz "github.com/OntoLedgy/domain_ontologies/code/data_models/musicbrainz"
	"github.com/OntoLedgy/storage_interop_services/code/services/database_transactions"
	"github.com/OntoLedgy/storage_interop_services/code/services/database_transactions/data_model"
	"github.com/OntoLedgy/storage_interop_services/code/services/databases"
	"testing"
)

func TestQuery(t *testing.T) {

	var queryArtists data_model.Queries = "SELECT * FROM artist WHERE Name = 'Thirdwave'"

	testMusicBrainz(queryArtists)

	testDiscogs(queryArtists)

}

func testDiscogs(queryArtists data_model.Queries) {

	discogs_database_factory := &databases.DatabaseFactory{
		"192.168.0.45",
		5432,
		"ladmin",
		"Numark234",
		"discogs",
		"postgres",
	}

	discogsDatabase :=
		discogs_database_factory.Create()

	discogsDatabase.ConnectDatabase()

	discogArtists := &[]discogs.Artist{}

	query := data_model.Queryx{
		queryArtists,
		[]interface{}{},
	}

	offset := 0
	count := 2

	var limitOptions = &data_model.LimitOptions{
		offset,
		count,
	}

	transactionFactory := &database_transactions.DatabaseTransactionFactory{
		discogsDatabase,
	}

	transaction := transactionFactory.Create()

	transaction.Selectx(
		discogArtists,
		query,
		limitOptions)

	fmt.Printf(
		"found %v records in discogs\n",
		len(*discogArtists))

	discogsDatabase.CloseDatabase()

}

func testMusicBrainz(queryArtists data_model.Queries) *databases.DatabaseFactory {
	//move to config.json
	musicbrainz_database_factory := &databases.DatabaseFactory{
		"192.168.0.45",
		5432,
		"musicbrainz",
		"musicbrainz",
		"musicbrainz_db",
		"postgres",
	}

	musicbrainzDatabase :=
		musicbrainz_database_factory.Create()

	musicbrainzDatabase.ConnectDatabase()

	musicbrainzArtists := &[]musicbrainz.Artist{}

	query := data_model.Queryx{
		queryArtists,
		[]interface{}{},
	}

	offset := 0
	count := 2

	var limitOptions = &data_model.LimitOptions{
		offset,
		count,
	}

	transactionFactory := &database_transactions.DatabaseTransactionFactory{
		musicbrainzDatabase,
	}

	transaction := transactionFactory.Create()

	transaction.Selectx(
		musicbrainzArtists,
		query,
		limitOptions)

	fmt.Printf(
		"found %v records in musicbrainz \n",
		len(*musicbrainzArtists))

	musicbrainzDatabase.CloseDatabase()
	return musicbrainz_database_factory
}
