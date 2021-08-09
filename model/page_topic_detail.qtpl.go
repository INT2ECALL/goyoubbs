// Code generated by qtc from "page_topic_detail.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// 文章详情页，继承 BasePage ，只修改 MainBody()
//

//line model/page_topic_detail.qtpl:3
package model

//line model/page_topic_detail.qtpl:3
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line model/page_topic_detail.qtpl:3
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line model/page_topic_detail.qtpl:3
func (p *TopicDetailPage) StreamMainBody(qw422016 *qt422016.Writer) {
//line model/page_topic_detail.qtpl:3
	qw422016.N().S(`
<div class="entry">

    <article role="article">

        <header>
            <a href="/member/`)
//line model/page_topic_detail.qtpl:9
	qw422016.N().DUL(p.TopicFmt.UserId)
//line model/page_topic_detail.qtpl:9
	qw422016.N().S(`" rel="nofollow"><img alt="`)
//line model/page_topic_detail.qtpl:9
	qw422016.E().S(p.TopicFmt.Name)
//line model/page_topic_detail.qtpl:9
	qw422016.N().S(` avatar" src="/static/avatar/`)
//line model/page_topic_detail.qtpl:9
	qw422016.N().DUL(p.TopicFmt.UserId)
//line model/page_topic_detail.qtpl:9
	qw422016.N().S(`.jpg" class="avatar"></a>
            <h1 class="entry-title">
                <a href="/t/`)
//line model/page_topic_detail.qtpl:11
	qw422016.N().DUL(p.TopicFmt.ID)
//line model/page_topic_detail.qtpl:11
	qw422016.N().S(`" rel="bookmark">`)
//line model/page_topic_detail.qtpl:11
	qw422016.E().S(p.TopicFmt.Title)
//line model/page_topic_detail.qtpl:11
	qw422016.N().S(`</a>
            </h1>
            <p class="meta">
                `)
//line model/page_topic_detail.qtpl:14
	qw422016.N().S(p.TopicFmt.ClockEmoji)
//line model/page_topic_detail.qtpl:14
	qw422016.N().S(` <time datetime="`)
//line model/page_topic_detail.qtpl:14
	qw422016.E().S(p.TopicFmt.AddTimeFmt)
//line model/page_topic_detail.qtpl:14
	qw422016.N().S(`" pubdate data-updated="true">`)
//line model/page_topic_detail.qtpl:14
	qw422016.E().S(p.TopicFmt.AddTimeFmt)
//line model/page_topic_detail.qtpl:14
	qw422016.N().S(`</time>
                by <a href="/member/`)
//line model/page_topic_detail.qtpl:15
	qw422016.N().DUL(p.TopicFmt.UserId)
//line model/page_topic_detail.qtpl:15
	qw422016.N().S(`" rel="nofollow">`)
//line model/page_topic_detail.qtpl:15
	qw422016.E().S(p.TopicFmt.Name)
//line model/page_topic_detail.qtpl:15
	qw422016.N().S(`</a>
                `)
//line model/page_topic_detail.qtpl:16
	if p.CurrentUser.Flag >= 99 {
//line model/page_topic_detail.qtpl:16
		qw422016.N().S(`
                &bull; <a rel="bookmark" href="/admin/topic/edit?id=`)
//line model/page_topic_detail.qtpl:17
		qw422016.N().DUL(p.TopicFmt.ID)
//line model/page_topic_detail.qtpl:17
		qw422016.N().S(`&back=here">Edit</a>
                `)
//line model/page_topic_detail.qtpl:18
	}
//line model/page_topic_detail.qtpl:18
	qw422016.N().S(`
            </p>
        </header>

        <div class="markdown-body entry-content">
            `)
//line model/page_topic_detail.qtpl:23
	qw422016.N().S(p.TopicFmt.ContentFmt)
//line model/page_topic_detail.qtpl:23
	qw422016.N().S(`
        </div>

        `)
//line model/page_topic_detail.qtpl:26
	if len(p.TopicFmt.Relative) > 0 {
//line model/page_topic_detail.qtpl:26
		qw422016.N().S(`
        <section>
            <h4 class="seealso-title">💘 相关文章</h4>
            <ul class="seealso">
                `)
//line model/page_topic_detail.qtpl:30
		for _, v := range p.TopicFmt.Relative {
//line model/page_topic_detail.qtpl:30
			qw422016.N().S(`
                <li><a href="/t/`)
//line model/page_topic_detail.qtpl:31
			qw422016.N().DUL(v.ID)
//line model/page_topic_detail.qtpl:31
			qw422016.N().S(`" rel="bookmark">`)
//line model/page_topic_detail.qtpl:31
			qw422016.E().S(v.Title)
//line model/page_topic_detail.qtpl:31
			qw422016.N().S(`</a></li>
                `)
//line model/page_topic_detail.qtpl:32
		}
//line model/page_topic_detail.qtpl:32
		qw422016.N().S(`
            </ul>
        </section>
        `)
//line model/page_topic_detail.qtpl:35
	}
//line model/page_topic_detail.qtpl:35
	qw422016.N().S(`

        <footer>

            <p class="meta gray">
                <span class="categories">
                    📁 Category: <a class="category" href="/n/`)
//line model/page_topic_detail.qtpl:41
	qw422016.N().DUL(p.TopicFmt.NodeId)
//line model/page_topic_detail.qtpl:41
	qw422016.N().S(`" rel="category tag">`)
//line model/page_topic_detail.qtpl:41
	qw422016.E().S(p.DefaultNode.Name)
//line model/page_topic_detail.qtpl:41
	qw422016.N().S(`</a>
                </span>
                `)
//line model/page_topic_detail.qtpl:43
	if len(p.TagLst) > 0 {
//line model/page_topic_detail.qtpl:43
		qw422016.N().S(`
                <span class="categories">🏷️ Tags:
                    `)
//line model/page_topic_detail.qtpl:45
		for _, tag := range p.TagLst {
//line model/page_topic_detail.qtpl:45
			qw422016.N().S(`
                    <a class="tag" href="/tag/`)
//line model/page_topic_detail.qtpl:46
			qw422016.N().U(tag.Name)
//line model/page_topic_detail.qtpl:46
			qw422016.N().S(`" rel="tag">`)
//line model/page_topic_detail.qtpl:46
			qw422016.E().S(tag.Name)
//line model/page_topic_detail.qtpl:46
			qw422016.N().S(`</a>
                    `)
//line model/page_topic_detail.qtpl:47
		}
//line model/page_topic_detail.qtpl:47
		qw422016.N().S(`
                </span>
                `)
//line model/page_topic_detail.qtpl:49
	}
//line model/page_topic_detail.qtpl:49
	qw422016.N().S(`
                <span class="categories">
                💬 <a href="#comments">Comments (`)
//line model/page_topic_detail.qtpl:51
	qw422016.N().DUL(p.TopicFmt.Comments)
//line model/page_topic_detail.qtpl:51
	qw422016.N().S(`)</a> 😊 PageView (`)
//line model/page_topic_detail.qtpl:51
	qw422016.N().DUL(p.TopicFmt.Views)
//line model/page_topic_detail.qtpl:51
	qw422016.N().S(`)
                </span>
            </p>

            <ul class="seealso">
                `)
//line model/page_topic_detail.qtpl:56
	if p.NewTopic.ID > 0 {
//line model/page_topic_detail.qtpl:56
		qw422016.N().S(`
                <li>上一篇 › <a class="next" href="/t/`)
//line model/page_topic_detail.qtpl:57
		qw422016.N().DUL(p.NewTopic.ID)
//line model/page_topic_detail.qtpl:57
		qw422016.N().S(`" rel="next">`)
//line model/page_topic_detail.qtpl:57
		qw422016.E().S(p.NewTopic.Title)
//line model/page_topic_detail.qtpl:57
		qw422016.N().S(`</a></li>
                `)
//line model/page_topic_detail.qtpl:58
	}
//line model/page_topic_detail.qtpl:58
	qw422016.N().S(`
                `)
//line model/page_topic_detail.qtpl:59
	if p.OldTopic.ID > 0 {
//line model/page_topic_detail.qtpl:59
		qw422016.N().S(`
                <li>下一篇 › <a class="prev" href="/t/`)
//line model/page_topic_detail.qtpl:60
		qw422016.N().DUL(p.OldTopic.ID)
//line model/page_topic_detail.qtpl:60
		qw422016.N().S(`" rel="prev">`)
//line model/page_topic_detail.qtpl:60
		qw422016.E().S(p.OldTopic.Title)
//line model/page_topic_detail.qtpl:60
		qw422016.N().S(`</a></li>
                `)
//line model/page_topic_detail.qtpl:61
	}
//line model/page_topic_detail.qtpl:61
	qw422016.N().S(`
            </ul>

        </footer>

    </article>

    `)
//line model/page_topic_detail.qtpl:68
	if len(p.CommentLst) > 0 {
//line model/page_topic_detail.qtpl:68
		qw422016.N().S(`
    <section>
        <h1 class="br-nav">评论</h1>
        <div id="comments">
            <h2 class="comments bot-line">共`)
//line model/page_topic_detail.qtpl:72
		qw422016.N().DUL(p.TopicFmt.Comments)
//line model/page_topic_detail.qtpl:72
		qw422016.N().S(`条关于"`)
//line model/page_topic_detail.qtpl:72
		qw422016.E().S(p.TopicFmt.Title)
//line model/page_topic_detail.qtpl:72
		qw422016.N().S(`"的评论</h2>
            `)
//line model/page_topic_detail.qtpl:73
		for _, item := range p.CommentLst {
//line model/page_topic_detail.qtpl:73
			qw422016.N().S(`
            <article id="r`)
//line model/page_topic_detail.qtpl:74
			qw422016.N().DUL(item.ID)
//line model/page_topic_detail.qtpl:74
			qw422016.N().S(`">
                <header>
                    <a href="/member/`)
//line model/page_topic_detail.qtpl:76
			qw422016.N().DUL(item.UserId)
//line model/page_topic_detail.qtpl:76
			qw422016.N().S(`" rel="nofollow"><img alt="`)
//line model/page_topic_detail.qtpl:76
			qw422016.E().S(item.Name)
//line model/page_topic_detail.qtpl:76
			qw422016.N().S(` avatar" src="/static/avatar/`)
//line model/page_topic_detail.qtpl:76
			qw422016.N().DUL(item.UserId)
//line model/page_topic_detail.qtpl:76
			qw422016.N().S(`.jpg" class="avatar"></a>
                    <div class="meta gray5">
                        <a href="`)
//line model/page_topic_detail.qtpl:78
			qw422016.E().S(item.Link)
//line model/page_topic_detail.qtpl:78
			qw422016.N().S(`" class="comment-count comment-id">#`)
//line model/page_topic_detail.qtpl:78
			qw422016.N().DUL(item.ID)
//line model/page_topic_detail.qtpl:78
			qw422016.N().S(`</a> <a href="/member/`)
//line model/page_topic_detail.qtpl:78
			qw422016.N().DUL(item.UserId)
//line model/page_topic_detail.qtpl:78
			qw422016.N().S(`" rel="nofollow">`)
//line model/page_topic_detail.qtpl:78
			qw422016.E().S(item.Name)
//line model/page_topic_detail.qtpl:78
			qw422016.N().S(`</a>
                        `)
//line model/page_topic_detail.qtpl:79
			if p.CurrentUser.Flag >= 99 {
//line model/page_topic_detail.qtpl:79
				qw422016.N().S(`
                        <a href="/admin/comment/edit?tid=`)
//line model/page_topic_detail.qtpl:80
				qw422016.N().DUL(item.TopicId)
//line model/page_topic_detail.qtpl:80
				qw422016.N().S(`&cid=`)
//line model/page_topic_detail.qtpl:80
				qw422016.N().DUL(item.ID)
//line model/page_topic_detail.qtpl:80
				qw422016.N().S(`&back=here">Edit</a>
                        `)
//line model/page_topic_detail.qtpl:81
			}
//line model/page_topic_detail.qtpl:81
			qw422016.N().S(`
                        <a rel="nofollow" class="right" href="#respond" onclick="replyTo('`)
//line model/page_topic_detail.qtpl:82
			qw422016.E().S(item.Name)
//line model/page_topic_detail.qtpl:82
			qw422016.N().S(`',`)
//line model/page_topic_detail.qtpl:82
			qw422016.N().DUL(item.ID)
//line model/page_topic_detail.qtpl:82
			qw422016.N().S(`)">回复</a>
                        <br>
                        <time datetime="`)
//line model/page_topic_detail.qtpl:84
			qw422016.N().DL(item.AddTime)
//line model/page_topic_detail.qtpl:84
			qw422016.N().S(`" pubdate data-updated="true"><em>`)
//line model/page_topic_detail.qtpl:84
			qw422016.E().S(item.AddTimeFmt)
//line model/page_topic_detail.qtpl:84
			qw422016.N().S(`</em></time>
                    </div>
                </header>
                <div class="markdown-body entry-content">
                    `)
//line model/page_topic_detail.qtpl:88
			qw422016.N().S(item.ContentFmt)
//line model/page_topic_detail.qtpl:88
			qw422016.N().S(`
                </div>
            </article>
            `)
//line model/page_topic_detail.qtpl:91
		}
//line model/page_topic_detail.qtpl:91
		qw422016.N().S(`
        </div>
    </section>
    `)
//line model/page_topic_detail.qtpl:94
	}
//line model/page_topic_detail.qtpl:94
	qw422016.N().S(`

    <section id="respond">
        <h2 class="br-nav">写一条评论</h2>
        <div class="write-comment pure-form">
            <form action="" method="post" id="commentform">
                <textarea name="comment" id="id-comment" class="pure-u-1" placeholder="* 字符限制 `)
//line model/page_topic_detail.qtpl:100
	qw422016.N().D(p.SiteCf.CommentConMaxLen)
//line model/page_topic_detail.qtpl:100
	qw422016.N().S(`"></textarea>
                `)
//line model/page_topic_detail.qtpl:101
	if p.CurrentUser.ID > 0 {
//line model/page_topic_detail.qtpl:101
		qw422016.N().S(`
                <div class="pure-button-group">
                    <input id="btn-preview" type="button" value="预览" name="submit" onclick="previewComment(); return false;" class="pure-button button-success" />
                    <input id="btn-submit" type="button" value="发表" name="submit" onclick="submitComment(); return false;" class="pure-button pure-button-primary" />
                </div>
                <span id="id-msg"></span>
                `)
//line model/page_topic_detail.qtpl:107
	} else {
//line model/page_topic_detail.qtpl:107
		qw422016.N().S(`
                <a href="/login" rel="nofollow" class="pure-button">登录发表评论</a>
                `)
//line model/page_topic_detail.qtpl:109
	}
//line model/page_topic_detail.qtpl:109
	qw422016.N().S(`
            </form>
        </div>
        <div id="id-preview" class="markdown-body entry-content"></div>

        <script>

            var toReplyId = 0;
            var conEle = document.getElementById("id-comment");
            var msgEle = document.getElementById("id-msg");
            var reviewEle = document.getElementById("id-preview");

            `)
//line model/page_topic_detail.qtpl:121
	if p.CurrentUser.ID > 0 {
//line model/page_topic_detail.qtpl:121
		qw422016.N().S(`
                function previewComment() {
                    var con = conEle.value.trim();
                    if (con === "") {
                        conEle.focus();
                        return
                    }
                    postAjax("/content/preview", JSON.stringify({Act: "commentPreview", Content: con}), function(data){
                        var obj = JSON.parse(data)
                        //console.log(obj);
                        if(obj.Code === 200) {
                            msgEle.style.display = "none";
                            reviewEle.innerHTML = obj.Html;
                            reviewEle.style.display = "block";
                        }else{
                            reviewEle.innerHTML = "";
                            reviewEle.style.display = "none";
                            msgEle.innerText = obj.Msg;
                        }
                    });
                }
                function submitComment() {
                    var con = conEle.value.trim();
                    if (con === "") {
                        conEle.focus();
                        return
                    }
                    postAjax("/t/`)
//line model/page_topic_detail.qtpl:148
		qw422016.N().DUL(p.TopicFmt.ID)
//line model/page_topic_detail.qtpl:148
		qw422016.N().S(`", JSON.stringify({Content: con, ReplyId: toReplyId}), function(data){
                        var obj = JSON.parse(data)
                        msgEle.innerText = obj.Msg;
                        conEle.focus();
                        conEle.value = "";
                        toReplyId = 0;
                        if(obj.Code === 200) {
                            window.location.href = "/t/`)
//line model/page_topic_detail.qtpl:155
		qw422016.N().DUL(p.TopicFmt.ID)
//line model/page_topic_detail.qtpl:155
		qw422016.N().S(`#r"+obj.Tid;
                            window.location.reload(true);
                            return;
                        } else if (obj.Code === 201) {
                            window.location.href = "/member/`)
//line model/page_topic_detail.qtpl:159
		qw422016.N().DUL(p.CurrentUser.ID)
//line model/page_topic_detail.qtpl:159
		qw422016.N().S(`?type=comment";
                            return;
                        }
                        reviewEle.style.display = "none";
                        msgEle.style.display = "block";
                    });
                }
                `)
//line model/page_topic_detail.qtpl:166
		if !p.SiteCf.UploadLimit || (p.SiteCf.UploadLimit && p.CurrentUser.Flag >= 99) {
//line model/page_topic_detail.qtpl:166
			qw422016.N().S(`
                document.addEventListener('paste', function (evt) {
                    var url = "/file/upload";
                    var items = evt.clipboardData && evt.clipboardData.items;
                    var file = null;
                    if(items && items.length) {
                        for(var i=0; i!==items.length; i++) {
                            if(items[i].type.indexOf('image') !== -1) {
                                file = items[i].getAsFile();
                                if(!!!file) {
                                    continue;
                                }

                                // upload file object.
                                var form = new FormData();
                                form.append('file', file);

                                postAjax("/file/upload", form, function(data){
                                    let obj = JSON.parse(data)
                                    //console.log(obj);
                                    if(obj.Code === 200) {
                                        let img_url = "\n" + obj.Url + "\n";
                                        let pos = conEle.selectionStart;
                                        let con = conEle.value;
                                        conEle.value = con.slice(0, pos) + img_url + con.slice(pos);
                                    }else{
                                        console.warn(obj.Msg);
                                    }
                                });
                            }
                        }
                    }

                });
                `)
//line model/page_topic_detail.qtpl:200
		}
//line model/page_topic_detail.qtpl:200
		qw422016.N().S(`
            `)
//line model/page_topic_detail.qtpl:201
	}
//line model/page_topic_detail.qtpl:201
	qw422016.N().S(`

            function replyTo(name, cid) {
                toReplyId = parseInt(cid, 10);
                var con = conEle.value;
                document.getElementsByTagName('textarea')[0].focus();
                conEle.value = " @"+name+" #" + cid + " " + con;
                return false
            }

            function linkClick() {
                //console.log("ele clicked:"+this.href);
                var curEle = this;
                postAjax("/get/link/count", JSON.stringify({Act: "set", Items: [this.href]}), function(data){
                    var obj = JSON.parse(data)
                    if(obj.Code===200) {
                        var num = obj.Num;
                        var nextEle = curEle.nextElementSibling;
                        if(nextEle && nextEle.nodeType === 1) {
                            nextEle.innerHTML = num;
                            nextEle.title = num + "次点击"
                        }else{
                            var newNode = document.createElement('span');
                            newNode.innerHTML = num;
                            newNode.classList.toggle("clicks");
                            newNode.title = num + "次点击"
                            curEle.after(newNode);
                        }
                    }
                });
            }

            function getContentLinkCount() {
                var urlEleDict = {};
                var linkDict = {};
                var conLst = document.querySelectorAll(".entry-content a");
                for (var i = 0, max = conLst.length; i < max; i++) {
                    var aEle = conLst[i];
                    if(!aEle.classList.contains('anchor') && !aEle.getAttribute("href").startsWith('/name/')) {
                        var fullLink = aEle.href;
                        linkDict[aEle.href] = null;
                        if (urlEleDict.hasOwnProperty(fullLink)) {
                            urlEleDict[fullLink].push(aEle)
                        }else{
                            urlEleDict[fullLink] = [aEle];
                        }
                        if(aEle.getAttribute("target")==="_blank"){
                            aEle.classList.toggle("external_link");
                        }
                        aEle.addEventListener("click", linkClick);
                    }
                }
                var linkArray = Object.keys(linkDict);
                if(linkArray.length > 0) {
                    postAjax("/get/link/count", JSON.stringify({Items: linkArray}), function(data){
                        var obj = JSON.parse(data)
                        if(obj.Code===200) {
                            Object.keys(obj.Info).forEach(function(key) {
                                var num = obj.Info[key];
                                if (urlEleDict.hasOwnProperty(key)) {
                                    for (var i = 0, max = urlEleDict[key].length; i < max; i++) {
                                        var newNode = document.createElement('span');
                                        newNode.innerHTML = num;
                                        newNode.classList.toggle("clicks");
                                        newNode.title = num + "次点击"
                                        var tmpEle = urlEleDict[key][i];
                                        tmpEle.after(newNode);
                                    }
                                }
                            });
                        }
                    });
                }
            }

            docReady(function() {
                getContentLinkCount();
            });

        </script>

    </section>

</div>

`)
//line model/page_topic_detail.qtpl:286
}

//line model/page_topic_detail.qtpl:286
func (p *TopicDetailPage) WriteMainBody(qq422016 qtio422016.Writer) {
//line model/page_topic_detail.qtpl:286
	qw422016 := qt422016.AcquireWriter(qq422016)
//line model/page_topic_detail.qtpl:286
	p.StreamMainBody(qw422016)
//line model/page_topic_detail.qtpl:286
	qt422016.ReleaseWriter(qw422016)
//line model/page_topic_detail.qtpl:286
}

//line model/page_topic_detail.qtpl:286
func (p *TopicDetailPage) MainBody() string {
//line model/page_topic_detail.qtpl:286
	qb422016 := qt422016.AcquireByteBuffer()
//line model/page_topic_detail.qtpl:286
	p.WriteMainBody(qb422016)
//line model/page_topic_detail.qtpl:286
	qs422016 := string(qb422016.B)
//line model/page_topic_detail.qtpl:286
	qt422016.ReleaseByteBuffer(qb422016)
//line model/page_topic_detail.qtpl:286
	return qs422016
//line model/page_topic_detail.qtpl:286
}
