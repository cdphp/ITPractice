<template>
<section>
<div class="wrapper gray-bg">
  <div class="container">
    <div class="row">
      <div class="col-sm-9">
        <div class="box">

          <div class="box-content article">
            <div class="title">
              {{article.title}}
            </div>
            <div class="about muted text-center">
              <span class="item">天府盐煎肉</span>
              <span class="item"><i class="el-icon-time"></i> 2017-07-11</span>
            </div>
            <div class="about muted text-right">
              <span class="item"><i class="el-icon-star-off"></i> 20</span>
              <span class="item"><span class="glyphicon glyphicon-eye-open"></span> 20</span>
            </div>
            <div class="content">
              {{article.content}}
            </div>
          </div>

        </div>

        <div class="box">
          <div class="box-header">
            发表评论
          </div>
          <div class="box-content">
          <div class="form-horizontal">

            <div class="form-group">
              <label for="content" class="col-sm-1 control-label">内容</label>
              <div class="col-sm-8">
                <textarea rows="4" class="form-control" id="content" placeholder="填写评论"></textarea>
              </div>
            </div>

            <div class="form-group">
              <div class="col-sm-offset-1 col-sm-10">
                <button  class="btn btn-blue">提交</button>
              </div>
            </div>
            </div>
          </div>
        </div>

        <div class="box">
          <div class="box-header">
            相关评论({{comments.length}})
          </div>
          <div class="box-content">
            <div class="none" v-if="comments.length==0">暂无内容</div>
            <ul class="comments" v-if="comments.length">
              <li v-for="(item,index) in comments">
              <div class="media">
                <div class="media-left">
                  <a href="#">
                    <img class="media-object img-circle" src="../../assets/imgs/user.png">
                  </a>
                </div>
                <div class="media-body">
                  <div class="media-heading"><span class="floor muted">#{{index+1}}</span><span>{{item.username}}</span></div>
                  {{item.content}}
                </div>
              </div>
              </li>
            </ul>
            <div class="more">查看更多</div>


          </div>
        </div>
      </div>
      <div class="col-sm-3">
        <div class="box">
          <div class="box-header">他的其他文章</div>
        </div>
      </div>
    </div>

  </div>
</div>
</section>
</template>
<script>
import {getArticleListPage} from '../../api/api'
export default {
  data() {
    return {
      article: {
        title: 'OpenStack大规模部署详解',
        content: 'this is content',
      },
      comments:[
        {
          username:'test001',
          avatar:'',
          content:'灵车漂移',
          comments:[
            {
            username:'test001',
            avatar:'',
            content:'坟头蹦迪',
            }
          ],
        },
        {
          username:'test002',
          avatar:'',
          content:'骨灰拌饭'
        }
      ]
    }
  },
}
</script>
<style scoped>

.article {}
.article .title {
  font-size:28px;
  text-align:center;
  margin:20px;
}
.article .about{padding:10px;}
.article .about span.item{margin:15px;}
.article .content {margin:30px;}
.comments li {
  margin-top:10px;
  margin-bottom:10px;
  padding:10px 0px;
  border-bottom:1px solid #ddd;
}
.comments li  span.floor{margin-right:10px;}
</style>
