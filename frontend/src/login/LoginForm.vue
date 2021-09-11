<template>
    <div class="container login-container">
        <div class="row">
            <div class="col-lg-10 offset-lg-1 login-header">
                Login
            </div>
        </div>
        <form id="login-form" @submit.prevent="submit" autocomplete="off">
            <div class="row mt-2">
                <div class="col-lg-10 offset-lg-1 mt-5 pword">
                    <input type="password" name="password" required />
                    <label for="password">password</label>
                </div>
            </div>
            <div class="row">
                <div class="col-lg-4 offset-lg-7 mt-2 forgot-password">
                    Forgot Password?
                </div>
            </div>
            <div class="row mt-3">
                <div class="mt-5">
                    <button class="log-in">Log In</button>
                </div>
            </div>
        </form>
    </div>
</template>

<script>
    import { defineComponent } from "@vue/runtime-core";

    export default defineComponent({
        name: "LoginForm"
    })
</script>

<style lang="scss" scoped>
    .login-container {
        height: 100%;
        border-radius: 15px;
    }
    .login-header {
        margin-top: 24%;
        margin-bottom: 3%;
        font: normal normal 200 34px Poppins;
        color: rgb(255, 255, 255);
    }
    .pword {
        display: flex;
        flex-flow: column-reverse;
        padding-top: 5px;     
    }
    input {
        width: 100%;
        padding-bottom: 0.3rem;
        border-width: 0px 0px 2px 0px;
        border-style: solid;
        border-color: rgb(255, 255, 255);
        font: normal normal 200 15px Poppins;
        color: rgb(255, 255, 255);
        background-color: transparent;
    }
    input:focus {
        outline: none;
    }
    input:focus ~ label,
    input:valid ~ label {
        font: normal normal 100 14px Poppins;
        color: rgba(255, 255, 255, 0.603);
        transform: translateY(-1.8rem);
    } 
    label {
        padding-bottom: 0.3rem;
        position: absolute;
        font: normal normal 100 17px Poppins;
        color: rgb(255, 255, 255);
        transition: 0.5s;
        pointer-events: none;
    }
    .forgot-password { 
        font: normal normal 200 12px Poppins;    
        text-align: right;
        color: rgb(255, 255, 255);
        cursor: pointer;
    }
    .log-in {
        height: 45px;
        width: 45%;
        min-width: 150px;
        padding: 3px;
        margin: 0 auto;
        line-height: 39px;    
        border-radius: 30px;
        border-width: 0px;
        display: block;
        font: normal normal normal 15px Poppins;
        text-decoration: none;
        text-align: center;
        color: rgb(2, 15, 9);
        background: linear-gradient(35deg, rgb(194, 33, 33), rgb(38, 194, 33));
        transition: box-shadow 0.3s ease-in-out, color 0.3s ease-in-out;      
    }
    .log-in:hover {
       color: rgb(255, 255, 255);
       box-shadow: -5px 0px 15px 0px rgba(194, 33, 33, 0.7), 5px 0px 15px 0px rgba(38, 194, 33, 0.7);
    }
</style>