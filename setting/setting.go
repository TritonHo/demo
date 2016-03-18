//a centralized location to store the configration from the environment variable

package setting

//a centralized location containing all the configuration needed by the system.
//so that it would avoid collusion of the config parameter name
const (
	DB_HOST     string = `DB_HOST`
	DB_USERNAME string = `DB_USERNAME`
	DB_NAME     string = `DB_NAME`
	DB_PASSWORD string = `DB_PASSWORD`
	DB_PORT     string = `DB_PORT`

	DB_MAX_IDLE_CONN string = `DB_MAX_IDLE_CONN`
	DB_MAX_OPEN_CONN string = `DB_MAX_OPEN_CONN`

	JWT_RSA_KEY_LOCATION     string = `JWT_RSA_KEY_LOCATION`
	JWT_OLD_RSA_KEY_LOCATION string = `JWT_OLD_RSA_KEY_LOCATION`

	//measured in minute, the lifetime of the issued jwt token
	JWT_TOKEN_LIFETIME string = `JWT_TOKEN_LIFETIME`

	DEFAULT_RECORD_PER_PAGE string = `DEFAULT_RECORD_PER_PAGE`
)
