<template>
<section>
<div class="wrapper">
  <div class="container">
    <div class="row">
      <div class="col-sm-8 col-sm-offset-2">


          <div class="article">
            <h1 class="title">
              <a href="#" class="text-blue">{{article.title}}</a>
            </h1>
            <div class="subtitle">
              <div class="media">
                <div class="media-left">
                  <a href="javascript:void(0)" v-on:click="viewUser(article.user_id)">
                    <img class="media-object img-circle-head" :src="article.avatar">
                  </a>
                </div>
                <div class="media-body">
                  <div class="media-heading">
                    <span class="label label-danger">作者</span>
                     {{article.author}}
                  </div>
                  <div class="muted font-small footer">
                    <span>发布于:{{formatTime(article.created_at)}}</span>
                    <span>阅读 {{article.clicks}}</span>
                    <span>评论 {{length}}</span>
                  </div>
                </div>
              </div>
            </div>

              <div v-html="article.content" class="content"></div>

          </div>



        <div class="box no-border">

          <div class="box-content ">
          <div class="form-horizontal">

            <div class="form-group">
              <div class="col-sm-12">
                <textarea rows="3" class="form-control" id="content" placeholder="填写评论" v-model="commentContent"></textarea>
              </div>
            </div>

            <div class="form-group">
              <div class="col-sm-12 text-right">
                <button  class="btn btn-blue" :disable="loading" v-on:click="addComment">评论</button>
              </div>
            </div>
            </div>
          </div>
        </div>

        <div class="box ">
          <div class="box-header no-border">
            相关评论({{length}})
          </div>
          <div class="box-content no-border">
            <div class="none" v-if="length==0">暂无内容</div>
            <div v-else>
            <ul class="comments" >
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

            </ul>

            <div class="more" v-if="nomore">没有啦</div>
            <div class="more" v-on:click="loadMore" v-else>查看更多</div>
            </div>



          </div>
        </div>
      </div>

    </div>

  </div>
</div>
</section>
</template>
<script>
import {getArticle, getCommentListPage,addComment,getUser} from '../../api/api'
import util from '../../common/js/util'

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
      isSelf: false,
      user:{},
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
          this.getUserInfo(this.article.user_id);
        }else {
          this.$router.push({ path: '/404' });
        }
      });
    },
    getUserInfo(id) {
      let para = {id :id};
      getUser(para).then(res => {

        if(res.errorNo == 0 ) {
          this.user = res.data;



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

    var user = localStorage.getItem('user');

    if (user) {
      user = JSON.parse(user);
      this.isSelf = user.user_id==id?true:false;

    }


  }
}
</script>
<style scoped>

.article {}
.article .title {
  font-size:30px;

  margin:20px 0px;
}
.article .title a {
  color:#303030;
  -webkit-transition:all .2s linear;
  -moz-transition:all .2s linear;
  -o-transition:all .2s linear;
}
.article .title a:hover {
  color:#069;
}
.article .subtitle .item{
  margin-right:20px;
}


.article .content {
margin:30px 0px;
padding: 10px 0px;
font-size: 20px;
  line-height: 1.8;
  word-wrap: break-word;
  color:#303030;

}


.comments li {
  margin-top:10px;
  margin-bottom:10px;
  padding:10px 0px;
  border-bottom:1px solid #ddd;
}
.comments li  span.floor{margin-right:10px;}
.img-circle-head {width:50px;height:50px;border-radius:50%;}
</style>
