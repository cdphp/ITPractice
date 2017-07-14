<template>
  <section>
    <div class="wrapper gray-bg">
      <div class="container">
        <div class="row">
          <div class="col-sm-9">
            <div class="box">
              <div class="box-header">
                撰写
              </div>
              <div class="box-content">
              <div class="form-horizontal">
                <div class="form-group">
                  <label for="inputEmail3" class="col-sm-2 control-label">标题</label>
                  <div class="col-sm-6">
                    <input type="text" class="form-control" id="title" placeholder="填写标题">
                  </div>
                </div>
                <div class="form-group">
                  <label for="content" class="col-sm-2 control-label">内容</label>
                  <div class="col-sm-10">
                    <Ueditor @ready="editorReady"></Ueditor>
                  </div>
                </div>

                <div class="form-group">
                  <div class="col-sm-offset-2 col-sm-10">
                    <button  class="btn btn-blue">确定</button>
                  </div>
                </div>
                </div>
              </div>
            </div>
          </div>

          <div class="col-sm-3">
            <div class="box">
              <div class="box-header">
                其他
              </div>
            </div>
          </div>

        </div>

      </div>
    </div>
  </section>
</template>
<script>
import Ueditor from '../../components/Ueditor'

export default {
  data(){
    return{
      content:''
    }
  },
  name: 'app',
  components: {
    Ueditor
  },
  methods: {
    editorReady (instance) {
      instance.setContent('填写内容');

      instance.addListener('contentChange', () => {
        this.content = instance.getContent();
      });
    },
  },
}
</script>
