<template>
    <div class="container login-container">
        <div class="row">
            <div class="col-lg-10 offset-lg-1 login-header">
                Unlock Keys
            </div>
        </div>
        <form id="login-form" @submit.prevent="submit" autocomplete="off">
            <div class="row mt-2">
                <div class="col-lg-10 offset-lg-1 mt-5 pword">
                    <input type="password" ref="pword" required />
                    <label v-bind:class="[invalidPassword ? 'invalid-password' : 'password-label']" for="password"> {{ passwordFieldMsg }}</label>
                </div>
            </div>
            <div class="row">
                <div class="col-lg-4 offset-lg-7 mt-2 forgot-password">
                    <!-- TODO: offer to delete keystore & reset profile keystore param. -->
                    <!-- For now put disc post showing how to do so -->
                    <!-- ensure done for beta -->
                    Forgot Password?
                </div>
            </div>
            <div class="row mt-3">
                <div class="mt-5">
                    <div v-bind:class="[unlockingKeys ? 'unlocking' : 'hidden']"> </div>
                    <button type="button" v-bind:class="[unlockingKeys ? 'hidden' : 'unlock']" v-on:click="logUserin()">UNLOCK</button>        
                </div>
            </div>
        </form>
    </div>
</template>

<script>
    import { defineComponent } from "@vue/runtime-core";

    export default defineComponent({
        name: "LoginForm",
        data() {
            return {
                passwordFieldMsg: "password",
                invalidPassword: false,
                unlockingKeys: false
            }
        },

        methods: {
            logUserin() {
                this.unlockingKeys = true;
                window.backend.Raverte.UnlockKeys(this.$refs.pword.value).then(() => {
                    this.$router.replace("/mainstage")
                }).catch((error) => {
                    if (error == "invalid password") {
                        // TODO: if x failed, offer to delete.
                        this.passwordFieldMsg = error
                        this.unlock = "Unlock";
                        this.unlockingKeys = false;
                        this.invalidPassword = true;
                    }else{
                        this.passwordFieldMsg = "password" // incase we entered incorrect password first, reset
                        this.invalidPassword = false;
                        console.log(error)
                        // TODO: error mechanism.
                    }
                });
            } 
        }
    });
</script>

<style lang="scss" scoped>
    .hidden {
        height: 0px;
        visibility: hidden;
    }
    
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
    input:focus ~ .password-label,
    input:valid ~ .password-label {
        font: normal normal 100 14px Poppins;
        color: rgba(255, 255, 255, 0.603);
        transform: translateY(-1.8rem);
    } 
    .invalid-password  {
        padding-bottom: 0.3rem;
        position: absolute;
        font: normal normal 300 17px Poppins;
        color: rgba(194, 33, 33, 0.9);
        transform: translateY(-1.8rem);
        text-transform: capitalize;
        pointer-events: none;
        
    }
    .password-label {
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
    .unlock {
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
    .unlock:hover {
       color: rgb(255, 255, 255);
       box-shadow: -5px 0px 15px 0px rgba(194, 33, 33, 0.7), 5px 0px 15px 0px rgba(38, 194, 33, 0.7);
    }
    .unlocking {
        // This hero: https://stackoverflow.com/a/67237610
        height: 45px;
        width: 45px;
        margin: 0 auto;
         --border-width: 8px;
        border-radius: 50%;
        --mask: radial-gradient(
            farthest-side, 
            transparent calc(100% - var(--border-width) - 0.5px), 
            #000 calc(100% - var(--border-width) + 0.5px)
        );
        -webkit-mask: var(--mask);
        mask: var(--mask);
        background: linear-gradient(to top, rgba(38, 194, 33, 0.9), rgba(194, 33, 33, 0.7)) 100% 0/50% 100% no-repeat,
              linear-gradient(rgba(194, 33, 33, 0.7) 50%, transparent 95%) 0 0/50% 100% no-repeat;
        animation: spin 1s linear infinite;
        visibility: visible;  
    }

    @keyframes spin {
        0% {
            transform: rotate(0deg);
        }
        100% {
            transform: rotate(360deg);
        }
    }
</style>