<template>
  <section>
  <div class="wrapper">
  <div class="container">
    <div class="row">
      <div class="col-md-10 col-md-offset-1 caption">
        <div class="col-md-6">
          <h1>学以致用</h1>
          <p>在知识的海洋里遨游，在兴趣的天空中飞翔.</p>

          <a class="btn btn-blue" href="#/master">开始学习</a>
        </div>
        <div class="col-md-6">
          <img class="img-responsive" src="../assets/imgs/herobanner-new.png" alt="One on One Check-ins to know the pulse and engagement of your team">
        </div>
      </div>
    </div>




  </div>
  <div class="container services">

  <ul class="clearfix">
    <li class="animated wow fadeInDown animated" style="visibility: visible; animation-name: fadeInDown;">
      <img class="icon" src="../assets/imgs/icon1.png" alt="">
      <span class="separator"></span>
      <h2>Officia Deserunt Mollit</h2>
      <p>sed do eiusmod tempor incididunt ut labore et dolore magna aliqua voluptate velit esse
      cillum dolore.</p>
    </li>
    <li class="animated wow fadeInDown animated" data-wow-delay=".2s" style="visibility: visible; animation-delay: 0.2s; animation-name: fadeInDown;">
      <img class="icon" src="../assets/imgs/icon2.png" alt="">
      <span class="separator"></span>
      <h2>Culpa Killum Dolore</h2>
      <p>aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.</p>
    </li>
    <li class="animated wow fadeInDown animated" data-wow-delay=".4s" style="visibility: visible; animation-delay: 0.4s; animation-name: fadeInDown;">
      <img class="icon" src="../assets/imgs/icon3.png" alt="">
      <span class="separator"></span>
      <h2>Elit Tempor Incididunt</h2>
      <p>nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat enim ad minim veniam.</p>
    </li>
  </ul>

  </div>
  </div>
  </section>
</template>
<script>
export default {
  data() {
    return {

    }
  },

}

</script>
<style scoped>

.caption {
  margin-top: 30px;

  color: #5d5d5d;
}

.graybg {
  background:#e7e7e7;
  padding-top:15px;
  padding-bottom:15px;
}
.services ul{
    margin-top: 100px;
}

.services ul li{
    list-style: none;
    display: block;
    width: 313px;
    float: left;
    margin-left: 80px;
    text-align: center;
}

.services ul li:first-child{
    margin-left: 0;
}

.services ul li .separator{
    display: block;
    width: 1px;
    height: 20px;
    background: #f6f6f6;
    margin: 20px auto;
}

.services ul li h2{
    color: #616161;
    font-family: "raleway-bold";
    font-size: 18px;
    font-weight: bold;
    margin: 0;
}

.services ul li p{
   color: #838181;
    font-family: "raleway-regular";
    font-size: 14px;
    line-height: 28px;
    margin-top: 40px;
    letter-spacing: .5px;
}

</style>
