<template>
<section>
<div class="wrapper gray-bg">
  <div class="container">
    <div class="row">
      <div class="col-sm-8 col-sm-offset-2">
          <div class="question box">
            <div class="box-content">
              <h1 class="title">
                <a href="#" class="text-blue">{{question.title}}</a>
              </h1>
              <div class="subtitle">
                <div class="media">
                  <div class="media-left">
                    <a href="javascript:void(0)" v-on:click="viewUser(question.user_id)">
                      <img class="media-object img-circle-head" :src="question.avatar">
                    </a>
                  </div>
                  <div class="media-body">
                    <div class="media-heading">
                      <span class="label label-danger">作者</span>
                       {{question.author}}
                    </div>
                    <div class="muted font-small footer">
                      <span>提问于{{formatTime(question.created_at)}}</span>
                      <span>阅读 {{question.clicks}}</span>
                      <span>回答 {{length}}</span>
                    </div>
                  </div>
                </div>
              </div>
                <div v-html="question.content" class="content"></div>
            </div>
          </div>


        <div class="box ">
          <div class="box-header no-border">
            相关回答({{length}})
          </div>
          <div class="none" v-if="length==0">暂无内容</div>
          <div class="box-content answers" v-else v-for="(item,index) in answers">



              <div class="media">
                <div class="media-left">

                      <a href="javascript:void(0)" v-on:click="evalute(item.id,1)" class="text-primary" data-toggle="tooltip" data-placement="top" title="这问题答得好，我要点赞">
                          <i class="fa fa-chevron-up"> </i>
                      </a>
                      <div class="text-center">{{item.approval}}</div>
                      <a href="javascript:void(0)" v-on:click="evalute(item.id,2)" class="text-primary"  data-toggle="tooltip" data-placement="bottom" title="不靠谱">
                          <i class="fa fa-chevron-down"> </i>
                      </a>

                </div>
                <div class="media-body">

                  <div v-html="item.content" class="content"></div>
                  <div class="footer muted">
                    <span>{{formatTime(item.created_at,'MM月dd')}}回答</span>
                    <span class="middle right"><img class="mini-head" :src="item.avatar"> {{item.author}}</span>
                  </div>
                </div>
              </div>
            </div>


            </div>
            <div class="more" v-if="nomore">没有啦</div>
            <div class="more" v-on:click="loadMore" v-else>查看更多</div>




        <div class="box no-border">
          <div class="box-content ">
          <div class="form-horizontal">

            <div class="form-group">
              <div class="col-sm-12">
                <vue-html5-editor :content="answerContent" @change="updateData" ref="editor" :height="200"></vue-html5-editor>
              </div>
            </div>

            <div class="form-group">
              <div class="col-sm-12 text-right">
                <button  class="btn btn-blue" :disable="loading" v-on:click="addAnswer">提交</button>
              </div>
            </div>
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
import {getQuestion, getAnswerListPage,addAnswer,getUser,evaluteAnswer} from '../../api/api'
import util from '../../common/js/util'

export default {

  data() {
    return {
      id:'',
      question: {
        title: '',
        content: '',
      },
      input: '# hello',
      length:0,
      answers:[],
      answerContent:'我要回答..',
      nomore: false,
      page: 1,
      loading: false,
      isSelf: false,
      user:{},
    }
  },
  
  methods: {
    updateData: function (data) {
        // sync content to component
        this.answerContent = data
    },
  loadMore() {
    this.page++;
    this.getAnswers();
  },
    getQuestionInfo(id) {
      let para = {id :id};
      getQuestion(para).then(res => {

        if(res.errorNo == 0 ) {
          this.question = res.data;

        }else {
          this.$router.push({ path: '/404' });
        }
      });
    },

    formatTime(unixTime,format='yy-MM-dd hh:mm') {
      return util.formatDate.format(new Date(unixTime*1000),format);
    },
    viewUser(id) {
      this.$router.push({ path: '/user?id='+id });
    },
    getAnswers() {
      let para = {question_id :this.id,page: this.page};
      getAnswerListPage(para).then(res => {

        if(res.errorNo == 0) {
          this.length = res.total;
          if(res.data) {
            this.answers.push.apply(this.answers, res.data);
          }else {
            this.nomore = true;
          }
        }
      });
    },
    evalute(id, type) {
      if(this.loading==true) {
        this.$message({ message: '正在操作，请等待', type: 'warning'});
      }
      this.loading = true;
      let para = {answer_id: id, type: type};
      evaluteAnswer(para).then(res => {
        if(res.errorNo == 0 ) {
          this.$message({ message: '点评成功', type: 'success'});
          this.page = 1;
          this.answers = [];
          this.getAnswers();

        }else {
        this.$message({ message: res.message, type: 'error' });
        }
        this.loading = false;

      });
    },

    addAnswer() {
      if(this.answerContent=='') {
        this.$message({
          message: '请先填写评论内容',
          type: 'warning'
        });
        return
      }
      this.loading = true;
      let para = {
        question_id :this.id,
        content: this.answerContent,
        type: "1",
        root_id: "0",
      };
      addAnswer(para).then(res => {

        if(res.errorNo == 0 ) {
          this.$message({
            message: '评论成功',
            type: 'success'
          });
          this.answerContent = '';
          this.page = 1;
          this.answers = [];
          this.getAnswers();
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
    }
  },
  mounted() {
    var id = this.$route.query.id;
    this.id = id;
    this.getQuestionInfo(id);
    this.getAnswers();

    var user = localStorage.getItem('user');

    if (user) {
      user = JSON.parse(user);
      this.isSelf = user.user_id==id?true:false;

    }


  },

}
</script>
<style scoped>

.question {}
.question .title {
  font-size:30px;
font-weight: 700;
  margin:20px 0px;
}
.question .title a {
  color:#303030;
  -webkit-transition:all .2s linear;
  -moz-transition:all .2s linear;
  -o-transition:all .2s linear;
}
.question .title a:hover {
  color:#069;
}
.question .subtitle .item{
  margin-right:20px;
}


.question .content {
margin:30px 0px;
padding: 10px 0px;
font-size: 20px;
  line-height: 1.8;
  word-wrap: break-word;
  color:#303030;

}

.answers {font-size:14px;}
.answers li {
  margin-top:10px;
  margin-bottom:10px;
  padding:10px 0px;
  border-bottom:1px solid #ddd;
}
.answers li  span.floor{margin-right:10px;}
.img-circle-head {width:50px;height:50px;border-radius:50%;}
</style>
