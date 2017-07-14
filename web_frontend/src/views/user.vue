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
                <img src="../assets/imgs/profile_big.jpg" />
              </div>
              </div>
              <div class="box-content">
              <h4>破晓</h4>
              <p class="muted text-indent">正义的伙伴,低级全栈</p>
              <h6>个人介绍</h6>
              <p class="muted text-indent">这个家伙很懒，什么都没有留下...</p>

              <button class="btn btn-red btn-block follow" v-if="canAttention">关注</button>
              <a href="#/article/add" class="btn btn-blue btn-block follow" v-if="isSelf">发表新文章</a>
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
                    <a href="#"><img class="img-circle" src="../assets/imgs/user.png" :title="item.username"/></a>
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
                    <a href="#"><img class="img-circle" src="../assets/imgs/user.png" :title="item.username"/></a>
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
export default {
  data() {
    return {
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

  },
  mounted() {
    var id = this.$route.query.id;
    console.log(id);
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
