// Code generated by qtc from "page_base.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// This is a base page template. All the other template pages implement this interface.
//

//line model/page_base.qtpl:3
package model

//line model/page_base.qtpl:3
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line model/page_base.qtpl:3
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line model/page_base.qtpl:4
type Page interface {
//line model/page_base.qtpl:4
	Header() string
//line model/page_base.qtpl:4
	StreamHeader(qw422016 *qt422016.Writer)
//line model/page_base.qtpl:4
	WriteHeader(qq422016 qtio422016.Writer)
//line model/page_base.qtpl:4
	Banner() string
//line model/page_base.qtpl:4
	StreamBanner(qw422016 *qt422016.Writer)
//line model/page_base.qtpl:4
	WriteBanner(qq422016 qtio422016.Writer)
//line model/page_base.qtpl:4
	MainBodyNav() string
//line model/page_base.qtpl:4
	StreamMainBodyNav(qw422016 *qt422016.Writer)
//line model/page_base.qtpl:4
	WriteMainBodyNav(qq422016 qtio422016.Writer)
//line model/page_base.qtpl:4
	MainBody() string
//line model/page_base.qtpl:4
	StreamMainBody(qw422016 *qt422016.Writer)
//line model/page_base.qtpl:4
	WriteMainBody(qq422016 qtio422016.Writer)
//line model/page_base.qtpl:4
	Aside() string
//line model/page_base.qtpl:4
	StreamAside(qw422016 *qt422016.Writer)
//line model/page_base.qtpl:4
	WriteAside(qq422016 qtio422016.Writer)
//line model/page_base.qtpl:4
	Footer() string
//line model/page_base.qtpl:4
	StreamFooter(qw422016 *qt422016.Writer)
//line model/page_base.qtpl:4
	WriteFooter(qq422016 qtio422016.Writer)
//line model/page_base.qtpl:4
}

// Page prints a page implementing Page interface.

//line model/page_base.qtpl:15
func StreamPageTemplate(qw422016 *qt422016.Writer, p Page) {
//line model/page_base.qtpl:15
	qw422016.N().S(`
<!DOCTYPE html>
<html lang="zh-cmn-Hans">
<head>
    <meta charset="utf-8">
    `)
//line model/page_base.qtpl:20
	p.StreamHeader(qw422016)
//line model/page_base.qtpl:20
	qw422016.N().S(`
    <meta name="HandheldFriendly" content="True">
    <meta name="viewport" content="width=device-width,minimum-scale=1,initial-scale=1" />
    <meta content="yes" name="apple-mobile-web-app-capable" />
    <meta content="black" name="apple-mobile-web-app-status-bar-style" />
    <meta name="format-detection" content="telephone=no" />
    <link rel="shortcut icon" href="/static/favicon.png" type="image/x-icon" />
    <link rel="top" title="Back to Top" href="#" />
    <link rel="stylesheet" href="/static/css/main.css" type="text/css">
    <script src="/static/js/main.js"></script>
    <!--[if lte IE 8]>
    <link rel="stylesheet" href="/static/css/grids-responsive-old-ie-min.css">
    <![endif]-->
    <!--[if gt IE 8]><!-->
    <link rel="stylesheet" href="/static/css/grids-responsive-min.css">
    <!--<![endif]-->
</head>

<body>
<a name="top"></a>
<header role="banner" class="header-wrap">
    <div class="home-menu pure-menu pure-menu-horizontal">
        `)
//line model/page_base.qtpl:42
	p.StreamBanner(qw422016)
//line model/page_base.qtpl:42
	qw422016.N().S(`
    </div>
</header>

<div id="main" class="main-wrap">
    <div id="content" class="pure-g">
        <div class="pure-u-1 pure-u-sm-17-24">
            <div class="main-left-body">

                `)
//line model/page_base.qtpl:51
	p.StreamMainBodyNav(qw422016)
//line model/page_base.qtpl:51
	qw422016.N().S(`

                `)
//line model/page_base.qtpl:53
	p.StreamMainBody(qw422016)
//line model/page_base.qtpl:53
	qw422016.N().S(`

            </div>
        </div>
        <div class="pure-u-1 pure-u-sm-7-24">
            <div class="main-right-body">
                `)
//line model/page_base.qtpl:59
	p.StreamAside(qw422016)
//line model/page_base.qtpl:59
	qw422016.N().S(`
            </div>
        </div>
    </div>
</div>

<footer role="contentinfo">
    `)
//line model/page_base.qtpl:66
	p.StreamFooter(qw422016)
//line model/page_base.qtpl:66
	qw422016.N().S(`
</footer>
<a style="display: none; " rel="nofollow" href="#top" id="go-to-top">▲</a>
<script>
document.addEventListener("scroll", handleScroll);
let scrollToTopBtn = document.getElementById("go-to-top");
scrollToTopBtn.addEventListener("click", scrollToTop);

function handleScroll() {
    if (document.documentElement.scrollTop > 300) {
        scrollToTopBtn.style.display = "block";
    } else if (document.documentElement.scrollTop < 300){
        scrollToTopBtn.style.display = "none";
    }
}
</script>
</body>
</html>
`)
//line model/page_base.qtpl:84
}

