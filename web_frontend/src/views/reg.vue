<template>
  <section class="login-bg">
  <el-form :model="ruleForm2" :rules="rules2" ref="ruleForm2" label-position="left" label-width="0px" class="demo-ruleForm login-container">
    <h4 class="title">用户注册</h4>
    <el-form-item prop="account">
      <el-input type="text" v-model="ruleForm2.account" auto-complete="off" placeholder="账号"></el-input>
    </el-form-item>
    <el-form-item prop="pass">
      <el-input type="password" v-model="ruleForm2.pass" auto-complete="off" placeholder="密码"></el-input>
    </el-form-item>
    <el-form-item prop="checkPass">
      <el-input type="password" v-model="ruleForm2.checkPass" auto-complete="off" placeholder="确认密码"></el-input>
    </el-form-item>

    <el-form-item>
      <el-button type="primary" style="width:100%;" @click.native.prevent="handleSubmit2" :loading="loading">注册</el-button>
    </el-form-item>
    <div class="login-footer">
      已有用户?<a href="#/login" class="text-gray">直接登录</a>

    </div>
    <div class="login-footer">

    </div>
  </el-form>
  </section>
</template>

<script>
  import { register } from '../api/api';
  //import NProgress from 'nprogress'
  export default {
    data() {
    var validatePass = (rule, value, callback) => {
        if (value === '') {
          callback(new Error('请输入密码'));
        } else {

          callback();
        }
      };
      var validatePass2 = (rule, value, callback) => {
        if (value === '') {
          callback(new Error('请再次输入密码'));
        } else if (value !== this.ruleForm2.pass) {
          callback(new Error('两次输入密码不一致!'));
        } else {
          callback();
        }
      };
      return {
        loading: false,
        ruleForm2: {
          account: '',
          pass:'',
          checkPass: ''
        },
        rules2: {
          account: [
            { required: true, message: '请输入账号', trigger: 'blur' },

          ],
          pass: [
           { validator: validatePass, trigger: 'blur' }
         ],
         checkPass: [
           { validator: validatePass2, trigger: 'blur' }
         ]
        },
        checked: true
      };
    },
    methods: {

      handleSubmit2(ev) {
        var _this = this;
        this.$refs.ruleForm2.validate((valid) => {
          if (valid) {

            this.loading = true;
            //NProgress.start();
            var params = { username: this.ruleForm2.account, password: this.ruleForm2.pass };
            register(params).then(data => {
              this.loading = false;
              //NProgress.done();

              if (data.errorNo !== 0) {
                this.$message({
                  message: data.errorMsg,
                  type: 'error'
                });
              } else {
                sessionStorage.setItem('user', JSON.stringify(data.data));
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
