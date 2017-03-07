/*
Package cyberark wraps the CyberArk Vault API to retrieve user names and passwords.
It is not affiliated with, nor endorsed by, CayberArk Software Ltd.

Basics

You must have a working CyberArk installation to talk to, and you must have credentials
in the vault that you're authorized to retrieve.

The pattern is to create a CypberArk client, get one of its services (currently, only
the GetPassword service is available), set its parameters, and call its Do() function.

Example (error checking omitted for brevity):

				// Host is required
				client, _ := cyberark.NewClient(
					cyberark.SetHost("cyberark.example.com"),
				)

				// App ID is required
				ret, _ := client.GetPassword().
					AppID("my_app_id").
					Safe("my_safe").
					Object("LDAP").
					Do()

				user := ret.UserName
				password := ret.Content

*/
package cyberark
