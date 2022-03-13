package postgresql

import (
	"fmt"
	discogs "github.com/OntoLedgy/domain_ontologies/code/data_models/discogs"
	musicbrainz "github.com/OntoLedgy/domain_ontologies/code/data_models/musicbrainz"
	"github.com/OntoLedgy/storage_interop_services/code/services/database_services/database_i_o_service"
	"github.com/OntoLedgy/storage_interop_services/code/services/database_services/database_i_o_service/object_model/configurations"
	"github.com/OntoLedgy/storage_interop_services/code/services/database_services/database_transactions_service"
	data_model2 "github.com/OntoLedgy/storage_interop_services/code/services/database_services/database_transactions_service/object_model"
	"testing"
)

func TestQuery(t *testing.T) {

	var queryArtists data_model2.Queries = "SELECT * FROM artist WHERE Name = 'Thirdwave'"

	testMusicBrainz(queryArtists)

	testDiscogs(queryArtists)

}

func testDiscogs(
	queryArtists data_model2.Queries) {

	discogs_database_factory := &database_i_o_service.DatabaseFactory{
		"192.168.0.45",
		5432,
		"ladmin",
		"Numark234",
		"discogs",
		"postgres",
		configurations.DbTypePostgresql,
		""}

	discogsDatabase :=
		discogs_database_factory.New()

	discogsDatabase.Connect()

	discogArtists := &[]discogs.Artist{}

	query := data_model2.Queryx{
		queryArtists,
		[]interface{}{},
	}

	offset := 0
	count := 2

	var limitOptions = &data_model2.LimitOptions{
		offset,
		count,
	}

	transactionFactory := &database_transactions_service.DatabaseTransactionFactory{
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

	discogsDatabase.Close()

}

func testMusicBrainz(
	queryArtists data_model2.Queries) {

	//move to config.json
	musicbrainz_database_factory := &database_i_o_service.DatabaseFactory{
		"192.168.0.45",
		5432,
		"musicbrainz",
		"musicbrainz",
		"musicbrainz_db",
		"postgres",
		configurations.DbTypePostgresql,
		""}

	musicbrainzDatabase :=
		musicbrainz_database_factory.New()

	musicbrainzDatabase.Connect()

	musicbrainzArtists := &[]musicbrainz.Artist{}

	query := data_model2.Queryx{
		queryArtists,
		[]interface{}{},
	}

	offset := 0
	count := 2

	var limitOptions = &data_model2.LimitOptions{
		offset,
		count,
	}

	transactionFactory := &database_transactions_service.DatabaseTransactionFactory{
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

	musicbrainzDatabase.Close()

}
