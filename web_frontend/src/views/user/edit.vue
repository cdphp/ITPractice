<template>
  <section>
    <div class="wrapper gray-bg">
      <div class="container">

      <div class="row">
        <div class="col-sm-3">
          <div class="user-sidebar box">
            <div class="box-header">
              <h5>基本信息</h5>

            </div>
            <div class="box-content no-padding">
            <div class="headpic">
              <img :src="user.avatar" />
            </div>
            </div>
            <div class="box-content">
            <h4>{{user.username}}</h4>
            </div>

          </div>
        </div>

        <div class="col-sm-9">
          <div class="box">
          <div class="box-header">
            <h5>个人信息</h5>
          </div>
            <div class="box-content">
            <div class="form-horizontal">
              <div class="form-group">
                <label for="email" class="col-sm-2 control-label">邮箱</label>
                <div class="col-sm-10">
                  <p class="form-control-static">{{user.email}}</p>
                </div>
              </div>
              <div class="form-group">
                <label for="labels" class="col-sm-2 control-label">标签</label>
                <div class="col-sm-6">
                <input type="text" class="form-control" id="labels" placeholder="..." v-model="user.labels">
                </div>
              </div>
              <div class="form-group">
                <label for="about" class="col-sm-2 control-label">个性签名</label>
                <div class="col-sm-10">
                  <textarea class="form-control" id="about" placeholder="自由发挥.." v-model="user.about"></textarea>
                </div>
              </div>

              <div class="form-group">
                <div class="col-sm-offset-2 col-sm-10">
                  <button class="btn btn-blue" v-on:click="updateInfo">保存</button>
                </div>
              </div>
              </div>
            </div>


          </div>

          <div class="box">
          <div class="box-header">
            <h5>上传头像</h5>

          </div>
          <div class="box-content">

          <div class="row">
            <div class="col-sm-8">
              <div class="avatar">
                <img id="image" :src="user.avatar" />
              </div>

            </div>
            <div class="col-sm-4">
            <button class="btn btn-blue fileinput-button">
              <span>选择图片</span>
              <input type="file" @change="onFileChange">
          </button>
          <button class="btn btn-blue" v-on:click="saveAvatar">
            保存头像
          </button>
            </div>
          </div>
          </div>
          </div>

        </div>
      </div>

      </div>
    </div>
  </section>
</template>
<script>
import {getArticleListPage,getUser,editUser} from '../../api/api'

export default {
  data() {

    return {
      user:{},
      option: {
        aspectRatio: 1 / 1,
        crop: function(e) {
          console.log(e);
        }
      },

    }
  },
  methods: {
    getUserInfo(id) {
      let para = {id :id};
      getUser(para).then(res => {

        if(res.errorNo == 0 ) {
          this.user = res.data;

        }else {
          this.$router.push({ path: '/404' });
        }
      });
    },
    updateInfo() {
      var that = this;
      let para = {
        id: this.user.id,
        labels: this.user.labels,
        about: this.user.about,
      }
      editUser(para).then(res => {

        if(res.errorNo == 0 ) {

          this.$router.push({ path: '/user?id='+ that.user.id});
        }else {
          this.$message({
          message: res.message,
          type: 'error'
        });
        }
      });

    },
    onFileChange(e) {


        var files = e.target.files || e.dataTransfer.files;
        if (!files.length)
          return;
        var file = files[0];
        var URL = window.URL || window.webkitURL;
        if (/^image\/\w+$/.test(file.type)) {
          var uploadedImageType = file.type;

          if (uploadedImageURL) {
            URL.revokeObjectURL(uploadedImageURL);
          }

          var uploadedImageURL = URL.createObjectURL(file);
          $("#image").cropper('destroy').attr('src', uploadedImageURL).cropper(this.option);

        } else {
          window.alert('Please choose an image file.');
        }
        /*
        var vm = this;

        var image = new Image();
      var reader = new FileReader();
      reader.onload = (e) => {
        console.log(e.target.result);
        vm.info.avatar = e.target.result;
      };
      reader.readAsDataURL(files[0]);
      */

    },
    saveAvatar() {
      var that = this;
      this.$message({
        message: "上传功能还未实现，先等等",
        type: 'warning'
      });
      $("#image").cropper('getCroppedCanvas').toBlob(function (blob) {

        var fd = new FormData();
        fd.append('file', blob);
        console.log(blob);

      });

    }

  },
  updated() {
  $("#image").cropper(this.option);
  },
  mounted() {

    var user = sessionStorage.getItem('user');

    if (user) {
      user = JSON.parse(user);
      this.getUserInfo(user.user_id);

    }else {
      this.$router.push({ path: '/login' });
    }


  }

}
</script>
<style scoped>
.fileinput-button {
            position: relative;
            display: inline-block;
            overflow: hidden;
        }

        .fileinput-button input{
            position:absolute;
            right: 0px;
            top: 0px;
            opacity: 0;
            -ms-filter: 'alpha(opacity=0)';
            font-size: 200px;
        }
        .avatar {
          width:400px;

          padding-left: 50px;
        }
</style>
