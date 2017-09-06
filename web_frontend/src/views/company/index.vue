<template>
<section>
<div class="wrapper gray-bg">
  <div class="container">
    <div class="search">
      <div class="box">
        <input type="text" class="form-control" v-model="name" @keyup.enter="handleSearch" placeholder="搜索">
      </div>

    </div>
    <div class="row">
      <div class="col-sm-9">
        <div class="box companys">
          <div class="box-content" v-for="(item,index) in companys">
            <div class="title">{{item.name}}<span class="right muted">{{formatTime(item.created_at)}}</span></div>
            <p class="description muted">{{item.description}}</p>
          </div>


        </div>
        <div class="more" v-if="nomore">没有啦</div>
        <div class="more" v-on:click="loadMore" v-else>查看更多</div>
      </div>
      <div class="col-sm-3">
        <div class="box">
          <div class="box-header">
            声明
          </div>
          <div class="box-content">
            该内容均为网友添加，与本网站无关
          </div>
        </div>
        <div class="box">
        <a href="#/company/add" class="btn btn-blue  btn-block">新增记录</a>
        </div>
      </div>

    </div>
  </div>
</div>
</section>
</template>
<script>
import {getCompanyListPage} from '../../api/api'
import util from '../../common/js/util'

export default {
  data() {
    return {
      companys:[],
      page: 1,
      loading: false,
      nomore: false,
      name:'',
    }
  },
  methods: {
    loadMore() {
      this.page++;
      this.getCompanys();
    },
    formatTime(unixTime) {
      return util.formatDate.format(new Date(unixTime*1000),'yy-MM-dd hh:mm');
    },

    //获取文章列表
    getCompanys: function () {
      let para = {
        page : this.page,
        name: this.name,

      };
      this.loading = true;

      getCompanyListPage(para).then(res => {
        if(res.errorNo == 0) {
          if(res.data) {
            this.companys.push.apply(this.companys, res.data);
          }else {
            this.nomore = true;
          }

          console.log("result:",this.companys);
          this.loading = false;
        }


        //NProgress.done();
      });
    },
    //获取文章列表
    handleSearch: function () {
      let para = {
        page : 1,
        name: this.name,

      };
      this.loading = true;

      getCompanyListPage(para).then(res => {
        if(res.errorNo == 0) {

          this.companys = res.data;

          this.loading = false;
        }

      });
    }
  },
  mounted() {
    this.getCompanys();
  }
}
</script>
<style scoped>
.companys .title {
  margin:15px 0px;
}
.companys .description {
  padding: 10px 0px;
  font-size: 16px;
    line-height: 1.8;

}
.companys .content p {
  padding: 10px 0px;
}
.companys .footer {
  margin:10px 0px;
  font-size:14px;
}
.companys .footer span{
  margin-right:15px;
}
.search {

}
.search .box {
  margin:20px auto;
  position: relative;
  width: 400px;
}
.search .box input {
  border-radius: 0px;
  padding:20px;
  border: none;
   text-align:center;
}
.search .box input:focus {
  box-shadow: none;
}


</style>