//line model/page_base.qtpl:84
func WritePageTemplate(qq422016 qtio422016.Writer, p Page) {
//line model/page_base.qtpl:84
	qw422016 := qt422016.AcquireWriter(qq422016)
//line model/page_base.qtpl:84
	StreamPageTemplate(qw422016, p)
//line model/page_base.qtpl:84
	qt422016.ReleaseWriter(qw422016)
//line model/page_base.qtpl:84
}

//line model/page_base.qtpl:84
func PageTemplate(p Page) string {
//line model/page_base.qtpl:84
	qb422016 := qt422016.AcquireByteBuffer()
//line model/page_base.qtpl:84
	WritePageTemplate(qb422016, p)
//line model/page_base.qtpl:84
	qs422016 := string(qb422016.B)
//line model/page_base.qtpl:84
	qt422016.ReleaseByteBuffer(qb422016)
//line model/page_base.qtpl:84
	return qs422016
//line model/page_base.qtpl:84
}

// Header bg

//line model/page_base.qtpl:88
func (p *BasePage) StreamHeader(qw422016 *qt422016.Writer) {
//line model/page_base.qtpl:88
	qw422016.N().S(`
<title>`)
//line model/page_base.qtpl:89
	qw422016.E().S(p.Title)
//line model/page_base.qtpl:89
	qw422016.N().S(`</title>
<meta name="description" content="`)
//line model/page_base.qtpl:90
	qw422016.E().S(p.Description)
//line model/page_base.qtpl:90
	qw422016.N().S(`">
<meta name="keywords" content="`)
//line model/page_base.qtpl:91
	qw422016.E().S(p.Keywords)
//line model/page_base.qtpl:91
	qw422016.N().S(`">
<link rel="canonical" href="`)
//line model/page_base.qtpl:92
	qw422016.E().S(p.Canonical)
//line model/page_base.qtpl:92
	qw422016.N().S(`">
<link href="/feed" rel="alternate" title="`)
//line model/page_base.qtpl:93
	qw422016.E().S(p.SiteCf.Name)
//line model/page_base.qtpl:93
	qw422016.N().S(`" type="application/atom+xml">

`)
//line model/page_base.qtpl:95
	if len(p.JsonLd) > 0 {
//line model/page_base.qtpl:95
		qw422016.N().S(`
<script type="application/ld+json">
`)
//line model/page_base.qtpl:97
		qw422016.N().S(p.JsonLd)
//line model/page_base.qtpl:97
		qw422016.N().S(`
</script>
`)
//line model/page_base.qtpl:99
	}
//line model/page_base.qtpl:99
	qw422016.N().S(`

`)
//line model/page_base.qtpl:101
	if p.ShowAutoAd && len(p.SiteCf.GoogleAutoAdJs) > 0 {
//line model/page_base.qtpl:101
		qw422016.N().S(`
`)
//line model/page_base.qtpl:102
		qw422016.N().S(p.SiteCf.GoogleAutoAdJs)
//line model/page_base.qtpl:102
		qw422016.N().S(`
`)
//line model/page_base.qtpl:103
	}
//line model/page_base.qtpl:103
	qw422016.N().S(`

`)
//line model/page_base.qtpl:105
	if len(p.SiteCf.HeaderPartCon) > 0 {
//line model/page_base.qtpl:105
		qw422016.N().S(`
`)
//line model/page_base.qtpl:106
		qw422016.N().S(p.SiteCf.HeaderPartCon)
//line model/page_base.qtpl:106
		qw422016.N().S(`
`)
//line model/page_base.qtpl:107
	}
//line model/page_base.qtpl:107
	qw422016.N().S(`

`)
//line model/page_base.qtpl:109
}

