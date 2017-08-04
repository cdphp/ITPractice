<template>
  <section class="login-bg">
  <el-form :model="ruleForm2" :rules="rules2" ref="ruleForm2" label-position="left" label-width="0px" class="demo-ruleForm login-container">
    <h4 class="title">用户登录</h4>
    <el-form-item prop="account">
      <el-input type="text" v-model="ruleForm2.account" auto-complete="off" placeholder="账号"></el-input>
    </el-form-item>
    <el-form-item prop="checkPass">
      <el-input type="password" v-model="ruleForm2.checkPass" auto-complete="off" placeholder="密码"></el-input>
    </el-form-item>
    <el-checkbox v-model="checked" checked class="remember">记住密码</el-checkbox>
    <el-form-item>
      <el-button type="primary" style="width:100%;" @click.native.prevent="handleSubmit2" :loading="logining">登录</el-button>
    </el-form-item>
    <div class="login-footer">
      <a href="#" class="text-gray">忘记密码</a> | <a href="#/reg" class="text-gray">直接注册</a>

    </div>
    <div class="login-footer">

    </div>
  </el-form>
  </section>
</template>

<script>
  import { requestLogin } from '../api/api';
  //import NProgress from 'nprogress'
  export default {
    data() {
      return {
        logining: false,
        ruleForm2: {
          account: 'user',
          checkPass: '123456'
        },
        rules2: {
          account: [
            { required: true, message: '请输入账号', trigger: 'blur' },
            //{ validator: validaePass }
          ],
          checkPass: [
            { required: true, message: '请输入密码', trigger: 'blur' },
            //{ validator: validaePass2 }
          ]
        },
        checked: true
      };
    },
    methods: {
      handleReset2() {
        this.$refs.ruleForm2.resetFields();
      },
      handleSubmit2(ev) {
        var _this = this;
        this.$refs.ruleForm2.validate((valid) => {
          if (valid) {

            this.logining = true;
            //NProgress.start();
            var loginParams = { username: this.ruleForm2.account, password: this.ruleForm2.checkPass };
            requestLogin(loginParams).then(res => {
              this.logining = false;
              //NProgress.done();

              if (res.errorNo != 0) {
                this.$message({
                  message: res.errorMsg,
                  type: 'error'
                });
              } else {
                sessionStorage.setItem('user', JSON.stringify(res.data));
                this.$router.push({ path: '/' });
              }
            });
          } else {
            console.log('error submit!!');
            return false;
          }
        });
      }
    }
  }

</script>

<style lang="scss" scoped>
  .login-bg {
    width:100%;
    height:100%;
    background-image:url('../assets/imgs/login_bg.jpg');
    background-size:cover;

  }
  .login-container {

    margin: 0px auto;
    top:180px;
    position: relative;
    width: 350px;
    padding: 35px 35px 15px 35px;
    background: #fff;


  }
  .login-footer {text-align:center;color:#888;}
  .title {
    margin: 0px auto 40px auto;
    text-align: center;
    color: #505458;
  }
  .remember {
    margin: 0px 0px 35px 0px;
  }
</style>
