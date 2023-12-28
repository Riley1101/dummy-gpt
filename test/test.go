
//r.Group(func(r chi.Router) {
//	auth := common.Auth{}
//	common.InitAuth(&auth)
//	r.Use(jwtauth.Verifier(auth.Token))
//	r.Use(jwtauth.Authenticator(auth.Token))
//	r.Get("/api/admin", func(w http.ResponseWriter, r *http.Request) {
//		_, claims, _ := jwtauth.FromContext(r.Context())
//		w.Write([]byte(fmt.Sprintf("protected area. hi %v", claims["user_id"])))
//	})
//})

//table := common.Table{
//	Name: "users",
//	Fields: []common.Field{
//		{
//			Name:     "id",
//			Datatype: "integer",
//		},
//		{
//			Name:     "name",
//			Datatype: "varchar(255)",
//		},
//		{
//			Name:     "email",
//			Datatype: "varchar(255)",
//		},
//	},
//}

//r.Get("/", func(w http.ResponseWriter, r *http.Request) {
//	const userid = "Arkar"
//	formBody := common.SchemaForm{
//		Schema: r.FormValue("schema"),
//	}
//	dbSchema := common.DbSchema{
//		Name: "users",
//	}
//	dbSchema.ParseSchema(formBody.Schema)
//	dbSchema.DescribeSchema()
//	dbSchema.GenerateQuery()
//	files := []string{
//		"templates/index.tmpl",
//		"templates/base.tmpl",
//	}
//	ts, err := template.ParseFiles(files...)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//	err = ts.ExecuteTemplate(w, "base", table)
//})
