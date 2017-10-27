<template>
    <div class='hello'>
        <img class='biaoti'
             src='../assets/logo2.png' />
        <h1>{{ msg }}</h1>
        <h2>网站正在建设中，请大家耐心等待！</h2>
        <a class='beian'
           href='http://www.miitbeian.gov.cn/'>
            蜀ICP备17036196号
        </a>
    </div>
</template>

<script>
  export default {
    name: 'HelloWorld',
    data () {
      return {
        msg: '欢迎各位同袍来到汉服心！'
      }
    }
  }
</script>

<!-- Add 'scoped' attribute to limit CSS to this component only -->
<style scoped>
    .hello {
        font-family: sans-serif;
        text-align: center;
        color: #2c3e50;
        width: 100%;
        height: 100%;
        background-size: cover;
        background: url(../../static/beijing.jpg) no-repeat;
        background-color: antiquewhite;
    }

    .beian {
        position: fixed;
        bottom: 0;
        text-align: center;
    }
</style>
