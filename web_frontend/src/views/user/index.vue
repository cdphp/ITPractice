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
                <img :src="info.avatar" />
              </div>
              </div>
              <div class="box-content">
              <h4>{{user.username}}</h4>
              <h6>标签</h6>
              <p class="muted text-indent" v-if="info.labels">{{info.labels}}</p>
              <p class="muted text-indent" v-else>暂无内容</p>
              <h6>个人介绍</h6>
              <p class="muted text-indent" v-if="info.about">{{info.about}}</p>
              <p class="muted text-indent" v-else>暂无内容</p>
              <a href="#/article/add" class="btn btn-blue btn-block follow" v-if="isSelf">发表新文章</a>
              <button class="btn btn-red btn-block follow" v-else>关注</button>

              </div>

            </div>
          </div>

          <div class="col-sm-9">
            <div class="box">
              <div class="box-header">
                <h5>他的文章</h5>
              </div>
              <div class="box-content">
                <div class="none" v-if="articles.length==0">暂无内容</div>
              <ul class="articles" v-if="articles.length">
                <li class="item" v-for="item in articles">
                  <a href="#/article" class="text-blue">
                    <div class="title"><i class="el-icon-information"></i>{{item.title}}</div>
                  </a>

                </li>

              </ul>
              </div>

            </div>

            <div class="box">
              <div class="box-header">
                <h5>谁在关注</h5>
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
                <h5>他关注的人</h5>
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
      info:{},
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
      articles:[
        {
          'title':'OpenStack大规模部署详解',
        },
        {
          'title':'Web前端知识体系精简',
        }
      ],
    }
  },
  methods: {
    getUserInfo(id) {
      let para = {id :id};
      getUser(para).then(res => {

        if(res.errorNo == 0 ) {
          this.user = res.data;
          this.info = res.data.info;
        }else {
          this.$router.push({ path: '/404' });
        }
      });
    },
    getArticles() {
      getArticleListPage().then(res=>{
        console.log(res);
      });
    }
  },
  mounted() {
    var id = this.$route.query.id;

    var user = sessionStorage.getItem('user');

    if (user) {
      user = JSON.parse(user);
      this.isSelf = user.user_id==id?true:false;

    }


    this.getArticles();
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
