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
              <div v-html="article.content"></div>
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
                <button  class="btn btn-blue" v-on:click="addComment">提交</button>
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
                    <img class="media-object img-circle-head" :src="item.author.info.avatar">
                  </a>
                </div>
                <div class="media-body">
                  <div class="media-heading"><span class="floor muted">#{{index+1}}</span><span>{{item.author.username}}</span></div>
                  {{item.content}}
                </div>
              </div>
              </li>
              <li>
                <div class="more">查看更多</div>
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
export default {
  data() {
    return {
      id:'',
      article: {
        title: '',
        content: '',
      },
      length:0,
      comments:[],
      commentContent:'',
    }
  },
  methods: {
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
    getComments(id) {
      let para = {target_id :id};
      getCommentListPage(para).then(res => {

        if(res.errorNo == 0 &&res.data!=null) {
          this.comments = res.data;
          this.length = this.comments.length;
        }
      });
    },

    addComment() {
      let para = {
        target_id :this.id,
        content: this.commentContent,
      };
      addComment(para).then(res => {

        if(res.errorNo == 0 ) {
          this.$message({
            message: '评论成功',
            type: 'success'
          });
          this.commentContent = '';
          this.getComments(para.target_id);
        }else {
        this.$message({
          message: res.errorMsg,
          type: 'error'
        });
        }
      });
    },
  },
  mounted() {
    var id = this.$route.query.id;
    this.id = id;
    this.getArticleInfo(id);
    this.getComments(id);
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
