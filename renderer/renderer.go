package renderer

import "bytes"

type BlockCodeFunc func(out *bytes.Buffer, text []byte, lang string)
type HeaderFunc func(out *bytes.Buffer, text func() bool, level int, id string)
type ListFunc func(out *bytes.Buffer, text func() bool, flags int)
type TableFunc func(out *bytes.Buffer, header []byte, bode []byte, columnData []int)
type FootnoteItemFunc func(out *bytes.Buffer, name, text []byte, flags int)
type AutoLinkFunc func(out *bytes.Buffer, link []byte, kind int)
type ImageFunc func(out *bytes.Buffer, link []byte, title []byte, alt []byte)
type LinkFunc func(out *bytes.Buffer, link []byte, title []byte, content []byte)
type RawHtmlTagFunc func(out *bytes.Buffer, tag []byte)
type FootnoteRefFunc func(out *bytes.Buffer, ref []byte, id int)
type EntityFunc func(out *bytes.Buffer, entity []byte)
type GetFlagsFunc func() int

type TextFunc func(out *bytes.Buffer, text []byte)
type TextFuncFunc func(out *bytes.Buffer, text func() bool)
type TextFlagsFunc func(out *bytes.Buffer, text []byte, flags int)
type EmptyFunc func(out *bytes.Buffer)

type CustomizableRenderer struct {
	BlockCodeHandler       BlockCodeFunc
	BlockQuoteHandler      TextFunc
	BlockHtmlHandler       TextFunc
	HeaderHandler          HeaderFunc
	HRuleHandler           EmptyFunc
	ListHandler            ListFunc
	ListItemHandler        TextFlagsFunc
	ParagraphHandler       TextFuncFunc
	TableHandler           TableFunc
	TableRowHandler        TextFunc
	TableHeaderCellHandler TextFlagsFunc
	TableCellHandler       TextFlagsFunc
	FootnotesHandler       TextFuncFunc
	FootnoteItemHandler    FootnoteItemFunc
	TitleBlockHandler      TextFunc
	AutoLinkHandler        AutoLinkFunc
	CodeSpanHandler        TextFunc
	DoubleEmphasisHandler  TextFunc
	EmphasisHandler        TextFunc
	ImageHandler           ImageFunc
	LineBreakHandler       EmptyFunc
	LinkHandler            LinkFunc
	RawHtmlTagHandler      RawHtmlTagFunc
	TripleEmphasisHandler  TextFunc
	StrikeThroughHandler   TextFunc
	FootnoteRefHandler     FootnoteRefFunc
	EntityHandler          EntityFunc
	NormalTextHandler      TextFunc
	DocumentHeaderHandler  EmptyFunc
	DocumentFooterHandler  EmptyFunc
	GetFlagsHandler        GetFlagsFunc
}

func (r *CustomizableRenderer) BlockCode(out *bytes.Buffer, text []byte, lang string) {
	r.BlockCodeHandler(out, text, lang)
}

func (r *CustomizableRenderer) BlockQuote(out *bytes.Buffer, text []byte) {
	r.BlockQuoteHandler(out, text)
}

func (r *CustomizableRenderer) BlockHtml(out *bytes.Buffer, text []byte) {
	r.BlockHtmlHandler(out, text)
}

func (r *CustomizableRenderer) Header(out *bytes.Buffer, text func() bool, level int, id string) {
	r.HeaderHandler(out, text, level, id)
}

func (r *CustomizableRenderer) HRule(out *bytes.Buffer) {
	r.HRuleHandler(out)
}

func (r *CustomizableRenderer) List(out *bytes.Buffer, text func() bool, flags int) {
	r.ListHandler(out, text, flags)
}

func (r *CustomizableRenderer) ListItem(out *bytes.Buffer, text []byte, flags int) {
	r.ListItemHandler(out, text, flags)
}

func (r *CustomizableRenderer) Paragraph(out *bytes.Buffer, text func() bool) {
	r.ParagraphHandler(out, text)
}

func (r *CustomizableRenderer) Table(out *bytes.Buffer, header []byte, body []byte, columnData []int) {
	r.TableHandler(out, header, body, columnData)
}

func (r *CustomizableRenderer) TableRow(out *bytes.Buffer, text []byte) {
	r.TableRowHandler(out, text)
}

func (r *CustomizableRenderer) TableHeaderCell(out *bytes.Buffer, text []byte, flags int) {
	r.TableHeaderCellHandler(out, text, flags)
}

func (r *CustomizableRenderer) TableCell(out *bytes.Buffer, text []byte, flags int) {
	r.TableCellHandler(out, text, flags)
}

func (r *CustomizableRenderer) Footnotes(out *bytes.Buffer, text func() bool) {
	r.FootnotesHandler(out, text)
}

func (r *CustomizableRenderer) FootnoteItem(out *bytes.Buffer, name, text []byte, flags int) {
	r.FootnoteItemHandler(out, name, text, flags)
}

func (r *CustomizableRenderer) TitleBlock(out *bytes.Buffer, text []byte) {
	r.TitleBlockHandler(out, text)
}

func (r *CustomizableRenderer) AutoLink(out *bytes.Buffer, link []byte, kind int) {
	r.AutoLinkHandler(out, link, kind)
}

func (r *CustomizableRenderer) CodeSpan(out *bytes.Buffer, text []byte) {
	r.CodeSpanHandler(out, text)
}

func (r *CustomizableRenderer) DoubleEmphasis(out *bytes.Buffer, text []byte) {
	r.DoubleEmphasisHandler(out, text)
}

func (r *CustomizableRenderer) Emphasis(out *bytes.Buffer, text []byte) {
	r.EmphasisHandler(out, text)
}

func (r *CustomizableRenderer) Image(out *bytes.Buffer, link []byte, title []byte, alt []byte) {
	r.ImageHandler(out, link, title, alt)
}

func (r *CustomizableRenderer) LineBreak(out *bytes.Buffer) {
	r.LineBreakHandler(out)
}

func (r *CustomizableRenderer) Link(out *bytes.Buffer, link []byte, title []byte, content []byte) {
	r.LinkHandler(out, link, title, content)
}

func (r *CustomizableRenderer) RawHtmlTag(out *bytes.Buffer, tag []byte) {
	r.RawHtmlTagHandler(out, tag)
}

func (r *CustomizableRenderer) TripleEmphasis(out *bytes.Buffer, text []byte) {
	r.TripleEmphasisHandler(out, text)
}

func (r *CustomizableRenderer) StrikeThrough(out *bytes.Buffer, text []byte) {
	r.StrikeThroughHandler(out, text)
}

func (r *CustomizableRenderer) FootnoteRef(out *bytes.Buffer, ref []byte, id int) {
	r.FootnoteRefHandler(out, ref, id)
}

func (r *CustomizableRenderer) Entity(out *bytes.Buffer, entity []byte) {
	r.EntityHandler(out, entity)
}

func (r *CustomizableRenderer) NormalText(out *bytes.Buffer, text []byte) {
	r.NormalTextHandler(out, text)
}

func (r *CustomizableRenderer) DocumentHeader(out *bytes.Buffer) {
	r.DocumentHeaderHandler(out)
}

func (r *CustomizableRenderer) DocumentFooter(out *bytes.Buffer) {
	r.DocumentFooterHandler(out)
}

func (r *CustomizableRenderer) GetFlags() int {
	return r.GetFlagsHandler()
}
