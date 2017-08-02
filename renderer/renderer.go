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

// CustomizableRenderer is a customizable Markdown renderer. It exposes all
// funtions invoked during the rendering process as fields on the renderer.
// To customize a single aspect of rendering, simply overwrite the associated
// function with your custom one.
type CustomizableRenderer struct {
	BlockCodeHandler  BlockCodeFunc
	BlockQuoteHandler TextFunc
	BlockHtmlHandler  TextFunc

	// HeaderHandler is invoked every time a header needs to be rendered.
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

// BlockCode invokes the BlockCodeHandler currently associated with the
// CustomizableRenderer.
func (r *CustomizableRenderer) BlockCode(out *bytes.Buffer, text []byte, lang string) {
	r.BlockCodeHandler(out, text, lang)
}

// BlockQuote invokes the BlockQuoteHandler currently associated with the
// CustomizableRenderer.
func (r *CustomizableRenderer) BlockQuote(out *bytes.Buffer, text []byte) {
	r.BlockQuoteHandler(out, text)
}

// BlockHtml invokes the BlockHtmlHandler currently associated with the
// CustomizableRenderer.
func (r *CustomizableRenderer) BlockHtml(out *bytes.Buffer, text []byte) {
	r.BlockHtmlHandler(out, text)
}

// Header invokes the HeaderHandler currently associated with the
// CustomizableRenderer.
func (r *CustomizableRenderer) Header(out *bytes.Buffer, text func() bool, level int, id string) {
	r.HeaderHandler(out, text, level, id)
}

// HRule invokes the HRuleHandler currently associated with the
// CustomizableRenderer.
func (r *CustomizableRenderer) HRule(out *bytes.Buffer) {
	r.HRuleHandler(out)
}

// List invokes the ListHandler currently associated with the
// CustomizableRenderer.
func (r *CustomizableRenderer) List(out *bytes.Buffer, text func() bool, flags int) {
	r.ListHandler(out, text, flags)
}

// ListItem invokes the ListItemHandler currently associated with the
// CustomizableRenderer.
func (r *CustomizableRenderer) ListItem(out *bytes.Buffer, text []byte, flags int) {
	r.ListItemHandler(out, text, flags)
}

// Paragraph invokes the ParagraphHandler currently associated with the
// CustomizableRenderer.
func (r *CustomizableRenderer) Paragraph(out *bytes.Buffer, text func() bool) {
	r.ParagraphHandler(out, text)
}

// Table invokes the TableHandler currently associated with the
// CustomizableRenderer.
func (r *CustomizableRenderer) Table(out *bytes.Buffer, header []byte, body []byte, columnData []int) {
	r.TableHandler(out, header, body, columnData)
}

// TableRow invokes the TableRowHandler currently associated with the
// CustomizableRenderer.
func (r *CustomizableRenderer) TableRow(out *bytes.Buffer, text []byte) {
	r.TableRowHandler(out, text)
}

// TableHeaderCell invokes the TableHeaderCellHandler currently associated
// with the CustomizableRenderer.
func (r *CustomizableRenderer) TableHeaderCell(out *bytes.Buffer, text []byte, flags int) {
	r.TableHeaderCellHandler(out, text, flags)
}

// TableCell invokes the TableCellHandler currently associated with the
// CustomizableRenderer.
func (r *CustomizableRenderer) TableCell(out *bytes.Buffer, text []byte, flags int) {
	r.TableCellHandler(out, text, flags)
}

// Footnotes invokes the FootnotesHandler currently associated with the
// CustomizableRenderer.
func (r *CustomizableRenderer) Footnotes(out *bytes.Buffer, text func() bool) {
	r.FootnotesHandler(out, text)
}

// FootnoteItem invokes the FootnoteItemHandler currently associated with the
// CustomizableRenderer.
func (r *CustomizableRenderer) FootnoteItem(out *bytes.Buffer, name, text []byte, flags int) {
	r.FootnoteItemHandler(out, name, text, flags)
}

// TitleBlock invokes the TitleBlockHandler currently associated with the
// CustomizableRenderer.
func (r *CustomizableRenderer) TitleBlock(out *bytes.Buffer, text []byte) {
	r.TitleBlockHandler(out, text)
}

// AutoLink invokes the AutoLinkHandler currently associated with the
// CustomizableRenderer.
func (r *CustomizableRenderer) AutoLink(out *bytes.Buffer, link []byte, kind int) {
	r.AutoLinkHandler(out, link, kind)
}

// CodeSpan invokes the CodeSpanHandler currently associated with the
// CustomizableRenderer.
func (r *CustomizableRenderer) CodeSpan(out *bytes.Buffer, text []byte) {
	r.CodeSpanHandler(out, text)
}

// DoubleEmphasis invokes the DoubleEmphasisHandler currently associated with
// the CustomizableRenderer.
func (r *CustomizableRenderer) DoubleEmphasis(out *bytes.Buffer, text []byte) {
	r.DoubleEmphasisHandler(out, text)
}

// Emphasis invokes the EmphasisHandler currently associated with the
// CustomizableRenderer.
func (r *CustomizableRenderer) Emphasis(out *bytes.Buffer, text []byte) {
	r.EmphasisHandler(out, text)
}

// Image invokes the ImageHandler currently associated with the
// CustomizableRenderer.
func (r *CustomizableRenderer) Image(out *bytes.Buffer, link []byte, title []byte, alt []byte) {
	r.ImageHandler(out, link, title, alt)
}

// LineBreak invokes the LineBreakHandler currently associated with the
// CustomizableRenderer.
func (r *CustomizableRenderer) LineBreak(out *bytes.Buffer) {
	r.LineBreakHandler(out)
}

// Link invokes the LinkHandler currently associated with the
// CustomizableRenderer.
func (r *CustomizableRenderer) Link(out *bytes.Buffer, link []byte, title []byte, content []byte) {
	r.LinkHandler(out, link, title, content)
}

// RawHtmlTag invokes the RawHtmlTagHandler currently associated with the
// CustomizableRenderer.
func (r *CustomizableRenderer) RawHtmlTag(out *bytes.Buffer, tag []byte) {
	r.RawHtmlTagHandler(out, tag)
}

// TripleEmphasis invokes the TripleEmphasisHandler currently associated with
// the CustomizableRenderer.
func (r *CustomizableRenderer) TripleEmphasis(out *bytes.Buffer, text []byte) {
	r.TripleEmphasisHandler(out, text)
}

// StrikeThrough invokes the StrikeThroughHandler currently associated with
// the CustomizableRenderer.
func (r *CustomizableRenderer) StrikeThrough(out *bytes.Buffer, text []byte) {
	r.StrikeThroughHandler(out, text)
}

// FootnoteRef invokes the FootnoteRefHandler currently associated with the
// CustomizableRenderer.
func (r *CustomizableRenderer) FootnoteRef(out *bytes.Buffer, ref []byte, id int) {
	r.FootnoteRefHandler(out, ref, id)
}

// Entity invokes the EntityHandler currently associated with the
// CustomizableRenderer.
func (r *CustomizableRenderer) Entity(out *bytes.Buffer, entity []byte) {
	r.EntityHandler(out, entity)
}

// NormalText invokes the NormalTextHandler currently associated with the
// CustomizableRenderer.
func (r *CustomizableRenderer) NormalText(out *bytes.Buffer, text []byte) {
	r.NormalTextHandler(out, text)
}

// DocumentHeader invokes the DocumentHeaderHandler currently associated with
// the CustomizableRenderer
func (r *CustomizableRenderer) DocumentHeader(out *bytes.Buffer) {
	r.DocumentHeaderHandler(out)
}

// DocumentFooter invokes the DocumentFooterHandler currently associated with
// the CustomizableRenderer.
func (r *CustomizableRenderer) DocumentFooter(out *bytes.Buffer) {
	r.DocumentFooterHandler(out)
}

// GetFlags invokes the GetFlagsHandler currently associated with the
// CustomizableRenderer.
func (r *CustomizableRenderer) GetFlags() int {
	return r.GetFlagsHandler()
}
