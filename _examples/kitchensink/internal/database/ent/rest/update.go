// Code generated by ent, DO NOT EDIT.

package rest

import (
	"context"
	"time"

	github "github.com/google/go-github/v63/github"
	"github.com/lrstanley/entrest/_examples/kitchensink/internal/database/ent"
	"github.com/lrstanley/entrest/_examples/kitchensink/internal/database/ent/category"
	"github.com/lrstanley/entrest/_examples/kitchensink/internal/database/ent/friendship"
	"github.com/lrstanley/entrest/_examples/kitchensink/internal/database/ent/pet"
	"github.com/lrstanley/entrest/_examples/kitchensink/internal/database/ent/settings"
	"github.com/lrstanley/entrest/_examples/kitchensink/internal/database/ent/user"
)

// UpdateCategoryParams defines parameters for updating a Category via a PATCH request.
type UpdateCategoryParams struct {
	Name       Option[string] `json:"name"`
	AddPets    Option[[]int]  `json:"add_pets,omitempty"`
	RemovePets Option[[]int]  `json:"remove_pets,omitempty"`
}

func (u *UpdateCategoryParams) ApplyInputs(builder *ent.CategoryUpdateOne) *ent.CategoryUpdateOne {
	if v, ok := u.Name.Get(); ok {
		builder.SetName(v)
	}

	if v, ok := u.AddPets.Get(); ok && v != nil {
		builder.AddPetIDs(v...)
	}
	if v, ok := u.RemovePets.Get(); ok && v != nil {
		builder.RemovePetIDs(v...)
	}
	return builder
}

// Exec wraps all logic (mapping all provided values to the build), updates the entity,
// and does another query (using provided query as base) to get the entity, with all eager
// loaded edges.
func (c *UpdateCategoryParams) Exec(ctx context.Context, builder *ent.CategoryUpdateOne, query *ent.CategoryQuery) (*ent.Category, error) {
	result, err := c.ApplyInputs(builder).Save(ctx)
	if err != nil {
		return nil, err
	}
	return EagerLoadCategory(query.Where(category.ID(result.ID))).Only(ctx)
}

// UpdateFriendshipParams defines parameters for updating a Friendship via a PATCH request.
type UpdateFriendshipParams struct {
	CreatedAt Option[time.Time] `json:"created_at"`
	UserID    Option[int]       `json:"user_id"`
	FriendID  Option[int]       `json:"friend_id"`
}

func (u *UpdateFriendshipParams) ApplyInputs(builder *ent.FriendshipUpdateOne) *ent.FriendshipUpdateOne {
	if v, ok := u.CreatedAt.Get(); ok {
		builder.SetCreatedAt(v)
	}
	if v, ok := u.UserID.Get(); ok {
		builder.SetUserID(v)
	}
	if v, ok := u.FriendID.Get(); ok {
		builder.SetFriendID(v)
	}

	return builder
}

// Exec wraps all logic (mapping all provided values to the build), updates the entity,
// and does another query (using provided query as base) to get the entity, with all eager
// loaded edges.
func (c *UpdateFriendshipParams) Exec(ctx context.Context, builder *ent.FriendshipUpdateOne, query *ent.FriendshipQuery) (*ent.Friendship, error) {
	result, err := c.ApplyInputs(builder).Save(ctx)
	if err != nil {
		return nil, err
	}
	return EagerLoadFriendship(query.Where(friendship.ID(result.ID))).Only(ctx)
}

// UpdatePetParams defines parameters for updating a Pet via a PATCH request.
type UpdatePetParams struct {
	Name      Option[string]   `json:"name"`
	Nicknames Option[[]string] `json:"nicknames,omitempty"`
	Age       Option[int]      `json:"age"`
	Type      Option[pet.Type] `json:"type"`
	// Categories that the pet belongs to.
	AddCategories Option[[]int] `json:"add_categories,omitempty"`
	// Categories that the pet belongs to.
	RemoveCategories Option[[]int] `json:"remove_categories,omitempty"`
	// Categories that the pet belongs to.
	Categories Option[[]int] `json:"categories,omitempty"`
	// The user that owns the pet.
	Owner Option[*int] `json:"owner,omitempty"`
	// Pets that this pet is friends with.
	AddFriends Option[[]int] `json:"add_friends,omitempty"`
	// Pets that this pet is friends with.
	RemoveFriends Option[[]int] `json:"remove_friends,omitempty"`
	// Users that this pet is followed by.
	AddFollowedBy Option[[]int] `json:"add_followed_by,omitempty"`
	// Users that this pet is followed by.
	RemoveFollowedBy Option[[]int] `json:"remove_followed_by,omitempty"`
}

