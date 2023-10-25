// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type Gzip struct {
	// Name of the cache condition controlling when this configuration applies.
	CacheCondition *string `form:"name=cache_condition"`
	// Space-separated list of content types to compress. If you omit this field a default list will be used.
	ContentTypes *string `form:"name=content_types"`
	// Space-separated list of file extensions to compress. If you omit this field a default list will be used.
	Extensions *string `form:"name=extensions"`
	// Name of the gzip configuration.
	Name *string `form:"name=name"`
}

func (o *Gzip) GetCacheCondition() *string {
	if o == nil {
		return nil
	}
	return o.CacheCondition
}

func (o *Gzip) GetContentTypes() *string {
	if o == nil {
		return nil
	}
	return o.ContentTypes
}

func (o *Gzip) GetExtensions() *string {
	if o == nil {
		return nil
	}
	return o.Extensions
}

func (o *Gzip) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}
