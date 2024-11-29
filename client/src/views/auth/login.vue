<template>
    <ion-page>
        <ion-content>
            <ion-grid class="login-container">
                <ion-row class="ion-justify-content-center ion-align-items-center">
                    <ion-col size="12" size-md="8" size-lg="4" size-xl="8">
                        <ion-card class="login-card">
                            <app-label>
                                <h4>
                                    Login
                                </h4>
                            </app-label>
                            <app-input 
                                id="email"
                                type="email"
                                label="Email"
                                label-position="floating"
                                :counter="true"
                                :min-length="10"
                                :max-length="40"
                                style-class="white-border white-label"
                                v-model:input="email"
                                @on-input="onInput($event)"
                            />

                            <app-input 
                                id="password"
                                label="Password"
                                type="password"
                                label-position="floating"
                                :counter="true"
                                :min-length="8"
                                :max-length="16"
                                style-class="white-border white-label"
                                v-model:input="password"
                                @on-input="onInput($event)"
                            />

                            <app-switch
                                :title="userTypeLoginMessage"
                                :checked="typeSwitchValue"
                                @change="onAppSwitchChange($event)"
                            />

                            <app-button 
                                id="login-btn"
                                title="Login"
                                expand="block"
                                :action="loginUser"
                            />
                        </ion-card>
                    </ion-col>
                </ion-row>
            </ion-grid>
        </ion-content>
    </ion-page>
</template>

<script>

import { IonCard, IonCol, IonContent, IonGrid, IonPage, IonRow } from '@ionic/vue';
import AppLabel from '../../components/AppLabel.vue';
import AppInput from '../../components/AppInput.vue';
import AppButton from '../../components/AppButton.vue';
import {
    mapActions
} from 'vuex';
import AppSwitch from '../../components/AppSwitch.vue';


export default {
    name: 'Login',
    components: {
        IonPage,
        IonContent,
        IonGrid,
        IonRow,
        IonCol,
        IonCard,
        AppLabel,
        AppInput,
        AppButton,
        AppSwitch
    },
    data() {
        return {
            email: '',
            password: '',
            type: 'staff',
            typeSwitchValue: false,
            userTypeLoginMessage: 'Log In as Admin'
        }
    },
    methods: {
        ...mapActions('authStore', {
            signIn: 'login',
            staffLogin: 'staffLogin'
        }),
        onInput(event) {
        },
        async loginUser(event) {
            // here call login
            if(this.typeSwitchValue) {
                await this.signIn({
                    data: {
                        email: this.email,
                        password: this.password
                    },
                    appOptions: {
                        idsToBeUnblocked: event
                    }
                });
            } else {
                await this.staffLogin({
                    data: {
                        email: this.email,
                        password: this.password
                    },
                    appOptions: {
                        idsToBeUnblocked: event
                    }
                });
            }
            await this.$router.push('/books')
        },
        onAppSwitchChange(event) {
            const val = event.target.checked
            if(val) {
                this.userTypeLoginMessage = 'Log In as Admin'
                this.typeSwitchValue = false
            } else {
                this.userTypeLoginMessage = 'Log In as Staff'
                this.typeSwitchValue = true
            }
        }
    }
}
</script>

<style lang="css" scoped>

.login-container {
    height: 100vh;
    display: flex;
    justify-content: center;
    align-items: center;
    margin-bottom: 20px;
}

.login-card {
    text-align: center;
    padding: 20px;
    border-radius: 10px;
    background: linear-gradient(180deg, #412fa5 0%, #703ed9 100%);
    box-shadow: 0px 4px 6px rgba(0, 0, 0, 0.1);
    color: white;
    min-width: 300px;
    max-width: 450px;
}

</style>