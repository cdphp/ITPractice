<template>
<section>
<div class="wrapper gray-bg">
  <div class="container">
    <div class="row">
      <div class="col-sm-9">
        <div class="box questions" v-for="(item,index) in questions">
          <div class="box-header no-border">
          <div class="title">
            <a href="javascript:void(0)" v-on:click="viewQuestion(item.id)" class="text-blue"><i class="fa fa-question-circle" aria-hidden="true"></i> {{item.title}}</a>
          </div>
          </div>
          <div class="box-content" >

          <div class="content text-gray" v-html="compiledMarkdown(item.content)"></div>

          <div class="footer">
          <span class="middle"><img class="img-circle-head mini-head" :src="item.avatar"> {{item.author}}</span>

          <span class="right">提问于：{{formatTime(item.created_at)}}</span>
          </div>


          </div>

        </div>
        <div class="more" v-if="nomore">没有啦</div>
        <div class="more" v-on:click="loadMore" v-else>查看更多</div>
      </div>

      <div class="col-sm-3">
        <div class="box">
          <div class="box-content">问题总计：20个</div>
        </div>
        <div class="box">
        <a href="#/question/add" class="btn btn-blue  btn-block">我要提问</a>
        </div>
      </div>
    </div>
  </div>
</div>
</section>
</template>
<script>
import {getQuestionListPage} from '../../api/api'
import util from '../../common/js/util'
import marked from 'marked'
export default {
  data() {
    return {
      questions:[],
      page: 1,
      loading: false,
      nomore: false,
    }
  },
  methods: {
    loadMore() {
      this.page++;
      this.getQuestions();
    },
    formatTime(unixTime) {
      return util.formatDate.format(new Date(unixTime*1000),'yy-MM-dd hh:mm');
    },

    //获取文章列表
    getQuestions: function () {
      let para = {
        page : this.page,

      };
      this.loading = true;

      getQuestionListPage(para).then(res => {
        if(res.errorNo == 0) {
          if(res.data) {
            this.questions.push.apply(this.questions, res.data);
          }else {
            this.nomore = true;
          }

          console.log("result:",this.questions);
          this.loading = false;
        }


        //NProgress.done();
      });
    },
    viewQuestion(id) {
      window.open('#/question/info?id='+id,'_blank');
    },
    compiledMarkdown(content) {
      return marked(content.substring(0,200), { sanitize: true })
    }
  },
  mounted() {
    this.getQuestions();
  }
}
</script>
<style scoped>

.questions .title a {
  font-size:24px;
}
.questions .content {
  padding: 10px 0px;
  font-size: 16px;
    line-height: 1.8;
    word-wrap: break-word;
}
.questions .content p {
  padding: 10px 0px;
}
.questions .footer {
  margin:10px 0px;
  font-size:14px;
}


</style>
