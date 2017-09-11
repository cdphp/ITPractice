<template>
  <section>
    <div class="wrapper gray-bg">
      <div class="container">

        <div class="row">
          <div class="col-sm-3">
            <div class="user-sidebar box">

              <div class="box-content no-border headpic">
              <div class=" userinfo">
                <img :src="user.avatar" />
              </div>
              <h4 class="text-center">{{user.username}}</h4>
              <p class="labels text-center">{{formatAuth(user.auth)}}</p>


              <div class="btns" v-if="isSelf">
                <a href="#/article/add" class="btn btn-blue btn-block btn-ellipse follow">发布文章</a>
                <a href="#/user/edit" class="btn btn-default btn-block btn-ellipse follow">修改信息</a>
              </div>
              <div class="btns" v-else>

                <button v-if="!isMaster && user.auth!='User'" class="btn btn-blue btn-block btn-ellipse follow" v-on:click="follow" :disable="following">拜 师 <i class="fa fa-plus" aria-hidden="true"></i></button>
                <button class="btn btn-default btn-block btn-ellipse follow">私 信 <i class="fa fa-envelope-o" aria-hidden="true"></i></button>
              </div>
              </div>
              <div class="box-content counts">

                  <div class="item">
                    <div class="num">{{totalArticle}}</div>
                    <div class="title">文章</div>
                  </div>
                  <div class="item">
                    <div class="num">{{totalPupil}}</div>
                    <div class="title">弟子</div>
                  </div>

                <div class="clearfix"></div>
              </div>
              <div class="box-content">
              <p class="about text-indent">{{user.labels}}</p>
              <p class="about text-indent" v-if="user.about">{{user.about}}</p>
              <p class="about text-indent" v-else>暂无内容</p>
              </div>
              <div class="box-content" v-if="user.github">

                  <p class="github"><a v-bind:href="user.github" target="_blank" class="text-blue" ><i class="fa fa-github" aria-hidden="true"></i> {{user.github}}</a></p>

              </div>

            </div>
          </div>

          <div class="col-sm-6">
            <div class="box">
              <div class="box-header no-border">
                <h5 class="box-title">文章</h5>
              </div>
              <div class="box-content"  v-if="articles.length==0">
                <div class="none ">暂无数据</div>
              </div>
              <div class="box-content articles" v-for="item in articles">
                <div class="item">
                  <a href="javascript:void(0)" v-on:click="viewArticle(item.id)" class="title text-blue">
                  <i class="fa fa-info-circle" aria-hidden="true"></i> {{item.title}}
                  <span class="muted right time">{{formatTime(item.created_at)}}</span>
                  </a>
                  <div class="footer">
                    <span class="muted">点击量：{{item.clicks}}</span>
                    <span class="muted right" v-if="isSelf">
                      <a href="javascript:void(0)" v-on:click="editArticle(item.id)" class="text-blue"><i class="fa fa-edit" aria-hidden="true"></i></a>
                      <a href="javascript:void(0)" v-on:click="delArticle(item.id)" class="text-red"><i class="fa fa-trash-o" aria-hidden="true"></i></a>
                    </span>
                  </div>
                </div>
              </div>

            </div>

          </div>
          <div class="col-sm-3">
          <div class="box">
            <div class="box-header no-border">

              <h5 class="box-title">徒弟</h5>
            </div>
            <div class="box-content">
              <div class="none" v-if="pupils.length==0">暂无内容</div>
              <ul class="follows" v-else>
                <li v-for="(item,index) in pupils">
                <div class="media">
                  <div class="media-left">
                    <a href="javascript:void(0)" >
                      <img class="media-object img-circle-head" :src="item.avatar">
                    </a>
                  </div>
                  <div class="media-body">
                    <div class="media-heading"><span class="username">{{item.username}}</span>
                    </div>
                    <div class="labels">{{item.labels}}</div>
                  </div>
                </div>
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
import {getArticleListPage,getUser,delArticle,getRelationListPage, addRelation} from '../../api/api'
import util from '../../common/js/util'

