// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"github.com/speakeasy-sdks/fastly-test-go/v2/internal/utils"
)

type VersionInput struct {
	// Whether this is the active version or not.
	Active *bool `default:"false" form:"name=active"`
	// A freeform descriptive note.
	Comment *string `form:"name=comment"`
	// Unused at this time.
	Deployed *bool `form:"name=deployed"`
	// Whether this version is locked or not. Objects can not be added or edited on locked versions.
	Locked *bool `default:"false" form:"name=locked"`
	// Unused at this time.
	Staging *bool `default:"false" form:"name=staging"`
	// Unused at this time.
	Testing *bool `default:"false" form:"name=testing"`
}

func (v VersionInput) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(v, "", false)
}

func (v *VersionInput) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &v, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *VersionInput) GetActive() *bool {
	if o == nil {
		return nil
	}
	return o.Active
}

func (o *VersionInput) GetComment() *string {
	if o == nil {
		return nil
	}
	return o.Comment
}

func (o *VersionInput) GetDeployed() *bool {
	if o == nil {
		return nil
	}
	return o.Deployed
}

func (o *VersionInput) GetLocked() *bool {
	if o == nil {
		return nil
	}
	return o.Locked
}

func (o *VersionInput) GetStaging() *bool {
	if o == nil {
		return nil
	}
	return o.Staging
}

func (o *VersionInput) GetTesting() *bool {
	if o == nil {
		return nil
	}
	return o.Testing
}
