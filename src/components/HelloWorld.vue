<template>
    <div class='hello'>
        <img class='biaoti'
             src='../assets/logo2.png' />
        <h1>{{ msg }}</h1>
        <h2>网站正在建设中，请大家耐心等待！</h2>
        <div class='beian'>
            <a target='_blank'
               href='http://www.beian.gov.cn/portal/registerSystemInfo?recordcode=51010702000866'
               style='display:inline-block;text-decoration:none;height:20px;line-height:20px;'>
                <img src='../../static/beian.png'
                     style='float:left;' />
                <p style='float:left;height:20px;line-height:20px;margin: 0px 0px 0px 5px; color:#939393;'>
                    川公网安备 51010702000866号
                </p>
            </a>
        </div>
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
        background: url(../../static/beijing.jpg) no-repeat;
        background-size: 100% 100%;
        /*background-color: antiquewhite;*/
    }

    .beian {
        position: fixed;
        bottom: 0;
        text-align: center;
        width: 300px;
        margin: 0 auto;
        padding: 20px 0;
    }
</style>
