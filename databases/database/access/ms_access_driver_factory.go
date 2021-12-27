package access

//import "database_manager/database"

func Create(database_filename string) *MsAccessDatabaseDrivers {

	ms_access_driver := new(MsAccessDatabaseDrivers)

	ms_access_driver.OleDb12ConnectionString = OleDb12ConnectionString_Prefix + database_filename

	return ms_access_driver

}