//line model/page_base.qtpl:109
func (p *BasePage) WriteHeader(qq422016 qtio422016.Writer) {
//line model/page_base.qtpl:109
	qw422016 := qt422016.AcquireWriter(qq422016)
//line model/page_base.qtpl:109
	p.StreamHeader(qw422016)
//line model/page_base.qtpl:109
	qt422016.ReleaseWriter(qw422016)
//line model/page_base.qtpl:109
}

//line model/page_base.qtpl:109
func (p *BasePage) Header() string {
//line model/page_base.qtpl:109
	qb422016 := qt422016.AcquireByteBuffer()
//line model/page_base.qtpl:109
	p.WriteHeader(qb422016)
//line model/page_base.qtpl:109
	qs422016 := string(qb422016.B)
//line model/page_base.qtpl:109
	qt422016.ReleaseByteBuffer(qb422016)
//line model/page_base.qtpl:109
	return qs422016
//line model/page_base.qtpl:109
}

// Header ed
//
// Banner bg

//line model/page_base.qtpl:113
func (p *BasePage) StreamBanner(qw422016 *qt422016.Writer) {
//line model/page_base.qtpl:113
	qw422016.N().S(`
<a href="/" class="pure-menu-heading pure-menu-link">`)
//line model/page_base.qtpl:114
	qw422016.E().S(p.SiteCf.Name)
//line model/page_base.qtpl:114
	qw422016.N().S(`</a>
<ul class="pure-menu-list">
    `)
//line model/page_base.qtpl:116
	if p.CurrentUser.ID > 0 {
//line model/page_base.qtpl:116
		qw422016.N().S(`
    <li class="pure-menu-item"><a href="/setting" class="pure-menu-link"><img src="/static/avatar/`)
//line model/page_base.qtpl:117
		qw422016.N().DUL(p.CurrentUser.ID)
//line model/page_base.qtpl:117
		qw422016.N().S(`.jpg" class="avatar" alt="`)
//line model/page_base.qtpl:117
		qw422016.E().S(p.CurrentUser.Name)
//line model/page_base.qtpl:117
		qw422016.N().S(` avatar"> `)
//line model/page_base.qtpl:117
		qw422016.E().S(p.CurrentUser.Name)
//line model/page_base.qtpl:117
		qw422016.N().S(`</a></li>
    <li class="pure-menu-item"><a href="/logout" class="pure-menu-link">退出</a></li>
    `)
//line model/page_base.qtpl:119
	} else {
//line model/page_base.qtpl:119
		qw422016.N().S(`
    <li class="pure-menu-item"><a href="/login" rel="nofollow" class="pure-menu-link">登录</a></li>
    <li class="pure-menu-item"><a href="/register" rel="nofollow" class="pure-menu-link">注册</a></li>
    `)
//line model/page_base.qtpl:122
	}
//line model/page_base.qtpl:122
	qw422016.N().S(`
</ul>
`)
//line model/page_base.qtpl:124
}

//line model/page_base.qtpl:124
func (p *BasePage) WriteBanner(qq422016 qtio422016.Writer) {
//line model/page_base.qtpl:124
	qw422016 := qt422016.AcquireWriter(qq422016)
//line model/page_base.qtpl:124
	p.StreamBanner(qw422016)
//line model/page_base.qtpl:124
	qt422016.ReleaseWriter(qw422016)
//line model/page_base.qtpl:124
}

//line model/page_base.qtpl:124
func (p *BasePage) Banner() string {
//line model/page_base.qtpl:124
	qb422016 := qt422016.AcquireByteBuffer()
//line model/page_base.qtpl:124
	p.WriteBanner(qb422016)
//line model/page_base.qtpl:124
	qs422016 := string(qb422016.B)
//line model/page_base.qtpl:124
	qt422016.ReleaseByteBuffer(qb422016)
//line model/page_base.qtpl:124
	return qs422016
//line model/page_base.qtpl:124
}

// Banner ed
//
// MainBodyNav bg

