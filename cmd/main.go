package main

import (
	"github.com/adzpm/jumoreski/internal/server"
	"github.com/adzpm/jumoreski/internal/storage"
)

const (
	token        = `vk1.a.11LJ7zU6iM252qTdfIK0Gs2aUDpo7Vn_SxcCM2gwxY1QLSTCVgIqspBTflZPh0hQAA-r2D4gzX2SE8XFOSIVZqyPma-DVbexRk25F_QOpxITLWYxkgLVFVH4gcd_gANsL58okBTYhM7D_dtZ9d-9obSxVGSO4cgphVlMLw3ya8r0DND7-PglQ1qMOUxOBqtu8ANIPhJI-o8aHbcQKdkVkg`
	groupDomain  = `jumoreski`
	databasePath = "/Users/dzpm/projects/my/jumoreski/assets/database.db"
	host         = "127.0.0.1"
	port         = "8080"
	frontPath    = "/Users/dzpm/projects/my/jumoreski/web/"
)

func main() {
	st, err := storage.NewStorage(&storage.Config{
		DatabasePath: databasePath,
	})

	if err != nil {
		panic(err)
	}

	srv := server.NewServer(
		&server.Config{
			Host:      host,
			Port:      port,
			FrontPath: frontPath,
		},
		st,
	)

	if err = srv.Start(); err != nil {
		panic(err)
	}

	//prs := parser.NewParser(
	//	&parser.Config{
	//		AccessToken: token,
	//		GroupDomain: groupDomain,
	//	},
	//	st,
	//)
	//
	//if err = prs.Start(); err != nil {
	//	log.Print("parsing error: ", err)
	//}

}
