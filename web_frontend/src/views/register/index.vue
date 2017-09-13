<template>
  <section>
    <div class="registerContainer">
      <div class="signup-step">
        <div class="step-wrapper" v-bind:class="{'active': step==1 }">
          <i class="fa fa-user-o" aria-hidden="true"></i>
          注册账号
        </div>
        <div class="step-wrapper" v-bind:class="{'active': step==2 }">
          <i class="fa fa-envelope-o" aria-hidden="true"></i>
          激活邮箱
        </div>
        <div class="step-wrapper" v-bind:class="{'active': step==3 }">
          <i class="fa fa-check-square-o" aria-hidden="true"></i>
          注册成功
        </div>
      </div>
      <div class="content" v-if="step==1">
      <el-form :model="ruleForm2" :rules="rules2" ref="ruleForm2" label-position="left" label-width="0px" class="demo-ruleForm login-container">

        <el-form-item prop="account">
          <el-input type="text" v-model="ruleForm2.account" auto-complete="off" placeholder="账号"></el-input>
        </el-form-item>
        <el-form-item prop="email">
          <el-input type="text" v-model="ruleForm2.email" auto-complete="off" placeholder="邮箱"></el-input>
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

      </el-form>
      </div>

      <div class="mail-content" v-if="step==2">

        <div class="msg-box" v-if="sendSuccess">
          <div class="msg-header">Hi,{{ruleForm2.email}}</div>
          <div class="msg-content">
            <p>感谢您注册修仙门派</p>

            <p>我们已经向您的邮箱发送一封注册邮箱，请登录您的邮箱完成验证。</p>
          </div>

          <a href="javascript:void(0)" class="btn btn-blue" v-on:click="goValidate">立刻登录邮箱完成验证</a>

        </div>
        <div class="tip" v-if="sendSuccess">
        若您没有收到邮件您可以：检查您的垃圾邮件中，是否包含验证邮件，或者<a href="javascript:void(0)" v-on:click="resend" class="text-blue">重发邮件</a>
        </div>
        <div class="tip" v-if="sendSuccess==false">
        邮件发送失败
        </div>
      </div>

      <div class="content" v-if="step==3">
        如果您已经完成验证，请直接：<a href="#/login" class="text-blue">登录</a>

      </div>
    </div>
  </section>

</template>

<script>
  import { register,sendMail } from '../../api/api';
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
        step : 1,
        loading: false,
        sendSuccess: true,
        ruleForm2: {
          account: '',
          email:'',
          pass:'',
          checkPass: ''
        },

        rules2: {
          account: [
            { required: true, message: '请输入账号', trigger: 'blur' },
          ],
          email: [
            { required: true, message: '请输入邮箱', trigger: 'blur' },
            { type: 'email', message: '请输入正确的邮箱地址', trigger: 'blur,change' }
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
      sendValidateMail() {
      var params = {email: this.ruleForm2.email,type:"register"};

      sendMail(params).then(res => {
        this.loading = false;
        //NProgress.done();
        if (res.errorNo !== 0) {
          this.$message({
            message: res.message,
            type: 'error'
          });
          this.sendSuccess = false;
        }

      });
      },
      goValidate() {
        var arr = this.ruleForm2.email.split("@");
        var server = arr[1];

        window.open("http://www."+server);
        this.step = 3;
      },
      resend() {
      this.$message({
        message: "该功能过几天加,2333",
        type: 'warning'
      });
      },
      handleSubmit2(ev) {
        var _this = this;
        this.$refs.ruleForm2.validate((valid) => {
          if (valid) {
            this.loading = true;
            //NProgress.start();
            var params = { username: this.ruleForm2.account,email:this.ruleForm2.email, password: this.ruleForm2.pass };
            register(params).then(data => {
              this.loading = false;
              //NProgress.done();
              if (data.errorNo !== 0) {
                this.$message({
                  message: data.message,
                  type: 'error'
                });
              } else {
                this.step = 2;
              }
            });
          } else {
            console.log('error submit!!');
            return false;
          }
        });
      }
    },
    mounted() {
      var ref = this.$route.query.ref;

      if(ref=='login') {
        this.step = 2;
        this.ruleForm2.email = this.$route.query.email;
        this.sendValidateMail();
      }
    }
  }

</script>

<style lang="scss" scoped>
  .signup-step {
    text-align: center;
    padding:40px 0px;
    color: #8f9bb3;
  }
  .signup-step .active {
    color: #1989fa;
  }
  .signup-step .step-wrapper {
    display: inline-block;
    padding: 0px 40px;
  }
  .signup-step .step-wrapper i {
  font-size: 24px;
  display: block;
  margin-bottom: 10px;
  }
.signup-step:before {
  content: "";
  position: absolute;
  height: 1px;
  width: 16%;
  background-color: #e1e6f0;
  top: 11%;
  left: 27%;
  }
  .signup-step:after {
    content: "";
    position: absolute;
    height: 1px;
    width: 16%;
    background-color: #e1e6f0;
    top: 11%;
    right:27%;
    }
    .registerContainer {
        .content {
        width: 350px;
        padding: 10px;
        margin: 0px auto;
        }
        .mail-content {
        width: 600px;
        padding: 10px;
        margin: 0px auto;
        }
    }

    .mail-content {
      padding: 10px;
      text-align:center;

      .msg-box {
        color:#197919;
        background:#eee;
        border:1px solod #ddd;
        border-radius: 5px;
        margin:20px 10px;
        padding:10px;
        .msg-header {
          margin:10px 0px;
        }
        .msg-content{
          margin:20px 0px;
          p {margin:10px 0px;}
        }
        .btn {
          margin:10px 0px;
        }
      }
    }
</style>