export default {
  data() {
    return {
      user:{},
      canAttention:true,
      isSelf: false,
      pupils: [],
      attentions: [],
      articles:[],
      following: false,
      isMaster: false,
      totalArticle:0,
      totalPupil:0,
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
  formatAuth(auth) {
    return util.getAuthName(auth);
  },
  delArticle(id) {
    this.$confirm('此操作将永久删除该文件, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        let para = {id: id};
        delArticle(para).then(res => {

          if(res.errorNo==0) {
          this.$message({
            type: 'success',
            message: '删除成功!'
          });
          }else {
          this.$message({
            type: 'error',
            message: res.message
          });
          }
        });

      }).catch(() => {
        this.$message({
          type: 'info',
          message: '已取消删除'
        });
      });
  },

    getUserInfo(id) {
      let para = {id :id};
      getUser(para).then(res => {

        if(res.errorNo == 0 ) {
          this.user = res.data;
          this.isMaster = this.user.is_master;
          this.getArticles(id);
          this.getPupils(id);

        }else {
          this.$router.push({ path: '/404' });
        }
      });
    },
    getArticles(id) {
      let para = {uid :id};
      getArticleListPage(para).then(res=>{
        if(res.errorNo == 0) {
          this.articles = res.data == null ? []:res.data;
          this.totalArticle = res.total;
        }
      });
    },
    getPupils(id) {
      let para = {uid: id};
      getRelationListPage(para).then(res=>{
        if(res.errorNo == 0) {
          this.pupils = res.data == null ? []:res.data;
          this.totalPupil = res.total;
        }
      });
    },
    viewArticle(id) {
      this.$router.push({ path: '/article/info?id='+id });
    },
    editArticle(id) {
      this.$router.push({ path: '/article/edit?id='+id });
    },

    follow() {
      this.following = true;
      this.$confirm('确定要拜其为师么?', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
        }).then(() => {
          let para = {master_id: this.user.id};
          addRelation(para).then(res => {
            this.following = false;
            if(res.errorNo==0) {
              this.$message({
                type: 'success',
                message: '拜师成功!'
              });
              this.isMaster = true;
              this.getPupils(this.user.id)
            }else {
              this.$message({
                type: 'error',
                message: res.message
              });
              }
          });

        }).catch(() => {
          this.following = false;
          this.$message({
            type: 'info',
            message: '已取消'
          });
        });

    },
    fetchData() {
      var id = this.$route.query.id;

      var user = localStorage.getItem('user');

      if (user) {
        user = JSON.parse(user);
        this.isSelf = user.user_id==id?true:false;

      }
      if(id==undefined) {
        id = user.user_id;
      }

      this.getUserInfo(id);
    }
  },
  mounted() {
    this.fetchData();

  },
  watch:{
      '$route':'fetchData'
  },


}
</script>
<style scoped>

.user-sidebar {
  width:100%;
  margin:10px;
}
.user-sidebar .headpic {
  padding-bottom: 0px;
  margin:10px 5px 5px 5px;
}

  .user-sidebar .btns{
    padding:0px 50px;

  }
  .user-sidebar .username {
    padding:10px 0px;
  }
  .userinfo {
    text-align: center;
  }

  .headpic {
    margin: 20px auto;

  }
  .headpic p {font-size:14px;}
  .headpic img {
    width:70px;
    height:70px;
    border-radius: 50%;
  }
  .labels, .about{
    margin: 10px 0px;
    font-size:14px;

  }
  .labels {
    color:#888;
  }
  .about {
    color:#333;
  }

  .user-sidebar .follow {
    margin:20px 0px;

  }


  .articles {
    width:100%;
  }

  .articles .item a.title {
    font-size:18px;
    padding:5px 0px;
    display:block;
    text-decoration:none;

  }
  .articles .item .time {
    font-size:14px;
  }

  .articles .title {
    font-size:18px;

  }
  .articles .title i {
    margin-right:10px;
  }
  .articles .footer {
    padding-top:10px;
    font-size:14px;
  }
  .followers li {
    width:60px;
    float:left;
    margin:10px;
  }
  .counts {
    padding: 0px;
    margin: 0px;
  }
  .counts .item {
    float:left;
    width:50%;
    height:100%;
    padding:12px 10px;
    text-align: center;
    border-right: 1px solid #eee;

  }
  .counts .item .num {
    padding-bottom:10px;
    color:#000;
    font-weight:blod;
    font-size:14px;
  }
  .counts .item .title {
    color:#c1c1c1;
    font-size:14px;
  }
  .attentions {
    margin-top:10px;

  }
  .attentions .title {
    color:#666;
    font-size: 14px;
  }
  .follows {
    padding:10px 0px;
  }

  .follows .username {
    font-size: 14px;
  }
  .follows .labels {
    color:#888;
  }
  .follows .follow {

  }
   p.github{
  overflow: hidden;
  text-overflow:ellipsis;
  white-space: nowrap;
  color:#00B1ED;
  }
  .github i {
    font-size:20px;
    margin-right:10px;
  }


</style>
