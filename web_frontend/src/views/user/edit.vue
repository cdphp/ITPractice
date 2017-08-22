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
                <label for="about" class="col-sm-2 control-label">Github</label>
                <div class="col-sm-10">
                  <textarea class="form-control" id="github" placeholder="你的github.." v-model="user.github"></textarea>
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
import {getArticleListPage,getUser,editUser,upload} from '../../api/api'

export default {
  data() {

    return {
      id: 0,
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
      this.id = id;
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
        github: this.user.github,
      }
      editUser(para).then(res => {

        if(res.errorNo == 0 ) {

          this.$router.push({ path: '/user?id='+ this.user.id});
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


    },
    saveAvatar() {
      var base64Data = $("#image").cropper('getCroppedCanvas').toDataURL('image/jpeg');
      let para = {type:"image", content: base64Data}
      upload(para).then(res => {
        if(res.errorNo!=0) {
          this.$message({
            type: 'error',
            message: '已取消上传'
          });
          return
        }

        this.$confirm('确实将此图片上传为头像么?', '提示', {
            confirmButtonText: '确定',
            cancelButtonText: '取消',
            type: 'warning'
          }).then(() => {
            let para = {id: this.id, avatar: res.url};
            editUser(para).then(res => {

              if(res.errorNo == 0 ) {

                this.$router.push({ path: '/user?id='+ this.user.id});
              }else {
                this.$message({
                message: res.message,
                type: 'error'
              });
              }
            });

          }).catch(() => {
            this.$message({
              type: 'info',
              message: '已取消上传'
            });
          });
      });




    }

  },
  updated() {
  $("#image").cropper(this.option);
  },
  mounted() {

    var user = localStorage.getItem('user');

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
            cursor:pointer;
        }
        .avatar {
          width:400px;

          padding-left: 50px;
        }
</style>
