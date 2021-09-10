<template>
  <div id=app>
    <router-view/>
  </div>
</template>

<style lang="scss">
/* check these */
    @font-face {
        font-family: "Poppins";
        font-style: normal;
        font-weight: normal;
        src: local("Poppins"),
        url(./fonts/Poppins/Poppins-Regular.ttf) format("truetype");
    }
    @font-face {
        font-family: "Poppins";
        font-style: normal;
        font-weight: 200;
        src: local("Poppins"),
        url(./fonts/Poppins/Poppins-Light.ttf) format("truetype");
    }
    @font-face {
        font-family: "Poppins";
        font-style: normal;
        font-weight: 100;
        src: local("Poppins"),
        url(./fonts/Poppins/Poppins-Thin.ttf) format("truetype");
    }
    @font-face {
      font-family: "Poppins";
      font-style: normal;
      font-weight: 300;
      src: local("Poppins"),
      url(./fonts/Poppins/Poppins-SemiBold.ttf) format("truetype");
    }
    #app {
      width: 100%;
      height: 100vh;
      margin: 0;
      padding: 0;
      background-color: rgb(11, 10, 13);
      font: normal normal 14px Poppins;
      color: rgb(255, 255, 255);
    }

</style>
