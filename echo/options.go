package echo

type markdownLink struct {
	Title   []byte
	Content []byte
	Link    []byte

	next *markdownLink
}

type options struct {
	links      *markdownLink
	linkPrefix string
}

func (opts *options) link(link []byte, title []byte, content []byte) int {
	if opts.links == nil {
		opts.links = &markdownLink{
			Title:   title,
			Content: content,
			Link:    link,
		}

		return 0
	}

	current := opts.links
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
