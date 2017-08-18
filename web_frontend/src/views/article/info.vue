<template>
<section>
<div class="wrapper gray-bg">
  <div class="container">
    <div class="row">
      <div class="col-sm-9">
        <div class="box">

          <div class="box-content article">
            <div class="title">
              {{article.title}}
            </div>
            <div class="about muted text-center">
              <span class="item">{{article.author_name}}</span>
              <span class="item"><i class="el-icon-time"></i> {{formatTime(article.created_at)}}</span>
            </div>
            <div class="about muted text-right">
              <span class="item"><i class="el-icon-star-off"></i> 20</span>
              <span class="item"><span class="glyphicon glyphicon-eye-open"></span> 20</span>
            </div>
            <div class="content">
              <div v-html="compiledMarkdown"></div>
            </div>
          </div>

        </div>

        <div class="box">
          <div class="box-header">
            发表评论
          </div>
          <div class="box-content">
          <div class="form-horizontal">

            <div class="form-group">
              <label for="content" class="col-sm-1 control-label">内容</label>
              <div class="col-sm-8">
                <textarea rows="4" class="form-control" id="content" placeholder="填写评论" v-model="commentContent"></textarea>
              </div>
            </div>

            <div class="form-group">
              <div class="col-sm-offset-1 col-sm-10">
                <button  class="btn btn-blue" :disable="loading" v-on:click="addComment">提交</button>
              </div>
            </div>
            </div>
          </div>
        </div>

        <div class="box">
          <div class="box-header">
            相关评论({{length}})
          </div>
          <div class="box-content">
            <div class="none" v-if="length==0">暂无内容</div>
            <ul class="comments" v-else>
              <li v-for="(item,index) in comments">
              <div class="media">
                <div class="media-left">
                  <a href="javascript:void(0)" v-on:click="viewUser(item.user_id)">
                    <img class="media-object img-circle-head" :src="item.avatar">
                  </a>
                </div>
                <div class="media-body">
                  <div class="media-heading"><span class="floor muted">#{{index+1}}</span><span>{{item.author.username}}</span></div>
                  {{item.content}}
                </div>
              </div>
              </li>
              <li>
              <div class="more" v-if="nomore">没有啦</div>
              <div class="more" v-on:click="loadMore" v-else>查看更多</div>
              </li>
            </ul>



          </div>
        </div>
      </div>
      <div class="col-sm-3">
        <div class="box">
          <div class="box-header">他的其他文章</div>
        </div>
      </div>
    </div>

  </div>
</div>
</section>
</template>
<script>
import {getArticle, getCommentListPage,addComment} from '../../api/api'
import util from '../../common/js/util'
import marked from 'marked'

export default {
  data() {
    return {
      id:'',
      article: {
        title: '',
        content: '',
      },
      input: '# hello',
      length:0,
      comments:[],
      commentContent:'',
      nomore: false,
      page: 1,
      loading: false,
    }
  },
  methods: {
  loadMore() {
    this.page++;
    this.getComments();
  },
    getArticleInfo(id) {
      let para = {id :id};
      getArticle(para).then(res => {

        if(res.errorNo == 0 ) {
          this.article = res.data;
        }else {
          this.$router.push({ path: '/404' });
        }
      });
    },
    formatTime(unixTime) {
      return util.formatDate.format(new Date(unixTime*1000),'yy-MM-dd');
    },
    viewUser(id) {
      this.$router.push({ path: '/user?id='+id });
    },
    getComments() {
      let para = {target_id :this.id,page: this.page};
      getCommentListPage(para).then(res => {

        if(res.errorNo == 0) {
          this.length = res.total;
          if(res.data) {
            this.comments.push.apply(this.comments, res.data);
          }else {
            this.nomore = true;
          }
        }
      });
    },

    addComment() {
      if(this.commentContent=='') {
        this.$message({
          message: '请先填写评论内容',
          type: 'warning'
        });
        return
      }
      this.loading = true;
      let para = {
        target_id :this.id,
        content: this.commentContent,
        type: "1",
        root_id: "0",
      };
      addComment(para).then(res => {

        if(res.errorNo == 0 ) {
          this.$message({
            message: '评论成功',
            type: 'success'
          });
          this.commentContent = '';
          this.page = 1;
          this.comments = [];
          this.getComments();
        }else {
        if (res.errorNo == 201 ) {
        this.$message({
          message: "请先登录",
          type: 'error'
        });
        }else {
        this.$message({
          message: res.message,
          type: 'error'
        });
        }

        }
        this.loading = false;
      });
    },
  },
  mounted() {
    var id = this.$route.query.id;
    this.id = id;
    this.getArticleInfo(id);
    this.getComments();
  },
  computed: {
    compiledMarkdown: function () {
      return marked(this.article.content, { sanitize: true })
    }
  },
}
</script>
<style scoped>

.article {}
.article .title {
  font-size:28px;
  text-align:center;
  margin:20px;
}
.article .about{padding:10px;}
.article .about span.item{margin:15px;}
.article .content {margin:30px;}
.comments li {
  margin-top:10px;
  margin-bottom:10px;
  padding:10px 0px;
  border-bottom:1px solid #ddd;
}
.comments li  span.floor{margin-right:10px;}
.img-circle-head {width:50px;height:50px;border-radius:50%;}
</style>
