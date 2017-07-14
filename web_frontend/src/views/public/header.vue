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
      <li><a href="#/invest">调查</a></li>

    </ul>

    <ul class="nav navbar-nav navbar-right">
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

      };
    },
    methods: {

    }
  }
</script>
<style scoped>
.navbar {
  margin-bottom:0px;
}
.navbar-default {
    padding:15px;
    background-color: transparent;

    border-style:none none solid none;

}
.navbar-default .navbar-nav>li a {
    color: #545e6b;
    font-weight: 600;
    font-size: 14px;
    -webkit-transition:all .2s linear;
    -moz-transition:all .2s linear;
    -o-transition:all .2s linear;
    color:#5d5d5d;
}
.navbar-default .navbar-nav>li a:hover {
  color:#F8B08D;
}
.navbar-brand {
  padding:7px 15px;
}
.navbar-brand img {
    width: 97px;
    height: 30px;
}
.navbar-default .navbar-nav>.active>a, .navbar-default .navbar-nav>.active>a:focus, .navbar-default .navbar-nav>.active>a:hover {
  background:#fff;
}
</style>