func (u *UpdatePetParams) ApplyInputs(builder *ent.PetUpdateOne) *ent.PetUpdateOne {
	if v, ok := u.Name.Get(); ok {
		builder.SetName(v)
	}
	if v, ok := u.Nicknames.Get(); ok {
		builder.SetNicknames(v)
	}
	if v, ok := u.Age.Get(); ok {
		builder.SetAge(v)
	}
	if v, ok := u.Type.Get(); ok {
		builder.SetType(v)
	}

	if v, ok := u.AddCategories.Get(); ok && v != nil {
		builder.AddCategoryIDs(v...)
	}
	if v, ok := u.RemoveCategories.Get(); ok && v != nil {
		builder.RemoveCategoryIDs(v...)
	}
	// If add_<edge> or remove_<edge> is provided, don't clear or use this field.
	if v, ok := u.Categories.Get(); ok && !u.AddCategories.Present() && !u.RemoveCategories.Present() {
		builder.ClearCategories()
		if v != nil {
			builder.AddCategoryIDs(v...)
		}
	}
	if v, ok := u.Owner.Get(); ok {
		if v != nil {
			builder.SetOwnerID(*v)
		} else {
			builder.ClearOwner()
		}
	}
	if v, ok := u.AddFriends.Get(); ok && v != nil {
		builder.AddFriendIDs(v...)
	}
	if v, ok := u.RemoveFriends.Get(); ok && v != nil {
		builder.RemoveFriendIDs(v...)
	}
	if v, ok := u.AddFollowedBy.Get(); ok && v != nil {
		builder.AddFollowedByIDs(v...)
	}
	if v, ok := u.RemoveFollowedBy.Get(); ok && v != nil {
		builder.RemoveFollowedByIDs(v...)
	}
	return builder
}

// Exec wraps all logic (mapping all provided values to the build), updates the entity,
// and does another query (using provided query as base) to get the entity, with all eager
// loaded edges.
func (c *UpdatePetParams) Exec(ctx context.Context, builder *ent.PetUpdateOne, query *ent.PetQuery) (*ent.Pet, error) {
	result, err := c.ApplyInputs(builder).Save(ctx)
	if err != nil {
		return nil, err
	}
	return EagerLoadPet(query.Where(pet.ID(result.ID))).Only(ctx)
}

// UpdateSettingParams defines parameters for updating a Setting via a PATCH request.
type UpdateSettingParams struct {
	// Global banner text to apply to the frontend.
	GlobalBanner Option[*string] `json:"global_banner,omitempty"`
	// Administrators for the platform.
	AddAdmins Option[[]int] `json:"add_admins,omitempty"`
	// Administrators for the platform.
	RemoveAdmins Option[[]int] `json:"remove_admins,omitempty"`
}

func (u *UpdateSettingParams) ApplyInputs(builder *ent.SettingsUpdateOne) *ent.SettingsUpdateOne {
	if v, ok := u.GlobalBanner.Get(); ok {
		if v != nil {
			builder.SetGlobalBanner(*v)
		} else {
			builder.ClearGlobalBanner()
		}
	}

	if v, ok := u.AddAdmins.Get(); ok && v != nil {
		builder.AddAdminIDs(v...)
	}
	if v, ok := u.RemoveAdmins.Get(); ok && v != nil {
		builder.RemoveAdminIDs(v...)
	}
	return builder
}

// Exec wraps all logic (mapping all provided values to the build), updates the entity,
// and does another query (using provided query as base) to get the entity, with all eager
// loaded edges.
func (c *UpdateSettingParams) Exec(ctx context.Context, builder *ent.SettingsUpdateOne, query *ent.SettingsQuery) (*ent.Settings, error) {
	result, err := c.ApplyInputs(builder).Save(ctx)
	if err != nil {
		return nil, err
	}
	return EagerLoadSetting(query.Where(settings.ID(result.ID))).Only(ctx)
}

