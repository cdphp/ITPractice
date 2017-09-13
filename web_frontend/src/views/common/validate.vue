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
  <div v-show="type=='register'">
    <div class="wait">
      <i class="fa fa-check" aria-hidden="true"></i>
    </div>
    <div class="content">
    <p>验证成功</p>
    </div>
  </div>
  <div v-show="type=='forget'">
    <div class="content">
    <h4 class="title">修改密码</h4>
    <el-form :model="ruleForm2" :rules="rules2" ref="ruleForm2" label-position="left" label-width="0px" class="demo-ruleForm login-container">

      <el-form-item prop="pass">
        <el-input type="password" v-model="ruleForm2.pass" auto-complete="off" placeholder="密码"></el-input>
      </el-form-item>
      <el-form-item prop="checkPass">
        <el-input type="password" v-model="ruleForm2.checkPass" auto-complete="off" placeholder="确认密码"></el-input>
      </el-form-item>

      <el-form-item>
        <el-button type="primary" style="width:100%;" @click.native.prevent="handleSubmit2" :loading="loading">提交</el-button>
      </el-form-item>


    </el-form>
  </div>

  </div>

</div>
</section>
</template>
<script>
import {validateEmail,resetPass} from '../../api/api'

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
      loading: true,
      type: '',
      key: '',
      ruleForm2: {
        pass:'',
        checkPass: ''
      },

      rules2: {

        pass: [
         { validator: validatePass, trigger: 'blur' }
       ],
       checkPass: [
         { validator: validatePass2, trigger: 'blur' }
       ]
      },
    }
  },
  methods: {

    //验证
    validate(type,key) {
      let para = { key : key};

      validateEmail(para).then(res => {
      if (res.errorNo == 0) {
        if (type == 'register') {
          this.$message({message: "邮箱验证成功,可以直接登录",type: 'success'});
          this.$router.push({ path: '/login' });
        }else if (type == 'forget') {
          this.loading = false;
        }
      } else {

        this.$message({ message: res.message, type: 'error'});

      }

      });
    },
    handleSubmit2(ev) {
      var _this = this;
      this.$refs.ruleForm2.validate((valid) => {
        if (valid) {
          this.loading = true;
          //NProgress.start();
          var params = { key: this.key, password: this.ruleForm2.pass };
          resetPass(params).then(data => {
            this.loading = false;
            //NProgress.done();
            if (data.errorNo !== 0) {
              this.$message({message: data.message,type: 'error'});
            } else {
              this.$message({message: '密码修改成功，请直接登录',type: 'success'});
              this.$router.push({ path: '/login' });
            }
          });
        } else {
          console.log('error submit!!');
          return false;
        }
      });
    },

  },
  mounted() {
    this.key = this.$route.query.key;
    this.type = this.$route.query.type;
    if(this.key && this.type) {
      this.validate(this.type, this.key);
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