//line model/page_base.qtpl:128
func (p *BasePage) StreamMainBodyNav(qw422016 *qt422016.Writer) {
//line model/page_base.qtpl:128
	qw422016.N().S(`
`)
//line model/page_base.qtpl:129
	if p.CurrentUser.ID > 0 {
//line model/page_base.qtpl:129
		qw422016.N().S(`
<div class="body-nav box bot-line">
    <div class="pure-button-group">
        `)
//line model/page_base.qtpl:132
		if p.CurrentUser.Flag == 0 {
//line model/page_base.qtpl:132
			qw422016.N().S(`
        <span class="pure-button button-warning">您已被管理员禁用</span>
        `)
//line model/page_base.qtpl:134
		} else {
//line model/page_base.qtpl:134
			qw422016.N().S(`
        `)
//line model/page_base.qtpl:135
			if p.CurrentUser.Flag == 1 {
//line model/page_base.qtpl:135
				qw422016.N().S(`
        <span class="pure-button button-warning">请等待管理员审核</span>
        `)
//line model/page_base.qtpl:137
			}
//line model/page_base.qtpl:137
			qw422016.N().S(`
        `)
//line model/page_base.qtpl:138
			if p.CurrentUser.Flag >= 5 {
//line model/page_base.qtpl:138
				qw422016.N().S(`
        `)
//line model/page_base.qtpl:139
				if p.HasMsg {
//line model/page_base.qtpl:139
					qw422016.N().S(`
        <a class="pure-button button-warning" href="/my/msg">未读信息</a>
        `)
//line model/page_base.qtpl:141
				}
//line model/page_base.qtpl:141
				qw422016.N().S(`
        <a class="pure-button pure-button-primary" href="/topic/add?nid=`)
//line model/page_base.qtpl:142
				qw422016.N().DUL(p.DefaultNode.ID)
//line model/page_base.qtpl:142
				qw422016.N().S(`">发帖</a>
        `)
//line model/page_base.qtpl:143
			}
//line model/page_base.qtpl:143
			qw422016.N().S(`
        `)
//line model/page_base.qtpl:144
			if p.CurrentUser.Flag >= 99 {
//line model/page_base.qtpl:144
				qw422016.N().S(`
            `)
//line model/page_base.qtpl:145
				if p.HasTopicReview {
//line model/page_base.qtpl:145
					qw422016.N().S(`
            <a class="pure-button button-warning" href="/admin/topic/review">审帖</a>
            `)
//line model/page_base.qtpl:147
				}
//line model/page_base.qtpl:147
				qw422016.N().S(`
            `)
//line model/page_base.qtpl:148
				if p.HasReplyReview {
//line model/page_base.qtpl:148
					qw422016.N().S(`
            <a class="pure-button button-warning" href="/admin/comment/review">审评</a>
            `)
//line model/page_base.qtpl:150
				}
//line model/page_base.qtpl:150
				qw422016.N().S(`
            <a class="pure-button" href="/admin/node">节点</a>
            <a class="pure-button" href="/admin/user">用户</a>
            <a class="pure-button" href="/admin/link">链接</a>
            <a class="pure-button" href="/admin/site/conf">设置</a>
        `)
//line model/page_base.qtpl:155
			}
//line model/page_base.qtpl:155
			qw422016.N().S(`
        `)
//line model/page_base.qtpl:156
		}
//line model/page_base.qtpl:156
		qw422016.N().S(`
    </div>
</div>
`)
//line model/page_base.qtpl:159
	}
//line model/page_base.qtpl:159
	qw422016.N().S(`
`)
//line model/page_base.qtpl:160
}

//line model/page_base.qtpl:160
func (p *BasePage) WriteMainBodyNav(qq422016 qtio422016.Writer) {
//line model/page_base.qtpl:160
	qw422016 := qt422016.AcquireWriter(qq422016)
//line model/page_base.qtpl:160
	p.StreamMainBodyNav(qw422016)
//line model/page_base.qtpl:160
	qt422016.ReleaseWriter(qw422016)
//line model/page_base.qtpl:160
}

//line model/page_base.qtpl:160
func (p *BasePage) MainBodyNav() string {
//line model/page_base.qtpl:160
	qb422016 := qt422016.AcquireByteBuffer()
//line model/page_base.qtpl:160
	p.WriteMainBodyNav(qb422016)
//line model/page_base.qtpl:160
	qs422016 := string(qb422016.B)
//line model/page_base.qtpl:160
	qt422016.ReleaseByteBuffer(qb422016)
//line model/page_base.qtpl:160
	return qs422016
//line model/page_base.qtpl:160
}

