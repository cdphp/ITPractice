<template>
<section>
<div class="registerContainer" v-if="loading">
  <div class="wait">
    <i class="fa fa-spinner fa-spin" aria-hidden="true"></i>
  </div>
  <div class="content">
  <p>正在验证...</p>
  </div>
</div>
<div class="registerContainer" v-else>

    <div class="wait">
      <i class="fa fa-check" aria-hidden="true"></i>
    </div>
    <div class="content">
    <p>验证成功</p>
    </div>

</div>
</section>
</template>
<script>
import {oauthGithub,getGithubUser} from '../../api/api'

export default {
  data() {

    return {
      loading: true,

      code: '',

    }
  },
  methods: {

    //验证
    oauth(code) {
      let para = {
        code : code,

      };

      oauthGithub(para).then(res => {
        if (res.errorNo==0) {

          if (res.need_login == 1) {
            this.$message({message: "用户关联成功,可以直接登录,默认密码是github的用户名",type: 'success'});
            this.$router.push({ path: '/login' });
          }else {
            //直接记录
            localStorage.setItem('user', JSON.stringify(res.data));
            this.$router.push({ path: '/' });
          }
        }else {
          this.$message({message: res.message,type: 'warning'});
        }
      });
    },
    getGithubUserinfo(access_token) {
      let para = {access_token: access_token};
      getGithubUser(para).then(res => {
        this.loading = false;
        console.log(res);
      });
    },

  },
  mounted() {
    this.code = this.$route.query.code;

    if(this.code) {
      this.oauth(this.code);
    }else {
      this.$message({ message: "请求参数错误", type: 'error'});
      this.$router.push({path: '/404'});
    }

  }
}
</script>
<style lang="scss" scoped>
.registerContainer {
    text-align: center;
    padding: 60px;
    height: 350px;
    .wait {

      margin: 30px 0px;
      font-size:40px;
      color:#20a0ff;

    }
    .content {
    width: 350px;
    padding: 10px;
    margin: 0px auto;
    font-size:20px;
    color:#666;
    }
}
</style>
