<template>
  <section>
    <div class="wrapper gray-bg">
      <div class="container">
        <div class="row">
          <div class="col-sm-12">
            <div class="box">
              <div class="box-header">
                撰写
              </div>
              <div class="box-content">
              <div class="form-horizontal">
                <div class="form-group">
                  <label for="title" class="col-sm-2 control-label">标题</label>
                  <div class="col-sm-6">
                    <input type="text" class="form-control" id="title" placeholder="填写标题" v-model="title">
                  </div>
                </div>
                <div class="form-group">
                  <label for="digest" class="col-sm-2 control-label">摘要</label>
                  <div class="col-sm-10">
                    <textarea type="text" rows="4" class="form-control" id="digest" placeholder="填写摘要" v-model="digest"></textarea>
                  </div>
                </div>
                <div class="form-group">
                  <label for="content" class="col-sm-2 control-label">内容</label>
                  <div class="col-sm-10">
                    <editor @imgAdd="imgAdd" @imgDel="imgDel" v-model="content"/>
                  </div>
                </div>

                <div class="form-group">
                  <div class="col-sm-offset-2 col-sm-10">
                    <button  class="btn btn-blue" :disable="loading" v-on:click="add">确定</button>
                  </div>
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
import {addArticle,upload} from '../../api/api'
import mavonEditor from 'mavon-editor'
import 'mavon-editor/dist/css/index.css'

export default {
  data(){
    return{
      title: '',
      digest: '',
      content:'',
      loading: false,
      img_file: {},
    }
  },
  components: {
    'editor': mavonEditor.mavonEditor
  },

  methods: {

    imgAdd(pos, $file){
          //this.img_file[pos] = $file;

          var reader = new FileReader();
        reader.readAsDataURL($file);
        let $vm = this.$children[0];
        reader.onload = function(e){
        let para = {type:"image", content: this.result};
        upload(para).then(res => {
          if(res.errorNo!=0) {
            this.$message({
              type: 'error',
              message: res.errorMsg,
            });
            return
          }


          $vm.$img2Url(pos,res.url);
        });
        }


      },
      imgDel(pos){
          delete this.img_file[pos];
      },
    add() {
      //console.log(this.$children[0]);return;
      if(this.title=='') {
        this.$message({
          message: '标题不能为空',
          type: 'warning'
        });
        return
      }

      if(this.digest=='') {
        this.$message({
          message: '摘要不能为空',
          type: 'warning'
        });
        return
      }

      if(this.content=='') {
        this.$message({
          message: '内容不能为空',
          type: 'warning'
        });
        return
      }
      this.loading = true;
      var para = {title: this.title,digest: this.digest, content: this.content};
      addArticle(para).then(res => {


        if (res.errorNo != 0) {
          this.$message({
            message: res.message,
            type: 'error'
          });
          this.loading = false;
        } else {

          this.$router.push({ path: '/article/info?id='+res.resourceId });
        }
      });
    }
  },
  mounted() {

  }
}
</script>
