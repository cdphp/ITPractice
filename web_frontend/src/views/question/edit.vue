<template>
  <section>
    <div class="wrapper gray-bg">
      <div class="container">
        <div class="row">
          <div class="col-sm-12">
            <div class="box">
              <div class="box-header">
                编辑
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
                    <mavon-editor v-model="content"/>
                  </div>
                </div>

                <div class="form-group">
                  <div class="col-sm-offset-2 col-sm-10">
                    <button  class="btn btn-blue" :disable="loading" v-on:click="edit">确定</button>
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
import {editArticle,getArticle} from '../../api/api'
import mavonEditor from 'mavon-editor'
import 'mavon-editor/dist/css/index.css'

export default {
  data(){
    return{
      id: '',
      title: '',
      digest: '',
      content:'',
      loading: false,
    }
  },
  components: {
    'mavon-editor': mavonEditor.mavonEditor
  },

  methods: {
  getArticleInfo(id) {
    let para = {id :id};
    getArticle(para).then(res => {

      if(res.errorNo == 0 ) {
        this.title = res.data.title;
        this.digest = res.data.digest;
        this.content = res.data.content;
      }else {
        this.$router.push({ path: '/404' });
      }
    });
  },
    edit() {
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
      var para = {id: this.id, title: this.title,digest: this.digest, content: this.content};

      editArticle(para).then(res => {


        if (res.errorNo != 0) {
          this.$message({
            message: res.message,
            type: 'error'
          });
          this.loading = false;
        } else {
          this.$message({
            message: "修改成功",
            type: 'info'
          });
          this.$router.push({ path: '/article/info?id='+this.id });
        }
      });
    }
  },
  mounted() {
    var id = this.$route.query.id;
    this.id = id;

    this.getArticleInfo(id);
  }
}
</script>
