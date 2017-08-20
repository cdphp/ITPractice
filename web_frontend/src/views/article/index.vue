<template>
<section>
<div class="wrapper gray-bg">
  <div class="container">
    <div class="row">
      <div class="col-sm-9">
        <div class="box articles">

          <div class="box-content" v-for="(item,index) in articles">
          <div class="title"><h1><a href="javascript:void(0)" v-on:click="viewArticle(item.id)" class="text-black">{{item.title}}</a></h1></div>

          <div class="content text-gray" v-html="compiledMarkdown(item.content)"></div>
          <div class="footer">
          <span class="muted">作者：{{item.author}}</span>
          <span class="muted right">发布于：{{formatTime(item.created_at)}}</span>
          </div>


          </div>
          <div class="more" v-if="nomore">没有啦</div>
          <div class="more" v-on:click="loadMore" v-else>查看更多</div>
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
import {getArticleListPage} from '../../api/api'
import util from '../../common/js/util'
import marked from 'marked'
export default {
  data() {
    return {
      articles:[],
      page: 1,
      loading: false,
      nomore: false,
    }
  },
  methods: {
    loadMore() {
      this.page++;
      this.getArticles();
    },
    formatTime(unixTime) {
      return util.formatDate.format(new Date(unixTime*1000),'yy-MM-dd hh:mm');
    },

    //获取文章列表
    getArticles: function () {
      let para = {
        page : this.page,

      };
      this.loading = true;

      getArticleListPage(para).then(res => {
        if(res.errorNo == 0) {
          if(res.data) {
            this.articles.push.apply(this.articles, res.data);
          }else {
            this.nomore = true;
          }

          console.log("result:",this.articles);
          this.loading = false;
        }


        //NProgress.done();
      });
    },
    viewArticle(id) {
      window.open('#/article/info?id='+id,'_blank');
    },
    compiledMarkdown(content) {
      return marked(content.substring(0,200), { sanitize: true })
    }
  },
  mounted() {
    this.getArticles();
  }
}
</script>
<style scoped>
.articles .title {
  margin:15px 0px;
}
.articles .content {
  padding: 10px 0px;
  font-size: 16px;
    line-height: 1.8;
    word-wrap: break-word;
}
.articles .content p {
  padding: 10px 0px;
}
.articles .footer {
  margin:10px 0px;
  font-size:14px;
}


</style>
