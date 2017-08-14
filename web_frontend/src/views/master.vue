<template>
<section>
<div class="wrapper gray-bg">
  <div class="info-header">
    <h1>以技会友</h1>
    <p>东半球第二的技术社交网站，让你感受程序员的世界..</p>
  </div>
  <div class="container">
    <div class="row">

      <div class="col-sm-4" v-for="item in users">
        <div class="master">
          <div class="bg" :style="bgFormat(item.info.bg)">

          </div>
          <div class="headpic">
            <img :src="item.info.avatar" />
          </div>
          <div class="content">
          <div class="username">{{item.username}}</div>
          <p class="digest" v-if="item.info.labels">{{item.info.labels}}</p>
          <p class="digest" v-else>暂无内容</p>
          <button class="btn btn-blue btn-block" v-on:click="viewUser(item.id)">查看更多</button>
          </div>

        </div>
      </div>



    </div>
    <div class="more" v-if="nomore">没有啦</div>
    <div class="more" v-on:click="loadMore" v-else>查看更多</div>

  </div>
</div>
</section>
</template>
<script>
import {getUserListPage} from '../api/api'

export default {
  data() {
    return {
      users: [],
      page: 1,
      nomore: false,
    }
  },
  methods: {
    loadMore() {
      this.page++;
      this.getUsers();
    },
    bgFormat(bg) {
      return {backgroundImage:'url('+bg+')'}
    },
    //获取用户列表
    getUsers: function () {
      let para = {
        page : this.page,
        row : 12,
      };
      this.loading = true;
      //NProgress.start();
      getUserListPage(para).then(res => {
        if(res.errorNo == 0) {
          if(res.data) {
            this.users.push.apply(this.users, res.data);
          }else {
            this.nomore = true;
          }

          console.log("result:",this.users);
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
    this.getUsers();
  }
}
</script>
<style scoped>

.info-header {
  padding: 20px;
  text-align:center;

}
.info-header p {
  color:#707070;
}
.master {
  background:#fff;
  border-radius:5px;
  margin-bottom:30px;
}
.master .bg {height:150px;background-size:cover;border-top-left-radius:5px;
border-top-right-radius:5px;}
.master .content {padding:10px 30px;}
.master .headpic {
  width:100%;
  text-align:center;
}
.master .headpic img {
  width:80px;
  height:80px;
  border-radius:50%;
  border: 3px solid #fff;
  box-shadow: 0 1px 1px rgba(0,0,0,0.1);
  margin-top:-40px;
}
.master .username {
  margin:15px;
  text-align:center;
  font-size:20px;
}
.master .digest {
  color:#707070;
  text-align:center;
  height:40px;
}
.master a {
margin-top:20px;
}

</style>
