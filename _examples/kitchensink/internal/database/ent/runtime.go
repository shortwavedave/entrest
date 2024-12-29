// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/lrstanley/entrest/_examples/kitchensink/internal/database/ent/category"
	"github.com/lrstanley/entrest/_examples/kitchensink/internal/database/ent/dog"
	"github.com/lrstanley/entrest/_examples/kitchensink/internal/database/ent/follows"
	"github.com/lrstanley/entrest/_examples/kitchensink/internal/database/ent/friendship"
	"github.com/lrstanley/entrest/_examples/kitchensink/internal/database/ent/pet"
	"github.com/lrstanley/entrest/_examples/kitchensink/internal/database/ent/settings"
	"github.com/lrstanley/entrest/_examples/kitchensink/internal/database/ent/user"
	"github.com/lrstanley/entrest/_examples/kitchensink/internal/database/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	categoryMixin := schema.Category{}.Mixin()
	categoryMixinFields0 := categoryMixin[0].Fields()
	_ = categoryMixinFields0
	categoryFields := schema.Category{}.Fields()
	_ = categoryFields
	// categoryDescCreatedAt is the schema descriptor for created_at field.
	categoryDescCreatedAt := categoryMixinFields0[0].Descriptor()
	// category.DefaultCreatedAt holds the default value on creation for the created_at field.
	category.DefaultCreatedAt = categoryDescCreatedAt.Default.(func() time.Time)
	// categoryDescUpdatedAt is the schema descriptor for updated_at field.
	categoryDescUpdatedAt := categoryMixinFields0[1].Descriptor()
	// category.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	category.DefaultUpdatedAt = categoryDescUpdatedAt.Default.(func() time.Time)
	// category.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	category.UpdateDefaultUpdatedAt = categoryDescUpdatedAt.UpdateDefault.(func() time.Time)
	// categoryDescNillable is the schema descriptor for nillable field.
	categoryDescNillable := categoryFields[3].Descriptor()
	// category.DefaultNillable holds the default value on creation for the nillable field.
	category.DefaultNillable = categoryDescNillable.Default.(string)
	dogFields := schema.Dog{}.Fields()
	_ = dogFields
	// dogDescID is the schema descriptor for id field.
	dogDescID := dogFields[0].Descriptor()
	// dog.DefaultID holds the default value on creation for the id field.
	dog.DefaultID = dogDescID.Default.(string)
	followsFields := schema.Follows{}.Fields()
	_ = followsFields
	// followsDescFollowedAt is the schema descriptor for followed_at field.
	followsDescFollowedAt := followsFields[0].Descriptor()
	// follows.DefaultFollowedAt holds the default value on creation for the followed_at field.
	follows.DefaultFollowedAt = followsDescFollowedAt.Default.(func() time.Time)
	friendshipFields := schema.Friendship{}.Fields()
	_ = friendshipFields
	// friendshipDescCreatedAt is the schema descriptor for created_at field.
	friendshipDescCreatedAt := friendshipFields[0].Descriptor()
	// friendship.DefaultCreatedAt holds the default value on creation for the created_at field.
	friendship.DefaultCreatedAt = friendshipDescCreatedAt.Default.(func() time.Time)
	petFields := schema.Pet{}.Fields()
	_ = petFields
	// petDescAge is the schema descriptor for age field.
	petDescAge := petFields[3].Descriptor()
	// pet.AgeValidator is a validator for the "age" field. It is called by the builders before save.
	pet.AgeValidator = func() func(int) error {
		validators := petDescAge.Validators
		fns := [...]func(int) error{
			validators[0].(func(int) error),
			validators[1].(func(int) error),
		}
		return func(age int) error {
			for _, fn := range fns {
				if err := fn(age); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	settingsMixin := schema.Settings{}.Mixin()
	settingsMixinFields0 := settingsMixin[0].Fields()
	_ = settingsMixinFields0
	settingsFields := schema.Settings{}.Fields()
	_ = settingsFields
	// settingsDescCreatedAt is the schema descriptor for created_at field.
	settingsDescCreatedAt := settingsMixinFields0[0].Descriptor()
	// settings.DefaultCreatedAt holds the default value on creation for the created_at field.
	settings.DefaultCreatedAt = settingsDescCreatedAt.Default.(func() time.Time)
	// settingsDescUpdatedAt is the schema descriptor for updated_at field.
	settingsDescUpdatedAt := settingsMixinFields0[1].Descriptor()
	// settings.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	settings.DefaultUpdatedAt = settingsDescUpdatedAt.Default.(func() time.Time)
	// settings.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	settings.UpdateDefaultUpdatedAt = settingsDescUpdatedAt.UpdateDefault.(func() time.Time)
	// settingsDescGlobalBanner is the schema descriptor for global_banner field.
	settingsDescGlobalBanner := settingsFields[0].Descriptor()
	// settings.GlobalBannerValidator is a validator for the "global_banner" field. It is called by the builders before save.
	settings.GlobalBannerValidator = func() func(string) error {
		validators := settingsDescGlobalBanner.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(global_banner string) error {
			for _, fn := range fns {
				if err := fn(global_banner); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	userMixin := schema.User{}.Mixin()
	userMixinFields0 := userMixin[0].Fields()
	_ = userMixinFields0
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userMixinFields0[0].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userMixinFields0[1].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
	// userDescDescription is the schema descriptor for description field.
	userDescDescription := userFields[3].Descriptor()
	// user.DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	user.DescriptionValidator = func() func(string) error {
		validators := userDescDescription.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(description string) error {
			for _, fn := range fns {
				if err := fn(description); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// userDescEnabled is the schema descriptor for enabled field.
	userDescEnabled := userFields[4].Descriptor()
	// user.DefaultEnabled holds the default value on creation for the enabled field.
	user.DefaultEnabled = userDescEnabled.Default.(bool)
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[5].Descriptor()
	// user.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	user.EmailValidator = func() func(string) error {
		validators := userDescEmail.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(email string) error {
			for _, fn := range fns {
				if err := fn(email); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// userDescAvatar is the schema descriptor for avatar field.
	userDescAvatar := userFields[6].Descriptor()
	// user.AvatarValidator is a validator for the "avatar" field. It is called by the builders before save.
	user.AvatarValidator = func() func([]byte) error {
		validators := userDescAvatar.Validators
		fns := [...]func([]byte) error{
			validators[0].(func([]byte) error),
			validators[1].(func([]byte) error),
		}
		return func(avatar []byte) error {
			for _, fn := range fns {
				if err := fn(avatar); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// userDescPasswordHashed is the schema descriptor for password_hashed field.
	userDescPasswordHashed := userFields[7].Descriptor()
	// user.PasswordHashedValidator is a validator for the "password_hashed" field. It is called by the builders before save.
	user.PasswordHashedValidator = userDescPasswordHashed.Validators[0].(func(string) error)
	// userDescProfileURL is the schema descriptor for profile_url field.
	userDescProfileURL := userFields[9].Descriptor()
	// user.DefaultProfileURL holds the default value on creation for the profile_url field.
	user.DefaultProfileURL = userDescProfileURL.Default.(*schema.ExampleValuer)
}
