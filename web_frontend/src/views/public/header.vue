<template>
<header>
<nav class="navbar navbar-default ">
<div class="container">
  <!-- Brand and toggle get grouped for better mobile display -->
  <div class="navbar-header">
    <button type="button" class="navbar-toggle collapsed" data-toggle="collapse" data-target="#bs-example-navbar-collapse-1" aria-expanded="false">
      <span class="sr-only">Toggle navigation</span>
      <span class="icon-bar"></span>
      <span class="icon-bar"></span>
      <span class="icon-bar"></span>
    </button>
    <a class="navbar-brand" href="#"><img class="logo" src="../../assets/imgs/logo1.png"></a>
  </div>

  <!-- Collect the nav links, forms, and other content for toggling -->
  <div class="collapse navbar-collapse" id="bs-example-navbar-collapse-1">
    <ul class="nav navbar-nav">
      <li class="active"><a href="#">门户 <span class="sr-only">(current)</span></a></li>
      <li><a href="#/master">传承</a></li>
      <li><a href="#/article">文章</a></li>
    </ul>


    <ul class="nav navbar-nav navbar-right" v-if="isLogin">
    <li class="dropdown">
        <a href="#" class="dropdown-toggle" data-toggle="dropdown" role="button" aria-haspopup="true" aria-expanded="false">{{username}} <span class="caret"></span></a>
        <ul class="dropdown-menu">

          <li><router-link :to="{ path: '/user', query: { id: user_id }}">个人中心</router-link></li>

          <li role="separator" class="divider"></li>
          <li><a href="javascript:void(0)" v-on:click="logout">退出登录</a></li>

        </ul>
      </li>
    </ul>
    <ul class="nav navbar-nav navbar-right" v-else>
      <li><a href="#/login">登录</a></li>
      <li><a href="#/reg">注册</a></li>

    </ul>
  </div><!-- /.navbar-collapse -->
</div><!-- /.container-fluid -->
</nav>
</header>
</template>
<script>
  export default {
    data() {
      return {
        username:'',
        user_id:'',
        isLogin:false,
      };
    },
    methods: {
      logout: function () {

        this.$confirm('确认退出吗?', '提示', {
          //type: 'warning'
        }).then(() => {
          this.isLogin = false
          sessionStorage.removeItem('user');
          this.$router.push('/login');
        }).catch(() => {

        });
        },
    },
    mounted() {
    var user = sessionStorage.getItem('user');

    if (user) {
      this.isLogin = true
      user = JSON.parse(user);

      this.username = user.username;
      this.user_id = user.user_id;
    }
    }
  }
</script>
<style scoped>

</style>
