<template>
  <section>
    <div class="wrapper gray-bg">
      <div class="container">

        <div class="row">
          <div class="col-sm-3">
            <div class="user-sidebar box">
              <div class="box-header">
                <h5>基本信息</h5>

              </div>
              <div class="box-content no-padding">
              <div class="headpic">
                <img :src="user.avatar" />
              </div>
              </div>
              <div class="box-content">
              <h4>{{user.username}}</h4>
              <h6>标签</h6>
              <p class="muted text-indent" v-if="user.labels">{{user.labels}}</p>
              <p class="muted text-indent" v-else>暂无内容</p>
              <h6>个人介绍</h6>
              <p class="muted text-indent" v-if="user.about">{{user.about}}</p>
              <p class="muted text-indent" v-else>暂无内容</p>
              <a href="#/user/edit" class="btn btn-danger btn-block follow" v-if="isSelf">修改信息</a>
              <button class="btn btn-red btn-block follow" v-else>关注</button>

              </div>

            </div>
          </div>

          <div class="col-sm-9">
            <div class="box">
              <div class="box-header">
                <h5 class="box-title" v-if="isSelf">我的文章</h5>
                <h5 class="box-title" v-else>他的文章</h5>
                <div class="box-tools">
                <a class="btn btn-blue" href="#/article/add" v-if="isSelf">发布文章</a>
                </div>
              </div>
              <div class="box-content">
                <div class="none" v-if="articles.length==0">暂无内容</div>
              <ul class="articles" v-if="articles.length">
                <li class="item" v-for="item in articles">
                  <a href="javascript:void(0)" v-on:click="viewArticle(item.id)" class="text-blue">
                    <div class="title"><i class="el-icon-information"></i>{{item.title}}</div>
                  </a>

                </li>

              </ul>
              </div>

            </div>

            <div class="box">
              <div class="box-header">
              <h5 v-if="isSelf">我的粉丝</h5>
              <h5 v-else>他的粉丝</h5>
              </div>
              <div class="box-content">
                <div class="none" v-if="followers.length==0">暂无内容</div>
                <ul class="followers" v-if="followers.length">
                  <li v-for="item in followers">
                    <a href="#"><img class="img-circle" src="../../assets/imgs/user.png" :title="item.username"/></a>
                  </li>
                </ul>
                <div class="clearfix"></div>
              </div>
            </div>

            <div class="box">
              <div class="box-header">
                <h5 v-if="isSelf">我的偶像</h5>
                <h5 v-else>他的偶像</h5>
              </div>
              <div class="box-content">
                <div class="none" v-if="attentions.length==0">暂无内容</div>
                <ul class="followers" v-if="attentions.length">
                  <li v-for="item in attentions">
                    <a href="#"><img class="img-circle" src="../../assets/imgs/user.png" :title="item.username"/></a>
                  </li>
                </ul>
                <div class="clearfix"></div>
              </div>
            </div>

          </div>
        </div>
      </div>
    </div>
  </section>
</template>
<script>
import {getArticleListPage,getUser} from '../../api/api'
export default {
  data() {
    return {
      user:{},
      canAttention:true,
      isSelf: false,
      followers: [
        {
          username:'hongker',
          avatar:''
        },
        {
          username:'test001',
          avatar:''
        }
      ],
      attentions: [],
      articles:[],
    }
  },
  methods: {
    getUserInfo(id) {
      let para = {id :id};
      getUser(para).then(res => {

        if(res.errorNo == 0 ) {
          this.user = res.data;

          this.getArticles(id);

        }else {
          this.$router.push({ path: '/404' });
        }
      });
    },
    getArticles(id) {
      let para = {uid :id};
      getArticleListPage(para).then(res=>{
        if(res.errorNo == 0 && res.data != null ) {
          this.articles = res.data;
        }
      });
    },
    viewArticle(id) {
      this.$router.push({ path: '/article/info?id='+id });
    }
  },
  mounted() {
    var id = this.$route.query.id;

    var user = sessionStorage.getItem('user');

    if (user) {
      user = JSON.parse(user);
      this.isSelf = user.user_id==id?true:false;

    }

    this.getUserInfo(id);

  }

}
</script>
<style scoped>

  .user-sidebar {
    width:100%;

  }
  .user-sidebar .username {
    padding:10px 0px;
  }
  .user-sidebar .follow {
    margin:20px 0px;
  }

  .headpic {

  }
  .headpic img {
    width:100%;
  }
  .articles {
    width:100%;
  }
  .articles li.item {
    padding:10px 20px;
  }
  .articles li.item a {
    display:block;
    text-decoration:none;
  }

  .articles .title {
    font-size:18px;

  }
  .articles .title i {
    margin-right:10px;
  }
  .followers li {
    width:60px;
    float:left;
    margin:10px;
  }

</style>
