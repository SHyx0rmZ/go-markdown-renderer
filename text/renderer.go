package text

import (
	"bytes"
	"html"
	"strings"

	"github.com/SHyx0rmZ/go-markdown-renderer/echo"
	"github.com/SHyx0rmZ/go-markdown-renderer/renderer"
)

// Renderer returns a Markdown renderer that emits plain text. This can be
// used to prepare a Markdown document for communication channels like chat
// messaging or email. To wrap at a certain amount of characters, set
// maximumLineLength to a value greater than 0. The renderer will take a best
// effort approach to try and make sure lines break before the specified line
// length.
func Renderer(maximumLineLength int) *renderer.Customizable {
	return &renderer.Customizable{
		echo.BlockCode,
		echo.BlockQuote,
		echo.BlockHtml,
		Header,
		echo.HRule,
		echo.List,
		ListItem,
		Paragraph(maximumLineLength),
		echo.Table,
		echo.TableRow,
		echo.TableHeaderCell,
		echo.TableCell,
		echo.Footnotes,
		echo.FootnoteItem,
		echo.TitleBlock,
		echo.AutoLink,
		echo.CodeSpan,
		DoubleEmphasis,
		Emphasis,
		echo.Image,
		echo.LineBreak,
		Link,
		RawHtmlTag,
		echo.TripleEmphasis,
		echo.StrikeThrough,
		echo.FootnoteRef,
		Entity,
		NormalText,
		echo.DocumentHeader,
		DocumentFooter,
		echo.GetFlags,
	}
}

func Header(out *bytes.Buffer, text func() bool, level int, id string) {
	marker := out.Len()

	if !text() {
		return
	}

	length := out.Len() - marker

	out.WriteByte('\n')

	underlineChar := byte('-')

	if level == 1 {
		underlineChar = byte('=')
	}

	for i := 0; i < length; i++ {
		out.WriteByte(underlineChar)
	}

	out.WriteByte('\n')
	out.WriteByte('\n')
}

func ListItem(out *bytes.Buffer, text []byte, flags int) {
	out.WriteByte('-')
	out.WriteByte(' ')
	out.Write(text)
	out.WriteByte('\n')
}

func Paragraph(maximumLineLength int) renderer.TextFuncFunc {
	return func(out *bytes.Buffer, text func() bool) {
		markerBefore := out.Len()

		if !text() {
			out.Truncate(markerBefore)
			return
		}

		markerAfter := out.Len()

		if maximumLineLength > 0 {
			modifiedText := out.Bytes()[markerBefore:markerAfter]

			lastNewline := 0
			index := 0

			for strings.Index(string(modifiedText[index:]), " ") != -1 {
				index += strings.Index(string(modifiedText[index:]), " ")

				if index-lastNewline >= maximumLineLength-1 {
					modifiedText[index] = '\n'

					lastNewline = index
				}

				index++
			}

			out.Truncate(markerBefore)
			out.WriteString(string(modifiedText))
		}

		out.WriteByte('\n')
		out.WriteByte('\n')
	}
}

func DoubleEmphasis(out *bytes.Buffer, text []byte) {
	out.Write(text)
}

func Emphasis(out *bytes.Buffer, text []byte) {
	out.Write(text)
}

func Link(out *bytes.Buffer, link []byte, title []byte, content []byte) {
	out.Write(content)
	out.WriteByte(' ')
	out.WriteByte('(')
	out.Write(link)
	out.WriteByte(')')
}

func RawHtmlTag(out *bytes.Buffer, tag []byte) {
	out.Write(tag)
}

func Entity(out *bytes.Buffer, entity []byte) {
	out.Write([]byte(html.UnescapeString(string(entity))))
}

func NormalText(out *bytes.Buffer, text []byte) {
	out.WriteString(strings.Replace(string(text), "\n", " ", -1))
}

func DocumentFooter(out *bytes.Buffer) {}
