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
import {validateEmail} from '../../api/api'

export default {
  data() {
    return {
      loading: true,
    }
  },
  methods: {

    //验证
    validate: function (key) {
      let para = {
        key : key,
      };


      validateEmail(para).then(res => {
      if (res.errorNo != 0) {
        this.$message({
          message: res.message,
          type: 'error'
        });
      } else {
        this.$message({
          message: "邮箱验证成功,可以直接登录",
          type: 'success'
        });
        this.$router.push({ path: '/login' });

      }

      });
    },

  },
  mounted() {
    var key = this.$route.query.key;
    if(key) {
      this.validate(key);
    }else {
      this.$message({
        message: "请求参数错误",
        type: 'error'
      });
      this.$router.push({path: '/404'});
    }

  }
}
</script>
<style lang="scss" scoped>
.registerContainer {
    text-align: center;
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
