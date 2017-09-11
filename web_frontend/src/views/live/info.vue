<template>
<section>
<div class="wrapper gray-bg">
  <div class="container">

    <div class="row">
      <div class="col-sm-6">
        <div class="box ">
          <div class="box-header">
            直播室
          </div>
          <div class="box-content chat-container">
          <ul>
            <li v-for="(item,index) in chats">
            <div class="media">
              <div class="media-left">
                <a href="javascript:void(0)" >
                  <img class="media-object img-circle-head" :src="item.avatar">
                </a>
              </div>
              <div class="media-body chat-message">
                <div class="media-heading"><span class="username text-gray">{{item.username}}</span>
                </div>
                <div class="content">{{item.content}}</div>
              </div>
            </div>
            </li>

          </ul>

          </div>
          <div class="msg-input">
          <input type="text" class="form-control" v-model="content" @keyup.enter="sendMsg" placeholder="说点什么...">
          </div>


        </div>

      </div>
      <div class="col-sm-6">
        <div class="box">
          <div class="box-header">
            话题
          </div>
          <div class="box-content" v-for="(item,index) in topics">
            <div :class="index==0?'text-blue':'text-gray'">
              {{item.title}}
              <span class="right muted">{{formatTime(item.created_at)}}</span>
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
import {getCompanyListPage} from '../../api/api'
import util from '../../common/js/util'

export default {
  data() {
    return {
      conn:{},
      topics:[],
      page: 1,
      loading: false,
      nomore: false,
      content:'',
      chats:[
        {
          avatar:'http://ouecw69lw.bkt.clouddn.com/347453f5566d328c19c7a1abee86bffb',
          username:'装逼的小黄牛',
          content:'这个so easy',
        },
        {
          avatar:'http://ouecw69lw.bkt.clouddn.com/790fc7c8d57ba8d120d5c3e38b503355',
          username:'无敌',
          content:'心里有点逼数么',
        },
      ]
    }
  },
  methods: {
    loadMore() {
      this.page++;
      this.getTopics();
    },
    formatTime(unixTime) {
      return util.formatDate.format(new Date(unixTime*1000),'yy-MM-dd hh:mm');
    },

    //获取话题列表
    getTopics: function () {
      let para = {
        page : this.page,


      };
      this.loading = true;

      this.topics = [
        {title:"谈谈怎么处理高并发业务",created_at:1504687135},
        {title:"mysql的索引有多少种",created_at:1504687135}
      ];



    },
    sendMsg() {
      console.log(this.content);
    },
  },
  mounted() {
    this.getTopics();

    if (window["WebSocket"]) {
        this.conn = new WebSocket("ws://127.0.0.1:8011/ws");
        this.conn.onclose = function (evt) {
            console.log("closes");
        };
        this.conn.onmessage = function (evt) {
            var messages = evt.data.split('\n');
            console.log(messages);
        };
    } else {
    this.$message({
      message: "您的浏览器不支持websocket",
      type: 'error'
    });
    }
  }
}
</script>
<style scoped>
.chat-container {
  background: #eee;
    padding: 15px;
    height: 400px;
    overflow-y: auto;
}
.chat-container ul li {
  margin-bottom:10px;
}
.chat-message {



}
.chat-message .username {
  font-size:14px;
}
.chat-message .content {

  background:#fff;
  padding: 10px;
  font-size:14px;
}
.msg-input input{
  border-radius:0px;
  padding:20px;
  border:1px solid #eee;;
}
.msg-input input:focus {
  border: 1px solid #00B1ED;
}




</style>
