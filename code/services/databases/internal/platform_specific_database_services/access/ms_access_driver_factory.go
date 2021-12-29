package access

//import "database_manager/platform_specific_database_services"

func Create(database_filename string) *MsAccessDatabaseDrivers {

	ms_access_driver := new(MsAccessDatabaseDrivers)

	ms_access_driver.OleDb12ConnectionString = OleDb12ConnectionString_Prefix + database_filename

	return ms_access_driver

}
