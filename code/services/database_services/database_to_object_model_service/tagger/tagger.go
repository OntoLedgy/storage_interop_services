package tagger

import (
	"github.com/OntoLedgy/storage_interop_services/code/services/database_services/contract"
	"github.com/OntoLedgy/storage_interop_services/code/services/database_services/database_i_o_service/object_model"
	"github.com/OntoLedgy/storage_interop_services/code/services/database_services/database_i_o_service/object_model/configurations"
	"strings"
	"sync"
)

const (
	tagsDisabled = 0

	// number is a ascending sequence of i*2 to determine which tags to generate later
	tagDb         = 1
	tagMastermind = 2
)

var stringPool = sync.Pool{
	New: func() interface{} {
		return new(strings.Builder)
	},
}

// Tagger interface for types of struct-tages
type Tagger interface {
	GenerateTag(db contract.IDatabases, column object_model.Column) string
}

// Taggers represents the supported tags to generate.
type Taggers struct {
	settings *configurations.Settings

	enabledTags int
	taggers     map[int]Tagger
}

// NewTaggers is the constructor function to create the supported taggers.
func NewTaggers(s *configurations.Settings) *Taggers {
	t := &Taggers{
		settings:    s,
		enabledTags: tagDb,
		taggers: map[int]Tagger{
			tagDb:         new(Db),
			tagMastermind: new(Mastermind),
		},
	}

	t.enableTags()

	return t
}

// enableTags enables the tags to generate determined by the settings.
// If multiple standlone tags where specified (the ones with "only" it its names)
// the last specified standalone tag wins.
func (t *Taggers) enableTags() {
	if t.settings.TagsNoDb {
		t.enabledTags = tagsDisabled
	}
	if t.settings.TagsMastermindStructable {
		t.enabledTags |= tagMastermind
	}
	if t.settings.TagsMastermindStructableOnly {
		t.enabledTags = tagsDisabled
		t.enabledTags |= tagMastermind
	}
}

// GenerateTag creates based on the enabled tags and the given database_i_o_service and column
// the tag for the struct field.
func (t *Taggers) GenerateTag(db contract.IDatabases, column object_model.Column) (tags string) {
	sb := stringPool.Get().(*strings.Builder)
	defer func() {
		sb.Reset()
		stringPool.Put(sb)
	}()

	for bit := 1; bit <= t.enabledTags; bit *= 2 {
		shouldTag := t.enabledTags&bit > 0
		if shouldTag {
			sb.WriteString(t.taggers[bit].GenerateTag(db, column))
			sb.WriteString(" ")
		}
	}

	tags = sb.String()

	if len(tags) > 0 {
		tags = "`" + strings.TrimSpace(tags) + "`"
	}

	return tags
}
