package format

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/princjef/gomarkdoc/format/formatcore"
	"github.com/princjef/gomarkdoc/lang"
)

// BitBucketMarkdown provides a Format which is compatible with GitHub
// Flavored Markdown's syntax and semantics. See GitHub's documentation for
// more details about their markdown format:
// https://guides.github.com/features/mastering-markdown/
type BitBucketMarkdown struct{}

// Bold converts the provided text to bold
func (f *BitBucketMarkdown) Bold(text string) (string, error) {
	return formatcore.Bold(text), nil
}

// CodeBlock wraps the provided code as a code block and tags it with the
// provided language (or no language if the empty string is provided).
func (f *BitBucketMarkdown) CodeBlock(language, code string) (string, error) {
	return formatcore.GFMCodeBlock(language, code), nil
}

// Header converts the provided text into a header of the provided level. The
// level is expected to be at least 1.
func (f *BitBucketMarkdown) Header(level int, text string) (string, error) {
	return formatcore.Header(level, formatcore.Escape(text))
}

// RawHeader converts the provided text into a header of the provided level
// without escaping the header text. The level is expected to be at least 1.
func (f *BitBucketMarkdown) RawHeader(level int, text string) (string, error) {
	return formatcore.Header(level, text)
}


// LocalHref generates an href for navigating to a header with the given
// headerText located within the same document as the href itself.
func (f *BitBucketMarkdown) LocalHref(headerText string) (string, error) {
	result := formatcore.PlainText(headerText)
	result = strings.ToLower(result)
	result = strings.TrimSpace(result)
	result = gfmWhitespaceRegex.ReplaceAllString(result, "-")
	result = gfmRemoveRegex.ReplaceAllString(result, "")

	return fmt.Sprintf("#%s", result), nil
}

// Link generates a link with the given text and href values.
func (f *BitBucketMarkdown) Link(text, href string) (string, error) {
	return formatcore.Link(text, href), nil
}

// CodeHref generates an href to the provided code entry.
func (f *BitBucketMarkdown) CodeHref(loc lang.Location) (string, error) {
	// If there's no repo, we can't compute an href
	if loc.Repo == nil {
		return "", nil
	}

	var (
		relative string
		err      error
	)
	if filepath.IsAbs(loc.Filepath) {
		relative, err = filepath.Rel(loc.WorkDir, loc.Filepath)
		if err != nil {
			return "", err
		}
	} else {
		relative = loc.Filepath
	}

	full := filepath.Join(loc.Repo.PathFromRoot, relative)
	p, err := filepath.Rel(string(filepath.Separator), full)
	if err != nil {
		return "", err
	}

	var locStr string
	if loc.Start.Line == loc.End.Line {
		locStr = fmt.Sprintf("lines-%d", loc.Start.Line)
	} else {
		locStr = fmt.Sprintf("lines-%d", loc.Start.Line)
	}

	return fmt.Sprintf(
		"%s/src/%s/%s#%s",
		loc.Repo.Remote,
		loc.Repo.DefaultBranch,
		filepath.ToSlash(p),
		locStr,
	), nil
}

// ListEntry generates an unordered list entry with the provided text at the
// provided zero-indexed depth. A depth of 0 is considered the topmost level of
// list.
func (f *BitBucketMarkdown) ListEntry(depth int, text string) (string, error) {
	return formatcore.ListEntry(depth, text), nil
}

// Accordion generates a collapsible content. The accordion's visible title
// while collapsed is the provided title and the expanded content is the body.
func (f *BitBucketMarkdown) Accordion(title, body string) (string, error) {
	return formatcore.GFMAccordion(title, body), nil
}

// AccordionHeader generates the header visible when an accordion is collapsed.
//
// The AccordionHeader is expected to be used in conjunction with
// AccordionTerminator() when the demands of the body's rendering requires it to
// be generated independently. The result looks conceptually like the following:
//
//	accordion := format.AccordionHeader("Accordion Title") + "Accordion Body" + format.AccordionTerminator()
func (f *BitBucketMarkdown) AccordionHeader(title string) (string, error) {
	return formatcore.GFMAccordionHeader(title), nil
}

// AccordionTerminator generates the code necessary to terminate an accordion
// after the body. It is expected to be used in conjunction with
// AccordionHeader(). See AccordionHeader for a full description.
func (f *BitBucketMarkdown) AccordionTerminator() (string, error) {
	return formatcore.GFMAccordionTerminator(), nil
}

// Paragraph formats a paragraph with the provided text as the contents.
func (f *BitBucketMarkdown) Paragraph(text string) (string, error) {
	return formatcore.Paragraph(text), nil
}

// Escape escapes special markdown characters from the provided text.
func (f *BitBucketMarkdown) Escape(text string) string {
	return formatcore.Escape(text)
}
