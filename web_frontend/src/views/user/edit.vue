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
              <img :src="info.avatar" />
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
                <input type="text" class="form-control" id="labels" placeholder="...">
                </div>
              </div>
              <div class="form-group">
                <label for="about" class="col-sm-2 control-label">个性签名</label>
                <div class="col-sm-10">
                  <textarea class="form-control" id="about" placeholder="自由发挥..">{{info.about}}</textarea>
                </div>
              </div>

              <div class="form-group">
                <div class="col-sm-offset-2 col-sm-10">
                  <button type="submit" class="btn btn-default">保存</button>
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
              <img id="image" :src="info.avatar" />
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
import {getArticleListPage,getUser} from '../../api/api'
export default {
  data() {
    return {
      user:{},
      info:{},

    }
  },
  methods: {
    getUserInfo(id) {
      let para = {id :id};
      getUser(para).then(res => {

        if(res.errorNo == 0 ) {
          this.user = res.data;
          this.info = res.data.info;
        }else {
          this.$router.push({ path: '/404' });
        }
      });
    },

  },
  updated() {
  $('#image').cropper({
aspectRatio: 16 / 9,
crop: function(e) {
  // Output the result data for cropping image.
  console.log(e.x);
  console.log(e.y);
  console.log(e.width);
  console.log(e.height);
  console.log(e.rotate);
  console.log(e.scaleX);
  console.log(e.scaleY);
}
});
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

</style>
