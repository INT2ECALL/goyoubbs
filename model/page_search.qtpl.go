// Code generated by qtc from "page_search.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// 搜索文章列表（同首页），继承 TopicLstPage ，只修改 MainBody()
//

//line model/page_search.qtpl:3
package model

//line model/page_search.qtpl:3
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line model/page_search.qtpl:3
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line model/page_search.qtpl:4
type SearchPage struct {
	TopicLstPage
}

//line model/page_search.qtpl:9
func (p *SearchPage) StreamMainBody(qw422016 *qt422016.Writer) {
//line model/page_search.qtpl:9
	qw422016.N().S(`

<div class="index">

    <article>
        <header>
            <h1 class="entry-title">Search: `)
//line model/page_search.qtpl:15
	qw422016.E().S(p.Q)
//line model/page_search.qtpl:15
	qw422016.N().S(`</h1>
        </header>
    </article>

    `)
//line model/page_search.qtpl:19
	for _, item := range p.TopicPageInfo.Items {
//line model/page_search.qtpl:19
		qw422016.N().S(`
    <article>

        <header>
            `)
//line model/page_search.qtpl:23
		if item.Comments > 0 {
//line model/page_search.qtpl:23
			qw422016.N().S(`
            <a href="/t/`)
//line model/page_search.qtpl:24
			qw422016.N().DUL(item.ID)
//line model/page_search.qtpl:24
			qw422016.N().S(`#r`)
//line model/page_search.qtpl:24
			qw422016.N().DUL(item.Comments)
//line model/page_search.qtpl:24
			qw422016.N().S(`"><img alt="`)
//line model/page_search.qtpl:24
			qw422016.E().S(item.Title)
//line model/page_search.qtpl:24
			qw422016.N().S(` icon" src="/icon/t/`)
//line model/page_search.qtpl:24
			qw422016.N().DUL(item.ID)
//line model/page_search.qtpl:24
			qw422016.N().S(`.jpg?r=`)
//line model/page_search.qtpl:24
			qw422016.N().DUL(item.Comments)
//line model/page_search.qtpl:24
			qw422016.N().S(`" class="avatar"></a>
            `)
//line model/page_search.qtpl:25
		} else {
//line model/page_search.qtpl:25
			qw422016.N().S(`
            <a href="/t/`)
//line model/page_search.qtpl:26
			qw422016.N().DUL(item.ID)
//line model/page_search.qtpl:26
			qw422016.N().S(`"><img alt="`)
//line model/page_search.qtpl:26
			qw422016.E().S(item.Title)
//line model/page_search.qtpl:26
			qw422016.N().S(` icon" src="/static/avatar/`)
//line model/page_search.qtpl:26
			qw422016.N().DUL(item.UserId)
//line model/page_search.qtpl:26
			qw422016.N().S(`.jpg" class="avatar"></a>
            `)
//line model/page_search.qtpl:27
		}
//line model/page_search.qtpl:27
		qw422016.N().S(`
            <h1 class="entry-title"><a href="/t/`)
//line model/page_search.qtpl:28
		qw422016.N().DUL(item.ID)
//line model/page_search.qtpl:28
		qw422016.N().S(`" rel="bookmark" title="Permanent Link to `)
//line model/page_search.qtpl:28
		qw422016.E().S(item.Title)
//line model/page_search.qtpl:28
		qw422016.N().S(`">`)
//line model/page_search.qtpl:28
		qw422016.E().S(item.Title)
//line model/page_search.qtpl:28
		qw422016.N().S(`</a></h1>
            <p class="meta">
                <time datetime="`)
//line model/page_search.qtpl:30
		qw422016.E().S(item.AddTimeFmt)
//line model/page_search.qtpl:30
		qw422016.N().S(`" pubdate data-updated="true">`)
//line model/page_search.qtpl:30
		qw422016.E().S(item.EditTimeFmt)
//line model/page_search.qtpl:30
		qw422016.N().S(`</time>
                in <a href="/n/`)
//line model/page_search.qtpl:31
		qw422016.N().DUL(item.NodeId)
//line model/page_search.qtpl:31
		qw422016.N().S(`" rel="bookmark">`)
//line model/page_search.qtpl:31
		qw422016.E().S(item.NodeName)
//line model/page_search.qtpl:31
		qw422016.N().S(`</a>
                by <a href="/member/`)
//line model/page_search.qtpl:32
		qw422016.N().DUL(item.UserId)
//line model/page_search.qtpl:32
		qw422016.N().S(`" rel="nofollow">`)
//line model/page_search.qtpl:32
		qw422016.E().S(item.AuthorName)
//line model/page_search.qtpl:32
		qw422016.N().S(`</a>
                `)
//line model/page_search.qtpl:33
		if item.Comments > 0 {
//line model/page_search.qtpl:33
			qw422016.N().S(`
                <a class="count" href="/t/`)
//line model/page_search.qtpl:34
			qw422016.N().DUL(item.ID)
//line model/page_search.qtpl:34
			qw422016.N().S(`#r`)
//line model/page_search.qtpl:34
			qw422016.N().DUL(item.Comments)
//line model/page_search.qtpl:34
			qw422016.N().S(`" title="Comment on `)
//line model/page_search.qtpl:34
			qw422016.E().S(item.Title)
//line model/page_search.qtpl:34
			qw422016.N().S(`">`)
//line model/page_search.qtpl:34
			qw422016.N().DUL(item.Comments)
//line model/page_search.qtpl:34
			qw422016.N().S(`</a>
                `)
//line model/page_search.qtpl:35
		}
//line model/page_search.qtpl:35
		qw422016.N().S(`
            </p>
        </header>

        <div class="entry-content">
            `)
//line model/page_search.qtpl:40
		qw422016.E().S(item.FirstCon)
//line model/page_search.qtpl:40
		qw422016.N().S(`
        </div>

    </article>

    `)
//line model/page_search.qtpl:45
	}
//line model/page_search.qtpl:45
	qw422016.N().S(`

</div>

`)
//line model/page_search.qtpl:49
}

//line model/page_search.qtpl:49
func (p *SearchPage) WriteMainBody(qq422016 qtio422016.Writer) {
//line model/page_search.qtpl:49
	qw422016 := qt422016.AcquireWriter(qq422016)
//line model/page_search.qtpl:49
	p.StreamMainBody(qw422016)
//line model/page_search.qtpl:49
	qt422016.ReleaseWriter(qw422016)
//line model/page_search.qtpl:49
}

//line model/page_search.qtpl:49
func (p *SearchPage) MainBody() string {
//line model/page_search.qtpl:49
	qb422016 := qt422016.AcquireByteBuffer()
//line model/page_search.qtpl:49
	p.WriteMainBody(qb422016)
//line model/page_search.qtpl:49
	qs422016 := string(qb422016.B)
//line model/page_search.qtpl:49
	qt422016.ReleaseByteBuffer(qb422016)
//line model/page_search.qtpl:49
	return qs422016
//line model/page_search.qtpl:49
}
