// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/fastly-test-go/v2/internal/utils"
)

// SnippetDynamic - Sets the snippet version.
type SnippetDynamic int64

const (
	SnippetDynamicZero SnippetDynamic = 0
	SnippetDynamicOne  SnippetDynamic = 1
)

func (e SnippetDynamic) ToPointer() *SnippetDynamic {
	return &e
}

func (e *SnippetDynamic) UnmarshalJSON(data []byte) error {
	var v int64
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case 0:
		fallthrough
	case 1:
		*e = SnippetDynamic(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SnippetDynamic: %v", v)
	}
}

// SnippetType - The location in generated VCL where the snippet should be placed.
type SnippetType string

const (
	SnippetTypeInit    SnippetType = "init"
	SnippetTypeRecv    SnippetType = "recv"
	SnippetTypeHash    SnippetType = "hash"
	SnippetTypeHit     SnippetType = "hit"
	SnippetTypeMiss    SnippetType = "miss"
	SnippetTypePass    SnippetType = "pass"
	SnippetTypeFetch   SnippetType = "fetch"
	SnippetTypeError   SnippetType = "error"
	SnippetTypeDeliver SnippetType = "deliver"
	SnippetTypeLog     SnippetType = "log"
	SnippetTypeNone    SnippetType = "none"
)

func (e SnippetType) ToPointer() *SnippetType {
	return &e
}

func (e *SnippetType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "init":
		fallthrough
	case "recv":
		fallthrough
	case "hash":
		fallthrough
	case "hit":
		fallthrough
	case "miss":
		fallthrough
	case "pass":
		fallthrough
	case "fetch":
		fallthrough
	case "error":
		fallthrough
	case "deliver":
		fallthrough
	case "log":
		fallthrough
	case "none":
		*e = SnippetType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SnippetType: %v", v)
	}
}

type Snippet struct {
	// The VCL code that specifies exactly what the snippet does.
	Content *string `form:"name=content"`
	// Sets the snippet version.
	Dynamic *SnippetDynamic `form:"name=dynamic"`
	// The name for the snippet.
	Name *string `form:"name=name"`
	// Priority determines execution order. Lower numbers execute first.
	Priority *string `default:"100" form:"name=priority"`
	// The location in generated VCL where the snippet should be placed.
	Type *SnippetType `form:"name=type"`
}

func (s Snippet) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *Snippet) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Snippet) GetContent() *string {
	if o == nil {
		return nil
	}
	return o.Content
}

func (o *Snippet) GetDynamic() *SnippetDynamic {
	if o == nil {
		return nil
	}
	return o.Dynamic
}

func (o *Snippet) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *Snippet) GetPriority() *string {
	if o == nil {
		return nil
	}
	return o.Priority
}

func (o *Snippet) GetType() *SnippetType {
	if o == nil {
		return nil
	}
	return o.Type
}
