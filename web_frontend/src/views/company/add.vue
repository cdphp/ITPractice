<template>
  <section>
    <div class="wrapper gray-bg">
      <div class="container">
        <div class="row">
          <div class="col-sm-8 col-sm-offset-2">
            <div class="box">
              <div class="box-header">
                添加公司信息
              </div>
              <div class="box-content">
              <div class="form-horizontal">
                <div class="form-group">
                  <label for="title" class="col-sm-2 control-label">名称</label>
                  <div class="col-sm-6">
                    <input type="text" class="form-control" id="name" placeholder="填写名称" v-model="name">
                  </div>
                </div>
                <div class="form-group">
                  <label for="digest" class="col-sm-2 control-label">说明</label>
                  <div class="col-sm-10">
                    <textarea type="text" rows="4" class="form-control" id="description" placeholder="填写相关说明" v-model="description"></textarea>
                  </div>
                </div>


                <div class="form-group">
                  <div class="col-sm-offset-2 col-sm-10">
                    <button  class="btn btn-blue" :disable="loading" v-on:click="add">提交</button>
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
import {addCompany} from '../../api/api'

export default {
  data(){
    return{
      name: '',
      description: '',
      loading: false,

    }
  },


  methods: {


    add() {

      if(this.name=='') {
        this.$message({
          message: '名称不能为空',
          type: 'warning'
        });
        return
      }

      if(this.description=='') {
        this.$message({
          message: '说明不能为空',
          type: 'warning'
        });
        return
      }


      this.loading = true;
      var para = {name: this.name,description: this.description};
      addCompany(para).then(res => {


        if (res.errorNo != 0) {
          this.$message({
            message: res.message,
            type: 'error'
          });
          this.loading = false;
        } else {
        this.$message({
          message: '添加成功',
          type: 'success'
        });
          this.$router.push({ path: '/company' });
        }
      });
    }
  },
  mounted() {

  }
}
</script>
<style scoped>
.container {
  min-height:430px;
}
</style>
