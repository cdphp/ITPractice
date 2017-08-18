<template>
<section>
<div class="wrapper gray-bg">
  <div class="container">
    <div class="row">
      <div class="col-sm-9">
        <div class="box articles">

          <div class="box-content" v-for="(item,index) in articles" v-key="index">
            <div class="media ">
              <div class="media-left">
                <a href="javascript:void(0)" v-on:click="">
                  <img class="media-object img-circle-head" :src="item.avatar">
                </a>
              </div>
              <div class="media-body">
                <div class="media-heading"><a href="#" class="text-primary">{{item.author}}</a></div>
                <div class="muted">{{formatTime(item.created_at)}}</div>

              </div>

            </div>

            <div class="title"><a href="#" class="text-black"><h4>{{item.title}}</h4></a></div>
            <div class="content text-gray substr">
              <div v-html="item.content"></div>
            </div>


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
import {getArticleListPage} from '../../api/api'
import util from '../../common/js/util'
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
      this.getUsers();
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
    viewUser(id) {
      this.$router.push({ path: '/user?id='+id });
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
  padding: 10px 5px;
}
.substr{
display: -webkit-box;
-webkit-box-orient: vertical;
-webkit-line-clamp: 2;
overflow: hidden;
text-overflow: ellipsis;
}

</style>