// MainBodyNav ed
//
// MainBody bg

//line model/page_base.qtpl:164
func (p *BasePage) StreamMainBody(qw422016 *qt422016.Writer) {
//line model/page_base.qtpl:164
	qw422016.N().S(`
This is a base MainBody
`)
//line model/page_base.qtpl:166
}

//line model/page_base.qtpl:166
func (p *BasePage) WriteMainBody(qq422016 qtio422016.Writer) {
//line model/page_base.qtpl:166
	qw422016 := qt422016.AcquireWriter(qq422016)
//line model/page_base.qtpl:166
	p.StreamMainBody(qw422016)
//line model/page_base.qtpl:166
	qt422016.ReleaseWriter(qw422016)
//line model/page_base.qtpl:166
}

//line model/page_base.qtpl:166
func (p *BasePage) MainBody() string {
//line model/page_base.qtpl:166
	qb422016 := qt422016.AcquireByteBuffer()
//line model/page_base.qtpl:166
	p.WriteMainBody(qb422016)
//line model/page_base.qtpl:166
	qs422016 := string(qb422016.B)
//line model/page_base.qtpl:166
	qt422016.ReleaseByteBuffer(qb422016)
//line model/page_base.qtpl:166
	return qs422016
//line model/page_base.qtpl:166
}

// MainBody ed
//
// Aside bg

