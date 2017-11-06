package echo

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/SHyx0rmZ/go-markdown-renderer/renderer"
)

// Renderer returns a Markdown renderer that will emit Markdown again. This
// can be used for things like modifying multiple Markdown documents, so you
// can merge them later.
func Renderer() *renderer.Customizable {
	opts := &options{}
	opts.linkPrefix = fmt.Sprintf("%p-", opts)

	return &renderer.Customizable{
		BlockCode,
		BlockQuote,
		BlockHtml,
		Header,
		HRule,
		List,
		ListItem,
		Paragraph,
		Table,
		TableRow,
		TableHeaderCell,
		TableCell,
		Footnotes,
		FootnoteItem,
		TitleBlock,
		AutoLink,
		CodeSpan,
		DoubleEmphasis,
		Emphasis,
		Image,
		LineBreak,
		Link(opts),
		RawHtmlTag,
		TripleEmphasis,
		StrikeThrough,
		FootnoteRef,
		Entity,
		NormalText,
		DocumentHeader,
		DocumentFooter(opts),
		GetFlags,
	}
}

func BlockCode(out *bytes.Buffer, text []byte, lang string) {
	panic("implement me")
}

func BlockQuote(out *bytes.Buffer, text []byte) {
	panic("implement me")
}

func BlockHtml(out *bytes.Buffer, text []byte) {
	panic("implement me")
}

func Header(out *bytes.Buffer, text func() bool, level int, id string) {
	marker := out.Len()

	for i := 0; i < level; i++ {
		out.WriteByte('#')
	}

	out.WriteByte(' ')

	if !text() {
		out.Truncate(marker)
		return
	}

	out.WriteByte('\n')
	out.WriteByte('\n')
}

func HRule(out *bytes.Buffer) {
	panic("implement me")
}

func List(out *bytes.Buffer, text func() bool, flags int) {
	if text() {
		out.WriteByte('\n')
	}
}

func ListItem(out *bytes.Buffer, text []byte, flags int) {
	out.WriteString("* ")
	out.Write(text)
	out.WriteByte('\n')
}

func Paragraph(out *bytes.Buffer, text func() bool) {
	marker := out.Len()

	if !text() {
		out.Truncate(marker)
		return
	}

	out.WriteByte('\n')
	out.WriteByte('\n')
}

func Table(out *bytes.Buffer, header []byte, body []byte, columnData []int) {
	panic("implement me")
}

func TableRow(out *bytes.Buffer, text []byte) {
	panic("implement me")
}

func TableHeaderCell(out *bytes.Buffer, text []byte, flags int) {
	panic("implement me")
}

func TableCell(out *bytes.Buffer, text []byte, flags int) {
	panic("implement me")
}

func Footnotes(out *bytes.Buffer, text func() bool) {
	panic("implement me")
}

func FootnoteItem(out *bytes.Buffer, name, text []byte, flags int) {
	panic("implement me")
}

func TitleBlock(out *bytes.Buffer, text []byte) {
	panic("implement me")
}

func AutoLink(out *bytes.Buffer, link []byte, kind int) {
	panic("implement me")
}

func CodeSpan(out *bytes.Buffer, text []byte) {
	out.WriteByte('`')
	out.Write(text)
	out.WriteByte('`')
}

func DoubleEmphasis(out *bytes.Buffer, text []byte) {
	out.WriteByte('*')
	out.WriteByte('*')
	out.Write(text)
	out.WriteByte('*')
	out.WriteByte('*')
}

func Emphasis(out *bytes.Buffer, text []byte) {
	out.WriteByte('_')
	out.Write(text)
	out.WriteByte('_')
}

func Image(out *bytes.Buffer, link []byte, title []byte, alt []byte) {
	panic("implement me")
}

func LineBreak(out *bytes.Buffer) {
	panic("implement me")
}

func Link(opts *options) renderer.LinkFunc {
	return func(out *bytes.Buffer, link []byte, title []byte, content []byte) {
		out.WriteByte('[')
		out.Write(content)
		out.WriteByte(']')
		out.WriteByte('[')
		out.WriteString(opts.linkPrefix + strconv.Itoa(opts.link(link, title, content)))
		out.WriteByte(']')
	}
}

func RawHtmlTag(out *bytes.Buffer, tag []byte) {
	panic("implement me")
}

func TripleEmphasis(out *bytes.Buffer, text []byte) {
	panic("implement me")
}

func StrikeThrough(out *bytes.Buffer, text []byte) {
	panic("implement me")
}

func FootnoteRef(out *bytes.Buffer, ref []byte, id int) {
	panic("implement me")
}

func Entity(out *bytes.Buffer, entity []byte) {
	out.Write(entity)
}

func NormalText(out *bytes.Buffer, text []byte) {
	out.Write(text)
}

func DocumentHeader(out *bytes.Buffer) {}

func DocumentFooter(opts *options) renderer.EmptyFunc {
	return func(out *bytes.Buffer) {
		current := opts.links
		i := 0

		for current != nil {
			out.WriteByte('[')
			out.WriteString(opts.linkPrefix + strconv.Itoa(i))
			out.WriteByte(']')
			out.WriteByte(':')
			out.WriteByte(' ')
			out.Write(current.Link)
			if len(current.Title) != 0 {
				out.WriteByte(' ')
				out.WriteByte('"')
				out.Write(current.Title)
				out.WriteByte('"')
			}
			out.WriteByte('\n')

			current = current.next
			i++
		}
	}
}

func GetFlags() int {
	return 0
}
