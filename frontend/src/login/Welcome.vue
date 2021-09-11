<template>
    <div id="login-wrapper">
        <div class="container container-mid">
            <div class="row px-0 justify-content-lg-center h-100">
                <div class="col-lg-2 container-shadow-red">
                </div>
                <div class="col-lg-2 container-shadow-green">
                </div>
                <div class="col-lg-4 content-layer">
                    <LoginForm />
                </div>
            </div>
        </div>
    </div>
</template>

<script>
// TODO: Check if keys set, if not, straight to charts.
    import { defineComponent } from "@vue/runtime-core"
    import LoginForm from "@/login/LoginForm.vue"

    export default defineComponent({
        components: {
            LoginForm
        },
        name: "Welcome",
    });
</script>

<style lang="scss" scoped>
    // TODO: Set some breakpoints.
    .login-wrapper {
        height: 100vh;
        background-color: rgb(2, 4, 15);
    }
    .container-mid {
        height: 580px;
        min-height: 500px;
        top: 50%;
        left: 50%;
        transform: translate(-50%, -50%);
        position: absolute;
    }
    // cannot overlap to create desired effect, due to limitation of webkit.
    // therefore create two seperate with an offset.
    .container-shadow-red {
        height: 90%;
        margin-top: auto;
        margin-bottom: auto;
        box-shadow: -15px 0px 65px 20px rgba(194, 33, 33, 0.7);
        background-color: transparent;
    }
    .container-shadow-green {
        height: 90%;
        margin-top: auto;
        margin-bottom: auto;
        box-shadow: 15px 0px 65px 20px rgba(38, 194, 33, 0.7);
        background-color: transparent;
    }
    .content-layer {
        height: 100%;
        border-radius: 25px;
        position: fixed;
        z-index: 9999;
        background-color: rgb(5, 12, 43);
    }
    </style>

