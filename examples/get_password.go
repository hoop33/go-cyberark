package main

import (
	"flag"
	"log"

	cyberark "github.com/hoop33/go-cyberark"
)

func main() {
	var (
		host                 = flag.String("host", "", "CyberArk host")
		appID                = flag.String("appID", "", "App ID")
		address              = flag.String("address", "", "Address")
		database             = flag.String("database", "", "Database")
		folder               = flag.String("folder", "", "Folder")
		object               = flag.String("object", "", "Object")
		policyID             = flag.String("policyID", "", "Policy ID")
		query                = flag.String("query", "", "Query")
		queryFormat          = flag.String("queryFormat", "", "Query Format")
		reason               = flag.String("reason", "", "Reason")
		safe                 = flag.String("safe", "", "Safe")
		timeout              = flag.Int("timeout", 30, "Timeout")
		userName             = flag.String("userName", "", "User Name")
		skipCertVerification = flag.Bool("insecure", false, "Skip Cert Verification")
	)
	flag.Parse()

	if *host == "" {
		log.Fatal("missing host")
	}
	if *appID == "" {
		log.Fatal("missing appID")
	}

	client, err := cyberark.NewClient(
		cyberark.SetHost(*host),
		cyberark.SetSkipCertVerification(*skipCertVerification),
	)
	if err != nil {
		log.Fatal(err.Error())
	}

	ret, err := client.GetPassword().
		AppID(*appID).
		Address(*address).
		Database(*database).
		Folder(*folder).
		Object(*object).
		PolicyID(*policyID).
		Query(*query).
		QueryFormat(*queryFormat).
		Reason(*reason).
		Safe(*safe).
		Timeout(*timeout).
		UserName(*userName).
		Do()
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Printf("HTTP %d\n", ret.StatusCode)

	if ret.ErrorCode != "" {
		log.Fatalf("Error! (%s): %s", ret.ErrorCode, ret.ErrorMsg)
	}
	log.Printf("User: %s; Password: %s", ret.UserName, ret.Content)
}