// UpdateUserParams defines parameters for updating a User via a PATCH request.
type UpdateUserParams struct {
	// Name of the user.
	Name Option[string] `json:"name"`
	// Type of object being defined (user or system which is for internal usecases).
	Type Option[user.Type] `json:"type"`
	// Full name if USER, otherwise null.
	Description Option[*string] `json:"description,omitempty"`
	// If the user is still in the source system.
	Enabled Option[bool] `json:"enabled"`
	// Email associated with the user. Note that not all users have an associated email address.
	Email Option[*string] `json:"email,omitempty"`
	// Avatar data for the user. This should generally only apply to the USER user type.
	Avatar Option[*[]byte] `json:"avatar,omitempty"`
	// Hashed password for the user, this shouldn't be readable in the spec anywhere.
	PasswordHashed Option[string] `json:"password_hashed"`
	// The github user raw JSON data.
	GithubData Option[*github.User] `json:"github_data,omitempty"`
	// Pets owned by the user.
	AddPets Option[[]int] `json:"add_pets,omitempty"`
	// Pets owned by the user.
	RemovePets Option[[]int] `json:"remove_pets,omitempty"`
	// Pets that the user is following.
	AddFollowedPets Option[[]int] `json:"add_followed_pets,omitempty"`
	// Pets that the user is following.
	RemoveFollowedPets Option[[]int] `json:"remove_followed_pets,omitempty"`
	// Friends of the user.
	AddFriends Option[[]int] `json:"add_friends,omitempty"`
	// Friends of the user.
	RemoveFriends     Option[[]int] `json:"remove_friends,omitempty"`
	AddFriendships    Option[[]int] `json:"add_friendships,omitempty"`
	RemoveFriendships Option[[]int] `json:"remove_friendships,omitempty"`
}

func (u *UpdateUserParams) ApplyInputs(builder *ent.UserUpdateOne) *ent.UserUpdateOne {
	if v, ok := u.Name.Get(); ok {
		builder.SetName(v)
	}
	if v, ok := u.Type.Get(); ok {
		builder.SetType(v)
	}
	if v, ok := u.Description.Get(); ok {
		if v != nil {
			builder.SetDescription(*v)
		} else {
			builder.ClearDescription()
		}
	}
	if v, ok := u.Enabled.Get(); ok {
		builder.SetEnabled(v)
	}
	if v, ok := u.Email.Get(); ok {
		if v != nil {
			builder.SetEmail(*v)
		} else {
			builder.ClearEmail()
		}
	}
	if v, ok := u.Avatar.Get(); ok {
		if v != nil {
			builder.SetAvatar(*v)
		} else {
			builder.ClearAvatar()
		}
	}
	if v, ok := u.PasswordHashed.Get(); ok {
		builder.SetPasswordHashed(v)
	}
	if v, ok := u.GithubData.Get(); ok {
		builder.SetGithubData(v)
	}

	if v, ok := u.AddPets.Get(); ok && v != nil {
		builder.AddPetIDs(v...)
	}
	if v, ok := u.RemovePets.Get(); ok && v != nil {
		builder.RemovePetIDs(v...)
	}
	if v, ok := u.AddFollowedPets.Get(); ok && v != nil {
		builder.AddFollowedPetIDs(v...)
	}
	if v, ok := u.RemoveFollowedPets.Get(); ok && v != nil {
		builder.RemoveFollowedPetIDs(v...)
	}
	if v, ok := u.AddFriends.Get(); ok && v != nil {
		builder.AddFriendIDs(v...)
	}
	if v, ok := u.RemoveFriends.Get(); ok && v != nil {
		builder.RemoveFriendIDs(v...)
	}
	if v, ok := u.AddFriendships.Get(); ok && v != nil {
		builder.AddFriendshipIDs(v...)
	}
	if v, ok := u.RemoveFriendships.Get(); ok && v != nil {
		builder.RemoveFriendshipIDs(v...)
	}
	return builder
}

// Exec wraps all logic (mapping all provided values to the build), updates the entity,
// and does another query (using provided query as base) to get the entity, with all eager
// loaded edges.
func (c *UpdateUserParams) Exec(ctx context.Context, builder *ent.UserUpdateOne, query *ent.UserQuery) (*ent.User, error) {
	result, err := c.ApplyInputs(builder).Save(ctx)
	if err != nil {
		return nil, err
	}
	return EagerLoadUser(query.Where(user.ID(result.ID))).Only(ctx)
}
