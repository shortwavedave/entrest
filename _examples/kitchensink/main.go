package main

import (
	"context"
	"database/sql"
	"net/http"

	"entgo.io/ent/dialect/sql/schema"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/lrstanley/entrest/_examples/kitchensink/internal/database/ent"
	"github.com/lrstanley/entrest/_examples/kitchensink/internal/database/ent/pet"
	"github.com/lrstanley/entrest/_examples/kitchensink/internal/database/ent/rest"
	_ "github.com/lrstanley/entrest/_examples/kitchensink/internal/database/ent/runtime" // Required by ent.
	"github.com/lrstanley/entrest/_examples/kitchensink/internal/database/ent/user"
	"modernc.org/sqlite"
)

func main() {
	sql.Register("sqlite3", &sqlite.Driver{})
	db, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_pragma=foreign_keys(1)&_busy_timeout=15")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	ctx := context.Background()

	err = db.Schema.Create(
		ctx,
		schema.WithDropColumn(true),
		schema.WithDropIndex(true),
		schema.WithGlobalUniqueID(true),
		schema.WithForeignKeys(true),
	)
	if err != nil {
		panic(err)
	}

	john := db.User.Create().
		SetName("john").
		SetEmail(gofakeit.Email()).
		SetType(user.TypeUser).
		SetEnabled(true).
		SetPasswordHashed(gofakeit.Password(true, true, true, true, true, 15)). // Not actually used.
		SaveX(ctx)

	for range 100 {
		db.User.Create().
			SetName(gofakeit.FirstName()).
			SetEmail(gofakeit.Email()).
			SetType(user.TypeUser).
			SetEnabled(true).
			SetPasswordHashed(gofakeit.Password(true, true, true, true, true, 15)). // Not actually used.
			ExecX(ctx)
	}

	oreo := db.Pet.Create().
		SetName("Riley").
		AddFollowedBy(john).
		SetAge(3).
		SetType(pet.TypeDog).
		SaveX(ctx)

	for range 100 {
		types := []pet.Type{
			pet.TypeDog,
			pet.TypeCat,
			pet.TypeBird,
			pet.TypeFish,
			pet.TypeAmphibian,
			pet.TypeReptile,
			pet.TypeOther,
		}

		db.Pet.Create().
			SetName(gofakeit.PetName()).
			SetOwner(john).
			AddFriends(oreo).
			AddFollowedBy(john).
			SetAge(gofakeit.Number(1, 15)).
			SetType(types[gofakeit.Number(0, len(types)-1)]).
			ExecX(ctx)
	}

	db.Dog.Create().SetName("George").SaveX(ctx)

	srv, err := rest.NewServer(db, &rest.ServerConfig{})
	if err != nil {
		panic(err)
	}

	// Example of using net/http serve-mux:
	//	mux := http.NewServeMux()
	//	mux.Handle("/", srv.Handler())
	//	http.ListenAndServe(":8080", mux)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Mount("/debug", middleware.Profiler())
	r.Mount("/", srv.Handler())
	http.ListenAndServe(":8080", r) //nolint:all
}
