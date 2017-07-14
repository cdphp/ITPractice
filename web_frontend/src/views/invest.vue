<template>
  <div class="container">
    <div class="row">

      <div class="col-sm-6">
        <div class="page-header"><h3>统计数据</h3></div>

        <div id="chartColumn" style="width:100%; height:400px;"></div>

      </div>
      <div class="col-sm-6">
        <div class="page-header"><h3>选择项</h3></div>
        <p>你可以选择自己感兴趣的三项内容</p>
        <ul class="list-group">
          <button type="button" class="list-group-item">Database</button>
          <button type="button" class="list-group-item">Design Pattern</button>
          <button type="button" class="list-group-item">Docker</button>
          <button type="button" class="list-group-item">Tensorflow</button>

        </ul>
        <div class="center">
        <el-button type="primary center">投票</el-button>
        </div>

      </div>
    </div>
  </div>
</template>
<script>
import echarts from 'echarts'
export default {
  data() {
    return {

    }
  },
  methods: {
  drawColumnChart() {
      this.chartColumn = echarts.init(document.getElementById('chartColumn'));
      var option = {
    color: ['#3398DB'],
    tooltip : {
        trigger: 'axis',
        axisPointer : {            // 坐标轴指示器，坐标轴触发有效
            type : 'shadow'        // 默认为直线，可选为：'line' | 'shadow'
        }
    },
    grid: {
        left: '3%',
        right: '4%',
        bottom: '3%',
        containLabel: true
    },
    xAxis : [
        {
            type : 'category',
            data : ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun'],
            axisTick: {
                alignWithLabel: true
            }
        }
    ],
    yAxis : [
        {
            type : 'value'
        }
    ],
    series : [
        {
            name:'直接访问',
            type:'bar',
            barWidth: '60%',
            data:[10, 52, 200, 334, 390, 330, 220]
        }
    ]
};
      this.chartColumn.setOption(option);
  },
  },
  mounted: function () {
      this.drawColumnChart();

  },

}
</script>
<style>
.list-group-item {
  margin:10px 5px;
}
a.list-group-item:focus, a.list-group-item:hover, button.list-group-item:focus, button.list-group-item:hover {
  background: #3398DB;
  color:#fff;
}
</style>
