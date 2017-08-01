package echo

import (
	"bytes"
	"strconv"

	"fmt"
	markdown "github.com/russross/blackfriday"
)

func Renderer() markdown.Renderer {
	m := &markdownMerger{}
	m.linkPrefix = fmt.Sprintf("%p-", m)
	return m
}

type markdownLink struct {
	Title   []byte
	Content []byte
	Link    []byte

	next *markdownLink
}

type markdownMerger struct {
	links      *markdownLink
	linkPrefix string
}

func (markdownMerger) BlockCode(out *bytes.Buffer, text []byte, lang string) {
	panic("implement me")
}

func (markdownMerger) BlockQuote(out *bytes.Buffer, text []byte) {
	panic("implement me")
}

func (markdownMerger) BlockHtml(out *bytes.Buffer, text []byte) {
	panic("implement me")
}

func (markdownMerger) Header(out *bytes.Buffer, text func() bool, level int, id string) {
	marker := out.Len()

	for i := 0; i <= level; i++ {
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

func (markdownMerger) HRule(out *bytes.Buffer) {
	panic("implement me")
}

func (markdownMerger) List(out *bytes.Buffer, text func() bool, flags int) {
	if text() {
		out.WriteByte('\n')
	}
}

func (markdownMerger) ListItem(out *bytes.Buffer, text []byte, flags int) {
	out.WriteString("* ")
	out.Write(text)
	out.WriteByte('\n')
}

func (markdownMerger) Paragraph(out *bytes.Buffer, text func() bool) {
	marker := out.Len()

	//out.WriteByte('\n')

	if !text() {
		out.Truncate(marker)
		return
	}

	out.WriteByte('\n')
	out.WriteByte('\n')
}

func (markdownMerger) Table(out *bytes.Buffer, header []byte, body []byte, columnData []int) {
	panic("implement me")
}

func (markdownMerger) TableRow(out *bytes.Buffer, text []byte) {
	panic("implement me")
}

func (markdownMerger) TableHeaderCell(out *bytes.Buffer, text []byte, flags int) {
	panic("implement me")
}

func (markdownMerger) TableCell(out *bytes.Buffer, text []byte, flags int) {
	panic("implement me")
}

func (markdownMerger) Footnotes(out *bytes.Buffer, text func() bool) {
	panic("implement me")
}

func (markdownMerger) FootnoteItem(out *bytes.Buffer, name, text []byte, flags int) {
	panic("implement me")
}

func (markdownMerger) TitleBlock(out *bytes.Buffer, text []byte) {
	panic("implement me")
}

func (markdownMerger) AutoLink(out *bytes.Buffer, link []byte, kind int) {
	panic("implement me")
}

func (markdownMerger) CodeSpan(out *bytes.Buffer, text []byte) {
	out.WriteByte('`')
	out.Write(text)
	out.WriteByte('`')
}

func (markdownMerger) DoubleEmphasis(out *bytes.Buffer, text []byte) {
	panic("implement me")
}

func (markdownMerger) Emphasis(out *bytes.Buffer, text []byte) {
	out.WriteByte('_')
	out.Write(text)
	out.WriteByte('_')
}

func (markdownMerger) Image(out *bytes.Buffer, link []byte, title []byte, alt []byte) {
	panic("implement me")
}

func (markdownMerger) LineBreak(out *bytes.Buffer) {
	panic("implement me")
}

func (m *markdownMerger) link(link []byte, title []byte, content []byte) int {
	if m.links == nil {
		m.links = &markdownLink{
			Title:   title,
			Content: content,
			Link:    link,
		}

		return 0
	}

	current := m.links
	i := 1

	for current.next != nil {
		i++
		current = current.next
	}

	current.next = &markdownLink{
		Title:   title,
		Content: content,
		Link:    link,
	}

	return i
}

func (m *markdownMerger) Link(out *bytes.Buffer, link []byte, title []byte, content []byte) {
	out.WriteByte('[')
	out.Write(content)
	out.WriteByte(']')
	out.WriteByte('[')
	out.WriteString(m.linkPrefix + strconv.Itoa(m.link(link, title, content)))
	out.WriteByte(']')
}

func (markdownMerger) RawHtmlTag(out *bytes.Buffer, tag []byte) {
	panic("implement me")
}

func (markdownMerger) TripleEmphasis(out *bytes.Buffer, text []byte) {
	panic("implement me")
}

func (markdownMerger) StrikeThrough(out *bytes.Buffer, text []byte) {
	panic("implement me")
}

func (markdownMerger) FootnoteRef(out *bytes.Buffer, ref []byte, id int) {
	panic("implement me")
}

func (markdownMerger) Entity(out *bytes.Buffer, entity []byte) {
	panic("implement me")
}

func (markdownMerger) NormalText(out *bytes.Buffer, text []byte) {
	out.Write(text)
}

func (markdownMerger) DocumentHeader(out *bytes.Buffer) {
}

func (m markdownMerger) DocumentFooter(out *bytes.Buffer) {
	current := m.links
	i := 0

	for current != nil {
		out.WriteByte('[')
		out.WriteString(m.linkPrefix + strconv.Itoa(i))
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

func (markdownMerger) GetFlags() int {
	return 0
}
