<template>
  <section>
  <div class="wrapper">
  <div class="container">
    <div class="row">
      <div class="col-md-10 col-md-offset-1 caption">
        <div class="col-md-6">
          <h1 class="title">Make a living off your passion</h1>
          <p class="content">Sell online courses and digital products from your own beautiful storefront in seconds.</p>

          <a class="btn btn-blue btn-lg" href="#/master">开始学习</a>
        </div>
        <div class="col-md-6">
          <img class="img-responsive" src="http://ouecw69lw.bkt.clouddn.com/profit-fbdeecd9d18c4e8a01edc3018d229e7df504e77153b64679d66cb03973eda702.svg" alt="One on One Check-ins to know the pulse and engagement of your team">
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
  margin-top: 0px;
  color: #5e6971;
}
.caption .title {
font-size: 36px;
margin-top:80px;
margin-bottom:20px;

color: #36434d;
}
.caption .content {
margin:30px 0px;
font-size: 20px;
line-height:1.4;

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