//line model/page_base.qtpl:170
func (p *BasePage) StreamAside(qw422016 *qt422016.Writer) {
//line model/page_base.qtpl:170
	qw422016.N().S(`
<aside class="sidebar">

    <section class="search-form">
        <form action="/q" method="get" class="pure-form">
            <input type="text" name="q" class="pure-input-rounded" placeholder="站内搜索" />
        </form>
    </section>

    `)
//line model/page_base.qtpl:179
	if len(p.RecentComment) > 0 {
//line model/page_base.qtpl:179
		qw422016.N().S(`
    <section>
        <h1>💬 最近评论</h1>
        <ul id="recent_comments">
            `)
//line model/page_base.qtpl:183
		for _, item := range p.RecentComment {
//line model/page_base.qtpl:183
			qw422016.N().S(`
            <li>
                <img alt="`)
//line model/page_base.qtpl:185
			qw422016.E().S(item.Name)
//line model/page_base.qtpl:185
			qw422016.N().S(` avatar" src="/static/avatar/`)
//line model/page_base.qtpl:185
			qw422016.N().DUL(item.UserId)
//line model/page_base.qtpl:185
			qw422016.N().S(`.jpg" class="avatar"> <a href="`)
//line model/page_base.qtpl:185
			qw422016.E().S(item.Link)
//line model/page_base.qtpl:185
			qw422016.N().S(`">`)
//line model/page_base.qtpl:185
			qw422016.E().S(item.Content)
//line model/page_base.qtpl:185
			qw422016.N().S(`</a>
            </li>
            `)
//line model/page_base.qtpl:187
		}
//line model/page_base.qtpl:187
		qw422016.N().S(`
        </ul>
    </section>
    `)
//line model/page_base.qtpl:190
	}
//line model/page_base.qtpl:190
	qw422016.N().S(`

    <section>
        <h1>📁 分类</h1>
        <div id="top-category-list">
            `)
//line model/page_base.qtpl:195
	for _, item := range p.NodeLst {
//line model/page_base.qtpl:195
		qw422016.N().S(`
            <a href="/n/`)
//line model/page_base.qtpl:196
		qw422016.N().DUL(item.ID)
//line model/page_base.qtpl:196
		qw422016.N().S(`">`)
//line model/page_base.qtpl:196
		qw422016.E().S(item.Name)
//line model/page_base.qtpl:196
		qw422016.N().S(`<span class="tag_meta">(`)
//line model/page_base.qtpl:196
		qw422016.N().DUL(item.TopicNum)
//line model/page_base.qtpl:196
		qw422016.N().S(`)</span></a>
            `)
//line model/page_base.qtpl:197
	}
//line model/page_base.qtpl:197
	qw422016.N().S(`
        </div>
        <a id="id-showMoreNode" href="#" class="more-node" onclick="ShowMoreNode();return false;">显示更多</a>
    </section>

    `)
//line model/page_base.qtpl:202
	if len(p.RangeTopicLst) > 0 {
//line model/page_base.qtpl:202
		qw422016.N().S(`
    <section>
        <h1>📝 最近浏览</h1>
        <ul id="recent_posts">
            `)
//line model/page_base.qtpl:206
		for _, item := range p.RangeTopicLst {
//line model/page_base.qtpl:206
			qw422016.N().S(`
            <li><a href="/t/`)
//line model/page_base.qtpl:207
			qw422016.N().DUL(item.ID)
//line model/page_base.qtpl:207
			qw422016.N().S(`">`)
//line model/page_base.qtpl:207
			qw422016.E().S(item.Title)
//line model/page_base.qtpl:207
			qw422016.N().S(`</a></li>
            `)
//line model/page_base.qtpl:208
		}
//line model/page_base.qtpl:208
		qw422016.N().S(`
        </ul>
    </section>
    `)
//line model/page_base.qtpl:211
	}
//line model/page_base.qtpl:211
	qw422016.N().S(`

    `)
//line model/page_base.qtpl:213
	if len(p.TagCloud) > 0 {
//line model/page_base.qtpl:213
		qw422016.N().S(`
    <section>
        <h1>🏷️ 标签</h1>
        <div id="tag-cloud">
            `)
//line model/page_base.qtpl:217
		for _, item := range p.TagCloud {
//line model/page_base.qtpl:217
			qw422016.N().S(`
            <a href="/tag/`)
//line model/page_base.qtpl:218
			qw422016.N().U(item.Name)
//line model/page_base.qtpl:218
			qw422016.N().S(`">`)
//line model/page_base.qtpl:218
			qw422016.E().S(item.Name)
//line model/page_base.qtpl:218
			qw422016.N().S(`<span class="tag-meta">(`)
//line model/page_base.qtpl:218
			qw422016.N().D(item.Size)
//line model/page_base.qtpl:218
			qw422016.N().S(`)</span></a>
            `)
//line model/page_base.qtpl:219
		}
//line model/page_base.qtpl:219
		qw422016.N().S(`
        </div>
    </section>
    `)
//line model/page_base.qtpl:222
	}
//line model/page_base.qtpl:222
	qw422016.N().S(`

    `)
//line model/page_base.qtpl:224
	if len(p.LinkLst) > 0 {
//line model/page_base.qtpl:224
		qw422016.N().S(`
    <section>
        <h1>🔗 链接</h1>
        <ul id="link-cloud">
            `)
//line model/page_base.qtpl:228
		for _, item := range p.LinkLst {
//line model/page_base.qtpl:228
			qw422016.N().S(`
            <li><a href="`)
//line model/page_base.qtpl:229
			qw422016.E().S(item.Url)
//line model/page_base.qtpl:229
			qw422016.N().S(`" target="_blank">`)
//line model/page_base.qtpl:229
			qw422016.E().S(item.Name)
//line model/page_base.qtpl:229
			qw422016.N().S(`</a></li>
            `)
//line model/page_base.qtpl:230
		}
//line model/page_base.qtpl:230
		qw422016.N().S(`
        </ul>
    </section>
    `)
//line model/page_base.qtpl:233
	}
//line model/page_base.qtpl:233
	qw422016.N().S(`

    `)
//line model/page_base.qtpl:235
	if p.SiteInfo.NodeNum > 0 {
//line model/page_base.qtpl:235
		qw422016.N().S(`
    <section>
        <h1>💡 本站已稳定运行 `)
//line model/page_base.qtpl:237
		qw422016.E().S(p.SiteInfo.Days)
//line model/page_base.qtpl:237
		qw422016.N().S(`</h1>
        <ul id="site-info">
            <li>会员: `)
//line model/page_base.qtpl:239
		qw422016.N().DUL(p.SiteInfo.UserNum)
//line model/page_base.qtpl:239
		qw422016.N().S(`</li>
            <li>帖子: `)
//line model/page_base.qtpl:240
		qw422016.N().DUL(p.SiteInfo.PostNum)
//line model/page_base.qtpl:240
		qw422016.N().S(`</li>
            <li>回复: `)
//line model/page_base.qtpl:241
		qw422016.N().DUL(p.SiteInfo.ReplyNum)
//line model/page_base.qtpl:241
		qw422016.N().S(`</li>
            <li>分类: `)
//line model/page_base.qtpl:242
		qw422016.N().DUL(p.SiteInfo.NodeNum)
//line model/page_base.qtpl:242
		qw422016.N().S(`</li>
            <li>标签: `)
//line model/page_base.qtpl:243
		qw422016.N().DUL(p.SiteInfo.TagNum)
//line model/page_base.qtpl:243
		qw422016.N().S(`</li>
        </ul>
    </section>
    `)
//line model/page_base.qtpl:246
	}
//line model/page_base.qtpl:246
	qw422016.N().S(`

</aside>

<script>
let show_node_btn = document.getElementById("id-showMoreNode");
let node_ul = document.getElementById("top-category-list");
let showAll = false;
function ShowMoreNode(){
    var showNum = 10;
    var liLst = node_ul.querySelectorAll("a");
    for (i = 0; i < liLst.length; i++) {
        if(showAll) {
            liLst[i].style.display = 'inline-block';
        } else {
            if(i<showNum){
                liLst[i].style.display = 'inline-block';
            }else{
                liLst[i].style.display = 'none';
            }
        }
    }
    showAll = !showAll;
    if(showAll) {
        show_node_btn.innerText = "↓ 显示更多 ↓";
    }else{
        show_node_btn.innerText = "↑ 显示更少 ↑";
    }
}
ShowMoreNode();
</script>

`)
//line model/page_base.qtpl:278
}

//line model/page_base.qtpl:278
func (p *BasePage) WriteAside(qq422016 qtio422016.Writer) {
//line model/page_base.qtpl:278
	qw422016 := qt422016.AcquireWriter(qq422016)
//line model/page_base.qtpl:278
	p.StreamAside(qw422016)
//line model/page_base.qtpl:278
	qt422016.ReleaseWriter(qw422016)
//line model/page_base.qtpl:278
}

//line model/page_base.qtpl:278
func (p *BasePage) Aside() string {
//line model/page_base.qtpl:278
	qb422016 := qt422016.AcquireByteBuffer()
//line model/page_base.qtpl:278
	p.WriteAside(qb422016)
//line model/page_base.qtpl:278
	qs422016 := string(qb422016.B)
//line model/page_base.qtpl:278
	qt422016.ReleaseByteBuffer(qb422016)
//line model/page_base.qtpl:278
	return qs422016
//line model/page_base.qtpl:278
}

// Aside ed
//
// Footer bg

//line model/page_base.qtpl:282
func (p *BasePage) StreamFooter(qw422016 *qt422016.Writer) {
//line model/page_base.qtpl:282
	qw422016.N().S(`
<p>
    Copyright &copy; <a href="/">`)
//line model/page_base.qtpl:284
	qw422016.E().S(p.SiteCf.Name)
//line model/page_base.qtpl:284
	qw422016.N().S(`</a> -
    <span class="credit">Powered by <a href="https://youbbs.org">goYouBBS</a> - <a href="#">&uarr;Go Top</a> </span>
</p>

`)
//line model/page_base.qtpl:288
	if len(p.SiteCf.FooterPartHtml) > 0 {
//line model/page_base.qtpl:288
		qw422016.N().S(`
`)
//line model/page_base.qtpl:289
		qw422016.N().S(p.SiteCf.FooterPartHtml)
//line model/page_base.qtpl:289
		qw422016.N().S(`
`)
//line model/page_base.qtpl:290
	}
//line model/page_base.qtpl:290
	qw422016.N().S(`

`)
//line model/page_base.qtpl:292
}

//line model/page_base.qtpl:292
func (p *BasePage) WriteFooter(qq422016 qtio422016.Writer) {
//line model/page_base.qtpl:292
	qw422016 := qt422016.AcquireWriter(qq422016)
//line model/page_base.qtpl:292
	p.StreamFooter(qw422016)
//line model/page_base.qtpl:292
	qt422016.ReleaseWriter(qw422016)
//line model/page_base.qtpl:292
}

//line model/page_base.qtpl:292
func (p *BasePage) Footer() string {
//line model/page_base.qtpl:292
	qb422016 := qt422016.AcquireByteBuffer()
//line model/page_base.qtpl:292
	p.WriteFooter(qb422016)
//line model/page_base.qtpl:292
	qs422016 := string(qb422016.B)
//line model/page_base.qtpl:292
	qt422016.ReleaseByteBuffer(qb422016)
//line model/page_base.qtpl:292
	return qs422016
//line model/page_base.qtpl:292
}

// Footer ed
